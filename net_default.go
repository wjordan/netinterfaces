//go:build !linux
// +build !linux

package netinterfaces

import "net"

func NetInterfaces() ([]net.Interface, error) {
	return net.Interfaces()
}
