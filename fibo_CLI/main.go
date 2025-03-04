package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// fibo code
func fibonacci(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("please enter a non-negative integer")
	}

	fib := make([]int, n)
	fib[0] = 0

	if n > 1 {
		fib[1] = 1
	}

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib, nil

}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "fibocli",
		Short: "fibocli is a fibonacci cli tool",
		Long:  "fibocli is a fibonacci cli tool",
	}

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "generate fibonacci sequence",
		Long:  "generate fibonacci sequence",
		Run: func(cmd *cobra.Command, args []string) {
			n, err := strconv.Atoi(args[0])

			if err != nil || n < 0 {
				fmt.Println("please enter a non-negative integer")
				os.Exit(1)
			}

			fib, err := fibonacci(n)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)

			}

			fmt.Println("Fibonacci Sequence: ", fib)

		},
	}

	rootCmd.AddCommand(generateCmd)
	rootCmd.Execute()
}
