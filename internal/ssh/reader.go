package ssh

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/hashicorp/go-set"
)

type KeysReader interface {
	ReadKeys(filename string) (*set.Set[Key], error)
}

func NewKeysReader() KeysReader {
	return new(reader)
}

type reader struct{}

func (r *reader) ReadKeys(filename string) (*set.Set[Key], error) {
	f, fErr := os.Open(filename)
	if fErr != nil {
		return nil, fErr
	}

	s := bufio.NewScanner(f)
	keys := set.New[Key](10)
	managed := false
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		switch {
		case isManaged(line):
			managed = true
		case isIgnorable(line):
			continue
		default:
			if managed {
				// managed keys will get renewed; we ignore them here so dead
				// keys get pruned as expected
				managed = false
				continue
			}
			key, err := ParseKey(line, managed)
			if err != nil {
				// log something
				continue
			}
			keys.Insert(key)
		}
	}

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
