package main

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
)

func main() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		log.Fatal("Get CPU failed ! : ", err.Error())
	}
	for _, cpuInfo := range cpuInfos {
		fmt.Println(cpuInfo)
	}
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Println("CPU Percent : ", percent)
	}
}
