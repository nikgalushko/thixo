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

func cidrsubnet(prefix string, newBits, num int) string {
	_, network, err := net.ParseCIDR(prefix)
	if err != nil {
		return ""
	}

	ip, err := cidr.SubnetBig(network, newBits, big.NewInt(int64(num)))
	if err != nil {
		return ""
	}

	return ip.String()
}

func cidrsubnets(prefix string, newBits ...int) []string {
	if len(newBits) == 0 {
		return nil
	}

	for _, n := range newBits {
		if n < 1 || n > 32 {
			return nil
		}
	}

	_, network, err := net.ParseCIDR(prefix)
	if err != nil {
		return nil
	}

	networkNumberSize := len(network.IP) * 8
	prefixLength, _ := network.Mask.Size()
	current, _ := cidr.PreviousSubnet(network, newBits[0]+prefixLength)

	ret := make([]string, len(newBits))
	for i, length := range newBits {
		length += prefixLength
		if length > networkNumberSize {
			return nil
		}

		next, overflow := cidr.NextSubnet(current, length)
		if overflow || !network.Contains(next.IP) {
			return nil
		}

		current = next
		ret[i] = current.String()
	}

	return ret
}
