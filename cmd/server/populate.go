package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gitlab.com/DGuedes/API-Exercise/parser"
	"gitlab.com/DGuedes/API-Exercise/storage/postgres"
)

var (
	path string
)

func init() {
	rootCmd.AddCommand(populateCmd)
	populateCmd.PersistentFlags().StringVar(&path, "path", "", "Csv file used to populate the database.")
}

var populateCmd = &cobra.Command{
	Use:   "populate",
	Short: "Populate the database.",
	Long: `Populate the current Postgres database with the titanic dataset.
Example:
	server populate --dbHost postgres --env dev --path /src/titanic
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := postgres.BuildConfig(environment)
		logrus.Infof("Connecting to postgres %s:%s", cfg.Host, cfg.Port)
		parser := parser.TitanicParser{}
		factory := postgres.BuildDAOFactory(&cfg)
		factory.Connect()

		logrus.Info("Running migrations")
		p := factory.NewPersonDAO()
		p.AutoMigrate()

		logrus.Infof("Parsing file %s", path)
		result := parser.Parse(path)

		logrus.Info("Saving parsing result to Postgres")
		err := p.BulkInsert(result)
		if err != nil {
			logrus.Info("Successfully populated the database.")
		}
	},
}
