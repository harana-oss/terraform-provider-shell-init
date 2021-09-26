Terraform Provider Shell Init
==========================================================
https://registry.terraform.io/providers/harana-oss/shell-init

This provider executes a sequence of shell scripts when it is initialised.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.12.x
-	[Go](https://golang.org/doc/install) >= 1.12

Building The Provider
---------------------

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```

Adding Dependencies
---------------------

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.


Using the Provider
----------------------

Set the following environment variable: TF_INIT_SCRIPTS to a list of script files e.g. TF_INIT_SCRIPTS="script1.sh, script2.sh".

Then configure the provider as follows

```
terraform {
  required_providers {
    shell = {
      source = "harana-oss/shell-init"
      version = "1.0"
    }
  }
}

provider "shell" {}
```