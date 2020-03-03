# DNSBrute-Go

A learning project for emulating the usefuleness of the python utility I used once
[interhack86/dnsbrute](https://github.com/interhack86/dnsbrute)

## Build

`make build` or `go build`

To build the Docker image:

`make build-docker`

## Run

```
➜  dnsbrute-go git:(master) ✗ ./dnsbrute-go -h
Usage of ./dnsbrute-go:
  -d string
        Domain Name (shorthand) (default "example.com")
  -domain string
        Domain Name (default "example.com")
  -o string
        Output file (shorthand) (default "/dev/null")
  -outputFile string
        Output file (default "/dev/null")
  -w string
        Subdomain wordlist (shorthand) (default "subdomain-lists/deepmagic.com-top500prefixes.txt")
  -wordlist string
        Subdomain wordlist (default "subdomain-lists/deepmagic.com-top500prefixes.txt")
```

`./dnsbrute-go -d surfline.com -w subdomain-lists/subdomains-top1mil-110000.txt -o something.txt`

Or, if using docker:

`docker run -it -v ${PWD}:/results dnsbrute -w /subdomain-lists/subdomains-top1mil-110000.txt -d example.com -o /results/results.txt`


## Learning Resources and References

- [Golang Standard Library](https://golang.org/pkg/#stdlib)
- [Golang DNS Lookup Blog Post](http://networkbit.ch/golang-dns-lookup/)
- [Go By Example: Writing Files](https://gobyexample.com/writing-files)

### Why write another tool?

I wanted to write something simple, and I used the python version of `dnsbrute` for a project. I realized this would be
pretty straightforward to code, just loop through the subdomains, run the queries, and save the results.

After taking a look at github, I see that there are a few other resources I should take a look at:
- [Q2h1Cg/dnsbrute](https://github.com/Q2h1Cg/dnsbrute)
- [DNS Library with Amazing Docs](https://github.com/miekg/dns/blob/418631f446c1e9f995f3a9a462a85f4c3cd9d811/doc.go)

I want to emulate what the user Q2h1Cg does and use goroutines for the queries to speed things up.

Once nice thing about this current implementation, is that the docker image size is a magnitude smaller in size than the
python project I was previously using.

```
➜  dnsbrute-go git:(master) ✗ docker images --no-trunc | grep dnsbrute
dnsbrute                                                                               latest
sha256:c2dae0fadad064195b58cd09756ee8cec876e724776cf2158d970c64a14802e2   16 minutes ago      7.38MB
interhack/dnsbrute                                                                     latest
sha256:56e05afcd493ff8e2e18dd9df50757244e97714a699813b4a1409e41e84aa092   22 months ago       90.7MB

```
