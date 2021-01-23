# polybar-weather

Adds simple weather to polybar. You need an API key from openweathermap.

## Configuration

Pull the code and edit the `wthrf.go` file in the `cmd/wthrf` folder. Place the filepath to your config file. Edit the default config file and fill in the variables. Build the bin by doing

```
go build cmd/wthrf/wthrf.go
```

Move it and the config file wherever you want as long as the filepath of the config file is the same as you're edit. Then add it to polybar

```
[module/weather]
type = custom/script
exec = /path/to/bin
exec-if = ping openweathermap.org -c 1
interval = 300
```
