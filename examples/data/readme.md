## Dealing With Blocks
The `ipfs add` command will create a merkle dag out of the data in the files you
specify, it follows the unixfs data format when doing this. what this means, is
that your files are broken down into blocks, and then arranged in a tree-like
structure using 'link nodes' to tie them together. a given files 'hash' is
actually the hash of the root (uppermost) node in the dag. for a given dag, you
can easily view the sub-blocks under it with `ipfs ls`.

For example:
```
# ensure this file is larger than 256k
ipfs add alargefile 
ipfs ls thathash
```

The above command should print out something like:
```
ipfs@earth ~> ipfs ls qms2hjwx8qejwm4nmwu7ze6ndam2sfums3x6idwz5myzbn
qmv8ndh7ageh9b24zngaextmuhj7aiuw3scc8hkczvjkww 7866189 
qmuvjja4s4cgyqyppozttssquvgcv2n2v8mae3gnkrxmol 7866189 
qmrgjmlhlddhvxuieveuuwkeci4ygx8z7ujunikzpfzjuk 7866189 
qmrolalcquyo5vu5v8bvqmgjcpzow16wukq3s3vrll2tdk 7866189 
qmwk51jygpchgwr3srdnmhyerheqd22qw3vvyamb3emhuw 5244129
```

What this is showing is all of the immediate sub-blocks of your file, and the
size of them and their children on disk.

### What to do With Blocks?
If you feel adventurous you can get a lot of different information out of these
different blocks. you can use the sub-block hashes as input to `ipfs cat` to
see only the data in any given sub-tree (the data of that block and its
children). to see just the data of a given block and not its children, use
`ipfs block get`. but be careful, as `ipfs block get` on an intermediate block
will print out the raw binary data of its dag structure to your screen.

`ipfs block stat` will tell you the exact size of a given block (without its
children), and `ipfs refs` will tell you all the children of that block. 
similarly, `ipfs ls` or `ipfs object links` will show you all children and
their sizes. `ipfs refs` is a more suitable command for scripting something
to run on each child block of a given object.

### Blocks vs Objects
In ipfs, a block refers to a single unit of data, identified by its key (hash).
a block can be any sort of data, and does not necessarily have any sort of 
format associated with it. an object, on the other hand, refers to a block that
follows the merkledag protobuf data format. it can be parsed and manipulated
via the `ipfs object` command. any given hash may represent an object or a block.

### Creating a block from scratch
Creating your own blocks is easy! simply put your data in a file and run
`ipfs block put <yourfile>` on it, or you can pipe your filedata into 
`ipfs block put`, like so:

```
$ echo "This is some data" | ipfs block put
QmfQ5QAjvg4GtA3wg3adpnDJug8ktA1BxurVqBD8rtgVjM
$ ipfs block get QmfQ5QAjvg4GtA3wg3adpnDJug8ktA1BxurVqBD8rtgVjM
This is some data
```
Note: When making your own block data, you wont be able to read the data with
`ipfs cat`, this is because you are inputting raw data without the unixfs data
format. To read raw blocks use `ipfs block get` as shown in the example.

By [whyrusleeping](http://github.com/whyrusleeping)
