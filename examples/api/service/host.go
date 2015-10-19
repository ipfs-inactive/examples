package main

import (
    "fmt"

    core "github.com/ipfs/go-ipfs/core"
    corenet "github.com/ipfs/go-ipfs/core/corenet"
    fsrepo "github.com/ipfs/go-ipfs/repo/fsrepo"

    "golang.org/x/net/context"
)

func main() {
    // Basic ipfsnode setup
    r, err := fsrepo.Open("~/.ipfs")
    if err != nil {
        panic(err)
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    cfg := new(core.BuildCfg)
    cfg.Repo = r
    cfg.Online = true

    nd, err := core.NewNode(ctx, cfg)

    if err != nil {
        panic(err)
    }

    list, err := corenet.Listen(nd, "/app/whyrusleeping")
    if err != nil {
        panic(err)
    }
    fmt.Printf("I am peer: %s\n", nd.Identity.Pretty())

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
