/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"strings"

	"github.com/attapon-th/go-vsql_export/pkg"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// csvCmd represents the csv command
var csvCmd = &cobra.Command{
	Use:   "csv [SQL]",
	Short: "export to csv",
	Long:  `Export SQL Query to CSV file.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("csv called")
		run(args[0])
	},
}

func init() {
	rootCmd.AddCommand(csvCmd)

	csvCmd.PersistentFlags().StringP("output", "o", "stdout", "Output file")
	csvCmd.PersistentFlags().StringP("compression", "z", "none", "Compression (gzip, zip, none)")
	csvCmd.PersistentFlags().IntP("level", "", 5, "Compression level (1-9)")

	viper.BindPFlag("csv.output", csvCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("csv.compression", csvCmd.PersistentFlags().Lookup("compression"))
	viper.BindPFlag("csv.level", csvCmd.PersistentFlags().Lookup("level"))
}

func run(sql string) {
	log.Debug().Str("sql", sql).Msg("SQL query")
	vdb, err := pkg.ConnectVerticaWithDSN(viper.GetString("db.dsn"))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}
	var fs io.Writer
	var output string

	output = viper.GetString("csv.output")
	if w, err := pkg.FileWriter(output); err != nil {
		log.Fatal().Err(err).Msg("Failed to create output file")
	} else {
		fs = w
		defer w.Close()
	}
	// log.Debug().Str("output", output).Msg("Output file")
	if viper.GetString("csv.compression") == "gzip" {
		w, err := pkg.CreateGzip(fs, viper.GetInt("csv.level"))
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create gzip file")
		}
		fs = w
		defer w.Close()
	} else if viper.GetString("csv.compression") == "zip" {
		output = strings.ToLower(output)
		if output == "" {
			output = "temp.csv"
		} else if strings.TrimSuffix(output, ".zip") == "" {
			if !strings.HasSuffix(output, ".csv") {
				output = output + ".csv"
			}
		}
		w, c, err := pkg.CreateZip(fs, output)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create zip file")
		}
		fs = w
		defer c.Close()
	}
	rows, err := vdb.Query(sql)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to query")
	}
	if err := pkg.ToCsv(rows, fs); err != nil {
		log.Fatal().Err(err).Msg("Failed to export csv")
	}
}
