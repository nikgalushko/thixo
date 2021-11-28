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
