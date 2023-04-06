package ssh

import (
	"testing"

	"github.com/shoenig/test/must"
)

func compareToFile(t *testing.T, filename string, expected []Key) {
	reader := NewKeysReader()
	keys, err := reader.ReadKeys(filename)
	must.NoError(t, err)
	must.SliceContainsAll(t, expected, keys.Slice())
}

func Test_read_1(t *testing.T) {
	expected := []Key{
		key(
			false,
			"alice",
			"h1",
			"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCwyDsTUwFCw==",
		),
	}

	compareToFile(t, "../../hack/tests/authorized_keys.1", expected)
}

func Test_read_2(t *testing.T) {
	expected := []Key{
		key(
			false,
			"alice",
			"h1",
			"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQCwyDsTUwFCw==",
		),
		key(
			false,
			"bob",
			"h1",
			"ssh-dss AAAAB3NzaC1kc3MAAACBAIY/WCgvvatRJG7vdh7lk==",
		),
		key(
			false,
			"bob",
			"h1",
			"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAADJKlkjkfjlkjff==",
		),
	}

	compareToFile(t, "../../hack/tests/authorized_keys.2", expected)
}
