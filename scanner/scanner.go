package scanner

import (
	"github.com/Ullaakut/nmap"
	"homemon/config"
	"homemon/repository"
	"log"
	"strings"
	"time"
)

func ScanNetwork(conf *config.ScanConfig) {
	c := make(chan *nmap.Run)
	for {
		for _, scan := range conf.Scans {
			go executeScan(scan, c)
		}

		for range conf.Scans {
			repository.StoreScan(<-c)
		}
		time.Sleep(time.Duration(conf.Frequency) * time.Millisecond)
	}
}

func executeScan(scan string, c chan<- *nmap.Run) {
	fmtScan := strings.Replace(scan, "nmap ", "", 1)
	scanner, err := nmap.NewScanner(
		nmap.WithCustomArguments(fmtScan),
	)
	if err != nil {
		log.Fatalf("Could not create nmap scanner: %v", err)
	}
	result, err := scanner.Run()
	if err != nil {
		log.Fatalf("Error in scan '%v' : %v", scan, err)
	}
	c <- result
}
