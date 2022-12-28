# go-webhook
Use the webhook written in go to provide convenience for cicd

# Installation

```bash
git clone https://github.com/cexll/go-webhook.git
go build -o webhook main.go
```

> download package [release](https://github.com/cexll/go-webhook/releases)
# Configure

```json
{
  "github": [
    {
      "name": "cexll/go-webhook",
      "password": "password",
      "ref": "refs/heads/master",
      "hook_name": "push",
      "cmds": [
        "git -C /yourpath/project pull"
      ]
    }
  ],
  "gitee": [
    {
      "name": "cexll/go-webhook",
      "password": "password",
      "ref": "refs/heads/master",
      "hook_name": "push_hooks",
      "cmds": [
        "git -C /yourpath/project pull"
      ]
    }
  ]
}
```
# License
Apache License Version 2.0, http://www.apache.org/licenses/ 