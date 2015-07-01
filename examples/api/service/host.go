package main

import (

	"fmt"

	core "github.com/ipfs/go-ipfs/core"
	corenet "github.com/ipfs/go-ipfs/core/corenet"
	fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"
	peer "github.com/ipfs/go-ipfs/p2p/peer"

	"code.google.com/p/go.net/context"
)

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