# Learning The Go Vault Client

This directory is dedicated to learning how to use the Go client for HashiCorp Vault.

## ~BEFORE USE THE GO VAULT CLIENT~

The Go Vault client requires the use go `go mod`.

If you do not use `go mod` with your implementation of the Go Vault client, you will receive errors that look very similar to this:

```bash
[akalaj@server vault]$ go get github.com/hashicorp/vault/api

cannot find package "github.com/cenkalti/backoff/v3" in any of:
	/usr/local/go/src/github.com/cenkalti/backoff/v3 (from $GOROOT)
	/home/akalaj/projects/src/src/github.com/cenkalti/backoff/v3 (from $GOPATH)
cannot find package "github.com/hashicorp/hcl/hcl/ast" in any of:
	/usr/local/go/src/github.com/hashicorp/hcl/hcl/ast (from $GOROOT)
	/home/akalaj/projects/src/src/github.com/hashicorp/hcl/hcl/ast (from $GOPATH)
```

### Example Of How To use `go mod` With Your Project

Below you'll find an example of how to initialize `go mod` and use it with your project.

```bash
[akalaj@server vault]$ mkdir vault
[akalaj@server vault]$ cd vault
[akalaj@server vault]$  go mod init vault
go: creating new go.mod: module vault
```

Once this is complete, you will be able to execute a `go run` against your code.

```bash
[akalaj@server vault]$  go run vault-init.go
go: finding module for package github.com/hashicorp/vault/api
go: downloading github.com/hashicorp/vault/api v1.1.1
go: downloading github.com/hashicorp/vault v1.7.3
go: found github.com/hashicorp/vault/api in github.com/hashicorp/vault/api v1.1.1
go: downloading github.com/hashicorp/go-cleanhttp v0.5.1
go: downloading golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1
go: downloading github.com/hashicorp/errwrap v1.0.0
go: downloading github.com/hashicorp/vault/sdk v0.2.1
go: downloading github.com/hashicorp/hcl v1.0.0
go: downloading golang.org/x/net v0.0.0-20200602114024-627f9648deb9
go: downloading github.com/hashicorp/go-retryablehttp v0.6.6
go: downloading gopkg.in/square/go-jose.v2 v2.5.1
go: downloading github.com/hashicorp/go-multierror v1.1.0
go: downloading github.com/hashicorp/go-rootcerts v1.0.2
go: downloading github.com/mitchellh/mapstructure v1.3.2
go: downloading github.com/cenkalti/backoff/v3 v3.0.0
go: downloading golang.org/x/crypto v0.0.0-20200604202706-70a84ac30bf9
go: downloading github.com/hashicorp/go-sockaddr v1.0.2
go: downloading github.com/ryanuber/go-glob v1.0.0
go: downloading github.com/golang/snappy v0.0.1
go: downloading github.com/pierrec/lz4 v2.5.2+incompatible
go: downloading golang.org/x/text v0.3.2
```

# Example of the Go Vault client

### vault-init.go - Check if Vault is initialized

```go
package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
)

func main() {

	vault_addr	:=	"http://127.0.0.1:8200"
	token		:=	"ROOT_TOKEN_VALUE"

	config := &api.Config{
		Address: vault_addr,
	}

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	client.SetToken(token)

	initStatus, err := client.Sys().InitStatus()
	if err != nil {
		fmt.Println(err)
		return
	}

    //Print returns "True" if initialized
	fmt.Println(initStatus)
}
```