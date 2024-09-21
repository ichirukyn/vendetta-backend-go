package utils

import (
	"net/url"
)

func NewURLConnectionString(proto, host, path, database, user, password string) string {
	const cDataBaseURLParameter = "database"

	var v = make(url.Values)
	if len(database) > 0 {
		v.Set(cDataBaseURLParameter, database)
	}

	var u = url.URL{
		Scheme:   proto,
		Host:     host,
		Path:     path,
		User:     url.UserPassword(user, password),
		RawQuery: v.Encode(),
	}

	return u.String()
}
