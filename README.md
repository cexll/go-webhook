# go-webhook
Use the webhook written in go to provide convenience for cicd

# Install

```bash
git clone https://github.com/cexll/go-webhook.git
go build -o server cmd/server.go
```

# Run

```bash
./server -p 8080
```

> download package [release](https://github.com/cexll/go-webhook/releases)
# Configure

> ./config.json
```json
{
  "github": [
    {
      "name": "cexll/go-webhook",
      "password": "password",
      "ref": "refs/heads/master",
      "hook_name": "push",
      "cmds": [
        "echo \"Hello WebHook\" && git -C /yourpath/project pull"
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
        "cd /Hello/WebHook && git -C /yourpath/project pull"
      ]
    }
  ]
}
```
# License
Apache License Version 2.0, http://www.apache.org/licenses/
