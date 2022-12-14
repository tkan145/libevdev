# libevdev

## Building
The build system uses a Docker container to generate the Go files directly from the source
checkouts of the kernel. This means that on any platform that support Docker, all the file
can be generated at once, and generated files will not change based on what the person running
the scripts has installed on their computer.

The OS specific files for the build system are located in the `${GOOS}` directory, and the build
is coordinated by the `${GOOS}/mkall.go` program. When the kernel or system library updates, modify
the Dockerfile at `${GOOS}/Dockerfile` to checkout the new release of the source.

To build all the file, you must be on an amd/linux system and have your GOOS and GOARCH set
accordingly. Running `build.sh` will then generate all of the files for all of the GOOS/GOARCH
pairs.

Requirements:

-   bash
-   go
-   Docker

## Component files

This section describes the various files used in the code generation process.

### mkall.go

This program generate all zecodes, zinput for all linux architectures. Append the targets array
within the file to add support for new architecture/OS

### mkecodes.go

This program takes in a list of header files containing the ecodes number declarations and parses
them to produce the corresponding list of Go numeric constants. See `zecodes_${GOOS}_${GOARCH}.go`
for the generated constants.

### types files

For each OS, there is a hand-written Go file at `${GOOS}/types.go`. This file included standard
C headers and create Go types alias to the corresponding C types. The file is then fed though
`godefs` to generate the Go compatible definitions. Finally, the generated code is fed through
`mkpost.go` to format the code and add any required import. The cleaned-up code is writtent to
`zinput_${GOOS}_${GOARCH}.go`

To add a new type, add in the necessary include statement at the top of the file and add in a type
alias line..

## Generated files

### `zecodes_${GOOS}_${GOARCH}.go`

A file containing all of the ecodes generated number.

### `zinput_${GOOS}_${GOARCH}.go`

A file containing Go types for parsing into evdev event. Generated by godefs and the
type files (see above).