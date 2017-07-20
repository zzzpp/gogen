# gogen

bootstrap a .go file and pipe to stdout.


## Installation
	$ go get github.com/zzzpp/gogen


# Usage

	$ gogen gen --i=fmt --i=os  a:=1 a+=1
	
outputs

```
package main
import (
	"fmt"
	"os"
)
func main() {
	a:=1
	a+=1
}
```


# License
MIT
