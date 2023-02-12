ssh-key-sync
============

Command `ssh-key-sync` is a tool for managing `authorized_key` files, by synchronizing
the contents with public keys managed by accounts on [GitHub](https://docs.github.com/en/rest/users/keys).

[![CI](https://github.com/shoenig/ssh-key-sync/actions/workflows/ci.yml/badge.svg)](https://github.com/shoenig/ssh-key-sync/actions/workflows/ci.yml)
[![GitHub](https://img.shields.io/github/license/shoenig/ssh-key-sync.svg)](LICENSE)

# Project Overview

Module `github.com/shoenig/ssh-key-sync` provides the command `ssh-key-sync`.

# Getting Started

The `ssh-key-sync` command can be downloaded from the [releases](https://github.com/shoenig/ssh-key-sync/releases) page.

Alternatively, you can use Go directly.
```shell
go install github.com/shoenig/ssh-key-sync@latest
```

#### Example Usage

There are a few arguments, but typically you should only need to specify `--github-user`.
The Linux user and `authorized_key` file are by default assumed to be of the user
running the command.

```shell
ssh-key-sync --github-user <user>
```

#### Configuration

```shell
ssh-key-sync -help
Usage of ./ssh-key-sync:
  -authorized-keys string
    	override the output authorized_keys file (default "/home/$USER/.ssh/authorized_keys")
  -github-api string
    	specify the GitHub API endpoint (default "https://api.github.com")
  -github-user string
    	specify the github user
  -system-user string
    	specify the unix system user (default "$USER")
  -prune
        delete all keys not found in github
  -verbose
    	print verbose logging
```

#### Systemd Timer
A great way to keep your `authorized_key` file up-to-date is to run `ssh-key-sync`
periodically via a systemd timer. To set this up, we will need two files - one service
file which represents execution of `ssh-key-sync`, and a timer file which represents
the schedule on which the service should be executed. Use the example below, modifying
paths to suite your needs. More examples can be found in this [blog post](https://jason.the-graham.com/2013/03/06/how-to-use-systemd-timers/).

##### The service file `/etc/systemd/system/ssh-key-sync.service`
```
[Unit]
Description=Synchronize ssh authorized keys with public keys from github.

[Service]
ExecStart=/opt/bin/ssh-key-sync -verbose -system-user <user> -github-user <user>
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

The `github.com/shoenig/ssh-key-sync` tool is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License
[MIT](https://raw.githubusercontent.com/shoenig/ssh-key-sync/master/LICENSE)
