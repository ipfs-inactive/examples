## init system integration

go-ipfs is relatively simple to integrate into your init system. 

Below are instructions for various systems:

- `systemd`
- `initd`

### `systemd`

For `systemd`, the best approach is to run the daemon in a user session. Here is a sample service file:

```systemd
[Unit]
Description=IPFS daemon
After=network.target

[Service]
ExecStart=/usr/bin/ipfs daemon

[Install]
WantedBy=multiuser.target
```

To run this in your user session, save it as `~/.config/systemd/user/ipfs.service` (creating directories as necessary). Once you run `ipfs init` to create your IPFS settings, you can control the daemon using the following commands:

* `systemctl --user start ipfs` - start the daemon
* `systemctl --user stop ipfs` - stop the daemon
* `systemctl --user status ipfs` - get status of the daemon
* `systemctl --user enable ipfs` - enable starting the daemon at boot
* `systemctl --user disable ipfs` - disable starting the daemon at boot

### `initd`

- Here is a full-featured sample service file: https://github.com/dylanPowers/ipfs-linux-service/blob/master/init.d/ipfs
- And below is a very basic sample service file. **Note the username jbenet**.

```
cat /etc/init/ipfs.conf
```
```initd
description "ipfs daemon"

start on (local-filesystems and net-device-up IFACE!=lo)
stop on runlevel [!2345]
limit nofile 524288 1048576
limit nproc 524288 1048576
chdir /home/jbenet
exec start-stop-daemon --start --chuid jbenet --exec /home/jbenet/go/bin/ipfs daemon
respawn
```

Install this file

```sh
ipfs cat /ipfs/QmbYCwVeA23vz6mzAiVQhJNa2JSiRH4ebef1v2e5EkDEZS/ipfs.conf >/etc/init/ipfs.conf
```

And edit it to replace all occurrences of `jbenet` with whatever user you want it to run as:

```sh
sed -i s/jbenet/<chosen-username>/ /etc/init/ipfs.conf
```

Once you run `ipfs init` to create your IPFS settings, you can control the daemon using the `init.d` commands:

```sh
sudo service ipfs start
sudo service ipfs stop
sudo service ipfs restart
...
```
