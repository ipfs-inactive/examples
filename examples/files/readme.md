## The ipfs files API

Ipfs has an exciting new (ish) API called the 'files' API. It allows you to
have a similar interface to traditional filesystems (mutable files and
directories) while being backed by ipfs. In this post, I will give a few useful
examples on what is possible.

Within each ipfs node is a virtual 'root' directory which you can access and
manipulate through the `ipfs files` subcommand.
Let's start out by making a couple directories and a few files.

```bash
$ ipfs files mkdir -p /foo/bar
$ ipfs files mkdir /baz
```

`mkdir` works as you would expect. To write a file:

```bash
$ echo "hello ipfs!" | ipfs files write --create /foo/hello
$ echo "merkledag" | ipfs files write -e /foo/bar/blob
```

Note: `-e` is an alias for `--create`.

You can list out directories:

```bash
$ ipfs files ls /
baz
foo
$ ipfs files ls /foo
bar
hello
```

You can copy files around:

```bash
$ ipfs files cp /foo/hello /baz/hi
$ ipfs files read /baz/hi
hello ipfs!
```

Copies here have a really cool feature in that they don't take up any extra
disk space (aside from a small bit of metadata). This means that you can have
multiple copies of *huge* files in different places of the files API without
having to worry about wasting space.

`cp` is also able to copy files from general ipfs hashes too:

```bash
$ echo "multi multi multi" | ipfs add -q
QmWChe9uQKJgZdjwuTjY1qcUebVDu7Jb5LpSZ2XXyvzdfh
$ ipfs files cp /ipfs/QmWChe9uQKJgZdjwuTjY1qcUebVDu7Jb5LpSZ2XXyvzdfh /foo/multi
$ ipfs files ls /foo
bar
hello
multi
```

This functionality of `cp` provides some subtle advantages. Data copied in this
way is not downloaded immediately. The cool thing here is that you could copy a
massive 100TB dataset into your namespace and operate on it as if it was local
without having to have the disk space for it.

Now, as you perform all these operations, the structure is synchronized into
ipfs for you. To get an ipfs hash for a given file or directory, use `stat`:

```bash
$ ipfs files stat /foo
QmPdZg2VXP4XcUPziqMMB4msjJeqcCVWdtA9251CRkwARC
Size: 0
CumulativeSize: 243
ChildBlocks: 3
Type: directory
```

You can then take that hash and investigate it to see that it matches what's in
the files API:

```bash
$ ipfs ls QmPdZg2VXP4XcUPziqMMB4msjJeqcCVWdtA9251CRkwARC
QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn 4  bar/
QmUe6akh4xc2H1yPgo7S5jQhvjGTWY6no9NgHSajDQPusR 70 hello
QmWChe9uQKJgZdjwuTjY1qcUebVDu7Jb5LpSZ2XXyvzdfh 26 multi
```

There are a few other commands in the files API command set, and you can check them out by running `ipfs files --help`. Hopefully you now have a good idea of what and how this API works and some of the things you can accomplish with it.

By [whyrusleeping](http://github.com/whyrusleeping)
