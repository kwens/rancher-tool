package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update  [ --host=IP:PORT ] [ --token= ] ",
	Short: "rancher 更新工具 ",
	Long:  "rancher 更新工具 ",
	Example: `	rancher 更新工具 
	命令如：
	rancher update --host=https://192.168.xx.xx:xx --token=xxx -p yourpro`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(opt.Host) == 0 {
			fmt.Println("invalid argument, flag: host is null")
			_ = cmd.Help()
			return
		}
		if len(opt.Project) == 0 {
			fmt.Println("invalid argument, flag: project is null")
			_ = cmd.Help()
			return
		}

		if len(opt.Deployment) == 0 {
			fmt.Println("invalid argument, flag: deployment is null")
			_ = cmd.Help()
			return
		}
		s := NewRnacherService(opt)
		s.Run()
	},
}

var (
	opt Options
)

func init() {
	var flags = UpdateCmd.PersistentFlags()
	flags.StringVar(&opt.Host, "host", "", "服务器地址")
	flags.StringVar(&opt.Token, "token", "", "访问token，请从rancher创建{ACCESS_KEY:ACCESS_SECRET}")
	flags.StringVarP(&opt.Project, "project", "p", "", "项目名")
	flags.StringVarP(&opt.Namespace, "namespace", "n", "", "命名空间名")
	flags.StringVarP(&opt.Deployment, "deployment", "d", "", "工作负载名称")
	flags.StringVarP(&opt.Container, "container", "c", "", "容器名称")
	flags.StringVarP(&opt.Tag, "tag", "t", "", "镜像tag")
}
