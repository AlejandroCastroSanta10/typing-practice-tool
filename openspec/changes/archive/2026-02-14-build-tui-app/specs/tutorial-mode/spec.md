## ADDED Requirements

### Requirement: Lesson selection menu
Tutorial mode SHALL display a list of available lessons grouped by keyboard row. The user SHALL select a lesson using arrow keys and Enter.

#### Scenario: Display lesson list
- **WHEN** user enters tutorial mode
- **THEN** a list of lessons is displayed, grouped as: Home Row, Top Row, Bottom Row, Numbers, Symbols

#### Scenario: Select a lesson
- **WHEN** user highlights a lesson and presses Enter
- **THEN** the selected lesson starts and the typing exercise is displayed

### Requirement: Progressive key introduction
Each lesson SHALL introduce a small set of new keys while reinforcing previously learned keys. The home row lesson SHALL be the first lesson and introduce keys: a, s, d, f, g, h, j, k, l, ;.

#### Scenario: Home row lesson content
- **WHEN** user starts the home row lesson
- **THEN** the exercise text contains only characters from the set: a, s, d, f, g, h, j, k, l, ;, and space

#### Scenario: Top row lesson includes home row keys
- **WHEN** user starts a top row lesson
- **THEN** the exercise text contains top row keys mixed with home row keys

### Requirement: Finger placement guidance
Each lesson SHALL display which finger to use for the currently expected key. The guidance SHALL appear above the typing area.

#### Scenario: Show finger hint
- **WHEN** user is expected to type the letter "f"
- **THEN** the guidance text indicates "Left index finger"

#### Scenario: Update hint per character
- **WHEN** the expected character changes as the user types
- **THEN** the finger placement guidance updates to match the new expected character

### Requirement: Real-time character feedback
Each typed character SHALL be highlighted immediately: green for correct, red for incorrect. The cursor SHALL advance to the next character after each keystroke.

#### Scenario: Correct keystroke
- **WHEN** user types the correct expected character
- **THEN** the character is displayed in green and the cursor advances

#### Scenario: Incorrect keystroke
- **WHEN** user types an incorrect character
- **THEN** the expected character is displayed in red and the cursor advances

### Requirement: Lesson completion
A lesson SHALL complete when the user has typed all characters in the exercise. Upon completion, the application SHALL display accuracy percentage and prompt to retry or return to the lesson list.

#### Scenario: Finish a lesson
- **WHEN** user types the last character of the exercise
- **THEN** the lesson ends and a summary is displayed showing accuracy percentage

#### Scenario: Retry or return after completion
- **WHEN** the lesson completion summary is displayed
- **THEN** user can press Enter to retry the same lesson or Escape to return to the lesson list
