package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"server/internal/service/entity"
)

// V1 : difficulty : requestID : requestTime : collapsedTime : requestSignature : nonce
const responseRule = "%d:%d:%d:%d:%d:%s:%d"

type validationSpec struct {
	version       int
	difficulty    int
	requestID     int64
	requestTime   int64
	collapsedTime int64
	signature     string
	nonce         int64
	challenge     string
}

func buildSpec(solution string) (*validationSpec, bool) {
	responseParams := strings.Split(solution, ":")

	v, err := strconv.Atoi(responseParams[0])
	if err != nil {
		return nil, false // invalid task format
	}

	d, err := strconv.Atoi(responseParams[1])
	if err != nil {
		return nil, false // invalid response format
	}

	reqID, err := strconv.ParseInt(responseParams[2], 10, 64)
	if err != nil {
		return nil, false // invalid response format
	}

	reqTime, err := strconv.ParseInt(responseParams[3], 10, 64)
	if err != nil {
		return nil, false // invalid response format
	}

	collapsedTime, err := strconv.ParseInt(responseParams[4], 10, 64)
	if err != nil {
		return nil, false // invalid response format
	}

	nonce, err := strconv.ParseInt(responseParams[6], 10, 64)
	if err != nil {
		return nil, false // invalid response format
	}

	return &validationSpec{
		version:       v,
		difficulty:    d,
		requestID:     reqID,
		requestTime:   reqTime,
		collapsedTime: collapsedTime,
		signature:     responseParams[5],
		nonce:         nonce,
	}, true
}

func (s *Service) Validate(solution string) bool {
	spec, ok := buildSpec(solution)
	if !ok {
		return false
	}

	challenge := s.Challenge(
		entity.Request{
			ID:        spec.requestID,
			CreatedAt: spec.requestTime,
		},
	)

	spec.challenge = challenge

	var validations []func() bool
	switch spec.version {
	case V1:
		validations = []func() bool{
			spec.timeValidator,
			spec.signatureValidator,
			spec.nonceValidator,
		}
	default:
		return false
	}

	for _, fnc := range validations {
		if ok = fnc(); !ok {
			return false
		}
	}

	return true
}

func (vp validationSpec) nonceValidator() bool {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%s%d", vp.challenge, vp.nonce)))
	res := hex.EncodeToString(hash.Sum(nil))

	return res[:vp.difficulty] == strings.Repeat("0", vp.difficulty)
}

func (vp validationSpec) signatureValidator() bool {
	requestSignature, _ := sign(vp.requestID, vp.requestTime, vp.difficulty)
	return vp.signature == requestSignature
}

func (vp validationSpec) timeValidator() bool {
	return time.Now().Unix() < vp.collapsedTime
}
