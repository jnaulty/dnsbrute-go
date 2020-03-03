
build:
	go build

.PHONY: build-docker
build-docker:
	docker build -t dnsbrute .

.PHONY: run
run: build
	./dnsbrute-go -d surfline.com -w subdomain-lists/subdomains-top1mil-110000.txt -o results.txt

.PHONY: run-docker
run-docker:
	docker run -it -v ${PWD}:/results dnsbrute -w /subdomain-lists/subdomains-top1mil-110000.txt -d example.com -o /results/results.txt
