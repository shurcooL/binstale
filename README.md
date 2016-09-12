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
		STALE: golang.org/x/tools/cmd/goimports (newer dependency)
	$ go install golang.org/x/tools/cmd/goimports
	$ binstale goimports
	goimports
		up to date: golang.org/x/tools/cmd/goimports

	$ binstale
	Go-Package-Store
		STALE: github.com/shurcooL/Go-Package-Store (newer source file)
	binstale
		up to date: github.com/shurcooL/binstale
	doc
		(no source package found)
	dump_args
		STALE: github.com/shurcooL/cmd/dump_args (build ID mismatch)
	dump_httpreq
		STALE: github.com/shurcooL/cmd/dump_httpreq (build ID mismatch)
	dupl
		STALE: github.com/mibk/dupl (build ID mismatch)
	git-branches
		STALE: github.com/shurcooL/cmd/git-branches (build ID mismatch)
	git-codereview
		STALE: golang.org/x/review/git-codereview (build ID mismatch)
	go-find-references
		STALE: github.com/lukehoban/go-find-references (build ID mismatch)
	go-outline
		up to date: github.com/lukehoban/go-outline
	gocode
		up to date: github.com/nsf/gocode
	godef
		STALE: github.com/rogpeppe/godef (build ID mismatch)
	godep
		STALE: github.com/tools/godep (build ID mismatch)
	goexec
		STALE: github.com/shurcooL/goexec (build ID mismatch)
	goimporters
		STALE: github.com/shurcooL/cmd/goimporters (build ID mismatch)
	goimportgraph
		STALE: github.com/shurcooL/cmd/goimportgraph (build ID mismatch)
	goimports
		up to date: golang.org/x/tools/cmd/goimports
	golint
		STALE: github.com/golang/lint/golint (build ID mismatch)
	gomobile
		STALE: golang.org/x/mobile/cmd/gomobile (build ID mismatch)
	gomvpkg
		STALE: golang.org/x/tools/cmd/gomvpkg (build ID mismatch)
	gopherjs
		STALE: github.com/gopherjs/gopherjs (build ID mismatch)
	gorename
		STALE: golang.org/x/tools/cmd/gorename (build ID mismatch)
	gorepogen
		STALE: github.com/shurcooL/cmd/gorepogen (build ID mismatch)
	gostatus
		STALE: github.com/shurcooL/gostatus (build ID mismatch)
	gostringer
		STALE: github.com/sourcegraph/gostringer (build ID mismatch)
	govers
		STALE: github.com/rogpeppe/govers (cannot stat install target)
	gtdo
		STALE: github.com/shurcooL/gtdo (build ID mismatch)
	implements
		up to date: honnef.co/go/implements
	jsonfmt
		up to date: github.com/shurcooL/cmd/jsonfmt
	markdownfmt
		up to date: github.com/shurcooL/markdownfmt
	staticcheck
		STALE: honnef.co/go/staticcheck/cmd/staticcheck (build ID mismatch)
	stringer
		STALE: golang.org/x/tools/cmd/stringer (build ID mismatch)
	unconvert
		up to date: github.com/mdempsky/unconvert
	unused
		STALE: honnef.co/go/unused/cmd/unused (build ID mismatch)
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
