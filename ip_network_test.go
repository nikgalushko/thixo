package thixo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCidrHost(t *testing.T) {
	require.Equal(t, "10.12.112.16", cidrhost("10.12.127.0/20", 16))
	require.Equal(t, "10.12.113.12", cidrhost("10.12.127.0/20", 268))
	require.Equal(t, "fd00:fd12:3456:7890::22", cidrhost("fd00:fd12:3456:7890:00a2::/72", 34))
}
