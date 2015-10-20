package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreunix"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"golang.org/x/net/context"
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
	r, err := fsrepo.Open("~/.ipfs")
	if err != nil {
		return nil, err
	}

	cfg := &core.BuildCfg{
		Repo:   r,
		Online: true,
	}

	return core.NewNode(context.Background(), cfg)
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

	read, err := coreunix.Cat(context.Background(), nd, keytofetch)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(CountChars(read))
}
