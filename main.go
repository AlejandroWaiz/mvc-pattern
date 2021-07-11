package main

import "os"

func main() {

}

func defineEnvVar() {

	os.Setenv("db_user", "db_user")
	os.Setenv("db_passwd", "db_passwd")
	os.Setenv("db_addr", "db_addr")
	os.Setenv("db_db", "db_db")
	os.Setenv("muxport", "muxport")

}
