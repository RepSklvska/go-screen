# go-screen
Run shell commands in GNU Screen using Golang.

## Requriements

- [GNU Screen](https://www.gnu.org/software/screen/)

## Install

    $ go get -u github.com/RepSklvska/go-screen

## Example

```go
package main

import "github.com/RepSklvska/go-screen"

func main() {
	var (
		command = []string{"emacs", "-nw"}
		tty     = "abc"
	)
	
	screen.Execute(tty, command...)
}
```
Then you can find the session with "screen -ls" command.
