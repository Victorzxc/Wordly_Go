package word

import (
	"fmt"
	"mygame/pkg/constants"
	"mygame/pkg/letter"
)

type Word struct {
	Letters []*letter.Letter
}

// Equals проверяет, совпадает ли слово с другим словом.
func (w *Word) Equals(other *Word) bool {
	for i, letter := range w.Letters {
		if letter.Char != other.Letters[i].Char {
			return false
		}
	}
	return true
}

// ChangeColor изменяет цвет буквы в слове по индексу.
func (w *Word) ChangeColor(index int, color string) {
	w.Letters[index].Color = color
}

// Print выводит слово в консоль.
func (w *Word) Print() {
	for _, letter := range w.Letters {
		fmt.Print(letter)
	}
	fmt.Println()
}

func NewWord(word string) *Word {
	w := &Word{}
	for _, char := range word {
		w.Letters = append(w.Letters, letter.NewLetter(char, constants.Gray))
	}
	return w
}
