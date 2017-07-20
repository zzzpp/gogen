# gogen

bootstrap a .go file and pipe to stdout.


## Installation

```
$ go get github.com/zzzpp/gogen
```


## Usage

```
$ gogen gen 

package main
import (
)
func main() {
}

$ gogen gen --i=fmt --i=os 's,_:=os.Hostname()' 'fmt.Println("hello ", s)'

package main
import (
	"fmt"
	"os"
)
func main() {
	s,_:=os.Hostname()
	fmt.Println("hello ", s)
}
```

or play it with shell pipe:

```
$ out=/tmp/gogenOut.go && \
	gogen gen --i=fmt --i=os \
	's,_:=os.Hostname()' 'fmt.Println("hello ", s)' \
	|cat >$out && go run $out

hello  ${the-hostname}
```

## License
MIT
