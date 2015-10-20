## Ipfs Demo Application
Lets take a look at making an application that uses ipfs. Lets say that our
fictional application needs to read in a file, and count all the letters in
it. But, being tech savvy as we are, we want the program to be able to take the
files in from ipfs, so that users from around the world need only send them a
hash in order for their files to be processed!

To Start off, lets get some imports:
```
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"code.google.com/p/go.net/context"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreunix"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)
```


Now, lets make a quick function to do a frequency count on characters:

```
func CountChars(r io.Reader) map[byte]int {
	m := make(map[byte]int)
	buf := bufio.NewReader(r)
	for {
		b, err := buf.ReadByte()
		if err != nil {
			return m
		}
		m[b]++
	}
}
```

Alright, now for the ipfs goodness:

```
func SetupIpfs() (*core.IpfsNode, error) {
	// Assume the user has run 'ipfs init'
	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		return nil, err
	}

	cfg := &core.BuildCfg{
		Repo:   r,
		Online: true,
	}

	cfg := new(core.BuildCfg)
	cfg.Repo = r
	cfg.Online = true

	return core.NewNode(context.Background(), cfg)
}
```

We've got our node construction out of the way now, lets move on to actually
doing something.

```
func main() {
	nd, err := SetupIpfs()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Please pass in an argument!")
		return
	}
	keytofetch := os.Args[1]

	read, err := coreunix.Cat(nd, keytofetch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(CountChars(read))
}
```

And thats it, the user passes in a file, and we read it from ipfs. If no such
file exists, we will error out from the `Cat` method. `Cat` returns a reader
that will manage retrieving the file specified by the given hash, whether its
stored locally on disk or if its pieces are split apart on multiple different
machines across the planet.

By [whyrusleeping](http://github.com/whyrusleeping)
