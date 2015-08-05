# WebApps on IPFS

This is a short guide on how to start making your own web applications on top of
ipfs.  

Ipfs provides an http api to interact with the daemon through, but by
default it has a few security features enabled that make it difficult for people
to use the api in their own apps. There are two different routes you could take
to use the API, first, is to run your own webserver, and make http requests to
localhost:5001. This is in my opinion the easier route for development. For this
to work, you'll need to disable cross origin check by running the daemon with
the environment variable `API_ORIGIN` set to `*`:

```
$ export API_ORIGIN="*"
$ ipfs daemon
```

Now all of the requests you make to the api will work just fine.

The second route is to publish your app to ipfs (via `ipfs add -r`) and load
it through the http server at `localhost:5001/ipfs/<hash of your app>`. This
does not require you to change the `API_ORIGIN` variable (as you'll be loading
the app from the correct origin) but you will need to tell the daemon that
you want it to allow you to load different apps through the server on the API port.
By default you can only load the ipfs webui through the API port, this is a
security measure to prevent random web pages from making api calls on your daemon.
To disable this for testing your app, run the daemon like:

```
$ ipfs daemon --unrestricted-api
```

So now that youre able to load your app, you probably want to start making
those api calls I mentioned (you *are* writing an ipfs app, after all). We
recommend everyone use the 
[browserified node-ipfs-api](https://github.com/ipfs/node-ipfs-api).

By
[whyrusleeping](https://github.com/whyrusleeping)
