package main

import (
	"github.com/susilo001/Wallet-System/user/config"
	"github.com/susilo001/Wallet-System/user/route"
)

func main() {
	config.InitDatabases()
	route.InitRoutes()
}
