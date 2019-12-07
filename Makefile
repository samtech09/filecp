APP_VER := "1.0.8"
GOOS=linux
GOARCH=amd64
GOARM=6

BUILDPATH=$(CURDIR)
BINPATH=$(BUILDPATH)/bin
EXENAME=filecp


.PHONY: clean
clean:
	@rm -rf $(BINPATH) || true


_mkdir_:
	@if [ ! -d $(BINPATH) ] ; then mkdir -p $(BINPATH) ; fi
	
	
.PHONY: build _mkdir_
build:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.appVer=$(APP_VER) -X main.buildVer=`date -u +b-%Y%m%d.%H%M%S`" -o $(BINPATH)/$(EXENAME) . || (echo "build failed $$?"; exit 1)
	@echo 'Build suceeded... done'


.PHONY: buildarm _mkdir_
buildarm: GOARCH=arm GOARM=6
buildarm:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -ldflags "-X main.appVer=$(APP_VER) -X main.buildVer=`date -u +b-%Y%m%d.%H%M%S`" -o $(BINPATH)/$(EXENAME)-arm . || (echo "build failed $$?"; exit 1)
	@echo 'ARM Build suceeded... done'

	
.PHONY: buildwin _mkdir_
buildwin: GOOS=windows GOARCH=386
buildwin:
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-X main.appVer=$(APP_VER) -X main.buildVer=`date -u +b-%Y%m%d.%H%M%S`" -o $(BINPATH)/$(EXENAME).exe . || (echo "build failed $$?"; exit 1)
	@echo 'Windows Build suceeded... done'
