## Some basics
To get started, we need to make sure ipfs has been initialized,
if you havent done this yet:
```
$ ipfs init
```

Now, run the daemon:
```
$ ipfs daemon
```

Now that we have the daemon up, lets have some fun.

Basic work with files in ipfs:
```
$ echo "welcome to ipfs!" > hello
$ ipfs add hello
```

That should have printed out something along the lines of:
```
added qmxzzpcazv6tw1tvicf9poare9kkb1fwmzbvamytdwvshe hello
```

That means that the file was successfully added into the ipfs datastore,
and may be accessed through ipfs now.

To check, try:
```
$ ipfs cat qmxzzpcazv6tw1tvicf9poare9kkb1fwmzbvamytdwvshe
```
(Note: if your files hash was different in the first step, use your
hash instead of mine)


If all went well, you should see the text from your file printed out to you!

Now, lets try out a directory.
```
$ mkdir foo
$ mkdir foo/bar
$ echo "hello" > foo/bar/baz
$ echo "hello" > foo/baz
$ ipfs add -r foo
```

View all the things!
```
$ ipfs ls <hash foo>
$ ipfs ls <hash foo>/bar
$ ipfs cat <hash foo>/bar/baz
$ ipfs cat <hash foo>/baz
```

So, that lets you explore the ipfs filesystem pretty much in the same way you
would explore a standard unix filesystem (like ext4 or zfs). Now, lets do a few
slightly more interesting things. `ipfs refs` will allow you to view blocks that
are associated with a given hash. Lets try it out with the `foo` directory 
structure we just made.

```
$ ipfs refs <hash foo>
$ ipfs refs -r <hash foo>
```
Note that the `-r` option output not just the direct children of foo, but all
of its decendants all the way down. `ipfs refs` has a few other really
interesting options, to learn more about them, run `ipfs refs --help`.


As you have seen `ipfs cat` is a great command to quickly retrieve and view
files, but if the file you are requesting contains binary data (such as an image
or movie) `ipfs get` might be more appropriate:
```
$ ipfs get -o cats.png <hashofcatpic>
```

This will create a file named 'cats.png' that contains the data from the
given hash.

