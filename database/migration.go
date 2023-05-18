package database

import "log"

func RunMigration(DB *SQL) {
	if err := DB.AutoMigrate(); err != nil {
		log.Fatal("Error migrating database: ", err.Error())
	}

}
