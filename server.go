package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/thinhhuynh/gin-base/bootstrap"
)

func main() {
	_ = godotenv.Load()
	bootstrap.RootApp.Execute()
}
