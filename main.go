package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	err := run3()
	if err != nil {
		log.Fatalln(err)
	}

	err = run()
	if err != nil {
		log.Fatalln(err)
	}
}

type importPathStatus struct {
	importPath string
	stale      bool
}

func (ips importPathStatus) String() string {
	switch ips.stale {
	case false:
		return "up to date: " + ips.importPath
	case true:
		return "\033[1m" + "\033[31m" + "STALE" + "\033[0m" + ": " + ips.importPath
	}
	panic("unreachable")
}

var commands = make(map[string][]importPathStatus) // Command name -> list of import paths with statuses.

func run3() error {
	out, err := exec.Command("go", "list", "-e", "-f", `{{if (and (not .Error) (not .Goroot) (eq .Name "main"))}}{{.ImportPath}}	{{.Stale}}{{end}}`, "all").Output()
	if err != nil {
		return fmt.Errorf("failed to run go list: %v", err)
	}

	br := bufio.NewReader(bytes.NewReader(out))
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		line = line[:len(line)-1] // Trim trailing newline.

		importPathAndStale := strings.Split(line, "\t")

		importPath := importPathAndStale[0]

		stale, err := strconv.ParseBool(importPathAndStale[1])
		if err != nil {
			return err
		}

		commandName := path.Base(importPath)

		commands[commandName] = append(commands[commandName],
			importPathStatus{
				importPath: importPath,
				stale:      stale,
			},
		)
	}

	return nil
}

/*func run2() error {
	goPackages := make(chan *gist7480523.GoPackage, 64)
	go gist8018045.GetGopathGoPackages(goPackages)
	for {
		goPackage, ok := <-goPackages
		if !ok {
			break
		}

		if goPackage.Bpkg.Name != "main" {
			continue
		}
		commandName := path.Base(goPackage.Bpkg.ImportPath)
		commands[commandName] = append(commands[commandName], importPathStatus{importPath: goPackage.Bpkg.ImportPath})
	}

	return nil
}*/

func run() error {
	workspaces := filepath.SplitList(build.Default.GOPATH)
	for _, workspace := range workspaces {
		gobin := filepath.Join(workspace, "bin")
		fmt.Println(gobin)

		fis, err := ioutil.ReadDir(gobin)
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return err
		}

		for _, fi := range fis {
			if fi.IsDir() {
				continue
			}
			if strings.HasPrefix(fi.Name(), ".") {
				continue
			}

			fmt.Println(" ", fi.Name())
			for _, importPathStatus := range commands[fi.Name()] {
				fmt.Println("   ", importPathStatus)
			}
		}
	}

	return nil
}
