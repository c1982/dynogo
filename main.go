package main

import (
	"flag"
	"fmt"
)

var EMAIL string
var TOKEN string
var DOMAIN string
var RECORD_NAME string

func init() {
	flag.StringVar(&EMAIL, "email", "", "CloudFlare email address")
	flag.StringVar(&TOKEN, "token", "", "CloudFlare API Key")
	flag.StringVar(&DOMAIN, "domain", "", "Domain name on CloudFlare")
	flag.StringVar(&RECORD_NAME, "name", "office", "DNS A Record Name")
}

func main() {
	flag.Parse()

	if EMAIL == "" || TOKEN == "" || DOMAIN == "" || RECORD_NAME == "" {
		flag.PrintDefaults()
		return
	}

	ipaddr, err := GetWANIP()

	if err != nil {
		fmt.Errorf(err.Error())
	}

	err = UpdateOrSaveRecord(ipaddr)

	if err != nil {
		fmt.Errorf(err.Error())
	}
}
