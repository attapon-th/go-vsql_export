package pkg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	verticago "github.com/vertica/vertica-sql-go"
)

// VerticaCTX return VerticaContext
func VerticaCTX(ctx context.Context) verticago.VerticaContext {
	return verticago.NewVerticaContext(ctx)
}

// ConnectVerticaWithDSN connect to Vertica with DSN URL
func ConnectVerticaWithDSN(dsn string) (db *sql.DB, err error) {
	err = errors.New("Error: Can't connect to Vertica")
	db, err = sql.Open("vertica", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return
}

// ConnectVertica connect to Vertica with user, password, host, dbname, port
func ConnectVertica(user, password, host, dbname, port string, optional map[string]string) (db *sql.DB, err error) {
	var rawQuery = url.Values{}
	if optional != nil {
		for k, v := range optional {
			rawQuery.Add(k, v)
		}
	}
	err = errors.New("Error: Can't connect to Vertica")
	u := url.URL{
		Scheme:   "vertica",
		User:     url.UserPassword(user, password),
		Host:     fmt.Sprintf("%s:%s", host, port),
		Path:     dbname,
		RawQuery: rawQuery.Encode(),
	}
	return ConnectVerticaWithDSN(u.String())
}
