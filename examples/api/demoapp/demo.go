package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"code.google.com/p/go.net/context"
	"github.com/jbenet/go-ipfs/core"
	"github.com/jbenet/go-ipfs/core/coreunix"
	"github.com/jbenet/go-ipfs/repo/fsrepo"
)

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

func SetupIpfs() (*core.IpfsNode, error) {
	// Assume the user has run 'ipfs init'
	r := fsrepo.At("~/.ipfs")
	if err := r.Open(); err != nil {
		return nil, err
	}

	builder := core.NewNodeBuilder().Online().SetRepo(r)
	return builder.Build(context.Background())
}

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
