//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	MACRO_REGEX = `^#define +((?:KEY|ABS|REL|SW|MSC|LED|BTN|REP|SND|ID|EV|BUS|SYN|FF)_\w+)\s+(\w+)`
)

type re struct {
	str string
	sub []string
}

// Match perform regular expression match
func (r *re) Match(exp string) bool {
	r.sub = regexp.MustCompile(exp).FindStringSubmatch(r.str)
	return r.sub != nil
}

func format(name string, num string) string {
	name = strings.ToUpper(name)
	return fmt.Sprintf("%s = %s\n", name, num)
}

func main() {
	goos := os.Getenv("GOOS")
	goarch := os.Getenv("GOARCH_TARGET")

	if goarch == "" {
		goarch = os.Getenv("GOARCH")
	}

	if goarch == "" || goos == "" {
		fmt.Fprintf(os.Stderr, "GOARCH or GOOS not defined in environment\n")
		os.Exit(1)
	}

	// Run the file thought the pre-processor so that we can get the sane file output.
	cc := os.Getenv("CC")
	if cc == "" {
		fmt.Fprintf(os.Stderr, "CC not defined in environment\n")
		os.Exit(1)
	}
	args := os.Args[1:]
	args = append([]string{"-E", "-dD"}, args...)
	cmd, err := exec.Command(cc, args...).Output() // Execute command and capture output
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't run %s", cmd)
		os.Exit(1)
	}

	text := ""
	s := bufio.NewScanner(strings.NewReader(string(cmd)))

	for s.Scan() {
		t := re{str: s.Text()}
		// Extract #define with regex
		if t.Match(MACRO_REGEX) {
			text += format(t.sub[1], t.sub[2])
		}
	}

	err = s.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	goBuildTags := fmt.Sprintf("%s && %s", goarch, goos)
	plusBuildTags := fmt.Sprintf("%s,%s", goarch, goos)
	fmt.Printf(template, "go run mkall.go", goBuildTags, plusBuildTags, text)
}

var template = `// %s
// Code generate by the command above; see README.md. DO NOT EDIT

//go:build %s
// +build %s

package libevdev

const(
%s)
`
