package errhandling

import (
	"fmt"
	"net"
)

func ListenAcc(host string, port uint16) (net.Listener, []error) {
	var er []error
	addr, addrErr := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
	if addrErr != nil {
		er = append(er, fmt.Errorf("Listen: %s", addrErr))
	}

	listener, listenError := net.ListenTCP("tcp", addr)
	if listenError != nil {
		er = append(er, fmt.Errorf("Listen: %s", listenError))
	}

	return listener, er
}
