## Basic API usage
using ipfs within your code is really quite simple!

At its simplest, you only need to create a node:
```
import "github.com/ipfs/go-ipfs/core"
.
.
// setup ctx
// setup cfg
.

node, err := core.NewNode(ctx, cfg)
```

The above code snippet is the simplest way to create an ipfs node. Below is explained how to get the context and the configuration objects in place.

### Configuration
Lets create a Node's build configuration:

```
cfg := &core.BuildCfg{
	Repo:    r,
	Online:  true,
	Routing: myRoutingOption,
}
```

A node created in 'Online' mode will start up bootstrapping, bitswap exchange,
and other network interfaces.

#### Repo
The ipfs 'repo' or repository represents all data that persists past a single
instance. This currently includes the configuration file and the local
datastore. By default, you will be given a blank config and an in memory
datastore. To set your own, call `SetRepo` with your own repo object.
The normal way to go about doing this is with an `FSRepo`, which represents
an on disk 'repository'. This looks a bit like:
```
import "github.com/ipfs/go-ipfs/repo/fsrepo"
.
.
.
r := fsrepo.Open("/path/to/.ipfs")
if err != nil {
	// Deal with the error
}
```

#### SetRouting
ipfs by default will use our DHT network for getting provider information and
ipns entries. If you wish to implement a separate routing system for your node
to get this information through, just make an object that implements the
IpfsRouting interface and pass the build configuration a RoutingOption for it.

### Context
If you have never dealt with contexts before, I highly recommend you first go read
[this wonderful explanation](https://blog.golang.org/context). Now, the context
we pass into the new `Node` we are creating is the "master" context to the entire
ipfs node, cancelling that context will shut down every single subprocess that ipfs
runs.

The easiest way to set up a context for an ipfs node is something like this:
```
ctx, cancel := context.WithCancel(context.Background)
```
This creates a context, and an anonymous function that can be called to cancel
the context, and by extension, all of the ipfs node.


