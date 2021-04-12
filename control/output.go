package control

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

type OutputFormatter struct {
	payloadChan chan ControlPayload
	format      string
	options     ControlReportingOptions
	wg          *sync.WaitGroup
}

func ensureOutputDirExists(outputDir string) {
	_, err := os.Stat(outputDir)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(outputDir, 0755)
		if errDir != nil {
			fmt.Printf(errDir.Error())
		}
	}
}

func getFileName(controlPack ControlPack, format string) string {
	fileName := fmt.Sprintf("sp_results_%d%02d%02dT%02d%02d%02d", controlPack.Timestamp.Year(), controlPack.Timestamp.Month(), controlPack.Timestamp.Day(), controlPack.Timestamp.Hour(), controlPack.Timestamp.Minute(), controlPack.Timestamp.Second())
	var extension string
	switch format {
	default:
		extension = format
	}
	return fmt.Sprintf("%s.%s", fileName, extension)
}

func outputFileResults(payloadChan chan ControlPayload, format string, outputDir string, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case payload := <-payloadChan:
		formattedResults := formatResults(payload.Pack, format)
		ensureOutputDirExists(outputDir)
		filePath := path.Join(outputDir, getFileName(payload.Pack, format))
		// TODO what file perms?
		// TODO what to do with file error?
		_ = ioutil.WriteFile(filePath, formattedResults, 0777)
	}
}
