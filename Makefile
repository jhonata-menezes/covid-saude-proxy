
build:
	CGO_ENABLED=0 go install -v -a --ldflags="-s -w" cmd/proxy/covid_proxy.go