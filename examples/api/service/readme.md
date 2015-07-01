## Making your own ipfs service
ipfs has a few default services that it runs by default, such as the dht,
bitswap, and the diagnostics service. Each of these simply registers a
handler on the ipfs PeerHost, and listens on it for new connections.  The
`corenet` package has a very clean interface to this functionality. So lets
try building an easy demo service to try this out!

Lets start by building the service host:
```
package main

import (

	"fmt"

	core "github.com/ipfs/go-ipfs/core"
	corenet "github.com/ipfs/go-ipfs/core/corenet"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"

	"code.google.com/p/go.net/context"
)
```

We dont need too many imports for this.
Now, the only other thing we need is our main function:

Set up an ipfsnode.

```
func main() {

	// Basic IPFS Node setup
	r, err := fsrepo.Open("~/.ipfs")
	if err!=nil {
	  panic(err)
	}

	nb := core.Online(r)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd, err := core.NewIPFSNode(ctx, nb)
	if err != nil {
		panic(err)
	}
```

Thats just the basic template of code to initiate a default ipfsnode from
the config in the users `~/.ipfs` directory.

Next, we are going to build our service.

```

	list, err := corenet.Listen(nd, "/app/zero")
	if err != nil {
		panic(err)
	}

	fmt.Printf("I am peer %s\n", peer.IDB58Encode(nd.Identity))

	for {
		con, err := list.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer con.Close()

		fmt.Fprintln(con, "ZERO IPFS service. Nothing to see here.")
		fmt.Printf("Connection from: %s\n", con.Conn().RemotePeer())
	}
}
```

And thats really all you need to write a service on top of ipfs. When a client
connects, we send them our greeting, print their peer ID to our log, and close
the session. This is the simplest possible service, and you can really write
anything you want to handle the connection.

Now we need a client to connect to us:

```
package main

import (

	"fmt"
	"io"
	"os"

	core "github.com/ipfs/go-ipfs/core"
	corenet "github.com/ipfs/go-ipfs/core/corenet"
	peer "github.com/ipfs/go-ipfs/p2p/peer"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"

	"code.google.com/p/go.net/context"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please give peer ID as an argument")
	}

	target, err := peer.IDB58Decode(os.Args[1])
	if err != nil {
		fmt.Println("Invalid peer ID")
		panic(err)
	}

	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		panic(err)
	}

	nb := core.NewNodeBuilder().Online()
	nb.SetRepo(r)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nd, err := nb.Build(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("I am peer %s dialing %s\n", nd.Identity, target)

	con, err := corenet.Dial(nd, target, "/app/zero")
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, con)

}
```

This client will set up their ipfs node (note: this is moderately expensive and
you normally wont just spin up an instance for a single connection) and dial the
service we just created.

To try it out, run the following on one computer:
```
$ ipfs init # if you havent already
$ go run host.go
```

That should print out that peers ID, copy it and use it on a second machine:
```
$ ipfs init # if you havent already
$ go run client.go <peerID>
```

It should print out `Hello! This is whyrusleepings awesome ipfs service`

Now, you might be asking yourself: "Why would I use this? How is it better than
the `net` package?". Well, here are the advantages:

1. You dial a specific peerID, no matter what their IP address happens to be at the moment.
2. You take advantage of the NAT traversal built into our net package.
3. Instead of a 'port' number, you get a much more meaningful protocol ID string.

By [whyrusleeping](http://github.com/whyrusleeping)
