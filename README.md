# Awk Implementation in Go

This repository provides an implementation of the Awk programming language in Go. Awk is a powerful text-processing tool commonly used for pattern scanning and processing.

## Overview

This implementation of Awk allows users to:

- Process text files using patterns and actions.
- Define custom functions.
- Use associative arrays for advanced data manipulation.
- Leverage built-in string functions for text processing.
- Execute shell commands and handle pipes.

The implementation is designed to be idiomatic, efficient, and extensible.

## Installation

To install the program, ensure you have Go installed on your system. Then, clone this repository and build the executable:

```bash
# Clone the repository
git clone https://github.com/iv4n-ga6l/awk.git
cd awk

# Build the executable
go build -o ccawk
```

## Usage

The program can be run using the following syntax:

```bash
ccawk [-F separator] [-RS separator] [-f programfile | 'program'] [-v var=value] [file ...]
```

### Command-Line Options

- `-F separator`: Specifies the field separator (default: tab `\t`).
- `-RS separator`: Specifies the record separator (default: newline `\n`).
- `-f programfile`: Specifies a file containing the Awk program.
- `-v var=value`: Assigns a value to a variable before execution.
- `[file ...]`: Specifies the input files to process (default: standard input).

### Examples

#### Example 1: Print the first field of each line

```bash
echo -e "John 25 London\nJane 30 New York" | ./ccawk -F ' ' '{ print $1 }'
```

Output:
```
John
Jane
```

#### Example 2: Use a program file

Create a file `program.awk`:

```awk
BEGIN { print "Name\tAge\tCity" }
{ print $1, $2, $3 }
END { print "Processing complete." }
```

Run the program:

```bash
./ccawk -F ' ' -f program.awk test.txt
```

Output:
```
Name    Age    City
John    25     London
Jane    30     New York
Bob     22     Paris
Alice   35     Tokyo
Charlie 28     Berlin
Processing complete.
```

#### Example 3: Using associative arrays

```bash
cat test.txt | ./ccawk -F ' ' '{ count[$3]++ } END { for (city in count) print city, count[city] }'
```

Output:
```
London 1
New York 1
Paris 1
Tokyo 1
Berlin 1
```

#### Example 4: String functions

```bash
echo "hello world" | ./ccawk '{ print toupper($0) }'
```

Output:
```
HELLO WORLD
```

## Features

### Custom Functions

You can define and call custom functions in your Awk programs. For example:

```awk
function greet(name) {
    print "Hello, " name "!"
}

BEGIN {
    greet("Alice")
    greet("Bob")
}
```

### Associative Arrays

Associative arrays are supported, allowing you to store and manipulate key-value pairs. For example:

```awk
{ count[$1]++ }
END {
    for (name in count) {
        print name, count[name]
    }
}
```

### Built-in String Functions

The following string functions are available:

- `length(s)`: Returns the length of the string `s`.
- `substr(s, start, length)`: Extracts a substring from `s` starting at `start` with the specified `length`.
- `index(s, substr)`: Returns the position of `substr` in `s`.
- `split(s, sep)`: Splits `s` into an array using `sep` as the delimiter.
- `sub(regex, replacement, s)`: Replaces the first match of `regex` in `s` with `replacement`.
- `gsub(regex, replacement, s)`: Replaces all matches of `regex` in `s` with `replacement`.
- `match(regex, s)`: Returns the position of the first match of `regex` in `s`.
- `sprintf(format, args...)`: Formats a string using the specified format.
- `tolower(s)`: Converts `s` to lowercase.
- `toupper(s)`: Converts `s` to uppercase.

### Control Flow Constructs

- `for (key in array)`: Iterates over the keys of an associative array.
- `delete array[key]`: Deletes a key-value pair from an associative array.

### External Commands and Pipes

You can execute shell commands and use pipes in your Awk programs. For example:

```awk
{ print $1 | "sort" }
```

You can also use the `close()` function to close a pipe:

```awk
{ print $1 | "sort"; close("sort") }
```

### Getline

The `getline` command allows you to read input from a file, command, or standard input:

```awk
BEGIN {
    while ((getline line < "test.txt") > 0) {
        print line
    }
    close("test.txt")
}
```

## Running Tests

To run the tests, use the following command:

```bash
go test ./...
```