package cmd

import (
	"fmt"
	"gcnt/internal/repository"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(migrateNewCmd)
	cmd.AddCommand(migrateUpCmd)
	cmd.AddCommand(migrateDownCmd)
}

var migrateNewCmd = &cobra.Command{
	Use:   "migratenew [migrations name]",
	Short: "Create new migrations file",
	Long:  `Create new migrations file on folder migrations with timestamp as prefix`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mDir := "migrations/"

		createMigrationFile(mDir, args[0])
	},
}

var migrateUpCmd = &cobra.Command{
	Use:   "migrateup",
	Short: "Migrate Up DB",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		mSource := getMigrateSource()
		doMigrate(mSource, "mysql", migrate.Up)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate down DB",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		mSource := getMigrateSource()
		doMigrate(mSource, "mysql", migrate.Down)
	},
}

func createMigrationFile(mDir string, mName string) error {
	var migrationContent = `-- +migrate Up

-- +migrate Down`

	datePrefix := time.Now().Format("200601021504")
	filename := fmt.Sprintf("%s_%s.sql", datePrefix, mName)
	filepath := fmt.Sprintf("%s%s", mDir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		log.Printf(fmt.Sprintf("Error create migrations file | %v", err))
		return err
	}
	defer f.Close()

	f.WriteString(migrationContent)
	f.Sync()

	log.Printf(fmt.Sprintf("New migrations file has been created: %s)", filepath))
	return nil
}

func getMigrateSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: "migrations",
	}

	return source
}

func doMigrate(mSource migrate.FileMigrationSource, dbDialect string, direction migrate.MigrationDirection) (err error) {
	db := &repository.Db{}
	err = db.InitDatabase("mysql")
	sqlDB, err := db.Mysql.DB()
	if err != nil {
		log.Fatal().Err(err).Caller().Send()
		return
	}
	defer sqlDB.Close()

	total, err := migrate.Exec(sqlDB, dbDialect, mSource, direction)
	if err != nil {
		log.Printf(fmt.Sprintf("Fail migrations | %v", err))
		return
	}

	log.Printf(fmt.Sprintf("Migrate Success, total migrated: %d", total))

	return
}
