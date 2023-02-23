package initializers

func LoadEnvVariables(){

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

}