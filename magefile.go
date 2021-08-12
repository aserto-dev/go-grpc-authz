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

	err = buf.Run(
		buf.AddArg("generate"),
		buf.AddArg("--template"),
		buf.AddArg(filepath.Join("buf", "buf.gen.yaml")),
		buf.AddArg(bufImage),
		buf.AddPaths(files),
	)
	if err != nil {
		return err
	}

	return filepath.Walk("aserto", replacePackageName)
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

	excludePattern := filepath.Join(bufExportDir, "aserto", "**", "*_internal.proto")
	authorizerFiles, err := fsutil.Glob(filepath.Join(bufExportDir, "aserto", "authorizer", "authorizer", "**", "*.proto"), excludePattern)
	if err != nil {
		return clientFiles, err
	}

	apiFiles, err := fsutil.Glob(filepath.Join(bufExportDir, "aserto", "api", "**", "*.proto"), excludePattern)
	if err != nil {
		return clientFiles, err
	}

	optionFiles, err := fsutil.Glob(filepath.Join(bufExportDir, "aserto", "options", "**", "*.proto"), excludePattern)
	if err != nil {
		return clientFiles, err
	}

	files := append(authorizerFiles, apiFiles...)
	files = append(files, optionFiles...)

	fmt.Printf("found: %v files \n", len(files))

	for _, f := range files {
		clientFiles = append(clientFiles, strings.TrimPrefix(f, bufExportDir+string(filepath.Separator)))
	}

	return clientFiles, nil
}

func replacePackageName(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if fi.IsDir() {
		return nil
	}

	matched, err := filepath.Match("*.go", fi.Name())

	if err != nil {
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		newContents := strings.Replace(string(read), "github.com/aserto-dev/proto/aserto", "github.com/aserto-dev/go-grpc-authz/aserto", -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			return err
		}

	}

	return nil
}

// Probably not needed
func Build() error {
	return nil
}
