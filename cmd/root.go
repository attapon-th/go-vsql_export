/*
Copyright © 2023 Attapon.TH
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vsql_export",
	Short: "Export Vertica SQL Query",
	Long: `Export Vertica SQL Query to CSV file.
Supports:
	- CSV`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug().Msg("run root")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.Flags().SortFlags = false
	rootCmd.PersistentFlags().SortFlags = false
	rootCmd.PersistentFlags().StringP("dsn", "d", "", "Veritca DSN (Example: vertica://dbadmin:password@localhost:5433/Vmart?sslmode=disable)")
	rootCmd.PersistentFlags().StringP("host", "H", "localhost", "Veritca Hostname")
	rootCmd.PersistentFlags().StringP("port", "", "5433", "Veritca Port")
	rootCmd.PersistentFlags().StringP("username", "U", "dbadmin", "Veritca Username")
	rootCmd.PersistentFlags().StringP("password", "P", "", "Veritca Password")
	rootCmd.PersistentFlags().StringP("db", "D", "VMart", "Veritca Database")

	rootCmd.PersistentFlags().String("log-level", "info", "Log level (debug, info, warn, error, fatal, panic)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// viperBindFlags()
}

func initConfig() {
	viper.AutomaticEnv()
	viper.BindPFlag("db.dsn", rootCmd.PersistentFlags().Lookup("dsn"))
	viper.BindPFlag("db.host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("db.port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("db.username", rootCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("db.password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("db.database", rootCmd.PersistentFlags().Lookup("database"))
	viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	llv, _ := zerolog.ParseLevel(viper.GetString("log.level"))
	zerolog.SetGlobalLevel(llv)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	// zerolog.NewConsoleWriter()
	log.Logger = zerolog.New(
		zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: "2006-01-02 15:04:05.000",
		}).With().Timestamp().Logger()

	for _, k := range viper.AllKeys() {
		log.Debug().Str(k, fmt.Sprint(viper.Get(k))).Msg("Config loaded")
	}
}
