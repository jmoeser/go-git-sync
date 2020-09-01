go-git-sync
===========

Simple application that will sync somthing when the source in Git changes.

Will initially aim for Consul and Vault (with a decrypt intermediate step for Vault).

## Usage

Basic usage:

```
$ go-git-sync -c localhost:8500 sync \
    -s https://github.com/jmoeser/go-git-sync.git \
    -f example/consul/sample.json
```
