ssh-key-sync
============

Command `ssh-key-sync` is a tool for managing `authorized_key` files, by synchronizing
the contents with public keys managed by accounts on github or gitlab instances.

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/cmds/ssh-key-sync)](https://goreportcard.com/report/gophers.dev/cmds/ssh-key-sync)
[![Build Status](https://travis-ci.com/shoenig/ssh-key-sync.svg?branch=master)](https://travis-ci.com/shoenig/ssh-key-sync)
[![GoDoc](https://godoc.org/gophers.dev/cmds/ssh-key-sync?status.svg)](https://godoc.org/gophers.dev/cmds/ssh-key-sync)
[![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/ssh-key-sync.svg)](OSSMETADATA)
[![GitHub](https://img.shields.io/github/license/shoenig/ssh-key-sync.svg)](LICENSE)

# Project Overview

Module `gophers.dev/cmds/ssh-key-sync` provides the command `ssh-key-sync`.

# Getting Started

The `ssh-key-sync` command can be installed by running
```
$ go install gophers.dev/cmds/ssh-key-sync@latest
```

#### Example Usage
There is only one argument, `--configfile` which specifies the location of the config file
that `ssh-key-sync` will read on startup.

```golang
ssh-key-sync --configfile /etc/ssh-key-sync.json
```

#### Configuration
In the configuration file, specify a list of system user accounts and associated SSH authorized_keys
files to manage. Also specify a set of github accounts, each with an associated github username and
local system user. The public SSH keys will be pulled from that github or gitlab account and unionized
with the keys in the specified authorized_keys file for that user. Keys which have been removed from github
or gitlab are automatically removed from the authorized_keys file. Keys which were added independent of
github or gitlab are left untouched.

Github SSH public keys are made available to the public. Gitlab SSH keys are accessible only from an
administrative account, so a service user with an API token will be required.

Use the following example as a template for creating a configuration file.

```json
{
    "system": [
        {"user": "clarkk", "authorized_keys_file": "/home/clarkk/.ssh/authorized_keys"},
        {"user": "bob", "authorized_keys_file": "/home/bob/.ssh/authorized_keys"}
    ],

    "github": {
        "url": "api.github.com",
        "accounts": [
            {"username": "superman", "system_user": "clarkk"}
        ]
    },

    "gitlab": {
        "url": "internal.gitlab.net",
        "token": "_jMr-KrDoy8GChTm998a",
        "accounts": [
            {"username":"billy", "system_user":"bob"}
        ]
    }
}
```

#### Systemd Timer
A great way to keep authorized_key files updated is to run `ssh-key-sync` periodically
via a systemd timer. To set this up, we will need two files - one service file which
represents execution of `ssh-key-sync`, and a timer file which represents the schedule
on which the service should be executed. Use the example below, modifying paths to
suite your needs. More examples can be found in this [blog post](https://jason.the-graham.com/2013/03/06/how-to-use-systemd-timers/).

##### The service file `/etc/systemd/system/ssh-key-sync.service`
```
[Unit]
Description=Synchronize ssh authorized keys with public keys from github.

[Service]
ExecStart=/opt/keys/ssh-key-sync --configfile /etc/ssh-key-sync.json
```

##### The timer file `/etc/systemd/system/ssh-key-sync.timer`
```
[Unit]
Description=Run ssk-key-sync every hour

[Timer]
OnBootSec=5min
OnUnitActiveSec=1h
Unit=ssh-key-sync.service

[Install]
WantedBy=timers.target
```

##### Enable the timer
```
$ systemctl enable ssh-key-sync.timer
```

# Contributing

The `gophers.dev/pkgs/regexplus` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License
[MIT](https://raw.githubusercontent.com/shoenig/ssh-key-sync/master/LICENSE)
