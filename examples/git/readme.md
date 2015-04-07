## Git, even more distributed

Have you ever said to yourself: "Man, my git server isn't distributed enough" or
"I wish I had an easy way to serve a static git repository worldwide". Well wish
no more, I have the solution for you!

In this article, I will be discussing how to serve a git repository through the
ipfs network. The end result will be a `git clone`able url served through ipfs!

To start, select a git repo you want to host, and do a bare clone of it:
```
$ git clone --bare git@myhost.io/myrepo
```

For those who aren't super git savvy, a bare repo means that it doesn't have
a working tree, and can be used as a server. They have a slightly different
format than your normal git repo.

Now, to get it ready to be cloned, you need to do the following:
```
$ cd myrepo
$ git update-server-info
```

Optionally, you can unpack all of gits objects:
```
$ cp objects/pack/*.pack .
$ git unpack-objects < ./*.pack
$ rm ./*.pack
```

Doing this breaks up gits large packfile into all of its individual objects.
This will allow ipfs to deduplicate objects if you add multiple versions of
this git repository.

Once youve done that, that repo is ready to be served. All thats left to do, is
to add it to ipfs:
```
$ pwd
/code/myrepo
$ ipfs add -r .
...
...
...
added QmX679gmfyaRkKMvPA4WGNWXj9PtpvKWGPgtXaF18etC95 .
```

Now, all thats left is to try cloning it:
```
$ cd /tmp
$ git clone http://localhost:8080/ipfs/QmX679gmfyaRkKMvPA4WGNWXj9PtpvKWGPgtXaF18etC95 myrepo
```

Note: make sure to change out that hash for yours.

Now, you may be asking "well what good is a git repo that I can't change anything on?"
Well let me tell you an awesome usecase! I tend to program in a language called Go,
for those who don't know go uses version control paths for its imports, i.e:
```go
import (
	"github.com/whyrusleeping/mycoollibrary"
)
```

This is a really nice feature, and solves a lot of problems, but often times, I run into
the issue where im using someones library, and they change the API, and it breaks my code.
Using what we've done above, you could clone the library, and add it into ipfs, so your import
paths will now look like:
```go
import (
	mylib "gateway.ipfs.io/ipfs/QmX679gmfyaRkKMvPA4WGNWXj9PtpvKWGPgtXaF18etC95"
)
```

And you will be guaranteed to have the same code every time!

By [whyrusleeping](http://github.com/whyrusleeping)
