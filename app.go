package main

import (
	"log"
	"rakamin-academy/database"
	c "rakamin-academy/src/controller"
	cr "rakamin-academy/src/repository/category_repo"
	mr "rakamin-academy/src/repository/merchant_repo"
	ur "rakamin-academy/src/repository/user_repo"
	cs "rakamin-academy/src/service/category_service"
	us "rakamin-academy/src/service/user_service"
)

func main() {

	sql, err := database.MySQLConnect()

	if err != nil {
		log.Fatal("Failed to initialize MySQL connection")
	}

	supabase := database.SupabaseConnect()

	database.RunMigration(sql)

	mr := mr.NewMerchantRepository(sql, supabase)

	ur := ur.NewUserRepository(sql, supabase)
	us := us.NewUserService(ur, mr)

	cr := cr.NewCategoryRepository(sql, supabase)
	cs := cs.NewCategoryService(cr, ur)

	c := c.NewController(
		us,
		cs,
	)

	c.RunServer()
}
