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
		AddBoolFlag(constants.ArgColor, "", true, "Add color to outputs in the terminal output.").
		AddBoolFlag(constants.ArgProgress, "", true, "Do not show progress of the operation in the terminal output.")

	return cmd
}

func runControlRunCmd(cmd *cobra.Command, args []string) {
	output := cmdconfig.Viper().GetString(constants.ArgOutput)
	outputFileDir := cmdconfig.Viper().GetString(constants.ArgOutputFileDir)
	outputFileFormats := cmdconfig.Viper().GetString(constants.ArgOutputFileFormat)
	withColor := cmdconfig.Viper().GetBool(constants.ArgColor)
	withProgress := cmdconfig.Viper().GetBool(constants.ArgProgress)

	//var isTerminal bool
	//if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
	//	isTerminal = true
	//} else {
	//	isTerminal = false
	//}

	splitFn := func(c rune) bool {
		return c == ','
	}

	reportingOptions := control.ControlReportingOptions{
		OutputFormat:        output,
		OutputFileFormats:   strings.FieldsFunc(outputFileFormats, splitFn),
		OutputFileDirectory: outputFileDir,
		WithColor:           withColor,
		WithProgress:        withProgress,
	}

	control.RunControl(reportingOptions)
}
