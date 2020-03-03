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

The difference between the container sizes of my project and `interhack86`'s is also one of the reasons I chose to
rewrite this in go.

```
➜  dnsbrute-go git:(master) ✗ docker images --no-trunc | grep dnsbrute
dnsbrute                                                                               latest
sha256:c2dae0fadad064195b58cd09756ee8cec876e724776cf2158d970c64a14802e2   16 minutes ago      7.38MB
interhack/dnsbrute                                                                     latest
sha256:56e05afcd493ff8e2e18dd9df50757244e97714a699813b4a1409e41e84aa092   22 months ago       90.7MB

```
