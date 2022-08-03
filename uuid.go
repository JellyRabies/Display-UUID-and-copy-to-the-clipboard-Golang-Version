package main

import (
	"fmt"

	"github.com/StackExchange/wmi"
	"github.com/atotto/clipboard"
)

type UUID struct {
	IdentifyingNumber string
	UUID              string
}

func main() {
	//win32_ComputerSystemProduct get uuid
	var UUID []UUID
	err := wmi.Query("Select name, uuid from win32_ComputerSystemProduct", &UUID)
	if err != nil {
		fmt.Println("There was an error - WMI Query Failed")
		return
	}
	PCName := UUID[0].IdentifyingNumber
	PCUUID := UUID[0].UUID
	strOutput := "NB" + PCName + "," + PCUUID
	fmt.Println()
	fmt.Println("    UUID Utility")
	fmt.Println("    EF-1.1 - 2022")
	fmt.Println()
	fmt.Println(strOutput)
	clipboard.WriteAll(strOutput)
}
