package main

import "github.com/liov/tiga/tools/protoc-gen-enum/command"

func main() {
	command.Write(command.Generate(command.Read()))
}
