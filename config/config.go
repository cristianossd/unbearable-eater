package config

import (
    "os"
)

func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
      return fallback
    }

    return value
}

var (
  CaloriesDataEndpoint = getEnv("CALORIES_DATA_ENDOINT", "https://caloriecontrol.org/wp-content/themes/ultimate-wp/calcontrol/searchFood.php")
)
