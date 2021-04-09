package control

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"time"

	"github.com/turbot/steampipe/constants"
)

func formatResults(controlPack ControlPack, format string) []byte {
	switch format {
	case constants.ArgCSV:
		return formatCSV(controlPack)
	case constants.ArgJSON:
		return formatJSON(controlPack)
	default:
		return nil
	}
}

func formatTimestampForCSV(timestamp time.Time) string {
	return timestamp.Format("2006-01-02T15:04:05")
}

func formatCSV(controlPack ControlPack) []byte {
	// TODO CSV structure
	outbuf := bytes.NewBufferString("")
	csvWriter := csv.NewWriter(outbuf)
	_ = csvWriter.Write([]string{
		"control_id",
		"status",
		"reason",
		"resource",
		"timestamp",
	})
	for _, controlRun := range controlPack.ControlRuns {
		for _, control := range controlRun.Results {
			_ = csvWriter.Write([]string{
				controlRun.Type.ControlID,
				control.Status,
				control.Reason,
				control.Resource,
				formatTimestampForCSV(controlPack.Timestamp),
			})
		}
	}
	csvWriter.Flush()
	return outbuf.Bytes()
}

func formatJSON(controlPack ControlPack) []byte {
	// TODO JSON structure
	res, _ := json.MarshalIndent(controlPack, "", "  ")
	return res
}
