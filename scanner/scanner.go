package scanner

import (
	"fmt"
	"github.com/Ullaakut/nmap"
	"homemon/config"
	"homemon/repository"
	"log"
	"strings"
	"time"
)

func ScanNetwork(conf *config.ScanConfig) {
	for {
		c := make(chan nmap.Run)
		results := make([]nmap.Run, len(conf.Scans))
		for _, scan := range conf.Scans {
			go executeScan(scan, c)
		}
		for range conf.Scans {
			results = append(results, <-c)
			fmt.Println("Got Result")
		}
		batch := parseBatchScan(results)
		repository.StoreScan(batch)
		fmt.Println("sleeping")
		time.Sleep(time.Duration(conf.Frequency) * time.Millisecond)
	}
}

func executeScan(scan string, c chan<- nmap.Run) {
	fmtScan := strings.Replace(scan, "nmap ", "", 1)
	args := strings.Split(fmtScan, " ")
	scanner, err := nmap.NewScanner(
		nmap.WithCustomArguments(args...),
	)
	if err != nil {
		log.Fatalf("Could not create nmap scanner: %v", err)
	}
	result, err := scanner.Run()
	if err != nil {
		log.Fatalf("Error in scan '%v' : %v", scan, err)
	}
	c <- *result
}

func parseBatchScan(scans []nmap.Run) repository.BatchScan {
	uHosts := parseUniqueHosts(scans)
	return repository.BatchScan{
		TimeCompleted: time.Now().Unix(),
		UniqueHosts:   uHosts,
		Scans:         scans,
	}
}

// Different scans in a batch may have picked up the same ip
// parse all hosts are return a slice of unique ips
func parseUniqueHosts(scans []nmap.Run) []string {
	hostSet := make(map[string]struct{})
	for _, scan := range scans {
		ips := getIpsFromHosts(scan.Hosts)
		for _, ip := range ips {
			hostSet[ip] = struct{}{}
		}
	}
	uHosts := make([]string, len(hostSet))
	i := 0
	for k := range hostSet {
		uHosts[i] = k
		i++
	}
	return uHosts
}

func getIpsFromHosts(hosts []nmap.Host) []string {
	var ips []string
	for _, h := range hosts {
		if h.Status.State != "up" {
			continue
		}
		ip, err := getIpFromAddresses(h.Addresses)
		if err != nil {
			fmt.Println("host Ip not found")
		}
		ips = append(ips, ip)
	}
	return ips
}

// return ipv4 address or ipv6 if ipv4 doesnt exist
func getIpFromAddresses(addrs []nmap.Address) (string, error) {
	if len(addrs) == 1 {
		return addrs[0].Addr, nil
	}
	var ipv6 string
	for _, a := range addrs {
		if a.AddrType == "ipv4" {
			return a.Addr, nil
		} else if a.AddrType == "ipv6" {
			ipv6 = a.Addr
		}
	}
	if ipv6 == "" {
		return ipv6, fmt.Errorf("host has no ipv4 or ipv6")
	}
	return ipv6, nil
}
