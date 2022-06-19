package tool

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbs *gorm.DB

func initConn() *gorm.DB {
	conf := getConfig()
	log.Printf("url:%s\n", conf.GetString("mysql.url"))
	db, err := gorm.Open(mysql.Open(conf.GetString("mysql.url")),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				//使用单数表名 Error 1146: Table 'test.users' doesn't exists
				SingularTable: true,
			},
		})
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	return db
}

func OpenDB() *gorm.DB {
	if dbs == nil {
		log.Println("init con mysql")
		dbs = initConn()
	}
	return dbs
}
