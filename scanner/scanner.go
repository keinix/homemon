package scanner

import (
	"fmt"
	"homemon/config"
	"homemon/repository"
	"time"
)

func ScanNetwork(conf *config.ScanConfig) {
	c := make(chan repository.ScanResult)
	for {
		for _, scan := range conf.Scans {
			go executeScan(scan, c)
		}

		for range conf.Scans {
			repository.StoreScan(<-c)
		}
		fmt.Println("Sleeping")
		time.Sleep(time.Duration(conf.Frequency) * time.Millisecond)
	}
}

func executeScan(scan string, c chan<- repository.ScanResult) {
	c <- repository.ScanResult{Temp: scan}
}
