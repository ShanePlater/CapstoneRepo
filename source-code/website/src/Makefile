NAME = lar-server
VERSION = 0.0.5
EXTRAVERSION = $(shell date "+%Y%m%d")git$(shell git log --pretty=format:'%h' -n 1)

default: binary

requires:
	@$(GOPATH)/bin/dep ensure -v

binary:
	@mkdir -p build
	@env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/$(NAME).exe -ldflags "-s -w -X main.mode=debug" -v

clean:
	@-rm -rf build

test:
	@go test -v *\.go

production: 
	@mkdir -p build
	@env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o build/$(NAME).exe -ldflags "-s -w -X main.mode=release" -v
	@-rm -rf build/$(NAME)-$(VERSION)-$(EXTRAVERSION).exe
	@echo -en "Checking for upx... "
	@echo -e "\e[1m\e[93m$(shell which upx 2>/dev/null)\e[0m\e[39m"
	@which upx >/dev/null 2>&1 || echo -e "\e[1m\e[91mno\e[0m\e[39m\nUPX compression disabled"
	@which upx >/dev/null 2>&1 || cp build/$(NAME).exe build/$(NAME)-$(VERSION)-$(EXTRAVERSION).exe
	@-upx --ultra-brute build/$(NAME).exe --color -v -o build/$(NAME)-$(VERSION)-$(EXTRAVERSION).exe 2>/dev/null ||:
	@echo -e "Output: $(PWD)/build/$(NAME)-$(VERSION)-$(EXTRAVERSION).exe"