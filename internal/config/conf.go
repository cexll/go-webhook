package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conf struct {
	GitHub []GithubConf `json:"github"`
	Gitee  []GiteeConf  `json:"gitee"`
}

type GithubConf struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Ref      string   `json:"ref"`
	HookName string   `json:"hook_name"`
	Cmds     []string `json:"cmds"`
}

type GiteeConf struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Ref      string   `json:"ref"`
	HookName string   `json:"hook_name"`
	Cmds     []string `json:"cmds"`
}

func FileGetContent(path string) (*Conf, error) {
	fmt.Println("file get path", path)
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var config Conf
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
