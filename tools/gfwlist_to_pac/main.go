package main

import (
	"flag"
	"io/ioutil"
	"fmt"
)

var (
	gfwlistFilePath = ""
	gfwlistData = ""
)

func init() {
	flag.StringVar(&gfwlistFilePath, "gfwlist.path", "", "gfwlist file path")
	flag.Parse()
	data, err := ioutil.ReadFile(gfwlistFilePath)
	if err != nil {
		panic(err)
	}
	gfwlistData = string(data)
}

func main() {

}