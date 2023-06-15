/*
Copyright © 2023 Kristin Laemmert kristin.laemmert@grafana.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mildwonkey/dashmig/internal/kinds/dashboard"
	"github.com/mildwonkey/dashmig/internal/legacydash"
	"github.com/spf13/cobra"
)

// schematizeCmd represents the schematize command
var schematizeCmd = &cobra.Command{
	Use:   "schematize",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		schematize(cmd, args)
	},
}

// schematizeCmd flags
var (
	inputFile  string
	outputFile string
)

func init() {
	schematizeCmd.Flags().StringVarP(&inputFile, "file", "f", "dashboard.json", "input json file to schematize")
	schematizeCmd.Flags().StringVarP(&outputFile, "out", "o", "", "out json file to schematize (defaults to stdout)")
	rootCmd.AddCommand(schematizeCmd)
}

func schematize(cmd *cobra.Command, args []string) error {
	src, err := os.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("opening input file failed: %s", err.Error())
	}

	// At this point we're implementing
	// https://github.com/grafana/grafana/blob/main/public/app/features/dashboard/state/DashboardMigrator.ts
	// in Go.
	newdash, err := readDash(src)
	if err != nil {
		fmt.Printf("error reading dashboard: %s\n", err.Error())
		return err
	}

	// write to stodout or a file
	j, err := json.Marshal(newdash)
	if err != nil {
		fmt.Printf("error marshalling json: %s\n", err.Error())
		return err
	}
	// TODO: actually capture errors here
	if outputFile == "" {
		fmt.Printf("writing to stdout\n")
		os.Stdout.Write(j)
	} else {
		fmt.Printf("writing to %s\n", outputFile)
		os.WriteFile(outputFile, j, 0644)
	}

	return nil
}

func readDash(src []byte) (*dashboard.Dashboard, error) {
	// get the legacy schema version
	legacyVersion, err := sniffSchemaVersion(src)
	if err != nil {
		return nil, err
	}

	switch legacyVersion {
	case 35:
		return legacydash.ReadDashv35(src)
	case 36:
		return legacydash.ReadDashv36(src)
	case 37:
		return legacydash.ReadDashv37(src)
	case 38:
		return legacydash.ReadDashv38(src)
	default:
		return nil, fmt.Errorf("unsupported legacy schema version %d", legacyVersion)
	}
}

func sniffSchemaVersion(src []byte) (uint64, error) {
	var sniff struct {
		SchemaVersion *uint64 `json:"schemaVersion"`
	}
	err := json.Unmarshal(src, &sniff)
	if err != nil {
		return 0, err
	}
	return *sniff.SchemaVersion, nil
}
