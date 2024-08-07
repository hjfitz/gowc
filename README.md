# wc in Go

This project provides a simple Go implementation of the classic Unix command wc, which counts lines, words, and bytes in a file.

## Features:

1. Counts lines, words, and bytes in a file.

## Getting Started

### Prerequisites:

1. Go version 1.18 or later: Download Go

### Clone the repository:

```bash
git clone https://github.com/hjfitz/gowc
cd gowc
```

### Build the project:

```bash
make build
```

This will create an executable file named wc in your project directory.

### Run the program:
```bash
./wc your_file.txt
```

This will print the number of lines, words, and bytes in the specified file.

## Usage

```
wc [FILE]
```

* FILE: The file path to be analyzed. If no file is provided, it will read from standard input.

### Example

```bash
./wc -l -m -c -w example.txt
```
This will output something similar to:

```
  10 20 50 100 example.txt
```

#### Explanation:

- **10:** The number of bytes in the file.
- **20:** The number of lines in the file.
- **50:** The number of characters in the file.
- **100:** The number of words in the file.

