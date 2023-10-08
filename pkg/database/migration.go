package database

func MakeMigration(models []interface{}) {
	for _, model := range models {
		err := DataBase.AutoMigrate(&model)
		if err != nil {
			panic(err)
		}
	}
}
