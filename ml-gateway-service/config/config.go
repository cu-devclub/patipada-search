package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var cfg *Config

type (
	Config struct {
		App      App
		MlConfig MlConfig
	}

	App struct {
		Port     int
		GRPCPort int
	}

	MlConfig struct {
		AmountAPIs int
		APIs       []*ExternalAPI
	}

	ExternalAPI struct {
		Name        string
		URL         string
		ScoreWeight float32
	}
)

func ReadConfig() {
	files := []string{"./app.env", "./ml-data/configs.env"}
	for _, file := range files {
		err := godotenv.Load(file)
		if err != nil {
			log.Fatalf("Error loading %s file, %s", file, err)
		}
	}

	serverPort := readStringENVAndConvertToInt("SERVER_PORT")
	grpcPort := readStringENVAndConvertToInt("GRPC_PORT")
	amountAPIs := readStringENVAndConvertToInt("AMOUNT_APIS")

	externalAPIs := make([]*ExternalAPI, amountAPIs)
	for i := 1; i <= amountAPIs; i++ {
		modelNameInEnv := "MODEL_" + strconv.Itoa(i) + "_NAME"
		modelName := os.Getenv(modelNameInEnv)
		envFileName := "./ml-data/" + modelName + "/" + modelName + ".env"
		err := godotenv.Load(envFileName)
		if err != nil {
			log.Fatalf("Error loading %s file, %s", envFileName, err)
		}

		scoreWeight := readStringENVAndConvertToFloat("SCORE_WEIGHT")

		externalAPIs[i-1] = &ExternalAPI{
			Name:        modelName,
			URL:         os.Getenv("API_URL"),
			ScoreWeight: scoreWeight,
		}

		os.Unsetenv("API_URL")
		os.Unsetenv("SCORE_WEIGHT")
	}

	cfg = &Config{
		App: App{
			Port:     serverPort,
			GRPCPort: grpcPort,
		},
		MlConfig: MlConfig{
			AmountAPIs: amountAPIs,
			APIs:       externalAPIs,
		},
	}
}

func GetConfig() Config {
	return *cfg
}

func readStringENVAndConvertToInt(env string) int {
	envString := os.Getenv(env)
	envInt, err := strconv.Atoi(envString)
	if err != nil {
		log.Fatalf("Error converting %s to int", env)
	}
	return envInt
}

func readStringENVAndConvertToFloat(env string) float32 {
	envString := os.Getenv(env)
	envFloat, err := strconv.ParseFloat(envString, 32)
	if err != nil {
		log.Fatalf("Error converting %s to float", env)
	}
	return float32(envFloat)
}
