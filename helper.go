package main

import (
	"errors"
	"fmt"
	cf "github.com/pearkes/cloudflare"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// Also you can use;
// bot.whatismyipaddress.com
// https://api.ipify.org
func GetWANIP() (string, error) {

	var hc http.Client
	var reader io.Reader

	req, err := http.NewRequest("GET", "http://icanhazip.com", reader)
	if err != nil {
		return "", err
	}

	hc.Timeout = time.Duration(4) * time.Second
	resp, err := hc.Do(req)

	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ipstr := strings.TrimSpace(string(body))

	if !isValidIPv4(ipstr) {
		return "", errors.New("Invalid IPv4 address: " + ipstr)
	}

	return ipstr, err
}

func UpdateOrSaveRecord(ipaddr string) error {

	client, err := cf.NewClient(EMAIL, TOKEN)

	if err != nil {
		return err
	}

	isExists, rc := isRecordExists(client, DOMAIN, RECORD_NAME)

	if isExists {

		update := cf.UpdateRecord{}
		update.Content = ipaddr
		update.Type = "A"
		update.Name = RECORD_NAME
		update.Ttl = "1"

		err = client.UpdateRecord(rc.Domain, rc.Id, &update)

		if err != nil {
			return err
		}

		fmt.Printf("Record Updated: %v.%v IN A %v", RECORD_NAME, DOMAIN, ipaddr)

	} else {

		create := cf.CreateRecord{}
		create.Name = RECORD_NAME
		create.Ttl = "1"
		create.Type = "A"
		create.Content = ipaddr

		_, err := client.CreateRecord(DOMAIN, &create)

		if err != nil {
			return err
		}

		fmt.Printf("Record Created: %v.%v IN A %v", RECORD_NAME, DOMAIN, ipaddr)
	}

	return err
}

func isRecordExists(client *cf.Client, domain string, name string) (bool, cf.Record) {

	var record cf.Record
	records, err := client.RetrieveRecordsByName(domain, name, false)

	if err != nil {
		log.Println(err.Error())
		return false, record
	}

	for _, rec := range records {
		if rec.Name == name {
			record = rec
			break
		}
	}

	return len(records) != 0, record
}

func isValidIPv4(ipStr string) bool {
	return net.ParseIP(ipStr).To4() != nil
}
