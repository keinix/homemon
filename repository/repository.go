package repository

import (
	"fmt"
	"github.com/Ullaakut/nmap"
)

func StoreScan(r *nmap.Run) {
	fmt.Printf("%+v\n", r.Hosts)
}
