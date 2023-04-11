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

# Install

The following steps should get `ssh-key-sync` running as a systemd-timer on any system with systemd.
You could also use cron or some other periodic task runner, but these instructions work with most major Linux distrobutions.

These examples use `linux` and `amd64` - be sure to use the correct version for your operating system and architecture.

#### Download Archive

```shell-session
$ wget https://github.com/shoenig/ssh-key-sync/releases/download/v1.7.1/ssh-key-sync_1.7.1_linux_amd64.tar.gz
```

#### Extract Archive

```shell-session
$ sudo tar -C /usr/local/bin -xf ssh-key-sync_1.7.1_linux_amd64.tar.gz
```

#### Create `authorized_keys`

```shell-session
$ mkdir ~/.ssh && chmod 700 ~/.ssh
$ touch ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys
```

#### Setup Systemd Timer
A great way to keep your `authorized_key` file up-to-date is to run `ssh-key-sync` periodically via a systemd timer.
To set this up, we will need two files - one service file which represents execution of `ssh-key-sync`, 
and a timer file which represents the schedule on which the service should be executed. Use the example below, 
modifying paths to suite your needs.

##### The service file `/etc/systemd/system/ssh-key-sync.service`

```
[Unit]
Description=Synchronize ssh authorized keys with public keys from github.

[Service]
ExecStart=/usr/local/bin/ssh-key-sync -verbose -system-user <user> -github-user <user>
```

##### The timer file `/etc/systemd/system/ssh-key-sync.timer`

```
[Unit]
Description=Run ssk-key-sync every 6 hours

[Timer]
OnBootSec=5min
OnUnitActiveSec=6h
Unit=ssh-key-sync.service

[Install]
WantedBy=timers.target
```

##### Enable the timer

```shell-session
$ sudo systemctl daemon-reload
$ sudo systemctl enable ssh-key-sync.timer
```

#### See if it works

The timer will run in the background on its schedule, but it's nice to run the service immediately so we can
see that it is working correctly. To do that, just start the service.

```shell-session
$ sudo systemctl start ssh-key-sync.service
```

And now we can view the status to make sure it ran correctly.

```shell-session
$ sudo systemctl status ssh-key-sync
```

You should see some output like, 

```logs
○ ssh-key-sync.service - Synchronize ssh authorized keys with public keys from github.
     Loaded: loaded (/etc/systemd/system/ssh-key-sync.service; static)
     Active: inactive (dead) since Tue 2023-04-11 15:29:32 CDT; 2s ago
   Duration: 236ms
TriggeredBy: ○ ssh-key-sync.timer
    Process: 2104 ExecStart=/usr/local/bin/ssh-key-sync -verbose -system-user shoenig -github-user shoenig (code=exited>
   Main PID: 2104 (code=exited, status=0/SUCCESS)
        CPU: 22ms

Apr 11 15:29:32 localhost.localdomain systemd[1]: Started Synchronize ssh authorized keys with public keys from github..
Apr 11 15:29:32 localhost.localdomain ssh-key-sync[2104]: 2023/04/11 15:29:32 using default output authorized_keys file>
Apr 11 15:29:32 localhost.localdomain ssh-key-sync[2104]: 2023/04/11 15:29:32 process local user shoenig from shoenig@g>
Apr 11 15:29:32 localhost.localdomain ssh-key-sync[2104]: 2023/04/11 15:29:32 loaded 0 existing keys for user "shoenig"
Apr 11 15:29:32 localhost.localdomain ssh-key-sync[2104]: 2023/04/11 15:29:32 acquire github keys from "https://api.git>
Apr 11 15:29:32 localhost.localdomain ssh-key-sync[2104]: 2023/04/11 15:29:32 retrieved 10 keys for github user: shoenig
Apr 11 15:29:32 localhost.localdomain systemd[1]: ssh-key-sync.service: Deactivated successfully.
```

It works! And inspecting the `~/.ssh/authorized_keys` file should reveal it contains your public keys from GitHub.

# Contributing

The `github.com/shoenig/ssh-key-sync` tool is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License
[MIT](https://raw.githubusercontent.com/shoenig/ssh-key-sync/master/LICENSE)
