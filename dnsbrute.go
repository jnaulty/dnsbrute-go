package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jnaulty/dnsbrute-go/dnshelper"
)

var domainType string
var outputFileType string
var wordlistType string

func init() {
	const (
		defaultDomain     = "example.com"
		usageDomain       = "Domain Name"
		defaultOutputFile = "/dev/null"
		usageOutputFile   = "Output file"
		defaultWordlist   = "subdomain-lists/deepmagic.com-top500prefixes.txt"
		usageWordlist     = "Subdomain wordlist"
	)

	flag.StringVar(&domainType, "domain", defaultDomain, usageDomain)
	flag.StringVar(&domainType, "d", defaultDomain, usageDomain+" (shorthand)")

	flag.StringVar(&outputFileType, "outputFile", defaultOutputFile, usageOutputFile)
	flag.StringVar(&outputFileType, "o", defaultOutputFile, usageOutputFile+" (shorthand)")

	flag.StringVar(&wordlistType, "wordlist", defaultWordlist, usageWordlist)
	flag.StringVar(&wordlistType, "w", defaultWordlist, usageWordlist+" (shorthand)")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func querySubdomains(wordlistPath string, domainName string, outputFile string) {
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // use scanwords

	var f *os.File
	// If the file doesn't exist, create it, or append to the file
	if outputFile != "/dev/null" {
		f, err = os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	for scanner.Scan() {
		var subdomain string = scanner.Text()
		fqdn := subdomain + "." + domainName
		cname, cnameErr := dnshelper.LookupCNAME(fqdn)
		ip, ipErr := dnshelper.LookupIP(fqdn)

		if cnameErr == nil && ipErr == nil {
			if outputFile != "/dev/null" {
				_, err := f.WriteString(subdomain + ", " + cname + ", " + ip + "\n")
				check(err)
				f.Sync()
			}
			fmt.Print(subdomain + ", ")
			fmt.Print(cname + ", ")
			fmt.Print(ip + "\n")
		}

	}

	if outputFile != "/dev/null" {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func main() {
	// basic dnsbrute functionality
	// print subdomain, cname, ip

	flag.Parse()
	// args are
	// -w for wordlist
	// -d for domain
	// -o for output

	fmt.Println("Welcome to DNSBrute")
	fmt.Println("")
	querySubdomains(wordlistType, domainType, outputFileType)

}
