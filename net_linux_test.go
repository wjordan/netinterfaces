package netinterfaces

import (
	"net"
	"testing"
)

func BenchmarkNetInterfaces(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ifs, err := NetInterfaces()
		if err != nil {
			b.Fatal(err)
		}
		for _, iface := range ifs {
			_, err := iface.Addrs()
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}

func BenchmarkOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ifs, err := net.Interfaces()
		if err != nil {
			b.Fatal(err)
		}
		for _, iface := range ifs {
			_, err := iface.Addrs()
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}
