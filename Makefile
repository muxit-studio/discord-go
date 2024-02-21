VERSION=v0.1.0
BINARY_NAME=discord-go

test:
	go test -v ./...

build-amd64:
	GOOS=linux GOARCH=amd64 go1.22.0 build -o ./build/$(BINARY_NAME)_amd64 ./cmd/discord/
build-arm64:
	GOOS=linux GOARCH=arm64 go1.22.0 build -o ./build/$(BINARY_NAME)_arm64 ./cmd/discord/

package-amd64:
	@echo "Packaging amd64 binary with README and LICENSE"
	mkdir -p ./release
	cp ./build/$(BINARY_NAME)_amd64 ./build/$(BINARY_NAME) # Copy and rename binary
	zip -j ./release/$(BINARY_NAME)_$(VERSION)_linux_amd64.zip ./build/$(BINARY_NAME) ./LICENSE
	rm ./build/$(BINARY_NAME) # Clean up

package-arm64:
	@echo "Packaging arm64 binary with README and LICENSE"
	mkdir -p ./release
	cp ./build/$(BINARY_NAME)_arm64 ./build/$(BINARY_NAME) # Copy and rename binary
	zip -j ./release/$(BINARY_NAME)_$(VERSION)_linux_arm64.zip ./build/$(BINARY_NAME) ./LICENSE
	rm ./build/$(BINARY_NAME) # Clean up

publish:
	@echo "Publishing release to GitHub"
	gh release create $(VERSION) ./release/* --generate-notes

cleanup:
	rm -rf ./build
	rm -rf ./release

build: cleanup build-amd64 build-arm64
package: build package-amd64 package-arm64
package-publish: package publish

.PHONY: build test watch package
