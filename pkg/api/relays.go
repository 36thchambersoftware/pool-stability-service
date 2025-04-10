package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pool-stability-service/pkg/blockfrost"
)

func GetPoolRelays(w http.ResponseWriter, r *http.Request) {
	poolID := "pool1fl0gdfcmzg4jf3hw839epdteuazm23nff44vddap00r0sm5gvvk"
	relays, err := blockfrost.PoolRelays(r.Context(), poolID)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("I couldn't find that pool's relays: %s", err)))
		return
	}

	b, err := json.Marshal(&relays)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("I couldn't convert the relays to json: %s", err)))
	}

	w.Write(b)
	return
}