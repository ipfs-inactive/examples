## The Inter-Planetary Naming System

ipns is a way to add a small amount of mutability to the permanent immutability
that is ipfs. It allows you to store a reference to an ipfs hash under the
namespace of your peerID ( hash of your public key ). The commands to set it up
are quite simple.

First, you'll need some content to publish:

```
$ echo "Let's have some mutable fun!" | ipfs add
```

note the hash that was output, and publish that hash out to the network:

```
$ ipfs name publish <that hash>
Published name <your peer ID> to <that hash>
```

Now, to test that it worked, you could try a couple of different things:

```
$ ipfs name resolve <your peer ID>
<that hash>
```

If you ran that on the same machine, it should return instantly, as you have
cached the entry locally; give it a shot on another computer running ipfs.

Another thing to try is viewing it on a gateway:

```
http://gateway.ipfs.io/ipns/<your peer ID>
```

So, now comes the fun part: Lets change things.

```
$ echo "Look! Things have changed!" | ipfs add
```

Next, take the hash from there and...
```
$ ipfs name publish <the new hash>
Published name <your peer ID> to <the new hash>
```

Viola! Now, if you resolve that entry again, you'll see your new object.

Congratulations! You just successfully published and resolved an ipns entry!
Now, there are a few things to note; first, right now, you can only
publish a single entry per ipfs node. This will change fairly soon. Second,
updating an ipns entry can "break links" because anything referencing an ipns
entry might no longer point to the content it expected. There is no way around 
this ( you know, mutability ), therefore, ipns links should be used carefully if
you want to ensure permanence. In the future, we may have ipns entries work as
a git commit chain, with each successive entry pointing back in time to other
values.

By [whyrusleeping](http://github.com/whyrusleeping)
