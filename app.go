package main

import (
	"log"
	"rakamin-academy/database"
	c "rakamin-academy/src/controller"
	ur "rakamin-academy/src/repository/user_repo"
	us "rakamin-academy/src/service/user_service"
)

func main() {

	sql, err := database.MySQLConnect()

	if err != nil {
		log.Fatal("Failed to initialize MySQL connection")
	}

	supabase := database.SupabaseConnect()

	database.RunMigration(sql)

	ur := ur.NewUserRepository(sql, supabase)
	us := us.NewUserService(ur)

	c := c.NewController(us)

	c.RunServer()
}
