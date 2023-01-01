package converter

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

		if found && strings.Count(strings.Trim(word, " "), " ") > longestWord-1 {
			word = fmt.Sprintf("%s:%s", word, defenition)
			found = false
		}

		if found {
			fmt.Printf("Found : seperator\n")
			if len(currentEntry[0]) > 0 {
				fmt.Printf("Found new entry - processing current entry\n")
				entry := createEntry(currentEntry[0], currentEntry[1])
				entries[entry.word] = entry
				fmt.Printf("New entry: %s\n", entries[currentEntry[0]])
			}
			currentEntry[0] = word
			currentEntry[1] = defenition
			fmt.Printf("New currentEntry:\n\tword: %s\n\tdefenition: %s\n", currentEntry[0], currentEntry[1])
		} else {
			fmt.Printf("No : seperator found\n")
			currentEntry[1] = fmt.Sprintf("%s%s%s", currentEntry[1], fmt.Sprintln(""), word)
			fmt.Printf("Updated currentEntry:\n\tword: %s\n\tdefenition: %s\n", currentEntry[0], currentEntry[1])
		}
		fmt.Printf("Finished processing line\n")
	}

	readFile.Close()

	fmt.Printf("Processing last entry\n")
	entry := createEntry(currentEntry[0], currentEntry[1])
	entries[entry.word] = entry
	fmt.Printf("New entry: %s\n", entries[currentEntry[0]])

	fmt.Printf("\n\nFinal Entries:\n %+v\n\n\n", entries)

	return entries, nil
}

func createEntry(key string, value string) Entry {
	pronunciation := ""
	word := key
	wordPronunciationRe := regexp.MustCompile(`(.*)\((.*)\)`)
	pronunciationMatch := wordPronunciationRe.FindStringSubmatch(key)

	var defenitions []Defenition

	for _, defenition := range regexp.MustCompile(`(?:\(\d+\))`).Split(value, -1) {
		defenition = strings.TrimSpace(defenition)
		if len(defenition) > 0 {
			defenitions = append(defenitions, Defenition{defenition: defenition})
		}
	}

	if len(pronunciationMatch) == 3 {
		word = pronunciationMatch[1]
		pronunciation = pronunciationMatch[2]
	}

	entry := Entry{
		word:          strings.TrimSpace(word),
		pronunciation: strings.TrimSpace(pronunciation),
		defenitions:   defenitions,
	}
	return entry
}
