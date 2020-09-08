go-git-sync
===========

Simple application that will sync somthing when the source in Git changes.

Will initially aim for Consul and Vault (with a decrypt intermediate step for Vault).

## Usage

Basic usage:

```
$ go-git-sync -c localhost:8500 sync \
    -s https://github.com/jmoeser/go-git-sync.git \
    -f example/consul/sample-json.json
```

If `example/consul/sample-json.json` in the repo `https://github.com/jmoeser/go-git-sync.git` changes the contents of it will be synced to Consul under `example/consul/sample-json`.

Currently the application will poll Git for changes every 3-5 minutes.

To Do:

- Github Webhooks so we don't need to poll
- Store state (last synced Git hash) in Consul instead of SQLite
- Check if hash in Git differs from recorded hash
- Perform sync if hash is different
- Record current Git hash in Consul
- Diff what's in Consul with what we got?
