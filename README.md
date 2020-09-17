go-git-sync
===========

go-git-sync will sync JSON or YAML files to Consul when they change in Git.

## Usage

Basic usage:

```
$ go-git-sync -c http://consul:8500 sync \
    -s https://github.com/jmoeser/go-git-sync.git \
    -f example/consul/sample-json.json
```

For example, given a directory `test` in the repo `https://github.com/jmoeser/test-gitops-repo` any JSON or YAML files will be synced to the path `test/` in Consul.

Contents of the repo

```
$ tree test/
test/
└── a.json
$ cat test/a.json
{
    "test": "yes",
    "data": "no"
}
```

Will produce the following in Consul:

```
$ consul kv  get test/a
{
    "test": "yes",
    "data": "no"
}
```

Changes to these files pushed to Git will update the values as they are in Consul.

Currently the application will poll Git for changes every 3-5 minutes.

### Config file

Config files are supported you can pass in all the values in the config file instead of passing them in as arguments.

```
$ go-git-sync --config example/config.yaml sync
```

### Environment variables

Instead of using a config file or command arguments you can use environment variables, just prefix them with `GGS_`. For example:

```
$ export GGS_FILE=example/consul/sample-json.json
$ export GGS_SOURCE=https://github.com/jmoeser/go-git-sync.git
$ export GGS_CONSUL="127.0.0.1:8500"
$ go-git-sync sync
```

To Do:

- Github Webhooks so we don't need to poll
- Store state (last synced Git hash) in Consul instead of SQLite
- Check if hash in Git differs from recorded hash
- Perform sync if hash is different
- Record current Git hash in Consul
- Diff what's in Consul with what we got?
- Option to auto-heal, if the value in Consul changes change it back to what's in Git.
- Vault support (with an intermediate decrypt step)
