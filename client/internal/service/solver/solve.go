package solver

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

	"client/pkg/logger"
)

const defaultRandomNonce = 1000000

var ErrInvalidFormat = errors.New("invalid challenge format")

type Solv struct {
	log logger.Logger
}

func New(log logger.Logger) *Solv {
	return &Solv{log: log}
}

func (*Solv) Solve(challenge string) (string, error) {
	challengeParams := strings.Split(challenge, ":") // version:difficulty:requestID:requestTime:requestTimeout:requestSignature

	difficulty, err := strconv.Atoi(challengeParams[1])
	if err != nil {
		return "", ErrInvalidFormat // invalid challenge format
	}

	timeout, err := strconv.ParseInt(challengeParams[4], 10, 64)
	if err != nil {
		return "", ErrInvalidFormat
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Unix(timeout, 0))
	defer cancel()

	resCh := make(chan int64, 1)

	go func() {
		i := findNonce(challenge, difficulty)
		resCh <- i
		close(resCh)
	}()

	select {
	case res := <-resCh:
		return buildSolution(challenge, res), nil
	case <-ctx.Done():
		return "", ctx.Err() // reached request timeout
	}
}

func findNonce(challenge string, difficulty int) int64 {
	nonce := rand.Intn(defaultRandomNonce)

	for {
		h := sha256.New()
		h.Write([]byte(challenge + strconv.Itoa(nonce)))
		hash := hex.EncodeToString(h.Sum(nil))

		if hash[:difficulty] == strings.Repeat("0", difficulty) {
			return int64(nonce)
		}

		h.Reset()
		nonce++
	}
}

func buildSolution(challenge string, nonce int64) string {
	return fmt.Sprintf("%s:%d", challenge, nonce)
}
