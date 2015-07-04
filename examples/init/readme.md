## init system integration
go-ipfs is relatively simple to integrate into your init system. For systemd, the best approach is to run
the daemon in a user session. Here is a sample service file:
```
[Unit]
Description=IPFS daemon
After=network.target

[Service]
ExecStart=/usr/bin/ipfs daemon

[Install]
WantedBy=multiuser.target
```
To run this in your user session, save it as `~/.config/systemd/user/ipfs.service` (creating directories 
as necessary). Once you run `ipfs init` to create your IPFS settings, you can control the daemon using the following
commands:
* `systemctl --user start ipfs` - start the daemon
* `systemctl --user stop ipfs` - stop the daemon
* `systemctl --user status ipfs` - get status of the daemon
* `systemctl --user enable ipfs` - enable starting the daemon at boot
* `systemctl --user disable ipfs` - disable starting the daemon at boot
