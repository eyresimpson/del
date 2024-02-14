package tools

import (
	"fmt"
	"net"
	"strconv"
)

func PortCheck(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", strconv.Itoa(port)))

	if err != nil {
		return false
	}
	defer l.Close()
	return true
}
