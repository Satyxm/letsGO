package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func parseJSON(input []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(input, &data)
	return data, err
}

func main() {
	var filePath string
	var pretty bool
	var queryKey string

	rootCmd := &cobra.Command{
		Use:   "jsonparser",
		Short: "A simple JSON parser CLI",
		Run: func(cmd *cobra.Command, args []string) {
			var input []byte
			var err error

			if filePath != "" {
				input, err = ioutil.ReadFile(filePath)
			} else {
				input, err = ioutil.ReadAll(os.Stdin)
			}

			if err != nil {
				fmt.Println("Error reading input:", err)
				os.Exit(1)
			}

			jsonData, err := parseJSON(input)
			if err != nil {
				fmt.Println("Invalid JSON:", err)
				os.Exit(1)
			}

			if queryKey != "" {
				if value, found := jsonData[queryKey]; found {
					fmt.Println(value)
				} else {
					fmt.Println("Key not found")
					os.Exit(1)
				}
				return
			}

			if pretty {
				prettyJSON, _ := json.MarshalIndent(jsonData, "", "  ")
				fmt.Println(string(prettyJSON))
			} else {
				fmt.Println(string(input))
			}
		},
	}

	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to JSON file")
	rootCmd.Flags().BoolVarP(&pretty, "pretty", "p", false, "Pretty-print JSON output")
	rootCmd.Flags().StringVarP(&queryKey, "query", "q", "", "Extract specific key from JSON")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
