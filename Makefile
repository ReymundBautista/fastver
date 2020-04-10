.PHONY: build export iterate clean

APPNAME = "fastver"

build:
	@docker-compose -p $(APPNAME) build --force-rm $(APPNAME)

clean:
	@docker rm $(APPNAME)

export:
	@docker run -d --name $(APPNAME) thebernank/$(APPNAME):latest;
	@docker cp $(APPNAME):/go/bin/darwin_amd64/$(APPNAME) . ;
	@mv $(APPNAME) $(GOPATH)/bin; docker rm $(APPNAME)

init:
	@go mod init github.com/reymundbautista/$(APPNAME)

iterate: build export

run:
	@$(APPNAME)