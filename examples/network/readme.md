## Playing With the Network
Ipfs is all about networking! included are a useful set of commands
to aid in observing that network.

See who youre directly connected to:
```
ipfs swarm peers
```

Get a listing of the entire network:
```
ipfs diag net
```

Manually connect to a specific peer:
```
ipfs swarm connect /ip4/104.236.176.52/tcp/4001/ipfs/qmsolnsgccfuzqjzradhn95w2crsfmzutddwp8hxahca9z
```

Search for a given peer on the network:
```
ipfs dht findpeer $peerid
```



