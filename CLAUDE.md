# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a [Hyperskill](https://hyperskill.org/projects/201) educational project: a console program in Go that reads a list of taboo (obscene) words from a file and censors them in user input.

## Running and Testing

Run the program:
```bash
cd "Obscene Vocabulary Checker/task"
go run main.go
```

Run the tests (requires Python with the `hstest` library from `requirements.txt`):
```bash
pip install -r requirements.txt
cd "Obscene Vocabulary Checker/task"
python tests.py
```

## Project Structure

All implementation goes in a single file: `Obscene Vocabulary Checker/task/main.go`.

Tests are in `Obscene Vocabulary Checker/task/tests.py` (Python, using Hyperskill's `hstest` framework — do not modify).

The four stages are defined in sub-directories under `Obscene Vocabulary Checker/`:
- `Find the bad words/` — Stage 1
- `Check the word/` — Stage 2
- `Censorship in action/` — Stage 3
- `Correct the sentence/` — Stage 4

Each stage's `task.html` describes the requirements. Only `task/main.go` is the learner-editable file.

## Stage Requirements Summary

**Stage 1 — Find the bad words:** Read a filename from stdin, read that file, print each word on a separate line.

**Stage 2 — Check the word:** After reading the taboo file, read one word from stdin. Print `True` if it's taboo (case-insensitive), `False` otherwise.

**Stage 3 — Censorship in action:** After reading the taboo file, loop reading words from stdin. Replace taboo words with `*` repeated to match the word's length (case-insensitive match). Print the (possibly censored) word each iteration. When `exit` is entered, print `Bye!` and stop.

**Stage 4 — Correct the sentence:** (See `Correct the sentence/task.html` for full spec.)

## Key Constraints

- A "word" is a sequence of Latin characters (upper or lower case).
- Taboo matching is case-insensitive; censorship uses `*` × len(word), not a fixed `******`.
- The active stage's tests are in `task/tests.py` — the test file is replaced per stage by the Hyperskill IDE plugin.
