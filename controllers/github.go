package controllers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	config2 "go_webhook/config"
	"os/exec"
	"strings"
)

type GithubHeader struct {
	XGithubEvent     string `header:"X-GitHub-Event"`
	XHubSignature    string `header:"X-Hub-Signature"`
	XHubSignature256 string `header:"X-Hub-Signature-256"`
}

type GithubPayload struct {
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	Ref string `json:"ref"`
}

func ParseGithub(c *gin.Context) {
	config, err := config2.FileGetContent("./config.json")
	if err != nil {
		c.String(200, "server error")
		return
	}

	var header GithubHeader
	body, err := c.GetRawData()
	if err != nil {
		c.String(200, "server error")
		return
	}

	err = c.BindHeader(&header)
	if err != nil {
		c.String(200, "server error")
		return
	}
	var payload GithubPayload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		c.String(200, "server error")
		return
	}

	for _, v := range config.GitHub {

		mac := hmac.New(sha256.New, []byte(v.Password))
		_, _ = mac.Write(body)
		expectedMAC := hex.EncodeToString(mac.Sum(nil))
		if !hmac.Equal([]byte(header.XHubSignature256[7:]), []byte(expectedMAC)) {
			continue
		}

		if v.Name != payload.Repository.FullName {
			continue
		}

		if v.Ref != payload.Ref {
			continue
		}

		if v.HookName != header.XGithubEvent {
			continue
		}

		if len(v.Cmds) == 0 {
			break
		}
		for _, cmd := range v.Cmds {
			arrCmd := strings.Split(cmd, " ")
			commandName := arrCmd[0]
			_ = exec.Command(commandName, arrCmd[1:]...).Run()
		}
	}
	c.String(200, "github hook")
}
