wxgo
====

Go bind of wxWidgets.

## Requirements:

1. go **1.2rc1** or above.
2. wxWidgets v2.9.5 **static library** build. 

## How to build:

	go get github.com/kevin-yuan/wxgo
	go run /your-local-repo-of-wxgo/config/main.go
	go install -ldflags=-linkmode=external github.com/kevin-yuan/wxgo/wx
