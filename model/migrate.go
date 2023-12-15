package model

func migration() {
	//自动迁移模型
	DB.Set("gorm:table_options", "charset=utf8mb4")
	DB.AutoMigrate(&User{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("failed to migrate")
	// }
	DB.AutoMigrate(&Task{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic("failed to migrate")
	// }
	//task关联到User                 外键名称                           task.uid外键字段名称     主表列名
	//ALERT (`task`) ADD CONSTRAINT task_uid_User_id_foreign FOREIGN KEY (`uid`) REFERENCES User(id) ON DELETE CASCADE ON UPDATE CASCADE;
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")

}
