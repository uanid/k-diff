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
	flag.StringVar(&args.DefaultNamespace, "default-namespace", "default", "new resource file")
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

	oldMap := manifest.Parse("---\n"+string(oldBuf), "default", false)
	newMap := manifest.Parse("---\n"+string(newBuf), "default", false)

	seenAnyChanges := diff.Manifests(oldMap, newMap, []string{}, false, -1, "diff", false, os.Stdout)
	if seenAnyChanges {
		os.Exit(2)
	}
}
