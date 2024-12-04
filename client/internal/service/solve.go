package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const defaultRandomNonce = 1000000

var ErrInvalidTaskFormat = errors.New("invalid task format")

// version : difficulty : requestID : requestTime : collapsedTime : requestSignature
const taskRule = "%d:%d:%d:%d:%d:%s"

func solve(task string) (string, error) {
	taskParams := strings.Split(task, ":")

	difficulty, err := strconv.Atoi(taskParams[1])
	if err != nil {
		return "", ErrInvalidTaskFormat // invalid task format
	}

	collapsedTime, err := strconv.ParseInt(taskParams[4], 10, 64)
	if err != nil {
		return "", ErrInvalidTaskFormat // invalid response format
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(collapsedTime, 0))
	defer cancel()

	resCh := make(chan int64, 1)

	go func() {
		i := findNonce(task, difficulty)
		resCh <- i
		close(resCh)
	}()

	select {
	case res := <-resCh:
		return composeResponse(task, res), nil
	case <-ctx.Done():
		return "", ctx.Err() // collapsedTime
	}
}

func findNonce(task string, difficulty int) int64 {
	nonce := rand.Intn(defaultRandomNonce)

	for {
		h := sha256.New()
		h.Write([]byte(task + strconv.Itoa(nonce)))
		hash := hex.EncodeToString(h.Sum(nil))

		if hash[:difficulty] == strings.Repeat("0", difficulty) {
			fmt.Println(hash)

			return int64(nonce)
		}

		h.Reset()
		nonce++
	}
}

func composeResponse(task string, nonce int64) string {
	return fmt.Sprintf("%s:%d", task, nonce)
}
