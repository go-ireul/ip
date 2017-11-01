package ip

import "net"

// IsReservedIP if ip is a reserved IP address
// reserved IP addresses are IP addresses you should not see from public internet
// if you are building a service that sends requests on user's demand, you can use this method to check if user inputed a malicious IP
func IsReservedIP(ip net.IP) bool {
	for i, ipnet := range ReservedIPNets {
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}
