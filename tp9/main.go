package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}

	for _, proc := range processes {
		name, err := proc.Name()
		if err != nil {
			continue
		}
		fmt.Println(proc.Pid, " ", name)
	}

	fmt.Println("combien:", len(processes))
}
