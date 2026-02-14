## Context

This is a greenfield Go TUI application for typing practice. There is no existing codebase beyond the project scaffold. The user is a developer who wants a terminal-native tool to learn proper touch typing, starting from zero. The app needs two modes: a structured tutorial and a free-form practice mode.

## Goals / Non-Goals

**Goals:**
- Build a responsive, visually clear TUI using Bubble Tea
- Provide a progressive tutorial that teaches home row first, then expands
- Track WPM and accuracy in real time during free practice
- Keep the architecture simple and easy to extend with new lessons or word lists

**Non-Goals:**
- Persistent statistics or progress tracking across sessions (future enhancement)
- Multiplayer or network features
- Custom themes or configuration files
- Support for non-English keyboards or languages

## Decisions

### 1. TUI Framework: Bubble Tea + Lip Gloss

**Choice**: Use `charmbracelet/bubbletea` for the Elm-architecture TUI loop and `charmbracelet/lipgloss` for styling.

**Why over alternatives**:
- `termbox-go` / `tcell`: Lower-level, requires manual event loops and rendering. Bubble Tea provides a clean Model-View-Update pattern that maps well to distinct screens (menu, tutorial, practice).
- `tview`: Widget-oriented, better for form-heavy apps. Our app is more about real-time keystroke capture and custom rendering.

### 2. Application Structure: Single Binary, Internal Packages

**Choice**: One `main.go` entry point with internal packages under `internal/`.

```
main.go
internal/
  ui/          # Bubble Tea models for each screen
    menu.go
    tutorial.go
    practice.go
  typing/      # Core typing logic (input matching, stats calculation)
  lesson/      # Tutorial lesson definitions and progression
  wordlist/    # Word lists for free practice mode
```

**Why**: Keeps the binary simple (`go build .`) while separating concerns. `internal/` prevents accidental external imports.

### 3. Tutorial Progression: Row-Based Lessons

**Choice**: Lessons progress through keyboard rows: home row → top row → bottom row → numbers → symbols. Each lesson focuses on a small set of new keys mixed with previously learned keys.

**Why over alternatives**:
- Random key introduction: No pedagogical structure, overwhelming for beginners.
- Full-word-first approach: Doesn't build proper finger-to-key muscle memory.

Lessons are defined as Go structs with the set of allowed characters and sample exercises. This makes adding new lessons trivial.

### 4. Typing Input Model: Character-by-Character Matching

**Choice**: Compare each keystroke against the expected character at the current position. Display correct characters in green, errors in red, and untyped characters dimmed.

**Why**: Immediate feedback per character helps build accuracy. This is simpler than word-level matching and more educational for beginners who need to see exactly which keys they're hitting wrong.

### 5. WPM Calculation: Standard 5-Character Word

**Choice**: WPM = (total correct characters / 5) / elapsed minutes. This is the standard "gross WPM" formula used by most typing tools.

**Why**: Industry standard, comparable with other tools. Net WPM (subtracting errors) can be added later but gross WPM is simpler and more encouraging for beginners.

### 6. Word List: Embedded Static List

**Choice**: Embed a curated list of common English words (~200-500 words) directly in Go source using `embed`. Free practice mode randomly selects from this list.

**Why over alternatives**:
- External file: Adds deployment complexity, file-not-found errors.
- API-based: Requires network, overkill for this use case.
- The embedded list keeps the tool as a single, self-contained binary.

## Risks / Trade-offs

- **Terminal compatibility**: Not all terminals render colors/styles identically. → Mitigation: Use Lip Gloss adaptive colors and test on common terminals (kitty, alacritty, gnome-terminal).
- **No persistence**: Users lose progress when they close the app. → Acceptable for v1; future enhancement can add JSON-based progress files.
- **Fixed word list**: Free practice may feel repetitive after extended use. → Mitigation: Include enough words (~500) and randomize order. Can expand later.
- **English only**: Non-English users can't benefit from the tutorial. → Explicitly a non-goal for v1; lesson structure supports future i18n.
