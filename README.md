# go-realpath

Realpath for Go. This finds the true path of a file or directory, resolving any symbolic links.

## Installation

```bash
go get github.com/gandarez/go-realpath
```

## Usage

```go
import "github.com/gandarez/go-realpath"

func main() {
    rp, err := realpath.Realpath("/some/path")
}
```
