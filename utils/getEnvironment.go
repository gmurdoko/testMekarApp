package utils

import "github.com/spf13/viper"

//ViperGetEnvironment app
func ViperGetEnvironment(key, defaultValue string) string {
	viper.AutomaticEnv()
	viper.SetConfigFile("./files/.env")
	viper.ReadInConfig()

	if environmentValue := viper.GetString(key); len(environmentValue) != 0 {
		return environmentValue
	}
	return defaultValue
}
