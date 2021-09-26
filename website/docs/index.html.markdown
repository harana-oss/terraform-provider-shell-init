---
layout: "shell"
page_title: "Provider: Shell Init"
sidebar_current: "docs-shell-index"
description: |-
  Terraform Provider Shell Init
---

# Shell Init Provider

This provider executes a sequence of shell scripts when it is initialised.

## Example Usage

Set the following environment variable: TF_INIT_SCRIPTS to a list of script files e.g. TF_INIT_SCRIPTS="script1.sh, script2.sh".

Then configure the provider as follows:

```hcl
terraform {
  required_providers {
    shell = {
      source = "harana-oss/shell-init"
      version = "1.0.1"
    }
  }
}

provider "shell" {}
```