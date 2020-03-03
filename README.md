# DNSBrute-Go

A learning project for emulating the usefuleness of the python utility I used once
[interhack86/dnsbrute](https://github.com/interhack86/dnsbrute)

## Build

`make build` or `go build`

To build the Docker image:

`make build-docker`

## Run

`./dnsbrute-go -d surfline.com -w subdomain-lists/subdomains-top1mil-110000.txt -o something.txt`

Or, if using docker:

`docker run -it -v ${PWD}:/results dnsbrute -w /subdomain-lists/subdomains-top1mil-110000.txt -d example.com -o /results/results.txt`


## Learning Resources and References

- [Golang Standard Library](https://golang.org/pkg/#stdlib)
- [Golang DNS Lookup Blog Post](http://networkbit.ch/golang-dns-lookup/)
- [Go By Example: Writing Files](https://gobyexample.com/writing-files)
