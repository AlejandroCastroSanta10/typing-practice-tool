## ADDED Requirements

### Requirement: Application entry point
The application SHALL start by displaying the main menu when launched with no arguments.

#### Scenario: Launch application
- **WHEN** user runs the binary with no arguments
- **THEN** the main menu screen is displayed in the terminal

### Requirement: Main menu navigation
The main menu SHALL display options for Tutorial Mode, Free Practice, and Quit. The user SHALL navigate options using arrow keys and select with Enter.

#### Scenario: Navigate and select tutorial
- **WHEN** the main menu is displayed
- **THEN** the user sees three options: "Tutorial", "Free Practice", and "Quit"

#### Scenario: Select menu option with Enter
- **WHEN** user highlights an option and presses Enter
- **THEN** the application transitions to the selected screen

#### Scenario: Quit from main menu
- **WHEN** user selects "Quit" or presses q
- **THEN** the application exits cleanly and restores the terminal state

### Requirement: Screen navigation with Escape
The user SHALL be able to return to the main menu from any mode by pressing Escape.

#### Scenario: Return to menu from tutorial
- **WHEN** user presses Escape during a tutorial lesson
- **THEN** the application returns to the main menu

#### Scenario: Return to menu from free practice
- **WHEN** user presses Escape during free practice
- **THEN** the application returns to the main menu

### Requirement: Terminal state management
The application SHALL enter Bubble Tea's alternate screen mode on start and restore the original terminal state on exit, including on unexpected termination via Ctrl+C.

#### Scenario: Clean exit on Ctrl+C
- **WHEN** user presses Ctrl+C at any point
- **THEN** the terminal state is restored and the application exits

### Requirement: Responsive layout
The application SHALL adapt its layout to the current terminal dimensions and handle terminal resize events.

#### Scenario: Terminal resize
- **WHEN** the terminal is resized while the application is running
- **THEN** the UI re-renders to fit the new dimensions
