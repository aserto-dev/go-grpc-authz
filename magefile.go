// +build mage

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aserto-dev/mage-loot/buf"
	"github.com/aserto-dev/mage-loot/deps"
	"github.com/aserto-dev/mage-loot/fsutil"
)

var bufImage = "buf.build/mitza/aserto"

func All() error {
	Deps()
	err := Clean()
	if err != nil {
		return err
	}
	err = Generate()
	if err != nil {
		return err
	}
	return Build()

}

// install required dependencies.
func Deps() {
	deps.GetAllDeps()
}

// Generate the code
func Generate() error {
	return GenerateAuthorizer()
}

// Generate the Authorizer code
func GenerateAuthorizer() error {

	files, err := getClientFiles()
	if err != nil {
		return err
	}

	oldPath := os.Getenv("PATH")
	currnetDirectory, err := os.Getwd()
	if err != nil {
		return err
	}

	// buf requires protoc in the path
	protocPath := filepath.Join(currnetDirectory, deps.BinPath("protoc"))
	err = os.Setenv("PATH", oldPath+string(os.PathListSeparator)+filepath.Dir(protocPath))
	if err != nil {
		return err
	}

	return buf.Run(
		buf.AddArg("generate"),
		buf.AddArg("--template"),
		buf.AddArg(filepath.Join("buf", "buf.gen.yaml")),
		buf.AddArg(bufImage),
		buf.AddPaths(files),
	)
}

func getClientFiles() ([]string, error) {
	var clientFiles []string

	bufExportDir, err := ioutil.TempDir("", "bufimage")
	if err != nil {
		return clientFiles, err
	}
	bufExportDir = filepath.Join(bufExportDir, "")

	defer os.RemoveAll(bufExportDir)
	err = buf.Run(
		buf.AddArg("export"),
		buf.AddArg(bufImage),
		buf.AddArg("-o"),
		buf.AddArg(bufExportDir),
	)
	if err != nil {
		return clientFiles, err
	}

	excludePattern := filepath.Join(bufExportDir, "aserto", "**", "private", "**", "*.proto")
	authorizerFiles, err := fsutil.Glob(filepath.Join(bufExportDir, "aserto", "authorizer", "authorizer", "**", "*.proto"), excludePattern)
	if err != nil {
		return clientFiles, err
	}

	fmt.Printf("found: %v files \n", len(authorizerFiles))

	for _, f := range authorizerFiles {
		clientFiles = append(clientFiles, strings.TrimPrefix(f, bufExportDir+string(filepath.Separator)))
	}

	return clientFiles, nil
}

// Removes generated files
func Clean() error {
	return os.RemoveAll("aserto")
}

// Probably not needed
func Build() error {
	return nil
}
