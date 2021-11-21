## Autobackup config
It's a small project for myself to backup config of alacritty & tmux

### How to use it?
1. You should edit `.env`, and follow this format
`backup_config_{what_your_service_is}={path}`(you need to change the string inside {} to what you want.)

2.
```
    go get
    go run *.go
```

3. Finished, it should backup all your path to the folder that names with your name.
