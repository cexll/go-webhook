package controllers

import (
	config2 "go_webhook/internal/config"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type GiteePayload struct {
	Ref      string `json:"ref"`
	HookName string `json:"hook_name"`
	Password string `json:"password"`
	Project  struct {
		PathWithNameSpace string `json:"path_with_namespace"`
	} `json:"project"`
}

func ParseGitee(c *gin.Context) {
	config, err := config2.FileGetContent("./config.json")
	if err != nil {
		c.String(200, "server error")
		return
	}
	var payload GiteePayload
	err = c.Bind(&payload)
	if err != nil {
		c.String(200, "server error")
		return
	}

	for _, v := range config.Gitee {
		if v.Name != payload.Project.PathWithNameSpace {
			continue
		}
		if v.Password != payload.Password {
			continue
		}
		if v.HookName != payload.HookName {
			continue
		}
		if v.Ref != payload.Ref {
			continue
		}
		if len(v.Cmds) == 0 {
			break
		}
		for _, cmd := range v.Cmds {
			_ = exec.Command("bash", "-c", cmd).Run()
		}
	}
	c.String(200, "gitee hook")
}
