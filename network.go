package thixo

import (
	"crypto/rand"
	"math/big"
	"net"
)

func getHostByName(name string) string {
	addrs, _ := net.LookupHost(name)
	i, _ := rand.Int(rand.Reader, big.NewInt(int64(len(addrs))))
	//TODO: add error handing when release comes out

	return addrs[i.Int64()]
}
