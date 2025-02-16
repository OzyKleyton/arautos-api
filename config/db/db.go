package db

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(uri string) (*gorm.DB, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	config := &gorm.Config{}

	switch u.Scheme {
	case "postgres":
		return gorm.Open(getPostgres(u), config)
	default:
		return nil, errors.New("invalid database engine")
	}
}

func getPostgres(url *url.URL) gorm.Dialector {
	template := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo"
	host, port, _ := net.SplitHostPort(url.Host)
	user := url.User.Username()
	pass, _ := url.User.Password()
	dbname := strings.Replace(url.Path, "/", "", 1)
	return postgres.Open(fmt.Sprintf(template, host, user, pass, dbname, port))
}