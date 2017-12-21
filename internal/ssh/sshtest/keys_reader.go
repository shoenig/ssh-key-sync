// Code autogenerated by mockery v2.0.0
//
// Do not manually edit the content of this file.

// Package sshtest contains autogenerated mocks.
package sshtest

import "github.com/stretchr/testify/mock"
import "github.com/shoenig/ssh-key-sync/internal/ssh"

// KeysReader is an autogenerated mock type for the KeysReader type
type KeysReader struct {
	mock.Mock
}

// ReadKeys provides a mock function with given fields: filename
func (mockerySelf *KeysReader) ReadKeys(filename string) ([]ssh.Key, error) {
	ret := mockerySelf.Called(filename)

	var r0 []ssh.Key
	if rf, ok := ret.Get(0).(func(string) []ssh.Key); ok {
		r0 = rf(filename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]ssh.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
