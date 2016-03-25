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
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: binstale [command names]")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// Populate filter, a set of binaries that user wants use.
	// Also keep track which filters have been matched (to print warnings at the end if not matched).
	filter := make(map[string]matched)
	if args := flag.Args(); len(args) != 0 {
		for _, arg := range args {
			filter[arg] = matched(false)
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
	for commandName, matched := range filter {
		binary := binaryName(commandName)
		if matched {
			continue
		}
		fmt.Fprintf(os.Stderr, "cannot find binary %q in any of:\n", binary)
		workspaces := filepath.SplitList(build.Default.GOPATH)
		for i, workspace := range workspaces {
			path := filepath.Join(workspace, "bin", binary)
			switch i {
			case 0:
				fmt.Fprintf(os.Stderr, "\t%s (from $GOPATH)\n", path)
			default:
				fmt.Fprintf(os.Stderr, "\t%s\n", path)
			}
		}
		if len(workspaces) == 0 {
			fmt.Fprintln(os.Stderr, "\t($GOPATH not set)")
		}
	}
	sort.Strings(binaries)
	for _, commandName := range binaries {
		fmt.Println(commandName)
		for _, importPathStatus := range commands[commandName] {
			fmt.Printf("\t%s\n", importPathStatus)
		}
		if len(commands[commandName]) == 0 {
			fmt.Printf("\t(no source package found)\n")
		}
	}

	// If any of the filters weren't matched, exit with code 1.
	for _, matched := range filter {
		if matched {
			continue
		}
		os.Exit(1)
	}
}

type matched bool

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

// TODO: There's a bit of a mismatch between name of func, binaries, and it returning "names of installed commands". Fix that.
// binaries finds binaries in GOPATH/bin directories, filtering results with filter if it's not empty.
func binaries(filter map[string]matched) ([]string, error) {
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
			commandName, ok := commandName(fi)
			if !ok {
				continue
			}

			// If user specified a list of binaries, filter out binaries that don't match.
			if len(filter) != 0 {
				if _, ok := filter[commandName]; !ok {
					continue
				}
				filter[commandName] = matched(true)
			}

			binaries = append(binaries, commandName)
		}
	}

	return binaries, nil
}

// commandName returns the name of Go command that would've resulted in this binary file, if possible.
func commandName(fi os.FileInfo) (commandName string, ok bool) {
	if fi.IsDir() {
		return "", false
	}
	if strings.HasPrefix(fi.Name(), ".") {
		return "", false
	}

	if "windows" == runtime.GOOS {
		if !strings.HasSuffix(fi.Name(), ".exe") {
			return "", false
		}
		return fi.Name()[:len(fi.Name())-4], true
	}
	return fi.Name(), true
}

// binaryName returns the name of binary for the given command name.
func binaryName(commandName string) string {
	if "windows" == runtime.GOOS {
		return commandName + ".exe"
	}
	return commandName
}
