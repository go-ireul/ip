package main

import (
	"fmt"
	"net"
	"strings"
)

// see https://en.wikipedia.org/wiki/Reserved_IP_addresses
var cidrs = []string{
	"0.0.0.0/8",
	"10.0.0.0/8",
	"100.64.0.0/10",
	"127.0.0.0/8",
	"169.254.0.0/16",
	"172.16.0.0/12",
	"192.0.0.0/24",
	"192.0.2.0/24",
	"192.88.99.0/24",
	"192.168.0.0/16",
	"198.18.0.0/15",
	"198.51.100.0/24",
	"203.0.113.0/24",
	"224.0.0.0/4",
	"240.0.0.0/4",
	"255.255.255.255/32",
	"::/128",
	"::1/128",
	//"::ffff:0:0/96", IPv4 mapped, ignored
	//"64:ff9b::/96", IPv4 translation, ignored
	"100::/64",
	//"2001::/32", Teredo tunneling, ignored
	"2001:10::/28",
	"2001:20::/28",
	"2001:db8::/32",
	//"2002::/16", 6to4, ignored
	"fc00::/7",
	"fe80::/10",
	"ff00::/8",
}

func ln(s ...interface{}) {
	fmt.Println(s...)
}

func main() {
	// Generate the code
	ln("// Generated by tools/create_constants.go")
	ln()
	ln("package ip")
	ln()
	ln("import \"net\"")
	ln()
	ln("// ReservedIPAddressCIDRs CIDR for reserved IP addresses")
	ln("var ReservedIPAddressCIDRs = []string{")
	for _, cidr := range cidrs {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(err)
		}
		if ipnet.String() != cidr {
			panic("Bad CIDR" + cidr)
		}
		ln("	\"" + cidr + "\",")
	}
	ln("}")
	ln()
	ln("// ReservedIPNets IPNet for reserved IP addresses")
	ln("var ReservedIPNets = []*net.IPNet{")
	for _, cidr := range cidrs {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(err)
		}
		ln("	{")
		hexs := []string{}
		for _, b := range ipnet.IP {
			hexs = append(hexs, fmt.Sprintf("0x%02x", b))
		}
		ln("		IP:   net.IP{" + strings.Join(hexs, ", ") + "},")
		hexs = []string{}
		for _, b := range ipnet.Mask {
			hexs = append(hexs, fmt.Sprintf("0x%02x", b))
		}
		ln("		Mask: net.IPMask{" + strings.Join(hexs, ", ") + "},")
		ln("	},")
	}
	ln("}")
}
