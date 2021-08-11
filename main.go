package main

import (
	"github.com/micro/micro/cmd"

	_ "github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	cmd.Init()
}
