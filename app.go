package main

import (
	"log"
	"rakamin-academy/database"
	c "rakamin-academy/src/controller"
	sr "rakamin-academy/src/repository/seller_repo"
	ss "rakamin-academy/src/service/seller_service"
)

func main() {

	sql, err := database.MySQLConnect()

	if err != nil {
		log.Fatal("Failed to initialize MySQL connection")
	}

	supabase := database.SupabaseConnect()

	database.RunMigration(sql)

	sr := sr.NewSellerRepository(sql, supabase)
	ss := ss.NewSellerService(sr)

	c := c.NewController(ss)

	c.RunServer()
}
