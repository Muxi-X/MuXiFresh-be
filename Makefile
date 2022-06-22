name=gateway

all: gotool
	@go build -o main -v .
vet:
	go vet ./...
clean:
	rm -f main
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}
gotool:
	gofmt -w .
tidy:
	go mod tidy
test:
	@go test -v -count=1  ./...
ca:
	openssl req -new -nodes -x509 -out conf/server.crt -keyout conf/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=127.0.0.1/emailAddress=xxxxx@qq.com"
github:
	git push origin && git push --tags origin
gitea:
	git push --tags muxi
tag:
	git tag release-${name}-${ver}
push: vet tag gitea github

help:
	@echo "make - compile the source code with local vendor"
	@echo "make build compile the source code with remote vendor"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make ca - generate ca files"
	@echo "make deploy name=gateway ver=1.0.x - push code and tag to github and gitea"

.PHONY: clean gotool ca help
