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

type validationSpec struct {
	version        int
	difficulty     int
	requestID      int64
	requestTime    int64
	requestTimeout int64
	signature      string
	nonce          int64
	challenge      string
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
	case v1:
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

func buildSpec(solution string) (*validationSpec, bool) {
	solutionParams := strings.Split(solution, ":") // version:difficulty:requestID:requestTime:requestTimeout:requestSignature:nonce

	v, err := strconv.Atoi(solutionParams[0])
	if err != nil {
		return nil, false
	}

	d, err := strconv.Atoi(solutionParams[1])
	if err != nil {
		return nil, false
	}

	reqID, err := strconv.ParseInt(solutionParams[2], 10, 64)
	if err != nil {
		return nil, false
	}

	reqTime, err := strconv.ParseInt(solutionParams[3], 10, 64)
	if err != nil {
		return nil, false
	}

	reqTimeout, err := strconv.ParseInt(solutionParams[4], 10, 64)
	if err != nil {
		return nil, false
	}

	nonce, err := strconv.ParseInt(solutionParams[6], 10, 64)
	if err != nil {
		return nil, false
	}

	return &validationSpec{
		version:        v,
		difficulty:     d,
		requestID:      reqID,
		requestTime:    reqTime,
		requestTimeout: reqTimeout,
		signature:      solutionParams[5],
		nonce:          nonce,
	}, true
}

func (vp *validationSpec) nonceValidator() bool {
	hash := sha256.New()
	if _, err := fmt.Fprintf(hash, "%s%d", vp.challenge, vp.nonce); err != nil {
		return false // TODO add error
	}
	res := hex.EncodeToString(hash.Sum(nil))

	return res[:vp.difficulty] == strings.Repeat("0", vp.difficulty)
}

func (vp *validationSpec) signatureValidator() bool {
	requestSignature, _ := sign(vp.requestID, vp.requestTime, vp.difficulty)
	return vp.signature == requestSignature
}

func (vp *validationSpec) timeValidator() bool {
	return time.Now().Unix() < vp.requestTimeout
}
