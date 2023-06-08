package config

import (
	"fmt"
	"log"
	"strconv"

	"github.com/andrefsilveira1/LoadEnv"
)

// This variables will be exported and used by Mysql connection
var (
	stringConnection = ""
	Port             = 0
)

func Config() string {
	var erro error
	Port, erro = strconv.Atoi(LoadEnv.LoadEnv("DB_PORT"))
	if erro != nil {
		log.Fatal("Something goes wrong...", erro)
	}

	user := LoadEnv.LoadEnv("DB_USER")
	password := LoadEnv.LoadEnv("DB_PASS")
	name := LoadEnv.LoadEnv("DB_NAME")
	fmt.Println("Using port:", Port)
	fmt.Println("Using user:", user)
	fmt.Println("Using password:", password)
	fmt.Println("Using name:", name)
	stringConnection = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		name,
	)
	fmt.Println("string:", stringConnection)
	return stringConnection
}
