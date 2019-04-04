// Usage of `ldflags`.
/*
Append version information from git tags

	VERSION="$(git describe --tags)"
	if [ $? -ne 0 ]; then
		VERSION="dev"
	fi
	go build -ldflags "-X main.Version=$VERSION"
	#or
	go run -ldflags "-X main.Version=$VERSION"
*/
package main

import "fmt"

// can not change by -ldflags
const Name = "cmdname"

// may to change by -ldflags "-X main.Commit=$COMMIT -X main.Date=$DATE"
var (
	// git describe --tags
	Version = "dev"
	// git rev-parse --abbrev-ref HEAD
	Branch = ""
	// git rev-parse --verify --short HEAD
	Commit = ""
	// date -u +%Y-%m-%d
	Date = ""
)

func main() {
	fmt.Printf("%s %s (%s %s %s)\n", Name, Version, Commit, Branch, Date)
}
