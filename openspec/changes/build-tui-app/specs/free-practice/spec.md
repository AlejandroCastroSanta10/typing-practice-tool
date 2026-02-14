## ADDED Requirements

### Requirement: Random word generation
Free practice mode SHALL display a sequence of random words selected from an embedded word list. Words SHALL be separated by spaces and wrap across multiple lines to fill the visible area.

#### Scenario: Display random words on start
- **WHEN** user enters free practice mode
- **THEN** a block of random words is displayed for typing

#### Scenario: Words are randomized each session
- **WHEN** user enters free practice mode multiple times
- **THEN** the word order differs between sessions

### Requirement: Real-time typing input
Each typed character SHALL be compared against the expected character at the current position. Correct characters SHALL display in green, incorrect in red, and untyped characters in a dimmed style.

#### Scenario: Correct character typed
- **WHEN** user types the correct expected character
- **THEN** the character turns green and the cursor advances

#### Scenario: Incorrect character typed
- **WHEN** user types an incorrect character
- **THEN** the character turns red and the cursor advances

#### Scenario: Untyped characters remain dimmed
- **WHEN** the exercise is in progress
- **THEN** characters not yet reached by the cursor are displayed in a dimmed color

### Requirement: Live WPM display
The application SHALL calculate and display the current WPM continuously during typing. WPM SHALL be calculated as (total correct characters / 5) / elapsed minutes. The timer SHALL start on the first keystroke.

#### Scenario: WPM updates during typing
- **WHEN** user is typing in free practice mode
- **THEN** the current WPM is displayed and updates after each keystroke

#### Scenario: Timer starts on first keystroke
- **WHEN** user has not typed any character yet
- **THEN** the WPM display shows 0 and the timer has not started

#### Scenario: First keystroke starts the timer
- **WHEN** user types the first character
- **THEN** the timer begins and WPM starts calculating

### Requirement: Live accuracy display
The application SHALL display accuracy as a percentage: (correct characters / total characters typed) * 100. Accuracy SHALL update after each keystroke.

#### Scenario: Accuracy display during typing
- **WHEN** user has typed 8 correct characters and 2 incorrect characters
- **THEN** the accuracy display shows 80%

#### Scenario: Initial accuracy display
- **WHEN** no characters have been typed yet
- **THEN** the accuracy display shows 100%

### Requirement: Continuous word flow
When the user finishes typing all displayed words, the application SHALL generate a new set of random words and continue the session without interruption.

#### Scenario: Words replenish automatically
- **WHEN** user finishes typing the last displayed word
- **THEN** a new set of random words appears and typing continues seamlessly

### Requirement: Session summary on exit
When the user presses Escape to leave free practice, the application SHALL display a brief summary showing final WPM, accuracy, and total characters typed before returning to the main menu.

#### Scenario: Display summary on Escape
- **WHEN** user presses Escape during free practice
- **THEN** a summary is shown with final WPM, accuracy percentage, and total characters typed
- **THEN** pressing any key returns to the main menu
