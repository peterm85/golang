package errhandling

import (
	"fmt"
	"net"
)

func Listen(host string, port uint16) (net.Listener, error) {
	addr, addrErr := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
	if addrErr != nil {
		return nil, fmt.Errorf("Listen: %s", addrErr)
	}

	listener, listenError := net.ListenTCP("tcp", addr)
	if listenError != nil {
		return nil, fmt.Errorf("Listen: %s", listenError)
	}

	return listener, nil
}
