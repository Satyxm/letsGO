# fibocli

## Overview
fibocli is a command-line tool for generating Fibonacci sequences. It is built using Go and utilizes the `cobra` package for command handling.

## Features
- Generate a Fibonacci sequence up to a specified number of terms.
- Input validation to ensure only non-negative integers are accepted.

## Installation
To use fibocli, you need to have Go installed. Clone the repository and build the executable:

```sh
git clone https://github.com/Satyxm/letsGO.git
cd fibocli
go build -o fibocli
```

## Usage
Run the tool from the command line:

```sh
./fibocli generate <n>
```
Where `<n>` is a non-negative integer representing the number of Fibonacci terms to generate.

### Example
```sh
./fibocli generate 10
```
Output:
```
Fibonacci Sequence: [0 1 1 2 3 5 8 13 21 34]
```

## Dependencies
- `github.com/spf13/cobra`