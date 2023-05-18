package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RancherService struct {
	opt Options
}

func NewRnacherService(opt Options) *RancherService {
	return &RancherService{
		opt: opt,
	}
}

func (s *RancherService) Run() {
	proData, err := s.getProject()
	if err != nil {
		fmt.Printf("get project data error")
		return
	}
	deploymentData, err := s.getDeployment(proData)
	if err != nil {
		fmt.Println("get deployment data error")
		return
	}
	updateData, err := s.getUpdateData(deploymentData)
	if err != nil {
		fmt.Println("get update data error")
		return
	}
	err = s.updateTag(deploymentData, updateData)
	if err != nil {
		fmt.Println("update error")
		return
	}
}

func (s *RancherService) getProject() (*ProjectData, error) {
	var resp GetProjectResponse
	reqUrl := fmt.Sprintf("%s%s", s.opt.Host, GetProjectV3Url)
	data, err := s.get(reqUrl)
	if err != nil {
		fmt.Printf("get project error:%+v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("parse get project response error:%+v", err)
		return nil, err
	}

	for _, d := range resp.Data {
		if d.Name == s.opt.Project {
			return &d, nil
		}
	}
	return nil, errors.New("no find project")
}

func (s *RancherService) getDeployment(projectData *ProjectData) (*DeploymentData, error) {
	if projectData == nil {
		return nil, errors.New("no find project")
	}
	var resp GetDeploymentResponse
	data, err := s.get(projectData.Links.Deployments)
	if err != nil {
		fmt.Printf("get deployment error:%+v", err)
		return nil, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		fmt.Printf("parse get deployment response error:%+v", err)
		return nil, err
	}
	for _, d := range resp.Data {
		if d.Name == s.opt.Deployment && d.NamespaceId == s.opt.Namespace {
			return &d, nil
		}
	}
	return nil, errors.New("no find deployment")
}

func (s *RancherService) getUpdateData(deploymentData *DeploymentData) ([]byte, error) {
	if deploymentData == nil {
		return nil, errors.New("no find deployment")
	}
	data, err := s.get(deploymentData.Links.Update)
	if err != nil {
		fmt.Printf("get update data error:%+v", err)
		return nil, err
	}
	return data, nil
}

func (s *RancherService) updateTag(deploymentData *DeploymentData, data []byte) error {
	if deploymentData == nil {
		return errors.New("no find deployment")
	}
	var (
		jsonMap          map[string]interface{}
		oriImage, oriTag string
	)
	err := json.Unmarshal(data, &jsonMap)
	if err != nil {
		fmt.Printf("parse update data error:%+v", err)
		return err
	}
	containers := jsonMap["containers"].([]interface{})
	for _, v := range containers {
		vMap := v.(map[string]interface{})
		cName := vMap["name"].(string)
		if cName != s.opt.Container {
			continue
		}
		imgUrl := vMap["image"].(string)
		imgSplice := strings.Split(imgUrl, ":")
		if len(imgSplice) != 2 {
			return errors.New("update images error")
		}
		oriImage = imgSplice[0]
		oriTag = imgSplice[1]
		if oriTag == s.opt.Tag {
			return nil
		}
		newImage := fmt.Sprintf("%s:%s", oriImage, s.opt.Tag)
		vMap["image"] = newImage
		break
	}
	var reqBody []byte
	reqBody, err = json.Marshal(jsonMap)
	if err != nil {
		fmt.Printf("parse update body error:%+v", err)
		return err
	}
	_, err = s.put(deploymentData.Links.Update, reqBody)
	if err != nil {
		fmt.Printf("update deployment error:%+v", err)
		return err
	}
	return nil
}

func (s *RancherService) get(addr string) ([]byte, error) {
	return s.req(addr, http.MethodGet, nil, 0)
}

func (s *RancherService) put(addr string, body []byte) ([]byte, error) {
	return s.req(addr, http.MethodPut, body, 0)
}

func (s *RancherService) req(addr, method string, data []byte, timeOut time.Duration) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableKeepAlives: false,
	}
	client := &http.Client{Transport: tr, Timeout: timeOut}

	reader := bytes.NewReader(data)

	req, err := http.NewRequest(method, addr, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.opt.Token))

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r, err := ioutil.ReadAll(resp.Body)
	return r, err
}
