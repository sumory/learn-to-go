package main

import (
	//"errors"
	"fmt"
	//"sync"
	//"time"
)

const (
	twepoch            = int64(1288834974657)
	workerIdBits       = uint(5)
	datacenterIdBits   = uint(5)
	maxWorkerId        = -1 ^ (-1 << workerIdBits)     //31
	maxDatacenterId    = -1 ^ (-1 << datacenterIdBits) //31
	sequenceBits       = uint(12)
	workerIdShift      = sequenceBits                                   //12
	datacenterIdShift  = sequenceBits + workerIdBits                    //17
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits //22
	sequenceMask       = -1 ^ (-1 << sequenceBits)                      //4095
)

func main() {
	fmt.Println(maxWorkerId, maxDatacenterId, sequenceMask)
	fmt.Println(-1 << 5)  //-32
	fmt.Println(1 ^ 32)   //33
	fmt.Println(-1 ^ 32)  //-33
	fmt.Println(-1 ^ -32) //31
}
