/*
binstale tells you whether the binaries in your GOPATH/bin are stale or up to date.

	$ binstale -help
	Usage: binstale [command names]

Example Output

This is an example of binstale usage.

	$ binstale goimports
	goimports
		STALE: golang.org/x/tools/cmd/goimports
	$ go install golang.org/x/tools/cmd/goimports
	$ binstale goimports
	goimports
		up to date: golang.org/x/tools/cmd/goimports
	$ binstale
	Go-Package-Store
		STALE: github.com/shurcooL/Go-Package-Store
	binstale
		up to date: github.com/shurcooL/binstale
	doc
		(no source package found)
	dump_args
		STALE: github.com/shurcooL/cmd/dump_args
	dump_glfw3_joysticks
		STALE: github.com/shurcooL/cmd/dump_glfw3_joysticks
	dump_httpreq
		STALE: github.com/shurcooL/cmd/dump_httpreq
	dupl
		STALE: github.com/mibk/dupl
	gencomponent
		STALE: github.com/neelance/dom/gencomponent
	git-branches
		STALE: github.com/shurcooL/cmd/git-branches
	git-codereview
		STALE: golang.org/x/review/git-codereview
	go-bindata
		STALE: github.com/jteeuwen/go-bindata/go-bindata
	go-find-references
		STALE: github.com/lukehoban/go-find-references
	go-outline
		up to date: github.com/lukehoban/go-outline
	gocode
		up to date: github.com/nsf/gocode
	godef
		STALE: github.com/rogpeppe/godef
	godep
		STALE: github.com/tools/godep
	goexec
		STALE: github.com/shurcooL/goexec
	goimporters
		STALE: github.com/shurcooL/cmd/goimporters
	goimportgraph
		STALE: github.com/shurcooL/cmd/goimportgraph
	goimports
		up to date: golang.org/x/tools/cmd/goimports
	golint
		STALE: github.com/golang/lint/golint
	gomobile
		STALE: golang.org/x/mobile/cmd/gomobile
	gomvpkg
		STALE: golang.org/x/tools/cmd/gomvpkg
	gopherjs
		STALE: github.com/gopherjs/gopherjs
	gorename
		STALE: golang.org/x/tools/cmd/gorename
	gorepogen
		STALE: github.com/shurcooL/cmd/gorepogen
	gostatus
		STALE: github.com/shurcooL/gostatus
	gostringer
		STALE: github.com/sourcegraph/gostringer
	govers
		STALE: github.com/rogpeppe/govers
	gtdo
		STALE: github.com/shurcooL/gtdo
	implements
		up to date: honnef.co/go/implements
	jsonfmt
		up to date: github.com/shurcooL/cmd/jsonfmt
	markdownfmt
		STALE: github.com/shurcooL/markdownfmt
	rune_stats
		STALE: github.com/shurcooL/cmd/rune_stats
*/
package main
