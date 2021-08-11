// +build mage

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aserto-dev/mage-loot/buf"
	"github.com/aserto-dev/mage-loot/deps"
)

// install required dependencies.
func Deps() {
	deps.GetAllDeps()
}

// Generate the code
func Generate() error {
	files := []string{
		filepath.Join("aserto"),
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
		buf.AddArg("buf.build/aserto/aserto"),
		buf.AddPaths(files),
	)
	if err != nil {
		return err
	}

	return filepath.Walk("aserto", replacePackageName)
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

		newContents := strings.Replace(string(read), "github.com/aserto-dev/proto/aserto", "github.com/aserto-dev/go-grpc-server/aserto", -1)

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
