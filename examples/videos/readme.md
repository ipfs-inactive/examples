## Playing Videos
ipfs can be used to store and share videos, if someone gives you the hash of
a video, you can view it a couple different ways.

On the command line:
```
ipfs cat $vidhash | mplayer -vo xv -
```

Via local gateway:
```
mplayer http://localhost:4001/ipfs/$vidhash

# or open it up in a tab in chrome (or firefox)

chromium http://localhost:4001/ipfs/$vidhash
```
(Note: the gateway method works with most video players and browsers)

By [whyrusleeping](http://github.com/whyrusleeping)
