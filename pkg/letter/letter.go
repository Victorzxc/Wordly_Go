package letter

import (
	"fmt"
)

// Структура для буквы со стилем отображения.
type Letter struct {
	Char  rune
	Color string
}

// Создание новой буквы с заданным символом и цветом.
func NewLetter(char rune, color string) *Letter {
	return &Letter{Char: char, Color: color}
}

// String возвращает строковое представление буквы с учетом цвета.
func (l *Letter) String() string {
	return fmt.Sprintf(l.Color, string(l.Char))
}
