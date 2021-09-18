---
layout: "shell"
page_title: "Provider: Shell Init"
sidebar_current: "docs-shell-index"
description: |-
  Terraform Provider Shell Init
---

# Shell Init Provider

This provider executes a shell script when it is initialised.

## Example Usage

```hcl
provider "shell" {
  shell_script = "${path.root}/init.sh" 

  sensitive_environment = {
    OAUTH_TOKEN = var.oauth_token
  }
}
```

Stdout and stderr stream to log files. You can get this by setting:

```
export TF_LOG=1
```
**Note:** if you are using sensitive_environment to set sensitive environment variables, these values won't show up in the logs