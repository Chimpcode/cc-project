package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"fmt"
	"path/filepath"
	"strings"
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



	/* ALPHA REVIEWS
	ft := token.NewFileSet()
	data, err := ioutil.ReadFile("template/u.go")
	if err != nil {
		panic(err)
	}

	f, err := parser.ParseFile(ft, "", data, 0)
	if err != nil {
		panic(err)
	}
	for k, d := range f.Scope.Objects {
		fmt.Println("Key:", k)
		fmt.Println(d)
	}
	*/


}