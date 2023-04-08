package main

import (
	"fmt"

	"github.com/MelihEmreGuler/go-sqlite-user-api/pkg/api"
	"github.com/MelihEmreGuler/go-sqlite-user-api/pkg/database"
	"github.com/MelihEmreGuler/go-sqlite-user-api/pkg/ui"
)

func main() {
	fmt.Println("...main")

	database.Database()
	api.Api()
	ui.Ui()
}
