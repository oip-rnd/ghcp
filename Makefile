TARGET := ghcp
OSARCH := darwin/amd64 linux/amd64 windows/amd64

.PHONY: check install release_bin release_homebrew release clean

all: $(TARGET)

check:
	go vet
	go test -v ./...

$(TARGET): check
	go build

install: check
	go install

dist/bin: check
	gox --osarch "$(OSARCH)" -output 'dist/bin/{{.Dir}}_{{.OS}}_{{.Arch}}'

release_bin: dist/bin
	ghr -u $(CIRCLE_PROJECT_USERNAME) -r $(CIRCLE_PROJECT_REPONAME) -b "$$(ghch -F markdown --latest)" $(CIRCLE_TAG) dist/bin

dist/$(TARGET).rb: dist/bin
	./homebrew.sh dist/bin/$(TARGET)_darwin_amd64 > dist/$(TARGET).rb

release_homebrew: dist/$(TARGET).rb install
	ghcp -u $(CIRCLE_PROJECT_USERNAME) -r homebrew-$(CIRCLE_PROJECT_REPONAME) -m $(CIRCLE_TAG) -C dist/ $(TARGET).rb

release: release_bin release_homebrew

clean:
	-rm $(TARGET)
	-rm -r dist
	-rm $$(go env GOPATH)/bin/ghcp
