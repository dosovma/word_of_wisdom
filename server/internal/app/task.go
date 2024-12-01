package app

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

const (
	V1 int = 1
)

const (
	masterKey        = 2
	difficulty       = 5
	collapsedTimeout = 24 * 60 * 60 // 24 часа
)

const (
	// masterKey : difficulty : requestID : requestTime : collapsedTime
	signatureRule = "%d:%d:%d:%d:%d"

	// version : difficulty : requestID : requestTime : collapsedTime : requestSignature
	taskRule = "%d:%d:%d:%d:%d:%s"
)

func TaskGenerator(requestID int64, requestTime int64) string {
	s, collapsedTime := signature(requestID, requestTime, difficulty)

	return fmt.Sprintf(taskRule, V1, difficulty, requestID, requestTime, collapsedTime, s)
}

func signature(requestID int64, requestTime int64, difficulty int) (string, int64) {
	hash := sha256.New()
	collapsedTime := requestTime + collapsedTimeout
	s := fmt.Sprintf(signatureRule, masterKey, requestID, requestTime, difficulty, collapsedTime)
	hash.Write([]byte(s))

	return hex.EncodeToString(hash.Sum(nil)), collapsedTime
}
