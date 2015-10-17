## Basic API usage
using ipfs within your code is really quite simple!

At its simplest, you only need to create a node:
```
import "github.com/ipfs/go-ipfs/core"
.
.
.
builder := core.NewNodeBuilder()
node, err := builder.Build(ctx)
```

The above code snippet is the simplest way to create an ipfs node. There are
a couple different things that I think deserve explanation.

#### NodeBuilder
The NodeBuilder is an object following the 'builder' pattern (who would have
guessed?). It can be used to configure the node before its actually constructed.
It has a few different setters and other options that we will discuss in a bit.

#### Contexts
If you've never dealt with contexts before, I highly recommend you first go read
[this wonderful explanation](https://blog.golang.org/context). Now, the context
we pass into `Build` is the "master" context to the entire ipfs node, cancelling
that context will shut down every single subprocess that ipfs runs.

The easiest way to set up a context for an ipfs node is something like this:
```
ctx, cancel := context.WithCancel(context.Background)
```
This creates a context, and an anonymous function that can be called to cancel
the context, and by extension, all of the ipfs node.

### Ipfs Options
So, now that all of that is out of the way, lets look at different configuration
options.

#### Online/Offline
The default state for a nodebuilder is 'Offline', so set it to 'Online' simply
call the `Online()` method.
```
builder := core.NewNodeBuilder().Online()
```
Or:
```
builder := core.NewNodeBuilder()
builder.Online()
```

A node created in 'Online' mode will start up bootstrapping, bitswap exchange,
and other network interfaces.

#### SetRepo
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
r := fsrepo.At("/path/to/.ipfs")
if err := r.Open(); err != nil {
	// Deal with the error
}

builder.SetRepo(r)
```

#### SetRouting
ipfs by default will use our DHT network for getting provider information and
ipns entries. If you wish to implement a separate routing system for your node
to get this information through, just make an object that implements the
IpfsRouting interface and pass the builder a RoutingOption for it.
```
builder := core.NewNodeBuilder().Online()
builder.SetRouting(myRoutingOption)
```

