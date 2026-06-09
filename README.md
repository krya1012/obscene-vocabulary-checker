# Obscene Vocabulary Checker

A console program in Go that reads a list of taboo words from a file and censors them in user input. Built as a [Hyperskill project](https://hyperskill.org/projects/201).

## Usage

```bash
cd "Obscene Vocabulary Checker/task"
go run main.go
```

The program reads a filename from stdin, opens that file, and processes its contents according to the current stage.

## Stages

| Stage | Description |
|-------|-------------|
| 1 | Read a taboo words file and print each word |
| 2 | Check whether a given word is taboo (case-insensitive) |
| 3 | Censor taboo words in a stream of input words |
| 4 | Correct a full sentence by censoring taboo words |

## Running Tests

Tests require Python 3 and the `hstest` library:

```bash
pip3 install -r requirements.txt
cd "Obscene Vocabulary Checker/task"
python3 tests.py
```
