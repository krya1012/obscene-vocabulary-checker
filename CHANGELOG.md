# Changelog

## [0.4.0] — 2026-06-09

### Added
- Stage 4: censor taboo words in full sentences — strips trailing punctuation for matching, preserves it in output

## [0.3.0] — 2026-06-09

### Added
- Stage 3: censor taboo words in a word-by-word input loop; replaces each match with `*` × len(word); exits on `exit`

## [0.2.0] — 2026-06-09

### Added
- Stage 2: case-insensitive taboo word lookup — reads one word from stdin, prints `True`/`False`

## [0.1.0] — 2026-06-09

### Added
- Stage 1: reads a filename from stdin, opens the taboo words file, and prints each word on its own line
- Whitelist `.gitignore` tracking only Go source files
