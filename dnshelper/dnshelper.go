package dnshelper

import (
	"fmt"
	"net"
    "strings"
)


func LookupAddr(ipAddr string) {
	names, err := net.LookupAddr(ipAddr)
	if err != nil {
		panic(err)
	}
	if len(names) == 0 {
		fmt.Printf("no record")
	}
	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
}


func LookupCNAME(domainName string) (cname string, err error) {
	cname, err = net.LookupCNAME(domainName)
	if err != nil {
        return "", err
	}
    return cname, nil
}

func LookupIP(domainName string) (ipString string, err error) {
	ips, err := net.LookupIP(domainName)
	if err != nil {
        return "", err
	}
	if len(ips) == 0 {
	}
    for _, ip := range ips {
        ipString += ip.String()+", "
    }
    ipString = strings.TrimSuffix(ipString, ", ")

    return ipString, nil

// 172.217.1.238
// 2607:f8b0:4000:80e::200e

}

func LookupTXT(domainName string) {
	txts, err := net.LookupTXT(domainName)
	if err != nil {
		panic(err)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		//dig +short gmail.com txt
		fmt.Printf("%s\n", txt)
	}

// v=spf1 redirect=_spf.google.com

}
