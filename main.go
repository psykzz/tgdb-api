package main

// DATABASE_HOST=direct.psykzz.com
// DATABASE_PORT=32823
// DATABASE_NAME=tgmc
// DATABASE_USER=root
// DATABASE_PASSWORD=paec3IeG5eequ6se0aeb1eikooquahla
// SECRET_KEY=Encode to Base64 format Encode to Base64 format
// GITHUB_TOKEN=f615ac281f178284041cf91e909608cfe5e70862

func main() {
	var cfg Config
	processConfig(&cfg)

	db := setupDatabase(&cfg)
	defer db.Close()

	// Run the webserver and handle routes
	r := handleRoutes(db)

	r.Run()
}
