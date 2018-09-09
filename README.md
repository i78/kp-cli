# Koalapaste Commandline Interface
This commandline tool allows uploading new pastes from the shell, i.E. for usage in scripts or other automated processes. By default, the stdin is used as source.



## Examples
Send the contents of your zsh config to paste.m9d.de with a default name
```
cat ~/.zshrc | ./koalapaste
```

Upload your docker running containers to your local koalapaste instance
```
docker ps | ./koalapaste --title "docker ps x" --stage http://paste.localhost:4200
```