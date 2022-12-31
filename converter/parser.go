package converter

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Defenition struct {
	partOfSpeech, defenition string
}

type Entry struct {
	word, pronunciation string
	seeAlso             []string
	defenitions         []Defenition
}

func Convert(glossaryPath string) (map[string]Entry, error) {
	const longestWord = 10

	entries := make(map[string]Entry)

	readFile, err := os.Open(glossaryPath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var currentEntry [2]string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Printf("\n\n\nProcessing line: %s\n", line)

		word, defenition, found := strings.Cut(line, ":")
		word = strings.TrimSpace(word)
		defenition = strings.TrimSpace(defenition)

		if found && strings.Count(word, " ") < longestWord-1 {
			fmt.Printf("Found : seperator\n")
			if len(currentEntry[0]) > 0 {
				fmt.Printf("Found new entry - processing current entry\n")
				entries[currentEntry[0]] = Entry{
					word: strings.TrimSpace(currentEntry[0]),
					defenitions: []Defenition{
						{
							defenition: strings.TrimSpace(currentEntry[1]),
						},
					},
				}
				fmt.Printf("New entry: %s\n", entries[currentEntry[0]])
			}
			currentEntry[0] = word
			currentEntry[1] = defenition
			fmt.Printf("New currentEntry:\n\tword: %s\n\tdefenition: %s\n", currentEntry[0], currentEntry[1])
		} else if found {
			fmt.Printf("Found seperator in defenition\n")
			currentEntry[1] = fmt.Sprintf("%s %s: %s", strings.TrimSpace(currentEntry[1]), strings.TrimSpace(word), strings.TrimSpace(defenition))
			fmt.Printf("Updated currentEntry:\n\tword: %s\n\tdefenition: %s\n", currentEntry[0], currentEntry[1])
		} else {
			fmt.Printf("No : seperator found\n")
			currentEntry[1] = fmt.Sprintf("%s %s", strings.TrimSpace(currentEntry[1]), strings.TrimSpace(word))
			fmt.Printf("Updated currentEntry:\n\tword: %s\n\tdefenition: %s\n", currentEntry[0], currentEntry[1])
		}
		fmt.Printf("Finished processing line\n")
	}

	readFile.Close()

	fmt.Printf("Processing last entry\n")
	entries[currentEntry[0]] = Entry{
		word: strings.TrimSpace(currentEntry[0]),
		defenitions: []Defenition{
			{
				defenition: strings.TrimSpace(currentEntry[1]),
			},
		},
	}
	fmt.Printf("New entry: %s\n", entries[currentEntry[0]])

	fmt.Printf("\n\nFinal Entries:\n %+v\n\n\n", entries)

	return entries, nil
}
