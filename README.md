# kubectl-dapply

This plugin wraps `kubectl apply` with showing differences that will be applied and asking user to proceed or not prior to apply.

[![GoDoc](https://godoc.org/github.com/ardaguclu/kubectl-dapply?status.svg)](https://godoc.org/github.com/ardaguclu/kubectl-dapply)
[![Go Report Card](https://goreportcard.com/badge/github.com/ardaguclu/kubectl-dapply.svc)](https://goreportcard.com/report/github.com/ardaguclu/kubectl-dapply)

## Details

Currenly, if user wants to see the differences before apply, user has to manually run `kubectl diff` command. This plugin provides a combined way of `kubectl apply` and `kubectl diff` commands 
with user interaction. It shows differences first and if user decides to proceed, it applies resources exactly the same way of `kubectl apply` does.

## Installation

Use [krew](https://sigs.k8s.io/krew) plugin manager to install,

```shell script
wget https://github.com/ardaguclu/kubectl-dapply/blob/main/.krew.yaml
kubectl krew install --manifest=.krew.yaml
kubectl dapply --help
```

Or manually clone the repo and install into your $GOPATH;

```shell script
go install cmd/kubectl-dapply.go
```

## Usage

It is fully compatible with `kubectl apply` command.

[![asciicast](https://asciinema.org/a/uLIzpkgsgXyEdroXDVpDuvaOm.svg)](https://asciinema.org/a/uLIzpkgsgXyEdroXDVpDuvaOm)
