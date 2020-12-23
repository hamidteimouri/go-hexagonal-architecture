package app

func StartApplication() {
	Migrate()
	Seed()
	Serve()
}

func Migrate() {
	// Migrate databases
}

func Seed() {
	// Seed database
}

func Serve() {
	// Serve the application
}
