# IPFS Examples

This is a simple viewer for ipfs examples. It is inspired by @mbostock's
gorgeous bl.ocks.org. It renders markdown for easy reading. You can check out
how it works, or contribute changes at:
[https://github.com/protocol/ipfs-app-examples/tree/master/ipfs-examples](https://github.com/protocol/ipfs-app-examples/tree/master/ipfs-examples)

Try it out!

## Writing your own examples

It is trivial to write out your own ipfs examples to use with this tool.

- [Step 1. Install IPFS](#step-1-install-ipfs)
- [Step 2. Bundle the Example](#step-2-bundle-the-example)
- [Setp 3. Publish!](#step-3-publish)

### Step 1. Install IPFS

In order to publish examples to the ipfs gateways, you need ipfs. Find out
how to install it here: [http://ipfs.io](http://ipfs.io). To check if you
have it installed, enter:

```go
> ipfs version
ipfs version 0.1.7
```

To view them on the web, you need to be running the daemon:

```go
> ipfs daemon
Initializing daemon...
API server listening on /ip4/127.0.0.1/tcp/5001
Gateway server listening on /ip4/127.0.0.1/tcp/8080
```

This will make it so you can view your examples in your browser. And, if you
are connected to the internet, so others can view them through the
ipfs gateways.

### Step 2. Bundle the Example

This viewer follows a very simple format. It's just a directory! Put all your
files into a directory. There are a couple of specially-named files that this
viewer will treat differently:

- `readme.md` this is the bulk of your example, in markdown.
- `thumbnail.png` this is the thumbnail image. it can also be thumbnail.jpg,
  or thumbnail.gif. the precedence is: `png -> jpg -> gif`

That's it! Endeavor to add any other files you need into the directory,
instead of making any external links. The idea is that an example should be
entirely self-contained.

At any time you can view how your example will be seen, following the same
steps in publish.

### Step 3. Publish!

Once you're ready to publish (or just view) your example, enter `ipfs add -r <your directory>`. For us:

```
> ipfs add -r my-example
added QmWWQSuPMS6aXCbZKpEjPHPUZN2NjB3YrhJTHsV4X3vb2t my-example/readme.md
added QmT4AeWE9Q9EaoyLJiqaZuYQ8mJeq4ZBncjjFH9dQ9uDVA my-example/thumbnail.jpg
added QmT9qk3CRYbFDWpDFYeAv8T8H1gnongwKhh5J68NLkLir6 my-example
```

Now, take the last hash you get -- for the directory -- and make a url
like this:

```
http://localhost:8080/ipfs/<hash-of-the-viewer>/example#/ipfs/<hash-you-got>
```

For example:

```
http://localhost:8080/ipfs/QmPDgUhqWE4WqRqAHHtjUTkTjWRiSKGtAHtf8YWcqUiUvA/example#/ipfs/QmT9qk3CRYbFDWpDFYeAv8T8H1gnongwKhh5J68NLkLir6
```

You can also view it on the public ipfs gateway, if you're connected:

```
http://ipfs.io/ipfs/QmPDgUhqWE4WqRqAHHtjUTkTjWRiSKGtAHtf8YWcqUiUvA/example#/ipfs/QmT9qk3CRYbFDWpDFYeAv8T8H1gnongwKhh5J68NLkLir6
```

#### Bonus: publish with a makefile

I like publishing my examples with a simple `Makefile`:

```
# Makefile that publishes this example

viewer = "QmPDgUhqWE4WqRqAHHtjUTkTjWRiSKGtAHtf8YWcqUiUvA"
local = "http://localhost:8080/ipfs/"
gway = "http://ipfs.io/ipfs/"

publish: $(shell find . )
  @hash=$(shell ipfs add -r -q . | tail -n1); \
    echo $(local)/$(viewer)/example#/ipfs/$$hash; \
    echo $(gway)/$(viewer)/example#/ipfs/$$hash

# we need ; and escaped newlines to capture the variable
```

Now you can just:

```sh
> make publish
http://localhost:8080/ipfs/QmPDgUhqWE4WqRqAHHtjUTkTjWRiSKGtAHtf8YWcqUiUvA/example#/ipfs/QmT9qk3CRYbFDWpDFYeAv8T8H1gnongwKhh5J68NLkLir6
http://ipfs.io/ipfs/QmPDgUhqWE4WqRqAHHtjUTkTjWRiSKGtAHtf8YWcqUiUvA/example#/ipfs/QmT9qk3CRYbFDWpDFYeAv8T8H1gnongwKhh5J68NLkLir6
```

:)

#### Important Note

Note that your example wont stay up on the ipfs web by itself. The ipfs
gateways cache things temporarily, but not forever. If you want to
make it stick around permanently, read up about it here (todo: add link).
