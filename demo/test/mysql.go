package test

import (
	"demo/com/cs/learn/model"
	"demo/com/cs/learn/tool"
	"log"

	"gorm.io/gorm"
)

var users []model.User

var user model.User

func create() {
	db := tool.OpenDB()
	user = model.User{
		ID:   25,
		Name: "舞蹈",
		Age:  20,
		Sex:  false,
	}
	//INSERT INTO `info` (`name`,`age`,`sex`,`id`) VALUES ('舞蹈',20,false,25)
	db.Debug().Create(&user)

	users = []model.User{
		// composite literal uses unkeyed fieldscomposites
		//省略key后顺序要跟结构体一致 ,后期结构体添加新属性会报错,你懂得
		0: model.User{27, "结算", 33, false},
		1: model.User{28, "来到", 17, true},
	}
	// res := db.Debug().CreateInBatches(&users, 1)
	// log.Println(res, res.Error)

}

func query() {
	db := tool.OpenDB()

	db.Debug().First(&user) //第一条记录

	// [{12 第三 11 false} {15 推广 19 false}]
	//db.Debug().Where("age IN (?)", []int{11, 19}).Find(&users)

	//[{15 推广 19 0} {21 发广告 21 0}]
	//db.Debug().Where("name like ?", "%广%").Find(&users)

	//等价select name,age from info
	//db.Debug().Select("name,age").Find(&users)

	// [{21 发广告 0 0}] 原生sql raw sacn
	//db.Debug().Raw("select id,name from info where sex=? and age>?", 1, 20).Scan(&users)

	//[{113 的人 0 0} {51 若非 0 0}]
	//result()

	lp()
}

func result(db *gorm.DB, users []model.User) {
	//select * 返回结果集  [{0  0 0} {0  0 0}]
	//rows, _ := db.Debug().Raw("select * from info where sex=?", 0).Rows()

	rows, _ := db.Debug().Raw("select id,name from info where sex=?", 0).Rows()
	for rows.Next() {
		var res model.User
		//查询字段要跟scan字段一致
		rows.Scan(&res.ID, &res.Name)
		users = append(users, res)
	}
	lp()
}

func update() {
	db := tool.OpenDB()

	//UPDATE `info` SET `name` = 'cs'  WHERE id=5 and age >10
	err := db.Debug().Model(&model.User{}).Where("id = ? and age >? ", 5, 10).Update("name", "cs").Error

	uu := model.User{ID: 12}
	//UPDATE `info` SET `name` = 'cs'  WHERE `id` = 12
	//更新多个值用结构体 SET name=cs,age=15
	err = db.Debug().Model(uu /*&uu 都可以*/).Update("name", "cs" /*model.User{"name":"cs","age",15}*/).Error

	if err != nil {
		log.Fatalln("err:", err)
	}
	lp()
}

func delete() {
	db := tool.OpenDB()

	err := db.Debug().Where("id=?", 25).Delete(&user).Error
	//匹配主键进行删除
	/*
		DELETE FROM `info` WHERE `info`.`id` = 25
		DELETE FROM `info` WHERE `info`.`id` IN (25, 30)
	*/
	err = db.Debug().Delete(users, 25).Error
	err = db.Debug().Delete(&users, []int{25, 30}).Error
	//	DELETE FROM `info` WHERE name like '%cs%'
	err = db.Debug().Delete(user, "name LIKE ?", "%cs%").Error

	/*model包含DeletedAt
	GORM会将DeletedAt的值设置为当前时间，并且使用常规查询方法无法再查找数据
	Unscoped()永久删除[Unscoped().Delete(&users,15)]或查询软删除[Unscoped().First(&user, 11)]匹配的记录
	*/
	//db.Debug().Unscoped().Delete(&users,15).Error

	if err != nil {
		log.Fatalln("err:", err)
	}
}

//报错回滚
func trans() {
	db := tool.OpenDB()
	db.Begin()

	err := db.Exec("UPDATE info SET age=17 WHERE id=15").Error
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := recover(); err != nil {
			db.Rollback()
			if val, ok := err.(error); ok {
				log.Fatalln(val)
			}
		} else {
			err := db.Commit().Error
			if err != nil {
				db.Rollback()
				log.Fatalln(err)
			}
		}
	}()

}

//打印结果
func lp() {
	if users != nil {
		for k, v := range users {
			log.Printf("%d,name:%s", k, v.Name)
		}
		log.Println("user[]:", users)
	}
	if user != (model.User{}) {
		log.Println("user:", user)
	}
}
