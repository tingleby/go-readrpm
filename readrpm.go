package main

import "os"
import "fmt"

import "github.com/sassoftware/go-rpmutils"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify RPM to read!")
		os.Exit(2)
	}
	inrpm := os.Args[1]

	f, err := os.Open(inrpm)
	if err != nil {
		fmt.Printf("Failed to open: %s\n", inrpm)
		os.Exit(3)
		panic(err)
	}

	rpm, err := rpmutils.ReadRpm(f)
	if err != nil {
		fmt.Println("Failed to parse RPM!")
		os.Exit(4)
	}

	// Get the name, epoch, version, release, and arch
	nevra, err := rpm.Header.GetNEVRA()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", nevra)

	// Reading the provides header
	provides, err := rpm.Header.GetStrings(rpmutils.PROVIDENAME)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Provides:\n")
	for _, p := range provides {
		fmt.Printf("\t%s\n", p)
	}

	depends, err := rpm.Header.GetStrings(rpmutils.REQUIRENAME)
        if err != nil {
                panic(err)
        }

        fmt.Printf("Depends:\n")
        for _, p := range depends {
                fmt.Printf("\t%s\n", p)
        }
}
