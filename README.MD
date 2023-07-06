# Vertica Export to CSV Tool


### Download Binary file

> download binary file and extract compressed file to use

- [x] [Windows (x86_64)](https://github.com/attapon-th/vsql_export/releases/download/v0.1.0/vsql_export_Windows_x86_64.zip)
- [x] [Linux (x86_64)](https://github.com/attapon-th/vsql_export/releases/download/v0.1.0/vsql_export_Linux_x86_64.tar.gz)
- [x] [MacOS (x86_64)](https://github.com/attapon-th/vsql_export/releases/download/v0.1.0/vsql_export_Darwin_x86_64.tar.gz)
- [x] [MacOS (Apple Silicon: arm64)](https://github.com/attapon-th/vsql_export/releases/download/v0.1.0/vsql_export_Linux_arm64.tar.gz)



### with `go install`
```
go install github.com/attapon-th/vsql_export@v0.1.0
```

## How to use:


### Shell
```shell
# set connection with DSN URI
export VERTICA_URI="vertica://user:password@localhost:5433/Vmart"

# output: stdout
vsql_export csv "SELECT * FROM public.test"

# output: csv file
vsql_export csv -o test.csv "SELECT * FROM public.test" 

# output: csv file and commpress with gzip
vsql_export csv -o test.csv.gz -z gzip "SELECT * FROM public.test"
```

#### help

```shell
vsql_export -h
```

