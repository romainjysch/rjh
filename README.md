# rjh cli tool

## how to start

### prerequisite

#### openweathermap

an openweathermap [api key](https://openweathermap.org/api)

```
export OWM_API_KEY=<OWM_API_KEY>
```

#### tasks

a csv file with the following structure

```
description,created,completed,deleted
first task,1753719339,0,0
second task,1753719339,0,0
```

#### config file structure

```
openweathermap:
  key: "<openweathermap_api_key>"

tasks:
  filepath: "</path/to/folder/rjh/internal/tasks/data/tasks.csv>"
```

#### config file const

```go
const PATH = "/path/to/folder/rjh/config/data/config.yml"
```

### binary

```
go build -o rjh
```

```
sudo mv rjh /usr/local/bin/rjh
```

### usage

```
rjh -h
```

## notes

- `network` command started as a cli discovery exercise;
- `weather` was the second command; adding forecasts could be a next step;
- `tasks` use a simple csv file as a mvp; future improvements could include a containerized db and gorm.

## license

mit
