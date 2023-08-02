# go-realpath

Realpath for Go. This finds the true path of a file or directory, resolving any symbolic links.

It adds some extra features to the yookoala library's [filepath](https://github.com/yookoala/realpath) package. Instead of manually iterating over the path components, it uses the `filepath.Abs`, `filepath.Clean` function to do it for you. It also evaluates symbolic links.

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
