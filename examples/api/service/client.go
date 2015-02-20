package main

import (
	"fmt"
	"io"
	"os"

	core "github.com/jbenet/go-ipfs/core"
	corenet "github.com/jbenet/go-ipfs/core/corenet"
	peer "github.com/jbenet/go-ipfs/p2p/peer"
	fsrepo "github.com/jbenet/go-ipfs/repo/fsrepo"

	"code.google.com/p/go.net/context"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please give a peer ID as an argument")
		return
	}
	target, err := peer.IDB58Decode(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Basic ipfsnode setup
	r := fsrepo.At("~/.go-ipfs")
	if err := r.Open(); err != nil {
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

	con, err := corenet.Dial(nd, target, "/app/whyrusleeping")
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, con)
}
