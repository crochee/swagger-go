// Copyright 2020, The Go Authors. All rights reserved.
// Author: OnlyOneFace
// Date: 2020/9/9

package main

import (
	"flag"
	"os"

	"github.com/swagger-go/generate"
)

func main() {
	routerPath := flag.String("rp", "routers", "router path")
	routerGo := flag.String("rg", "router.go", "router.go file")
	outPutJsonPath := flag.String("oj", "swagger/swagger.json", "set output json file path")
	outPutYmlPath := flag.String("oy", "swagger/swagger.yml", "set output yml file path")

	currpath, _ := os.Getwd()
	generate.GenerateDocs(generate.NewParm(currpath, *routerPath, *routerGo, *outPutJsonPath, *outPutYmlPath))
}
