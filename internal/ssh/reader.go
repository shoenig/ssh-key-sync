// Author hoenig

package ssh

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strings"
)

//go:generate mockery -interface=KeysReader -package sshtest

type KeysReader interface {
	ReadKeys(filename string) ([]Key, error)
}

func NewKeysReader() KeysReader {
	return &keysReader{}
}

type keysReader struct{}

func (kr *keysReader) ReadKeys(filename string) ([]Key, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	s := bufio.NewScanner(f)
	keys := make([]Key, 0, 10)
	managed := false
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		switch {
		case isManaged(line):
			managed = true
		case isIgnorable(line):
			continue
		default:
			key, err := ParseKey(line, managed)
			if err != nil {
				return nil, err
			}
			keys = append(keys, key)
			managed = false
		}
	}

	sort.Sort(KeySorter(keys))

	return keys, s.Err()
}

func isManaged(line string) bool {
	return strings.HasPrefix(line, "# managed by ssh-key-sync")
}

func isIgnorable(line string) bool {
	return line == "" || strings.HasPrefix(line, "#")
}

func ParseKey(line string, managed bool) (Key, error) {
	parts := strings.Fields(line)
	if len(parts) < 2 || len(parts) > 3 {
		return Key{}, errors.New("key format is not well formed")
	}

	value := parts[0] + " " + parts[1]
	var user string
	var host string
	if len(parts) == 3 {
		metadata := strings.Split(parts[2], "@")
		if len(metadata) == 2 {
			user = metadata[0]
			host = metadata[1]
		}
	}

	return Key{
		Managed: managed,
		Value:   value,
		User:    user,
		Host:    host,
	}, nil
}
