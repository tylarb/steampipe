package control

import (
	"fmt"
	spinner2 "github.com/briandowns/spinner"
	"github.com/turbot/steampipe/constants"
	"sync"

	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/turbot/steampipe/utils"
)

func getControlStatusText(status string, options ControlReportingOptions) string {
	// Don't add color if the user has opted out
	if !options.WithColor {
		return fmt.Sprintf(status)
	}

	switch status {
	case ControlAlarm:
		fallthrough
	case ControlError:
		return text.FgRed.Sprint(status)
	case ControlOK:
		return text.FgGreen.Sprint(status)
	case ControlInfo:
		return text.FgHiBlue.Sprint(status)
	case ControlSkipped:
		fallthrough
	default:
		return status
	}
}

func getControlStatusTotalText(status string, total int, options ControlReportingOptions) string {
	// If 0 total, don't display
	if total == 0 {
		return ""
	}

	// Don't add color if the user has opted out
	if !options.WithColor {
		return fmt.Sprintf("%d %s. ", total, status)
	}

	switch status {
	case ControlAlarm:
		fallthrough
	case ControlError:
		return fmt.Sprintf("%s. ", text.FgRed.Sprint(fmt.Sprintf("%d %s", total, status)))
	case ControlOK:
		return fmt.Sprintf("%s. ", text.FgGreen.Sprint(fmt.Sprintf("%d %s", total, status)))
	case ControlInfo:
		return fmt.Sprintf("%s. ", text.FgHiBlue.Sprint(fmt.Sprintf("%d %s", total, status)))
	case ControlSkipped:
		fallthrough
	default:
		return fmt.Sprintf("%d %s. ", total, status)
	}
}

func displayControlResults(stdoutChan chan ControlPayload, reportingOptions ControlReportingOptions, wg *sync.WaitGroup) {
	var spinner *spinner2.Spinner

	if reportingOptions.WithProgress {
		spinner = utils.ShowSpinner("Running controls")
	}

	select {
	case payload := <-stdoutChan:
		utils.StopSpinner(spinner)
		switch reportingOptions.OutputFormat {
		case constants.ArgDynamic:
			displayDynamicControlDetails(payload.Pack, reportingOptions)
			displayDynamicControlsSummary(payload.Pack, reportingOptions)
			break
		default:
			displayResults(string(formatResults(payload.Pack, reportingOptions.OutputFormat)))
		}
		wg.Done()
	}
}

//func displayControlsTable(controlPack ControlPack, options ControlReportingOptions) {
//	// the buffer to put the output data in
//	outbuf := bytes.NewBufferString("")
//
//	// the table
//	t := table.NewWriter()
//	t.SetOutputMirror(outbuf)
//	t.SetStyle(table.StyleDefault)
//	t.Style().Format.Header = text.FormatDefault
//
//	headers := table.Row{
//		"Status",
//		"ID",
//		"Name",
//		"Description",
//	}
//	colConfigs := []table.ColumnConfig{
//		{
//			Name:     "Status",
//			Number:   1,
//			WidthMax: constants.MaxColumnWidth,
//		},
//		{
//			Name:     "ID",
//			Number:   2,
//			WidthMax: constants.MaxColumnWidth,
//		},
//		{
//			Name:     "Name",
//			Number:   3,
//			WidthMax: constants.MaxColumnWidth,
//		},
//		{
//			Name:     "Description",
//			Number:   4,
//			WidthMax: constants.MaxColumnWidth,
//		},
//	}
//
//	t.SetColumnConfigs(colConfigs)
//	t.AppendHeader(headers)
//
//	for _, control := range controlPack.ControlRuns {
//		for _, result := range control.Results {
//			row := table.Row{
//				getControlStatusTotalText(result.Status, options),
//				control.Type.ControlID,
//				control.Type.Title,
//				control.Type.Description,
//			}
//			t.AppendRow(row)
//		}
//	}
//
//	t.Render()
//	fmt.Println("Control details")
//	fmt.Println("")
//	display.ShowPaged(outbuf.String())
//	fmt.Println("")
//}

func displayDynamicControlsSummary(controlPack ControlPack, options ControlReportingOptions) {
	// the buffer to put the output data in
	//outbuf := bytes.NewBufferString("")

	alarmTotal := 0
	errorTotal := 0
	okTotal := 0
	infoTotal := 0
	skippedTotal := 0
	totalControls := 0

	for _, control := range controlPack.ControlRuns {
		for _, result := range control.Results {
			totalControls += 1
			switch result.Status {
			case ControlAlarm:
				alarmTotal += 1
			case ControlError:
				errorTotal += 1
			case ControlInfo:
				infoTotal += 1
			case ControlOK:
				okTotal += 1
			case ControlSkipped:
				skippedTotal += 1
			}
		}
	}

	alarmText := getControlStatusTotalText(ControlAlarm, alarmTotal, options)
	errorText := getControlStatusTotalText(ControlError, errorTotal, options)
	okText := getControlStatusTotalText(ControlOK, okTotal, options)
	infoText := getControlStatusTotalText(ControlInfo, infoTotal, options)
	skippedText := getControlStatusTotalText(ControlSkipped, skippedTotal, options)
	totalText := getControlStatusTotalText("total", totalControls, options)

	fmt.Println(fmt.Sprintf("Controls: %s%s%s%s%s%s\n", totalText, alarmText, errorText, okText, infoText, skippedText))

	//fmt.Println(fmt.Sprintf("%d %s", totalControls, getPluralisedControlsText(totalControls)))
	//fmt.Println("")
	//display.ShowPaged(outbuf.String())
	//fmt.Println("")
}

func displayDynamicControlDetails(controlPack ControlPack, options ControlReportingOptions) {
	for _, control := range controlPack.ControlRuns {
		fmt.Println(fmt.Sprintf("Control: %s: \"%s\"", control.Type.ControlID, control.Type.Title))
		for _, result := range control.Results {
			switch result.Status {
			case ControlAlarm:
				fmt.Println("")
				fmt.Println(fmt.Sprintf("  %s for resource: %s", getControlStatusText(ControlAlarm, options), result.Resource))
				fmt.Println("")
			}
		}
	}
}

func displayResults(results string) {
	fmt.Println(results)
}
