# Cat Program

## Description

This Go program mimics the behavior of the Unix `cat` command. It reads the contents of one or more text files and prints them to the standard output. The program also supports optional flags for numbering lines and suppressing repeated empty lines.

## Usage

You can run the program from the command line as follows:

```cmd
cat.exe [options] [filename]
```

- `[options]`: One or more optional flags described below.
- `file1`, `file2`, ...: One or more file paths that you want to read.

## Options

- `-n`: Number all output lines.
- `-s`: Suppress repeated empty lines.
