## Random Ipfs Objects
During the development, ive frequently found the need to just get a hash of some
random ipfs object. At first, I would just ask in irc "can someone give me a
hash?". But I decided I could do better, So I decided to make it a service. In
this article, im going to go over how I did that (hint: its really simple!)

First, lets get some imports:
```
package main

import (
	"context"
	"io"
	"net/http"

	u "github.com/ipfs/go-ipfs-util"
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreunix"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)
```

This just pulls in some basic ipfs packages, and the default golang http server.
Now, since im lazy, im going to have a global for our ipfsnode.

```
var gnode *core.IpfsNode
```

Now, lets write the http handler func for generating our random objects.

```
func ServeIpfsRand(w http.ResponseWriter, r *http.Request) {
	read := io.LimitReader(u.NewTimeSeededRand(), 2048)

	str, err := coreunix.Add(gnode, read)
	if err != nil {
		w.WriteHeader(504)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(str))
	}
}
```

And now, lets tie it all together in a main function.

Set up our node configuration, and use the users standard ipfs configuration directory.

```
func main() {
	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		panic(err)
	}
```

Now we need to set up our context

```
	// Make our 'master' context and defer cancelling it
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
```

Then create a configuration and finally create our node!

```
	cfg := &core.BuildCfg{
		Repo:   r,
		Online: true,
	}

	node, err := core.NewNode(ctx, cfg)
	if err != nil {
		panic(err)
	}

	// Set the global node for access in the handler
	gnode = node

	http.HandleFunc("/ipfsobject", ServeIpfsRand)
	http.ListenAndServe(":8080", nil)
}
```

By [whyrusleeping](http://github.com/whyrusleeping)
