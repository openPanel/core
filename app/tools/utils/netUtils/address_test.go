package netUtils

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/openPanel/core/app/constant"
)

// TestCheckPublicIp tests the CheckPublicIp function with various types of IP addresses
func TestCheckPublicIp(t *testing.T) {
	testCases := []struct {
		name        string
		ip          string
		expectError error
	}{
		{"Valid public IP", "93.184.216.34", nil},
		{"ISP reserved IP", "100.64.0.1", errors.New("IP address is not global unicast address 100.64.0.1")},
		{"localhost", "127.0.0.1", errors.New("IP address is private 127.0.0.1")},
		{"Private IP", "192.168.0.1", errors.New("IP address is private 192.168.0.1")},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parsedIP := net.ParseIP(tc.ip)
			err := CheckPublicIp(parsedIP)
			if err != nil && err.Error() != tc.expectError.Error() {
				t.Errorf("Expected error '%v', got '%v'", tc.expectError, err)
			}
		})
	}
}

func TestAssertPublicAddress(t *testing.T) {
	type testCase struct {
		address        string
		expectedIP     net.IP
		expectedPort   int
		shouldPanic    bool
		panicMessage   string
		portInTestCase int
	}

	testCases := []testCase{
		// valid input without defining a port
		{
			address:      "8.8.8.8",
			expectedIP:   net.ParseIP("8.8.8.8"),
			expectedPort: constant.DefaultListenPort,
		},
		// valid input with a port
		{
			address:      "8.8.8.8:12345",
			expectedIP:   net.ParseIP("8.8.8.8"),
			expectedPort: 12345,
		},
		// invalid address without a colon
		{
			address:      "example",
			shouldPanic:  true,
			panicMessage: "Invalid IP address example",
		},
		// invalid IP address
		{
			address:      "300.300.300.300:12345",
			shouldPanic:  true,
			panicMessage: "Invalid IP address 300.300.300.300",
		},
		// invalid port number
		{
			address:      "8.8.8.8:70000",
			shouldPanic:  true,
			panicMessage: "Invalid port 70000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.address, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r != nil {
						err := fmt.Sprintf("%v", r)
						if err != tc.panicMessage {
							t.Errorf("Expected panic with message '%v', got '%v'", tc.panicMessage, err)
						}
					} else {
						t.Errorf("Expected a panic, but function did not panic")
					}
				}()
			}

			// Call the target function
			resultIp, resultPort := AssertPublicAddress(tc.address)

			// Check if the results match the expected ones
			if !resultIp.Equal(tc.expectedIP) {
				t.Errorf("Expected IP %s, got %s", tc.expectedIP.String(), resultIp.String())
			}

			if resultPort != tc.expectedPort {
				t.Errorf("Expected port %d, got %d", tc.expectedPort, resultPort)
			}
		})
	}
}