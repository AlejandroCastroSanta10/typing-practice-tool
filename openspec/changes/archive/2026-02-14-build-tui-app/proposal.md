## Why

Learning to type correctly requires consistent practice with proper finger placement and technique. A terminal-based typing tutor provides a distraction-free, developer-friendly environment for building muscle memory - accessible anywhere with a terminal, no browser needed.

## What Changes

- Add a Go TUI application using Bubble Tea for terminal UI rendering
- Implement a **tutorial mode** that teaches correct finger placement and progressively introduces keys row by row (home row first, then top/bottom rows, numbers, and symbols)
- Implement a **free practice mode** that displays random words for the user to type, tracking speed (WPM) and accuracy in real time
- Add a main menu to switch between modes
- Include real-time visual feedback: correct/incorrect character highlighting, current WPM, and accuracy percentage

## Capabilities

### New Capabilities

- `tui-core`: Terminal UI framework setup, main menu, screen navigation, and shared rendering components (Bubble Tea model/view/update architecture)
- `tutorial-mode`: Guided typing lessons with progressive key introduction, finger placement guidance, and per-lesson completion criteria
- `free-practice`: Random word generation, real-time typing input with WPM and accuracy tracking, and session summary

### Modified Capabilities

_(none - greenfield project)_

## Impact

- **Dependencies**: Go modules - `charmbracelet/bubbletea`, `charmbracelet/lipgloss` for TUI rendering
- **Code**: Entirely new Go codebase under the project root (`main.go`, `internal/` packages)
- **Build**: Requires Go 1.22+ toolchain
- **Systems**: Terminal only, no external services or APIs
