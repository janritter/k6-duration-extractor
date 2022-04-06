package cmd

import (
	"encoding/csv"
	"testing"

	"os"

	"github.com/stretchr/testify/assert"
)

func TestRemoveSliceElement(t *testing.T) {
	test := []string{"a", "b", "c"}
	result := removeSliceElement(test, 1)

	assert.Equal(t, []string{"a", "c"}, result)
}

func TestGetNameFromFilename(t *testing.T) {
	test := "test1.csv"
	result := getNameFromFilename(test)

	assert.Equal(t, "test1", result)
}

func TestReadCSVFile(t *testing.T) {
	filename := "../tests/short.csv"
	result := readCSVFile(filename)

	assert.Equal(t, [][]string{
		{
			"metric_name",
			"timestamp",
			"metric_value",
			"check",
			"error",
			"error_code",
			"group",
			"method",
			"name",
			"proto",
			"scenario",
			"status",
			"subproto",
			"tls_version",
			"url",
			"extra_tags",
		},
		{
			"http_reqs",
			"1595325560",
			"1.000000",
			"",
			"",
			"",
			"",
			"GET",
			"http://test.k6.io",
			"HTTP/1.1",
			"default",
			"200",
			"",
			"",
			"http://test.k6.io",
			"",
		},
	},
		result)
}

func TestWriteCSVFile(t *testing.T) {
	filename := "../tests/write.csv"
	data := [][]string{
		{
			"header1",
			"header2",
		},
		{
			"data1",
			"data2",
		},
	}

	writeCSVFile(filename, data)

	file, err := os.Open(filename)
	assert.Nil(t, err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	assert.Nil(t, err)

	assert.Equal(t, data, records)

	err = os.Remove(filename)
	assert.Nil(t, err)
}

func TestRunCommand(t *testing.T) {
	filename := "../tests/result.csv"
	filenameExtracted := "../tests/result-extracted.csv"
	rootCmd.Run(nil, []string{filename})

	file, err := os.Open(filenameExtracted)
	assert.Nil(t, err)
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	assert.Nil(t, err)

	assert.Equal(t, [][]string{
		{
			"timestamp",
			"duration",
			"check",
			"error",
			"error_code",
			"group",
			"method",
			"name",
			"proto",
			"scenario",
			"status",
			"subproto",
			"tls_version",
			"url",
			"extra_tags",
		},
		{
			"1595325560",
			"221.899000",
			"",
			"",
			"",
			"",
			"GET",
			"http://test.k6.io",
			"HTTP/1.1",
			"default",
			"200",
			"",
			"",
			"http://test.k6.io",
			"",
		},
		{
			"1595325560",
			"221.899000",
			"",
			"",
			"",
			"",
			"GET",
			"http://test.k6.io",
			"HTTP/1.1",
			"default",
			"200",
			"",
			"",
			"http://test.k6.io",
			"",
		},
	},
		records)

	err = os.Remove(filenameExtracted)
	assert.Nil(t, err)
}
