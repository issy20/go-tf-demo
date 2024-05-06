package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "app",
	Short: "app is a CLI tool to interact with the app API",
	Run: func(cmd *cobra.Command, args []string) {
		log.Print("root command executed")
	},
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()

	cobra.OnInitialize()
	RootCmd.Flags().String("db-name", "postgresql", "database name")
	RootCmd.Flags().String("stage", "dev", "stage")

	err := viper.BindPFlags(RootCmd.Flags())
	if err != nil {
		panic(err)
	}
}
