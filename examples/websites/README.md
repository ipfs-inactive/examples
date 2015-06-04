#ipfs for websites
### A short guide to hosting your site on ipfs

Adding your static website to ipfs is quite simple! Simply turn on your daemon:
```bash
$ ipfs daemon
```

And add the directory containing your website:
```bash
$ ls mysite
index.html
$ ipfs add -r mysite
added QmbRyftGAtNjXs7hWZUTLyp27qrpgRUXKEwXCVPBueZGek mysite/index.html
added QmYAMp9ptyn3Bv8ijoYoYAAQHFCfpZ1NKRNedkEWrR9DPH mysite/
```

The very last hash next to the folder name is the one you want, lets call it
`$SITE_HASH` for now.  

Now, you can test it out locally by opening `http://localhost:8080/ipfs/$SITE_HASH`
in your web browser! Next, to view it coming from another ipfs node, you can try
`http://gateway.ipfs.io/ipfs/$SITE_HASH`. Cool, right?  But those hashes are
kinda ugly. Lets look at some ways to get rid of them.

First, you can do a simple DNS TXT record, containing `dnslink=/ipfs/$SITE_HASH`.
Once that record propogates, you should be able to view your site at
`http://localhost:8080/ipns/your.domain`. Now thats quite a bit cleaner.

Next, you might be asking "well what if i want to change my website, DNS is slow!"
Well let me tell you about this little thing called Ipns. Ipns is the Interplanetary
Naming System, you might have noticed the above link has `/ipns/` instead of `/ipfs/`.
Ipns is used for mutable content in the ipfs network, it's relatively easy to 
use, and will allow you to change your website without updating the dns record
every time! So how do you use it?

After adding your webpage, simply do:
```bash
$ ipfs name publish $SITE_HASH
Published to <your peer id>: /ipfs/$SITE_HASH
```

Now, you can test that it worked by viewing: `http://localhost:8080/ipns/<your peer id>`.
And also try the same link on the public gateway. Once youre convinced that works,
lets again hide the hash. Change your DNS TXT record to `dnslink=/ipns/<your peer id>`,
wait for that record to propogate, and then try accessing `http://localhost:8080/ipns/your.domain`.

Happy Hacking!

By:
[Whyrusleeping](https://github.com/whyrusleeping)


