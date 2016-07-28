# Implementing a new libp2p transport

Libp2p is the modular networking stack used by ipfs. It is designed in a way that anyone should be able to make it work with little effort over any given transport. In this example, i'm going to go through how to make libp2p (and consequently ipfs) work over websockets. 

To implement a transport on the go side of libp2p, all you need to do is satisfy the `Transport` interface in [go-libp2p-transport](https://github.com/ipfs/go-libp2p-transport). To do this, I'm going to use an existing websockets library from [golang.org/x/net/websocket](https://golang.org/x/net/websocket) and create my own package `ws-transport`.

## Multiaddr Handling
The first thing we're going to need is a multiaddr parser to be able to reference our websocket addresses. 

Lets create a multiaddr protocol for this:
```go
var WsProtocol = ma.Protocol{
	Code:  477,
	Name:  "ws",
	VCode: ma.CodeToVarint(477),
}
```

Here we select `ws` as the tag for websockets, and `477` as the binary code for websockets.  We also need to implement a few extra things as shims so we can reuse a lot of the code already written in [go-multiaddr-net](https://github.com/jbenet/go-multiaddr-net).

Next, a `multiaddr-net` shim:

```go
var WsCodec = &manet.NetCodec{
	NetAddrNetworks:  []string{"websocket"},
	ProtocolName:     "ws",
	ConvertMultiaddr: ConvertWebsocketMultiaddrToNetAddr,
	ParseNetAddr:     ParseWebsocketNetAddr,
}
```

This defines a codec to allow libp2p to translate conventional address types to multiaddr types. The two functions referenced there implement those translations, and you can check out the code [here](https://github.com/whyrusleeping/ws-transport/blob/master/websocket.go#L43).

Now, a formatter object to parse and validate websocket multiaddrs:

```go
var WsFmt = mafmt.And(mafmt.TCP, mafmt.Base(WsProtocol.Code))
```

The above statement declares a websocket address as a `TCP` address with a `/ws/` protocol tag at the end, for example: `/ip4/1.2.3.4/tcp/8005/ws`

Now that we have those, we should register them with the multiaddr library in our package initialization function:

```go
func init() {
	err := ma.AddProtocol(WsProtocol)
	if err != nil {
		log.Fatalf("error registering websocket protocol: %s", err)
	}

	manet.RegisterNetCodec(WsCodec)
}
```

## The Transport

Now that all that nonsense is out of the way, let's get to actually implementing the Transport. To do this, we need four main types, the `Transport`, the `Dialer`, the `Listener` and a `Conn` type.

The transport is used to manage the creation of dialers and listeners. The reason we have an object for that instead of just using constructor functions is that some transports want to be able to reuse certain resources, for example TCP with reuseport, or UTP reusing the same underlying socket for dialing out. That transport struct in our case is very simple, and looks like:

```go
type WebsocketTransport struct{}

func (t *WebsocketTransport) Matches(a ma.Multiaddr) bool {
	return WsFmt.Matches(a)
}

func (t *WebsocketTransport) Dialer(_ ma.Multiaddr, opts ...tpt.DialOpt) (tpt.Dialer, error) {
	return &dialer{}, nil
}

func (t *WebsocketTransport) Listen(a ma.Multiaddr) (tpt.Listener, error) {
	list, err := manet.Listen(a)
	if err != nil {
		return nil, err
	}

	tlist := t.wrapListener(list)

	u, err := url.Parse("ws://" + list.Addr().String())
	if err != nil {
		return nil, err
	}

	s := &ws.Server{
		Handler: tlist.handleWsConn,
		Config:  ws.Config{Origin: u},
	}

	go http.Serve(list.NetListener(), s)

	return tlist, nil
}
```

In this code, we define a transport object and three methods on it. 

First, we have `Matches`. This method is used to determine if a transport can handle a given multiaddr. At first glance, this is just a way to match TCP addresses to TCP transports and so on, but it can also be used for other more interesting things. For example, you could use a different transport for TCP connections over the localhost vs TCP connections over WAN or LAN.

The final two methods are pretty self explanatory, `Listen` and `Dialer`. `Listen` creates a listener object to accept incoming connections from the specified multiaddr, and `Dialer` creates an object (with options) that can be used to create outgoing connections.

For now, we aren't supporting any special dialing options, so the implementation of `Dialer` is trivial. The implementation of `Listen` sets up a websockets server and will return connections that get made to it.

### The Listener

The interface for the listener is somewhat simple as well:

```go
type Listener interface {
    Accept() (Conn, error)
    Close() error
    Addr() net.Addr
    Multiaddr() ma.Multiaddr
}
```

The most complicated method here is `Accept`.  It just needs to return the next incoming connection. `Close` needs to shut down the listener, and `Addr` and `Multiaddr` need to return the address (in respective formats) that this listener is listening on.

### The Dialer

```go
type Dialer interface {
    Dial(ma.Multiaddr) (Conn, error)
    Matches(ma.Multiaddr) bool
}
```

Even smaller, the `Dialer` simply exists to open connections to peers over this transport. The `Matches` method serves the same purpose as discussed earlier on the transport.

### The Conn

The final piece of the puzzle, the `Conn` object is a bit more involved than the others. Its interface as described in `go-libp2p-transport` looks like this:

```go
type Conn interface {
    manet.Conn
    Transport() Transport
}
```

This says that we need to implement all the methods of a `manet.Conn` plus the additional `Transport` method which is used to identify where a given `Conn` came from. You can manually implement these methods (defined [here](https://github.com/jbenet/go-multiaddr-net/blob/master/net.go#L14)) or, if you already have an object that implements the `net.Conn` interface (from the go stdlib `net` package) you can use `go-multiaddr-net`'s  `WrapNetConn` method to fulfill the rest. 

### Plugging it in

Once you have that code down, you will have successfully written a go-libp2p transport! The next step is to actually integrate that into go-libp2p so you can use it in ipfs. This involves editing the [libp2p swarm code](https://github.com/ipfs/go-libp2p/blob/master/p2p/net/swarm/swarm.go#L118), the baseline diff to accomplish this for websockets looks like:

```diff
diff --git a/p2p/net/swarm/swarm.go b/p2p/net/swarm/swarm.go
index 100a982..5c06a10 100644
--- a/p2p/net/swarm/swarm.go
+++ b/p2p/net/swarm/swarm.go
@@ -29,6 +29,7 @@ import (
        spdy "github.com/whyrusleeping/go-smux-spdystream"
        yamux "github.com/whyrusleeping/go-smux-yamux"
        mafilter "github.com/whyrusleeping/multiaddr-filter"
+       ws "github.com/whyrusleeping/ws-transport"
        context "golang.org/x/net/context"
 )
 
@@ -118,6 +119,7 @@ func NewSwarm(ctx context.Context, listenAddrs []ma.Multiaddr,
                transports: []transport.Transport{
                        transport.NewTCPTransport(),
                        transport.NewUtpTransport(),
+                       new(ws.WebsocketTransport),
                },
                bwc:         bwc,
                fdRateLimit: make(chan struct{}, concurrentFdDials),
```

You will also have to let libp2p's address parser know that your address type is okay, This is accomplished by adding the protocol path strings you want to support to the `SupportedTransportStrings` array in [`p2p/net/swarm/addr/addr.go`](https://github.com/ipfs/go-libp2p/blob/master/p2p/net/swarm/addr/addr.go#L17).

```diff
diff --git a/p2p/net/swarm/addr/addr.go b/p2p/net/swarm/addr/addr.go
index d9ba872..a505097 100644
--- a/p2p/net/swarm/addr/addr.go
+++ b/p2p/net/swarm/addr/addr.go
@@ -7,6 +7,8 @@ import (
        ma "github.com/jbenet/go-multiaddr"
        manet "github.com/jbenet/go-multiaddr-net"
        context "golang.org/x/net/context"
+
+       _ "github.com/whyrusleeping/ws-transport"
 )
 
 var log = logging.Logger("github.com/ipfs/go-libp2p/p2p/net/swarm/addr")
@@ -19,6 +21,8 @@ var SupportedTransportStrings = []string{
        "/ip6/tcp",
        "/ip4/udp/utp",
        "/ip6/udp/utp",
+       "/ip4/tcp/ws",
+       "/ip6/tcp/ws",
        // "/ip4/udp/udt", disabled because the lib doesnt work on arm
        // "/ip6/udp/udt", disabled because the lib doesnt work on arm
 }
```

And that's it! You now have a libp2p that can communicate over websockets! If you have any questions, please hit me up on IRC, send me an email, or file an issue on the [`ipfs/examples`](https://github.com/ipfs/examples) repo!

By [whyrusleeping](http://github.com/whyrusleeping)
