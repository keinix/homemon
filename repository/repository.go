package repository

import (
	"fmt"
	"github.com/Ullaakut/nmap"
)

type BatchScan struct {
	TimeCompleted int64
	UniqueHosts   []string
	Scans         []nmap.Run
}

func StoreScan(b BatchScan) {
	fmt.Printf("%+v\n", b.UniqueHosts)
}
