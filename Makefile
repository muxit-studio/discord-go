build:
	go build -o your_cmd ./cmd/your_cmd/main
test:
	# TODO: use what you would like here
watch:
	# TODO: use what you would like here
init:
	@read -p "Enter the new command name: " newCmdName && \
	find . -type f -name '*.go' ! -path "./cmd/$$newCmdName/*" -exec sed -i'' -e "s/package cmd/package $$newCmdName/g" {} + && \
	sed -i'' -e "s/module github.com\/your_cmd/module github.com\/$$newCmdName/g" go.mod && \
	sed -i'' -e "s/github.com\/your_cmd\/cmd/github.com\/$$newCmdName\/cmd/g" cmd/your_cmd/main.go && \
	[ -d ./cmd/your_cmd ] && mv ./cmd/your_cmd ./cmd/$$newCmdName

.PHONY: build test watch
