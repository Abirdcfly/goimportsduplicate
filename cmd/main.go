package main

import (
	"github.com/abirdcfly/importsduplicate"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(importsduplicate.Analyzer)
}
