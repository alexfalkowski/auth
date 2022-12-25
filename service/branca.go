package service

import (
	"time"

	"github.com/essentialkaos/branca"
)

func generateBrancaToken(params TokenParams) (string, error) {
	brc, err := branca.NewBranca([]byte(params.Branca))
	if err != nil {
		return "", err
	}

	t := time.Now()
	brc.SetTTL(uint32(t.Add(params.Service.Duration).Unix()))

	return brc.EncodeToString([]byte(params.Service.ID))
}
