package main

import (
	"franigen-example/api/app"
	"os"
)

func main() {

	os.Setenv("ENV", "develop")
	os.Setenv("PORT", "5000")
	os.Setenv("Local", "true")
	os.Setenv("JWT_SECRET", "]d^'._g.%,;Yr?x")
	os.Setenv("MS_NAME", "franigen-example")
	os.Setenv("DB_NAME", "franigen")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "")

	app.Start()
}
