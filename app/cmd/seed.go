package cmd

import (
	"gohub/pkg/console"
	"gohub/pkg/database/seeders"
	"gohub/pkg/seed"

	"github.com/spf13/cobra"
)

var CmdDBSeed = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with dummy data",
	Run:   runSeeders,
	Args:  cobra.MaximumNArgs(1),
}

func runSeeders(cmd *cobra.Command, args []string) {
	seeders.Initialize()
	if len(args) > 0 {
		// 有参数的情况
		name := args[0]
		seeder := seed.GetSeeder(name)
		if len(seeder.Name) > 0 {
			seed.RunSeeder(name)
		} else {
			console.Error("Seeder not found: " + name)
		}
	} else {
		// 没有参数的情况，默认全部执行
		seed.RunAll()
		console.Success("Database seeded!")
	}
}
