package balancer

import (
	"hash/fnv"
	"strings"
)

func GetAddressNode(addr string, nodes ...string) string {
	h := hash(addr)
	i := int(h % uint32(len(nodes)))
	a := strings.Split(nodes[i], "|")
	return a[1]
}

func hash(s string) uint32 {
	h := fnv.New32()
	h.Write([]byte(s))
	return h.Sum32()
}
