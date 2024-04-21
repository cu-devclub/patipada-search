package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	MLConfig struct {
		AmountAPIs       int
		TfIDFScoreWeight float64
		APIs             []*ExternalAPI
	}

	ExternalAPI struct {
		Name        string
		CsvFilePath string
	}
)

func (cfg *Config) ReadMlConfig() {
	file := "./ml-data/configs.env"

	err := godotenv.Load(file)
	if err != nil {
		log.Fatalf("Error loading %s file, %s", file, err)
	}

	tfIdfScoreWeight := readStringENVAndConvertToFloat("TF_IDF_SCORE_WEIGHT")
	amountAPIs := readStringENVAndConvertToInt("AMOUNT_APIS")

	externalAPIs := make([]*ExternalAPI, amountAPIs)
	for i := 1; i <= amountAPIs; i++ {
		modelNameInEnv := "MODEL_" + strconv.Itoa(i) + "_NAME"
		modelName := os.Getenv(modelNameInEnv)
		csvFilePath := "./ml-data/" + modelName + "/" + modelName + "-vector.csv"

		externalAPIs[i-1] = &ExternalAPI{
			Name:        modelName,
			CsvFilePath: csvFilePath,
		}
	}

	cfg.MlConfig = MLConfig{
		AmountAPIs:       amountAPIs,
		TfIDFScoreWeight: tfIdfScoreWeight,
		APIs:             externalAPIs,
	}
}

func (mlConfig *MLConfig) ToString() string {
	s := ""
	s += "API amount:" + strconv.Itoa(mlConfig.AmountAPIs) + "\n"
	for i, api := range mlConfig.APIs {
		s += "API " + strconv.Itoa(i+1) + ": " + api.Name + "\n"
		s += "CSV file path: " + api.CsvFilePath + "\n"
	}
	return s
}

func readStringENVAndConvertToInt(env string) int {
	envString := os.Getenv(env)
	envInt, err := strconv.Atoi(envString)
	if err != nil {
		log.Fatalf("Error converting %s to int", env)
	}
	return envInt
}

func readStringENVAndConvertToFloat(env string) float64 {
	envString := os.Getenv(env)
	envFloat, err := strconv.ParseFloat(envString, 64)
	if err != nil {
		log.Fatalf("Error converting %s to float", env)
	}
	return envFloat
}
