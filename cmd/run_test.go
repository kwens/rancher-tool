/**
 * @Author: kwens
 * @Date: 2023-12-26 14:53:26
 * @Description:
 */
package cmd

import "testing"

func TestRun(t *testing.T) {
	srv := NewRnacherService(Options{
		Host:       "https://192.168.xx.xx:xxxx",
		Token:      "token-xx:xxxx",
		Project:    "awsome-manage",
		Namespace:  "server",
		Deployment: "test",
		Container:  "test",
		Tag:        "v1.0.16",
	})
	srv.Run()
}
