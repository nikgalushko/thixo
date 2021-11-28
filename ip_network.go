package thixo

import (
	"math/big"
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
)

func cidrhost(prefix string, num int) string {
	_, network, err := net.ParseCIDR(prefix)
	if err != nil {
		return ""
	}

	ip, err := cidr.HostBig(network, big.NewInt(int64(num)))
	if err != nil {
		return ""
	}

	return ip.String()
}

func cidrnetmask(prefix string) string {
	_, network, err := net.ParseCIDR(prefix)
	if err != nil {
		return ""
	}

	return net.IP(network.Mask).String()
}
