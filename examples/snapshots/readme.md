## Snapshots
Lets take a quick look at how ipfs can be used to take basic snapshots.

Save your directory:
```
$ ipfs add -r ~/code/myproject
```

Note the hash:
```
$ echo $hash `date` >> snapshots
```


Or all at once:
```
$ echo `ipfs add -q -r ~/code/myproject | tail -n1` `date` >> snapshots
```
(Note: the `-q` makes the output only contain the hashes, piping through
`tail -n1` ensures only the hash of the top folder is output.)

Make sure to have the placeholders for the mount points:
```
$ sudo mkdir /ipfs /ipns
$ sudo chown `whoami` /ipfs /ipns
```

View the snapshots live:
```
$ ipfs mount
$ ls /ipfs/$hash/

# can also

$ cd /ipfs/$hash/
$ ls
```

Through the fuse interface, youll be able to access your files exactly as
they were when you took the snapshot.

By [whyrusleeping](http://github.com/whyrusleeping)
