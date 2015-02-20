## Backups
Lets take a quick look at how ipfs can be used to keep basic backups.

Save your directory:
```
$ ipfs add -r ~/code/myproject
```

Note the hash:
```
$ echo $hash `date` >> backups
```


Or all at once:
```
$ echo `ipfs add -q -r ~/code/myproject | tail -n1` `date` >> backups
```
(Note: the `-q` makes the output only contain the hashes, piping through
`tail -n1` ensures only the hash of the top folder is output.)


View the backups live:
```
$ ipfs mount
$ ls /ipfs/$hash/

# can also

$ cd /ipfs/$hash/
$ ls
```

Through the fuse interface, youll be able to access your files exactly as
they were when you took the backup.

By [whyrusleeping](http://github.com/whyrusleeping)
