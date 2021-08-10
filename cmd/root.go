/*
Copyright © 2021 Jan Ritter <git@janrtr.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

//metric_name timestamp duration check error error_code group method name proto scenario status subproto tls_version url extra_tags
type record struct {
	timestamp     string
	duration  string
	check         string
	error_message string
	error_code    string
	group         string
	method        string
	name          string
	proto         string
	scenario      string
	status        string
	subproto      string
	tls_version   string
	url           string
	extra_tags    string
}

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "k6-duration-extractor",
	Short: "Tool to extract the duration information per request from the K6 csv output",
	Long:  `Tool which extracts the duration information per request from the K6 csv output and stores it in a new csv file.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide the filename of the K6 output csv.")
			os.Exit(1)
		}

		filename := args[0]

		fmt.Printf("Reading K6 output from %s\n", filename)

		records := readCSVFile(filename)

		var result []record

		for k, v := range records {
			if k != 0 && v[0] == "http_req_duration" {
				data := record{
					timestamp:     v[1],
					duration:  	   v[2],
					check:         v[3],
					error_message: v[4],
					error_code:    v[5],
					group:         v[6],
					method:        v[7],
					name:          v[8],
					proto:         v[9],
					scenario:      v[10],
					status:        v[11],
					subproto:      v[12],
					tls_version:   v[13],
					url:           v[14],
					extra_tags:    v[15],
				}

				result = append(result, data)
			}
		}

		var output [][]string

		output = append(output,[]string{"timestamp", "duration", "check", "error_message", "error_code", "group", "method", "name", "proto", "scenario", "status", "subproto", "tls_version", "url", "extra_tags"})

		for _, v := range result {
			output = append(output, recordStructToSlice(v))
		}

		writeCSVFile("k6-duration-extractor.csv", output)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {}

func readCSVFile(filenme string) [][]string {
	file, err := os.Open(filenme)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records
}

func writeCSVFile(filename string, records [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(file)
	writer.WriteAll(records)
	writer.Flush()

	fmt.Println("Wrote " + strconv.Itoa(len(records)) + " records to " + filename)

}


func recordStructToSlice(record record) []string {
	return []string{
		record.timestamp,
		record.duration,
		record.check,
		record.error_message,
		record.error_code,
		record.group,
		record.method,
		record.name,
		record.proto,
		record.scenario,
		record.status,
		record.subproto,
		record.tls_version,
		record.url,
		record.extra_tags,
	}
}
