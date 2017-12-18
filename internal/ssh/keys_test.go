// Author hoenig
// License MIT

package ssh

import (
	"fmt"
	"sort"
	"testing"
)

func key(managed bool, user, host, value string) Key {
	return Key{
		Managed: managed,
		User:    user,
		Host:    host,
		Value:   value,
	}
}

func compareKeys(a, b []Key) error {
	if len(a) != len(b) {
		return fmt.Errorf("len a != b (%d != %d)", len(a), len(b))
	}

	for i := range a {
		keyA := a[i]
		keyB := b[i]
		if keyA != keyB {
			return fmt.Errorf("keys at %d do not match, a: %s, b: %s", i, keyA, keyB)
		}
	}
	return nil
}

func Test_Key_String_full(t *testing.T) {
	k := Key{
		Managed: true,
		Value:   "abc123",
		User:    "bob",
		Host:    "host1",
	}

	expected := "[managed:true user:bob host:host1 value:abc123]"
	s := k.String()
	if s != expected {
		t.Fatalf("key string is not as expected, s: %s, expected: %s", s, expected)
	}
}

func Test_Key_String_empty(t *testing.T) {
	k := Key{
		Managed: false,
		Value:   "abc123",
	}

	expected := "[managed:false user: host: value:abc123]"
	s := k.String()
	if s != expected {
		t.Fatalf("key string is not as expected, s: %s, expected: %s", s, expected)
	}
}

func Test_sortByMetadata(t *testing.T) {

	keys := []Key{
		key(true, "ned", "h5", "xcvwe"),
		key(false, "xavior", "h4", "lkdsf"),
		key(true, "bob", "h1", "abcdefg"),
		key(true, "bob", "h2", "eioqije"),
		key(false, "bob", "h3", "oiwejre"),
		key(true, "ned", "h5", "zzzzz"),
		key(false, "ned", "h5", "aaaaa"),
		key(true, "alice", "h1", "klsdjfd"),
		key(false, "alice", "h2", "ioweffs"),
		key(false, "alice", "h1", "iznei"),
	}

	sort.Sort(KeySorter(keys))

	expected := []Key{
		key(false, "alice", "h1", "iznei"),
		key(false, "alice", "h2", "ioweffs"),
		key(false, "bob", "h3", "oiwejre"),
		key(false, "ned", "h5", "aaaaa"),
		key(false, "xavior", "h4", "lkdsf"),
		key(true, "alice", "h1", "klsdjfd"),
		key(true, "bob", "h1", "abcdefg"),
		key(true, "bob", "h2", "eioqije"),
		key(true, "ned", "h5", "xcvwe"),
		key(true, "ned", "h5", "zzzzz"),
	}

	if err := compareKeys(expected, keys); err != nil {
		t.Fatal(err)
	}
}
