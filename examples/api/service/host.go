package main

import (
	"fmt"

	core "github.com/jbenet/go-ipfs/core"
	corenet "github.com/jbenet/go-ipfs/core/corenet"
	fsrepo "github.com/jbenet/go-ipfs/repo/fsrepo"

	"code.google.com/p/go.net/context"
)

func main() {
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

	list, err := corenet.Listen(nd, "/app/whyrusleeping")
	if err != nil {
		panic(err)
	}
	fmt.Printf("I am peer: %s\n", nd.Identity)

	for {
		con, err := list.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer con.Close()

		fmt.Fprintln(con, "Hello! This is whyrusleepings awesome ipfs service")
		fmt.Printf("Connection from: %s\n", con.Conn().RemotePeer())
	}
}
