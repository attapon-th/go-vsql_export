# Vertica Export to CSV Tool


## How to installation:

### with `go install`
```
go install github.com/attapon-th/vsql_export@latest
```

## How to use:


### Export to CSV (stdout)

#### tool help

```shell
vsql_export -h
```

```shell
VERTICA_DSN="vertica://user:password@localhost:5433/Vmart"
vsql_export -d $VERTICA_DSN "SELECT * FROM public.test"
```

