.PHONY: build export iterate clean

APPNAME = "testapp"

build:
	@docker-compose -p $(APPNAME) build --force-rm $(APPNAME)

clean:
	@docker rm $(APPNAME)

export:
	@docker cp $(APPNAME):/go/bin/darwin_amd64/$(APPNAME) . ;
	@mv $(APPNAME) $(GOPATH)/bin; docker rm $(APPNAME)

init:
	@go mod init github.com/reymundbautista/$(APPNAME)

iterate: build export

run:
	@$(APPNAME)