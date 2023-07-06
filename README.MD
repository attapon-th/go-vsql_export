# Vertica Export to CSV Tool


## How to installation:

### Download Binary file
- [x] [Windows (x86_64)](https://github.com/attapon-th/vsql_export/releases/download/v0.1.0/vsql_export_Windows_x86_64.zip)
- [x] [Linux (x86_64)](https://github.com/attapon-th/vsql_export/releases/download/v0.1.0/vsql_export_Linux_x86_64.tar.gz)
- [x] [MacOS (x86_64)](https://github.com/attapon-th/vsql_export/releases/download/v0.1.0/vsql_export_Darwin_x86_64.tar.gz)
- [x] [MacOS (Apple Silicon: arm64)](https://github.com/attapon-th/vsql_export/releases/download/v0.1.0/vsql_export_Linux_arm64.tar.gz)

### with `go install`
```
go install github.com/attapon-th/vsql_export@main
```

## How to use:


### Export to CSV (stdout)

#### tool help

```shell
vsql_export -h
```

#### Use with DSN connection
```shell
export VERTICA_URI="vertica://user:password@localhost:5433/Vmart"
vsql_export csv "SELECT * FROM public.test"
```
