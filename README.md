ssh-key-sync
============

`ssh-key-sync` is a tool written in Go for managing `authorized_key` files, by synchronizing
the contents with public keys managed by accounts on https://github.com

[![Go Report Card](https://goreportcard.com/badge/github.com/shoenig/ssh-key-sync)](https://goreportcard.com/report/github.com/shoenig/ssh-key-sync)

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
        {"user": "clarkk", "authorized_keys_file": "/home/clarkk/.ssh/authorized_keys"},
    ],

    "github": {
        "url": "github.com",
        "accounts": [
            {"username": "superman", "system_user": "clarkk"},
        ]
    }
}
```

### Systemd Timer
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
ExecStart=/opt/keys/ssh-key-sync --configfile /opt/keys/config.json
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
systemctl enable ssh-key-sync.timer
```

### License
[MIT](https://raw.githubusercontent.com/shoenig/ssh-key-sync/master/LICENSE)