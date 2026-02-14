package lesson

import (
	"math/rand"
	"strings"
)

// Lesson defines a typing lesson with a set of allowed keys.
type Lesson struct {
	Name  string
	Group string
	Keys  []rune
}

// FingerName returns the finger used for a given key.
func FingerName(ch rune) string {
	fingers := map[rune]string{
		// Home row
		'a': "Left pinky", 's': "Left ring", 'd': "Left middle", 'f': "Left index",
		'g': "Left index", 'h': "Right index", 'j': "Right index", 'k': "Right middle",
		'l': "Right ring", ';': "Right pinky",
		// Top row
		'q': "Left pinky", 'w': "Left ring", 'e': "Left middle", 'r': "Left index",
		't': "Left index", 'y': "Right index", 'u': "Right index", 'i': "Right middle",
		'o': "Right ring", 'p': "Right pinky",
		// Bottom row
		'z': "Left pinky", 'x': "Left ring", 'c': "Left middle", 'v': "Left index",
		'b': "Left index", 'n': "Right index", 'm': "Right index", ',': "Right middle",
		'.': "Right ring", '/': "Right pinky",
		// Numbers
		'1': "Left pinky", '2': "Left ring", '3': "Left middle", '4': "Left index",
		'5': "Left index", '6': "Right index", '7': "Right index", '8': "Right middle",
		'9': "Right ring", '0': "Right pinky",
		// Space
		' ': "Thumb",
	}
	if name, ok := fingers[ch]; ok {
		return name
	}
	return ""
}

// GenerateExercise creates a random exercise string using only the allowed keys.
// It generates groups of 3-5 characters separated by spaces.
func (l *Lesson) GenerateExercise(length int) string {
	var words []string
	chars := 0
	for chars < length {
		wordLen := 3 + rand.Intn(3) // 3-5 chars
		if chars+wordLen > length {
			wordLen = length - chars
		}
		var word strings.Builder
		for j := 0; j < wordLen; j++ {
			word.WriteRune(l.Keys[rand.Intn(len(l.Keys))])
		}
		words = append(words, word.String())
		chars += wordLen + 1 // +1 for space
	}
	return strings.Join(words, " ")
}

// AllLessons returns the full list of lessons in order.
func AllLessons() []Lesson {
	homeRow := []rune("asdfghjkl;")
	topRow := []rune("qwertyuiop")
	bottomRow := []rune("zxcvbnm,./")
	numbers := []rune("1234567890")

	return []Lesson{
		{Name: "Home Row Basics", Group: "Home Row", Keys: homeRow},
		{Name: "Home Row Words", Group: "Home Row", Keys: append([]rune{' '}, homeRow...)},
		{Name: "Top Row Basics", Group: "Top Row", Keys: topRow},
		{Name: "Top + Home Row", Group: "Top Row", Keys: append(append([]rune{' '}, homeRow...), topRow...)},
		{Name: "Bottom Row Basics", Group: "Bottom Row", Keys: bottomRow},
		{Name: "Bottom + Home Row", Group: "Bottom Row", Keys: append(append([]rune{' '}, homeRow...), bottomRow...)},
		{Name: "All Letters", Group: "Bottom Row", Keys: append(append(append([]rune{' '}, homeRow...), topRow...), bottomRow...)},
		{Name: "Number Row", Group: "Numbers", Keys: numbers},
		{Name: "Numbers + Letters", Group: "Numbers", Keys: append(append(append(append([]rune{' '}, homeRow...), topRow...), bottomRow...), numbers...)},
	}
}
