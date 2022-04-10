# Minecraft reverse copy tool

This simple tool had been created for my needs to archive minecraft server world save and send it to my WebDAV server

## Config example
```toml
world_path = "/opt/minecraft/world"

[webdav]
webdav_host = "http://localhost"
webdav_save_path = "sample/backup_dir/"
use_auth = true

[webdav.auth]
username = "username"
password = "password"
```

## How to build and run
Just run:
```bash
go build github.com/k0tletka/minecraft-reserve-copy
```
And then:
```bash
./minecraft-reserve-copy -c /path/to/configuration.toml
```
