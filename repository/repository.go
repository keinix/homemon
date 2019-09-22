package repository

import (
	"fmt"
)

type ScanResult struct {
	Temp string
}

func StoreScan(r ScanResult) {
	fmt.Println(r)
}
