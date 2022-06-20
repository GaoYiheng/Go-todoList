package model

func migration() {
	DB.Set(`grom:table_options`, "charset=utf8mb4").AutoMigrate(&User{})
}
