package model

func migration() {
	_ = DB.AutoMigrate(&User{})
}
