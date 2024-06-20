package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"mygame/pkg/constants"
	"mygame/pkg/word"
	"os"
	"strings"
)

type Game struct {
	dictionaryFile string
	attempts       []*word.Word
}

// NewGame создает новую игру и загружает словарь.
func NewGame(dictionaryFile string) *Game {
	return &Game{dictionaryFile: dictionaryFile}
}

// Start запускает игру.
func (g *Game) Start() {
	fmt.Print("\033[H\033[2J")
	attemptsCount := 0
	randomWord := g.getRandomWord()

	fmt.Println("\033[1;37mОтгадайте слово из 5 букв.\033")
	for attemptsCount < constants.AttemptsNumber {
		fmt.Printf("\033[1;37mПопытка %d из %d\n", attemptsCount+1, constants.AttemptsNumber)
		fmt.Print("\033[1;37mВведите слово: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		fmt.Print("\033[H\033[2J")

		if !g.isWordValid(input) {
			fmt.Printf("\033[1;31mСлово '%s' не найдено в словаре или не соответствует длине в %d букв!\033\n", input, constants.WordLength)
			g.printAttempts()
			continue
		}

		current := word.NewWord(input)
		g.compare(current, word.NewWord(randomWord))
		g.attempts = append(g.attempts, current)
		g.printAttempts()

		if current.Equals(word.NewWord(randomWord)) {
			fmt.Printf("\033[1;32mПоздравляем, вы отгадали слово: %s\033\n", randomWord)
			break
		}

		attemptsCount++
		if attemptsCount == constants.AttemptsNumber {
			fmt.Printf("\033[1;31mСлово не угадано. Загаданное слово было: %s\033\n", randomWord)
		}
	}
}

// Проверяем на норм слово
func (g *Game) isWordValid(word string) bool {
	file, err := os.Open(g.dictionaryFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == word {
			return true
		}
	}
	return false
}

// Жёсткий рандомайзер слов
func (g *Game) getRandomWord() string {
	var currentWord string

	file, err := os.Open(g.dictionaryFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	randomIndex := rand.Intn(constants.WordQuantity)

	scanner := bufio.NewScanner(file)
	for i := 0; i <= randomIndex; i++ {
		scanner.Scan()
		currentWord = scanner.Text()
	}
	return currentWord
}

// Выводим попытки
func (g *Game) printAttempts() {
	fmt.Println("\033[1;31mВаши попытки:\033")
	for _, attempt := range g.attempts {
		attempt.Print()
	}
	fmt.Println()
}

// compare сравнивает текущее слово с правильным и изменяет цвет букв.
func (g *Game) compare(current, correct *word.Word) {
	usedIndices := make(map[int]bool)

	// Сначала отмечаем зелёные буквы
	for i, letter := range current.Letters {
		if letter.Char == correct.Letters[i].Char {
			current.ChangeColor(i, constants.Green)
			usedIndices[i] = true
		}
	}

	// Затем отмечаем жёлтые буквы
	for i, letter := range current.Letters {
		if current.Letters[i].Color == constants.Gray {
			for j, correctLetter := range correct.Letters {
				if letter.Char == correctLetter.Char && !usedIndices[j] {
					current.ChangeColor(i, constants.Yellow)
					usedIndices[j] = true
					break
				}
			}
		}
	}
}
