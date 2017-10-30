package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"fmt"
	"strings"
	"path/filepath"
)

var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	configFile    = kingpin.Arg("config", "Path of chimpcode config file.").String()
)

func main() {
	kingpin.Parse()
	if strings.EqualFold(*configFile, "") {
		var err error
		configFile, err = GetCurrentDir()
		*configFile = filepath.Join(*configFile, "ccseed.json")
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("%v, %s\n", *verbose, *configFile)

	cc, err := ParseSeed(*configFile)
	if err != nil {
		panic(err)
	}
	err = ProcessChimpcodeSeed(cc)
	if err != nil {
		panic(err)
	}
	fmt.Println(cc)
}