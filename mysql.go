package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2/config"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
}

func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {
	mysqlConfig := &MysqlConfig{}
	err := config.Get(path...).Scan(mysqlConfig)
	if err != nil {
		return nil
	}
	return mysqlConfig
}

func (m *MysqlConfig) Open() (*gorm.DB, error) {
	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&loc=Local", m.User, m.Pwd,
		m.Host, m.Port, m.Database)
	return gorm.Open("mysql", connectStr)
}
