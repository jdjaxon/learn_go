## Chapter 1 Notes
- `go run`: executes the given file without creating a binary
- `go build`: creates binary from the provided file(s)
    - ldflags
        - `-w`: turns off DWARF debugging info; you won't be able to use gdb or pprof profiling
        - `-s`: turns off generation of the Go symbol table; won't be able to use `go tool nm` to list symbols in the binary
        - `go build -ldflags "-w -s"` reduces the binary size by approximately 30%

    - cross-compiling
        ```
        GOOS="linux" GOARCH="amd64" go build hello.go
        ```
