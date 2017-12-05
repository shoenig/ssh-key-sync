// Author hoenig

package ssh

import "fmt"

type Key struct {
	Managed bool
	Value   string
	User    string
	Host    string
}

func (k Key) String() string {
	return fmt.Sprintf(
		"[managed:%t user:%s host:%s value:%s]",
		k.Managed,
		k.User,
		k.Host,
		k.Value,
	)
}

type sortByMetadata []Key

func (s sortByMetadata) Len() int {
	return len(s)
}

func (s sortByMetadata) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s sortByMetadata) Less(a, b int) bool {
	// unmanaged keys should be listed before managed keys
	if s[a].Managed && !s[b].Managed {
		return false
	} else if !s[a].Managed && s[b].Managed {
		return true
	}

	return s.metaLess(a, b)
}

func (s sortByMetadata) metaLess(a, b int) bool {
	if s[a].User < s[b].User {
		return true
	} else if s[a].User > s[b].User {
		return false
	}

	if s[a].Host < s[b].Host {
		return true
	} else if s[a].Host > s[b].Host {
		return false
	}

	return s[a].Value < s[b].Value
}
