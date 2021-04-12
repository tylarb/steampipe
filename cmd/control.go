package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/turbot/steampipe/cmdconfig"
	"github.com/turbot/steampipe/constants"
	"github.com/turbot/steampipe/control"
)

// ControlCmd :: Steampipe control management
func ControlCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "control [command]",
		Args:  cobra.NoArgs,
		Short: "Steampipe control management",
		// TODO(nw) expand long description
		Long: `Steampipe control management.

Run Steampipe control packs and output in a variety of formats.`,
	}

	cmd.AddCommand(ControlRunCmd())

	return cmd
}

// ControlRunCmd :: handler for control run
func ControlRunCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "run",
		Args:  cobra.ArbitraryArgs,
		Run:   runControlRunCmd,
		Short: "Steampipe control runner",
		Long: `Steampipe control runner.

Run a  as a local service, exposing it as a database endpoint for
connection from any Postgres compatible database client.`,
	}

	cmdconfig.
		OnCmd(cmd).
		AddStringFlag(constants.ArgOutput, "", "dynamic", "The format(s) to output the report to. Must be one or more of csv, html or json in a comma-separated list").
		AddStringFlag(constants.ArgOutputFileDir, "", "", "The directory to output the control results to. Defaults to ./control-runs.").
		AddStringFlag(constants.ArgOutputFileFormat, "", "", "The directory to output the control results to.").
		AddBoolFlag(constants.ArgNoColor, "", false, "Do not add color to outputs in the terminal output.").
		AddBoolFlag(constants.ArgNoProgress, "", false, "Do not show progress of the operation in the terminal output.")

	return cmd
}

func runControlRunCmd(cmd *cobra.Command, args []string) {
	output := cmdconfig.Viper().GetString(constants.ArgOutput)
	outputFileDir := cmdconfig.Viper().GetString(constants.ArgOutputFileDir)
	outputFileFormat := cmdconfig.Viper().GetString(constants.ArgOutputFileFormat)
	noColor := cmdconfig.Viper().GetBool(constants.ArgNoColor)
	noProgress := cmdconfig.Viper().GetBool(constants.ArgNoProgress)

	reportingOptions := control.ControlReportingOptions{
		OutputFormat:        output,
		OutputFileFormats:   strings.Split(outputFileFormat, ","),
		OutputFileDirectory: outputFileDir,
		WithColor:           !noColor,
		WithProgress:        !noProgress,
	}

	control.RunControl(reportingOptions)
}
