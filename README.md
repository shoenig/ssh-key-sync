ssh-key-sync
============

`ssh-key-sync` is a tool written in Go for managing `authorized_key` files, by synchronizing
the contents with public keys managed by accounts on https://github.com

### Install
Currently, `ssh-key-sync` must be compiled and installed manually. With a typical Go workspace,
run `go get github.com/shoenig/ssh-key-sync` to produce a binary. Copy that binary to the destination
server somewhere on `$PATH`.

### Run
There is only one argument, `--configfile` which specifies the location of the config file
that `ssh-key-sync` will read on startup.

### Configuration
In the configuration file, specify a list of system user accounts and associated SSH authorized_keys
files to manage. Also specify a set of github accounts, each with an associated github username and
local system user. The public SSH keys will be pulled from that github account and unionized with the
keys in the specified authorized_keys file for that user. Keys which have been removed from github are
automatically removed from the authorized_keys file. Keys which were added independent of github are left
untouched.

Use the following example as a template for creating a configuration file.
```
{
    "system": [
        {"user": "clarkk", "authorized_keys_file": "/home/clarkk/authorized_keys"},
    ],

    "github": {
        "url": "github.com",
        "accounts": [
            {"username": "superman", "system_user": "clarkk"},
        ]
    }
}
```