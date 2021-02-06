package network

import (
	"net"
)

// IsPortAvailable check if a port on localhost is available
func IsPortAvailable(port string) bool {

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return false
	}

	_ = ln.Close()

	return true
}
