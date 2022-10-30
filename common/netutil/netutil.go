package netutil

import (
	"log"
	"myvpn/common/counter"
	"net"
	"os/exec"
	"strings"
	"time"
)

// ExecuteCommand executes the given command
// ExecuteCommand executes the given command
func ExecCmd(c string, args ...string) string {
	log.Printf("exec %v %v", c, args)
	cmd := exec.Command(c, args...)
	out, err := cmd.Output()
	if err != nil {
		log.Println("failed to exec cmd:", err)
	}
	if len(out) == 0 {
		return ""
	}
	s := string(out)
	return strings.ReplaceAll(s, "\n", "")
}

// isPhysicalInterface returns true if the interface is physical
// isPhysicalInterface returns true if the interface is physical
func isPhysicalInterface(addr string) bool {
	prefixArray := []string{"ens", "enp", "enx", "eno", "eth", "en0", "wlan", "wlp", "wlo", "wlx", "wifi0", "lan0"}
	for _, pref := range prefixArray {
		if strings.HasPrefix(strings.ToLower(addr), pref) {
			return true
		}
	}
	return false
}

// getAllInterfaces returns all interfaces
// getAllInterfaces returns all interfaces
func getAllInterfaces() []net.Interface {
	var outInterfaces []net.Interface

	interfaces, err := net.Interfaces()
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, interf := range interfaces {
		if interf.Flags&net.FlagLoopback == 0 && interf.Flags&net.FlagUp == 1 && isPhysicalInterface(interf.Name) {
			netAddrs, _ := interf.Addrs()
			if len(netAddrs) > 0 {
				outInterfaces = append(outInterfaces, interf)
			}
		}
	}
	return outInterfaces
}

// GetInterfaceName returns the name of interfaces
// GetInterfaceName returns the name of interfaces
func GetInterface() (name string) {
	ifaces := getAllInterfaces()
	if len(ifaces) == 0 {
		return ""
	}
	netAddrs, _ := ifaces[0].Addrs()
	for _, addr := range netAddrs {
		ip, ok := addr.(*net.IPNet)
		if ok && ip.IP.To4() != nil && !ip.IP.IsLoopback() {
			name = ifaces[0].Name
			break
		}
	}
	return name
}

// Lookup IP address of the given hostname
// Lookup IP address of the given hostname
func LookupIP(domain string) net.IP {
	ips, err := net.LookupIP(domain)
	if err != nil || len(ips) == 0 {
		log.Println(err)
		return nil
	}
	return ips[0]
}

// LookupServerAddrIP returns the IP of server address
// LookupServerAddrIP returns the IP of server address
func LookupServerAddrIP(serverAddr string) net.IP {
	host, _, err := net.SplitHostPort(serverAddr)
	if err != nil {
		log.Panic("error server address")
		return nil
	}
	ip := LookupIP(host)
	return ip
}

// PrintStats returns the stats info
// PrintStats returns the stats info
func PrintStats(enableVerbose bool) {
	if !enableVerbose {
		return
	}
	go func() {
		for {
			time.Sleep(30 * time.Second)
			log.Printf("stats:%v", counter.PrintBytes())
		}
	}()
}

// PrintErr returns the error log
// PrintErr returns the error log
func PrintErr(err error, enableVerbose bool) {
	if !enableVerbose {
		return
	}
	log.Printf("error:%v", err)
}

// IsIPv4 returns true if the packet is IPv4s
// IsIPv4 returns true if the packet is IPv4s
func IsIPv4(packet []byte) bool {
	flag := packet[0] >> 4
	return flag == 4
}

// GEtIPv4Dst returns the IPv4 destination address of the packet
func GetIPv4Dst(packet []byte) net.IP {
	return net.IPv4(packet[16], packet[17], packet[18], packet[19])
}

// GetdstKey returns the destination key of the packets
// GetdstKey returns the destination key of the packets
func GetDstKey(packet []byte) string {
	key := ""

	if IsIPv4(packet) && len(packet) >= 20 {
		//To4().String()返回类型为string而不是net.IP
		key = GetIPv4Dst(packet).To4().String()
	}

	return key
}

// GetIPv4Src returns the IPv4 source address of the packet
// GetIPv4Src returns the IPv4 source address of the packet
func GetIPv4Src(packet []byte) net.IP {
	return net.IPv4(packet[12], packet[13], packet[14], packet[15])
}

// GetSrcKey returns the source key of the packet
// GetSrcKey returns the source key of the packet
func GetSrcKey(packet []byte) string {
	key := ""
	if IsIPv4(packet) && len(packet) >= 20 {
		key = GetIPv4Src(packet).To4().String()
	}

	return key
}
