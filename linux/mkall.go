//     go run mkall.go <linux_dir>

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// These will be paths to the appropriate source directories.
var LinuxDir string

const TempDir = "/tmp"
const IncludeDir = TempDir + "/include" // To hold our C headers
const BuildDir = TempDir + "/build"     // To hold intermediate build files

const GOOS = "linux"      // Only for Linux targets
const BuildArch = "amd64" // Must be built on this architecture

type target struct {
	GoArch     string // Architecture name according to Go
	LinuxArch  string // Architecture name according to the Linux Kernel
	GNUArch    string // Architecture name according to GNU tools (https://wiki.debian.org/Multiarch/Tuples)
	SignedChar bool   // Is -fsigned-char needed (default no)
	Bits       int
}

// List of all Linux targets we need to support.
var targets = []target{
	{
		GoArch:    "amd64",
		LinuxArch: "x86",
		GNUArch:   "x86_64-linux-gnu",
		Bits:      64,
	},
	{
		GoArch:    "arm",
		LinuxArch: "arm",
		GNUArch:   "arm-linux-gnueabi",
		Bits:      32,
	},
	{
		GoArch:     "arm64",
		LinuxArch:  "arm64",
		GNUArch:    "aarch64-linux-gnu",
		SignedChar: true,
		Bits:       64,
	},
}

func makeCommand(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	return cmd
}

func (t *target) setTargetBuildArch(cmd *exec.Cmd) {
	cmd.Env = append(os.Environ(), "GOARCH_TARGET="+t.GoArch)
	for i, s := range cmd.Env {
		if strings.HasPrefix(s, "GOARCH=") {
			cmd.Env[i] = "GOARCH=" + BuildArch
		}
	}
}

// Run runs the command, pipes output to a formatter, and write to a file
func (t *target) commandFormatOutput(formatter string, outputFile string, name string, args ...string) error {
	mainCmd := makeCommand(name, args...)
	var err error
	if name == "ecodes" {
		args = append([]string{"run", "linux/mkecodes.go"}, args...)
		mainCmd = makeCommand("go", args...)
		t.setTargetBuildArch(mainCmd)
	}

	fmtCmd := makeCommand(formatter)
	if formatter == "mkpost" {
		fmtCmd = makeCommand("go", "run", "linux/mkpost.go")
		t.setTargetBuildArch(fmtCmd)
	}
	if fmtCmd.Stdin, err = mainCmd.StdoutPipe(); err != nil {
		return err
	}
	if fmtCmd.Stdout, err = os.Create(outputFile); err != nil {
		return err
	}

	if err = fmtCmd.Start(); err != nil {
		return err
	}

	defer func() {
		fmtErr := fmtCmd.Wait()
		if err == nil {
			err = fmtErr
		}
	}()

	fmt.Printf("\tGOOS:   %v\n", os.Getenv("GOOS"))
	fmt.Printf("\tGOARCH: %v\n", os.Getenv("GOARCH"))
	fmt.Printf("\tCommand Env: %v\n", mainCmd.Env)
	fmt.Printf("\tRunning %s\n", mainCmd.String())

	return mainCmd.Run()
}

// Generate all the files for a Linux target
func (t *target) generateFiles() error {
	os.Setenv("GOOS", GOOS)
	os.Setenv("GOARCH", t.GoArch)

	// Get appropriate compiler
	if t.LinuxArch != "x86" {
		compiler := t.GNUArch + "-gcc"
		if _, err := exec.LookPath(compiler); err != nil {
			return err
		}
		os.Setenv("CC", compiler)
	} else {
		os.Setenv("CC", "gcc")
	}

	if err := os.MkdirAll(IncludeDir, os.ModePerm); err != nil {
		return err
	}
	defer os.RemoveAll(IncludeDir)

	// Make headers file
	fmt.Println("Generating header files ...")
	if err := t.makeHeaders(); err != nil {
		return fmt.Errorf("could not make headers file: %v", err)
	}
	fmt.Println("Header files generated")

	fmt.Printf("Generating zecodes_%s_%s ...\n", GOOS, t.LinuxArch)
	if err := t.makeEcodesFile(); err != nil {
		return fmt.Errorf("could not make ecodes file: %v", err)
	}
	fmt.Printf("zecodes_%s_%s generated\n", GOOS, t.LinuxArch)

	fmt.Printf("Generating zinput_%s_%s ...\n", GOOS, t.LinuxArch)
	if err := t.makeInputFile(); err != nil {
		return fmt.Errorf("could not make input file: %v", err)
	}
	fmt.Printf("zinput_%s_%s generated\n", GOOS, t.LinuxArch)

	return nil
}

func (t *target) makeEcodesFile() error {
	zecodesFile := fmt.Sprintf("zecodes_linux_%s.go", t.GoArch)
	ecodesFile := filepath.Join(IncludeDir, "linux/input-event-codes.h")

	args := []string{ecodesFile}
	return t.commandFormatOutput("gofmt", zecodesFile, "ecodes", args...)
}

func (t *target) makeInputFile() error {
	inputFile := fmt.Sprintf("zinput_linux_%s.go", t.GoArch)

	args := []string{"tool", "cgo", "-godefs", "--"}
	args = append(args, t.cFlags()...)
	args = append(args, "linux/types.go")
	return t.commandFormatOutput("mkpost", inputFile, "go", args...)
}

func (t *target) makeHeaders() error {
	linuxMade := makeCommand("make", "headers_install", "ARCH="+t.LinuxArch, "INSTALL_HDR_PATH="+TempDir)
	linuxMade.Dir = LinuxDir
	fmt.Printf("\tRunning %s\n", linuxMade.String())
	if err := linuxMade.Run(); err != nil {
		return err
	}

	if err := os.MkdirAll(BuildDir, os.ModePerm); err != nil {
		return err
	}
	defer os.RemoveAll(BuildDir)

	return nil
}

func (t *target) cFlags() []string {
	flags := []string{"-Wall", "-Werror", "-static", "-I" + IncludeDir}

	if t.LinuxArch == "x86" {
		flags = append(flags, fmt.Sprintf("-m%d", t.Bits))
	}

	return flags
}

func main() {
	if runtime.GOOS != GOOS || runtime.GOARCH != BuildArch {
		fmt.Printf("Build system has GOOS_GOARCH = %s_%s, need %s_%s\n",
			runtime.GOOS, runtime.GOARCH, GOOS, BuildArch)
		return
	}

	// Parse the command line options
	if len(os.Args) != 2 {
		fmt.Println("USAGE: go run mkall.go <linux_dir> <glibc_dir>")
		return
	}
	LinuxDir = os.Args[1]
	for _, t := range targets {
		fmt.Printf("----- GENERATING: %s -----\n", t.GoArch)
		if err := t.generateFiles(); err != nil {
			fmt.Printf("%v\n***** FAILURE:    %s *****\n\n", err, t.GoArch)
		} else {
			fmt.Printf("----- SUCCESS:    %s -----\n\n", t.GoArch)
		}
	}
}
