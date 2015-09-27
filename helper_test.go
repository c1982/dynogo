package main

import (
	"net"
	"testing"
)

func Test_GetWANIP(t *testing.T) {

	ipStr, err := GetWANIP()

	if err != nil {
		t.Log("Function Error!")
		t.Error(err.Error())
	}

	if ipStr == "" {
		t.Error("IP Address not determine")
	}

	IpAddr := net.ParseIP(ipStr)

	if IpAddr.To4() == nil {
		t.Errorf("Invalid IPv4 Address %v", ipStr)
	} else {
		t.Logf("WAN IP Address is %v", ipStr)
	}
}
