// Author hoenig
// License MIT

package ssh

import "testing"

func compareToFile(filename string, expected []Key) error {
	reader := NewKeysReader()
	keys, err := reader.ReadKeys(filename)
	if err != nil {
		return err
	}

	return compareKeys(expected, keys)
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

	if err := compareToFile("../../hack/tests/authorized_keys.1", expected); err != nil {
		t.Fatal(err)
	}
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
		key(
			true,
			"alice",
			"h1",
			"ssh-rsa AAAAB3NzaD1yc2EAAAADAQABAAAAgQCoiewjflsdf==",
		),
		key(
			true,
			"alice",
			"h2",
			"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQDQ9Tf0lycls==",
		),
		key(
			true,
			"bob",
			"h2",
			"ssh-edsa AAAAB3NzaC1yc2EAdlksjflksdjflksdjflsdkdj==",
		),
	}

	if err := compareToFile("../../hack/tests/authorized_keys.2", expected); err != nil {
		t.Fatal(err)
	}
}
