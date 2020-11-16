package main

import (
	"fmt"
	"my5G-RANTester/internal/templates"
)

func main() {

	// testing attach and ping with a number of UEs.
	// fmt.Println(templates.TestMultiAttachUesInQueue(3))

	// testing concurrent UEs registration with GNBs.
	// fmt.Println( templates.TestMultiAttachUesInConcurrencyWithGNBs(20) )

	// testing concurrent UEs registration with some SCTPs associations.
	fmt.Println(templates.TestMultiAttachUesInConcurrencyWithTNLAs(3))

	// testing multiple GNBs authentication(control plane only)-> NGAP Request and response tester.
	//fmt.Println(testMultiAttachGnbInQueue(100))
	// fmt.Println(templates.TestMultiAttachGnbInConcurrency(100))
}
