package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: binstale [binaries]")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// Populate filter, a set of binaries that user wants use.
	filter := make(map[string]struct{})
	if args := flag.Args(); len(args) != 0 {
		for _, arg := range args {
			filter[arg] = struct{}{}
		}
	}

	// Find all commands and determine if they're stale or up to date.
	commands, err := commands()
	if err != nil {
		log.Fatalln(err)
	}

	// Find binaries in GOPATH/bin directories.
	binaries, err := binaries(filter)
	if err != nil {
		log.Fatalln(err)
	}

	// Print output.
	sort.Strings(binaries)
	for _, binary := range binaries {
		fmt.Println(" ", binary)
		for _, importPathStatus := range commands[binary] {
			fmt.Println("   ", importPathStatus)
		}
		if len(commands[binary]) == 0 {
			fmt.Println("    (no source package found)")
		}
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

// commands finds all commands in all GOPATH workspaces (not GOROOT), determines if they're stale or up to date,
// and returns the results.
func commands() (map[string][]importPathStatus, error) {
	var commands = make(map[string][]importPathStatus) // Command name -> list of import paths with statuses.

	out, err := exec.Command("go", "list", "-e", "-f", `{{if (and (not .Error) (not .Goroot) (eq .Name "main"))}}{{.ImportPath}}	{{.Stale}}{{end}}`, "all").Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run go list: %v", err)
	}

	br := bufio.NewReader(bytes.NewReader(out))
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		line = line[:len(line)-1] // Trim trailing newline.

		importPathAndStale := strings.Split(line, "\t")

		importPath := importPathAndStale[0]

		stale, err := strconv.ParseBool(importPathAndStale[1])
		if err != nil {
			return nil, err
		}

		commandName := path.Base(importPath)

		commands[commandName] = append(commands[commandName],
			importPathStatus{
				importPath: importPath,
				stale:      stale,
			},
		)
	}

	return commands, nil
}

// binaries finds binaries in GOPATH/bin directories, filtering results with filter if it's not empty.
func binaries(filter map[string]struct{}) ([]string, error) {
	var binaries []string // Binaries that were found and not filtered out.

	workspaces := filepath.SplitList(build.Default.GOPATH)
	for _, workspace := range workspaces {
		gobin := filepath.Join(workspace, "bin")

		fis, err := ioutil.ReadDir(gobin)
		if os.IsNotExist(err) {
			continue
		} else if err != nil {
			return nil, err
		}

		for _, fi := range fis {
			if fi.IsDir() {
				continue
			}
			if strings.HasPrefix(fi.Name(), ".") {
				continue
			}

			// If user specified a list of binaries, filter out binaries that don't match.
			if len(filter) != 0 {
				if _, ok := filter[fi.Name()]; !ok {
					continue
				}
			}

			binaries = append(binaries, fi.Name())
		}
	}

	return binaries, nil
}
