go-test  module  2020-09-28
=======

Source Code for Go In Action examples
Modified to Test

> cd /Volumes/development/git/golang/
> mkdir go-test
> cd go-test


----------------------------------------------
GO module
----------------------------------------------

SEE: https://golang.org/cmd/go/#hdr-GOPATH_and_Modules
     https://blog.golang.org/using-go-modules
     https://golang.org/doc/code.html


For more fine-grained control, the go command continues to respect a temporary environment variable, GO111MODULE, 
which can be set to one of three string values: off, on, or auto (the default). 
If GO111MODULE=on, then the go command requires the use of modules, never consulting GOPATH.
If GO111MODULE=off, then the go command never uses module support.
If GO111MODULE=auto or is unset, then the go command enables or disables module support based on the current directory. 
   Module support is enabled only when the current directory contains a go.mod file or is below a directory containing a go.mod file.

In module-aware mode, GOPATH no longer defines the meaning of imports during a build, 
but it still stores downloaded dependencies (in GOPATH/pkg/mod) and installed commands (in GOPATH/bin, unless GOBIN is set).

-- check with:
> go env

-- create the module with github.com/bethmage/go-test

> go mod init github.com/bethmage/go-test

The resulting file ./go.mod

module github.com/bethmage/go-test
go 1.15



--------------------------------------------------------------------------------
NOTE: the $GOPATH shoud not contain the  /Volumes/development/git/golang/go-test 
      because we are now working with modules NOT with the  GOPATH !
      (if is was set before ... unset GOPATH

      If we decided to work with the $GOPATH (not with go.mod)
      Then the project path should be $GOPATH/src/github.com/bethmage/go-test
      this allows the go build to find packages locally.
--------------------------------------------------------------------------------


-- to compile or build must cd into the directory : go-test  
   VSCode go-test for sample.

> cd go-test
├── README.md
├── go.mod
├── hello.go
└── morestrings
    └── reverse.go

Code in or below a directory named "morestrings" is importable 
only by code in the directory tree rooted at the parent of "morestrings". 

The code in reverse.go is imported as "github.com/bethmage/go-test/morestrings", 
but that import statement can only appear in source files in the subtree rooted at go-test (parent of more strings). 

Just for now I Think parent is the directory which contains the go.mod file.

SEE: https://golang.org/cmd/go/#hdr-GOPATH_and_Modules




-- See the GO environment
> go env

-- set GOBIN place to put the binaries
   SEE: https://golang.org/doc/code.html#Command 

> go env -w GOBIN=/Volumes/development/git/golang/go-test/bin

-- check the GOBIN 
> dirname $(go list -f '{{.Target}}' .)

----------------------------------------------------------------------------
!!! go env ´settings remains!! also after Terminal exit.
    so we must Change or unset it !! OR use $GOBIN env-var insted of go env.
-----------------------------------------------------------------------------
-- When need to unset then use:
   -- To unset a variable previously set by go env -w, use go env -u:
   >  go env -u GOBIN


-- add it to the path
> export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))


-- create a file named hello.go  in directory go-test

package main

import "fmt"

func main() {
	fmt.Println("Hello, one world.")
}

-- build the binary into this folder
$ go build .

-- build and install that program with the go tool:

$ go install  github.com/bethmage/go-test

The install directory is controlled by the GOPATH and GOBIN environment variables. 
If GOBIN is set, binaries are installed to that directory. 
If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. 
Otherwise, binaries are installed to the bin subdirectory of the default GOPATH  $HOME/go

in my case to: /Volumes/development/git/golang/go-test/bin
-------------

Commands like go install apply within the context of the module containing the current working directory. 
If the working directory is not within the go-test module, go install may fail.

For convenience, go commands accept paths relative to the working directory, 
and default to the package in the current working directory if no other path is given. 
So in our working directory: go-test , the following commands are all equivalent:

$ go install github.com/bethmage/go-test
$ go install .
$ go install

-- run the binary .. $GOBIN is already added to the $PATH above.

> /Volumes/development/git/golang/go-test/bin/hello
or
> hello

Hello, one world.

DONE.



-----------------------------------------------------------
HOW-TO RUN godoc
-----------------------------------------------------------

option 1:
> go help doc 
> go doc 
Sample : go doc fmt.printf

option 2:

install godoc with following command:

> go get golang.org/x/tools/cmd/godoc

Start godoc server:
> $GOPATH/bin/godoc -http=:6060

In your browser, visit:
http://localhost:6060

Info: 
  HowTo Document my own go Code : https://blog.golang.org/godoc  

  Any Go packages installed inside $GOROOT/src/pkg and any GOPATH work spaces will already be accessible via godoc's command-line and HTTP interfaces, 
  By default, godoc looks at the packages it finds via $GOROOT and $GOPATH (if set). 

  See the [godoc documentation](https://godoc.org/golang.org/x/tools/cmd/godoc) for more details.
  
  if you run : $GOPATH/bin/godoc  -http=localhost:6060 -play           (-play show the playground))
  insiide of yor project root (where go.mod lives) directory then the package Doc are under the "Third party"
  
  Howto write Testable Examples in Go ... which can be tested in godoc -play RUN
  See: https://blog.golang.org/examples



-----------------------------------------------------------
HOW-TO VSCode 
-----------------------------------------------------------

Install gopls (“go please”) in VSCode
SEE: https://github.com/golang/vscode-go

