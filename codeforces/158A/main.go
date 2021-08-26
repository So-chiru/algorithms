package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// run fails
// wrong answer 1st numbers differ - expected: '5', found: '4'
// only on codeforces, why?

func main() {
	var (
		participants int
		minPlaces    int
		advances     int
	)

	fmt.Scanf("%d %d\n", &participants, &minPlaces)

	reader := bufio.NewReader(os.Stdin)
	scoreString, _ := reader.ReadString('\n')

	var scores []string = strings.Split(strings.ReplaceAll(scoreString, "\n", ""), " ")

	for i := 0; i < participants; i++ {
		score, _ := strconv.Atoi(scores[i])

		if i < minPlaces {
			if score != 0 {
				advances++
				continue
			}
		}

		if i == 0 {
			if score != 0 {
				advances++
				continue
			}

			break
		}

		previousScore, _ := strconv.Atoi(scores[i-1])

		if previousScore == score {
			advances++
			continue
		}

		break
	}

	fmt.Println(advances)
}
