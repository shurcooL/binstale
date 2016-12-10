binstale
========

[![Build Status](https://travis-ci.org/shurcooL/binstale.svg?branch=master)](https://travis-ci.org/shurcooL/binstale) [![GoDoc](https://godoc.org/github.com/shurcooL/binstale?status.svg)](https://godoc.org/github.com/shurcooL/binstale)

binstale tells you whether the binaries in your GOPATH/bin are stale or up to date.

	$ binstale -help
	Usage: binstale [command names]

Example Output

This is an example of binstale usage.

	$ binstale goimports
	goimports
		stale: golang.org/x/tools/cmd/goimports (newer dependency)
	$ go install golang.org/x/tools/cmd/goimports
	$ binstale goimports
	goimports
		up to date: golang.org/x/tools/cmd/goimports

	$ binstale
	Go-Package-Store
		stale: github.com/shurcooL/Go-Package-Store (newer source file)
	binstale
		up to date: github.com/shurcooL/binstale
	doc
		(no source package found)
	dump_args
		stale: github.com/shurcooL/cmd/dump_args (build ID mismatch)
	dump_httpreq
		stale: github.com/shurcooL/cmd/dump_httpreq (build ID mismatch)
	dupl
		stale: github.com/mibk/dupl (build ID mismatch)
	git-branches
		stale: github.com/shurcooL/cmd/git-branches (build ID mismatch)
	git-codereview
		stale: golang.org/x/review/git-codereview (build ID mismatch)
	go-find-references
		stale: github.com/lukehoban/go-find-references (build ID mismatch)
	go-outline
		up to date: github.com/lukehoban/go-outline
	gocode
		up to date: github.com/nsf/gocode
	godef
		stale: github.com/rogpeppe/godef (build ID mismatch)
	godep
		stale: github.com/tools/godep (build ID mismatch)
	goexec
		stale: github.com/shurcooL/goexec (build ID mismatch)
	goimporters
		stale: github.com/shurcooL/cmd/goimporters (build ID mismatch)
	goimportgraph
		stale: github.com/shurcooL/cmd/goimportgraph (build ID mismatch)
	goimports
		up to date: golang.org/x/tools/cmd/goimports
	golint
		stale: github.com/golang/lint/golint (build ID mismatch)
	gomobile
		stale: golang.org/x/mobile/cmd/gomobile (build ID mismatch)
	gomvpkg
		stale: golang.org/x/tools/cmd/gomvpkg (build ID mismatch)
	gopherjs
		stale: github.com/gopherjs/gopherjs (build ID mismatch)
	gorename
		stale: golang.org/x/tools/cmd/gorename (build ID mismatch)
	gorepogen
		stale: github.com/shurcooL/cmd/gorepogen (build ID mismatch)
	gostatus
		stale: github.com/shurcooL/gostatus (build ID mismatch)
	gostringer
		stale: github.com/sourcegraph/gostringer (build ID mismatch)
	govers
		stale: github.com/rogpeppe/govers (cannot stat install target)
	gtdo
		stale: github.com/shurcooL/gtdo (build ID mismatch)
	implements
		up to date: honnef.co/go/implements
	jsonfmt
		up to date: github.com/shurcooL/cmd/jsonfmt
	markdownfmt
		up to date: github.com/shurcooL/markdownfmt
	staticcheck
		stale: honnef.co/go/staticcheck/cmd/staticcheck (build ID mismatch)
	stringer
		stale: golang.org/x/tools/cmd/stringer (build ID mismatch)
	unconvert
		up to date: github.com/mdempsky/unconvert
	unused
		stale: honnef.co/go/unused/cmd/unused (build ID mismatch)
	vfsgendev
		up to date: github.com/shurcooL/vfsgen/cmd/vfsgendev

Installation
------------

```bash
go get -u github.com/shurcooL/binstale
```

License
-------

-	[MIT License](https://opensource.org/licenses/mit-license.php)
