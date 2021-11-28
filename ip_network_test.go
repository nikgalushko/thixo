package thixo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCidrHost(t *testing.T) {
	tests := map[string]struct {
		prefix string
		num    int
	}{
		"10.12.112.16":            {"10.12.127.0/20", 16},
		"10.12.113.12":            {"10.12.127.0/20", 268},
		"fd00:fd12:3456:7890::22": {"fd00:fd12:3456:7890:00a2::/72", 34},
	}

	for expected, in := range tests {
		require.Equal(t, expected, cidrhost(in.prefix, in.num))
	}
}

func TestCidrNetMask(t *testing.T) {
	require.Equal(t, "255.240.0.0", cidrnetmask("172.16.0.0/12"))
}

func TestCidrSubnet(t *testing.T) {
	tests := map[string]struct {
		prefix       string
		newBits, num int
	}{
		"172.18.0.0/16":                 {"172.16.0.0/12", 4, 2},
		"10.1.2.240/28":                 {"10.1.2.0/24", 4, 15},
		"fd00:fd12:3456:7800:a200::/72": {"fd00:fd12:3456:7890::/56", 16, 162},
	}

	for expected, in := range tests {
		require.Equal(t, expected, cidrsubnet(in.prefix, in.newBits, in.num))
	}
}
