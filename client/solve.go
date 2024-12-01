package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

const defaultRandomNonce = 1000000

type Task struct {
	version       int
	difficulty    int
	collapsedTine string
	requestHash   string
}

func solve(task string) string {
	taskParams := strings.Split(task, ":")
	//
	//if len(taskParams) == 4 {
	//	return "" // invalid task format
	//}
	//
	//t := Task{
	//	collapsedTine: taskParams[3],
	//	requestHash:   taskParams[3],
	//}
	//
	//ver, err := strconv.Atoi(taskParams[0])
	//if err != nil {
	//	return "" // invalid task format
	//}

	difficulty, err := strconv.Atoi(taskParams[1])
	if err != nil {
		return "" // invalid task format
	}

	nonce := findNonce(task, difficulty)

	return composeResponse(task, nonce)

}

func findNonce(task string, difficulty int) int {
	nonce := rand.Intn(defaultRandomNonce)

	for {
		h := sha256.New()
		h.Write([]byte(task + strconv.Itoa(nonce)))
		hash := hex.EncodeToString(h.Sum(nil))

		if hash[:difficulty] == strings.Repeat("0", difficulty) {
			fmt.Println(hash)

			return nonce
		}

		h.Reset()
		nonce++
	}
}

func composeResponse(task string, nonce int) string {
	//h := sha256.New()
	//h.Write([]byte(fmt.Sprintf("%s:%s", task, strconv.Itoa(nonce))))
	//fmt.Println(hex.EncodeToString(h.Sum(nil)))
	return fmt.Sprintf("%s:%s", task, strconv.Itoa(nonce))
}
