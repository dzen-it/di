package db

var Client ClientIface

type ClientIface interface {
	Keys(string) ([]string, error)
	Del(...string) error
	Set(string, string) error
}
