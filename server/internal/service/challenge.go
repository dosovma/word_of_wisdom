package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"server/internal/service/entity"
)

const (
	v1 int = 1
)

const (
	masterKey  = 2           // TODO set by envs
	difficulty = 5           // TODO set by envs
	timeout    = 1 * 60 * 60 // 1 час // TODO set by envs
)

const (
	// masterKey:requestID:requestTime:difficulty:requestTimeout
	signatureRule = "%d:%d:%d:%d:%d"

	// version:difficulty:requestID:requestTime:requestTimeout:requestSignature
	challengeRule = "%d:%d:%d:%d:%d:%s"
)

func (*Service) Challenge(r entity.Request) string {
	signature, reqTimeout := sign(r.ID, r.CreatedAt, difficulty)

	return fmt.Sprintf(challengeRule, v1, difficulty, r.ID, r.CreatedAt, reqTimeout, signature)
}

func sign(requestID, requestTime int64, difficulty int) (string, int64) {
	hash := sha256.New()
	reqTimeout := requestTime + timeout

	if _, err := fmt.Fprintf(hash, signatureRule, masterKey, requestID, requestTime, difficulty, reqTimeout); err != nil {
		return "", 0 // TODO add error and handling
	}

	return hex.EncodeToString(hash.Sum(nil)), reqTimeout
}
