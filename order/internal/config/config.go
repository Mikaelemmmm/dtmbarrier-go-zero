package config

import "github.com/tal-tech/go-zero/rest"

type Config struct {
	rest.RestConf
	DataSource  string          //mysql
	BarrierDataSource  string          //barrier mysql
}
