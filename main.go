package main

import (
	"flag"
	"github.com/databus23/helm-diff/v3/diff"
	"github.com/databus23/helm-diff/v3/manifest"
	"io/ioutil"
	"log"
	"os"
)

type CommandArguments struct {
	OldResources     string
	NewResources     string
	DefaultNamespace string
	SuppressedKinds  []string
	ShowSecrets      bool
	Context          int
	Output           string
	StripTrailingCr  bool
}

func parseFlag() *CommandArguments {
	args := CommandArguments{}
	flag.StringVar(&args.OldResources, "old", "", "old resource file")
	flag.StringVar(&args.NewResources, "new", "", "new resource file")
	flag.StringVar(&args.DefaultNamespace, "default-namespace", "default", "new resource file")
	flag.BoolVar(&args.ShowSecrets, "show-secrets", false, "show secrets")
	flag.IntVar(&args.Context, "context", -1, "show diff lines")
	flag.StringVar(&args.Output, "output", "diff", "output mode, one of 'simple', 'template', 'json' 'diff'")
	flag.BoolVar(&args.StripTrailingCr, "strip-trailing-cr", false, "strip trailing CR")
	flag.Parse()
	return &args
}

func main() {
	args := parseFlag()

	oldBuf, err := ioutil.ReadFile(args.OldResources)
	if err != nil {
		log.Fatal(err)
	}

	newBuf, err := ioutil.ReadFile(args.NewResources)
	if err != nil {
		log.Fatal(err)
	}

	oldMap := manifest.Parse("---\n"+string(oldBuf), args.DefaultNamespace, false)
	newMap := manifest.Parse("---\n"+string(newBuf), args.DefaultNamespace, false)

	seenAnyChanges := diff.Manifests(oldMap, newMap, []string{}, args.ShowSecrets, args.Context, args.Output, args.StripTrailingCr, os.Stdout)
	if seenAnyChanges {
		os.Exit(2)
	}
}
