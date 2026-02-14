## 1. Project Setup

- [x] 1.1 Initialize Go module (`go mod init`), add Bubble Tea and Lip Gloss dependencies
- [x] 1.2 Create project directory structure: `main.go`, `internal/ui/`, `internal/typing/`, `internal/lesson/`, `internal/wordlist/`

## 2. TUI Core

- [x] 2.1 Implement root Bubble Tea model in `main.go` with alternate screen mode and Ctrl+C handling
- [x] 2.2 Implement main menu model (`internal/ui/menu.go`) with Tutorial, Free Practice, and Quit options using arrow key navigation and Enter to select
- [x] 2.3 Implement screen switching: main menu dispatches to tutorial/practice models, Escape returns to menu
- [x] 2.4 Handle terminal resize events and pass window size to child models

## 3. Typing Engine

- [x] 3.1 Implement character-by-character input matcher (`internal/typing/`) that tracks position, correct/incorrect per character
- [x] 3.2 Implement WPM calculation: (correct chars / 5) / elapsed minutes, timer starts on first keystroke
- [x] 3.3 Implement accuracy calculation: (correct / total typed) * 100
- [x] 3.4 Implement styled text rendering: green for correct, red for incorrect, dimmed for untyped characters

## 4. Tutorial Mode

- [x] 4.1 Define lesson data structures and lesson list (`internal/lesson/`): home row, top row, bottom row, numbers, symbols
- [x] 4.2 Implement exercise text generator that produces random character sequences from the allowed key set for each lesson
- [x] 4.3 Implement finger placement mapping: each key maps to a finger name (e.g., "f" → "Left index finger")
- [x] 4.4 Implement tutorial UI model (`internal/ui/tutorial.go`): lesson selection list, grouped by keyboard row
- [x] 4.5 Implement tutorial typing screen: display exercise text, finger hint above typing area, real-time green/red character feedback
- [x] 4.6 Implement lesson completion: show accuracy percentage, Enter to retry, Escape to return to lesson list

## 5. Free Practice Mode

- [x] 5.1 Create embedded word list (`internal/wordlist/`): ~500 common English words using Go `embed`
- [x] 5.2 Implement random word sequence generator: shuffle and concatenate words with spaces, wrap to fill visible area
- [x] 5.3 Implement free practice UI model (`internal/ui/practice.go`): display word block with live character feedback, WPM, and accuracy
- [x] 5.4 Implement continuous word flow: generate new words when user finishes the current set
- [x] 5.5 Implement session summary on Escape: show final WPM, accuracy, total characters typed, any key to return to menu

## 6. Polish and Testing

- [x] 6.1 Add Lip Gloss styles: consistent color scheme, borders, padding for all screens
- [ ] 6.2 Manual end-to-end testing: verify all navigation flows, tutorial lessons, free practice, and edge cases (empty input, rapid typing, resize)
