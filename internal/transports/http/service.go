package http

import (
	"encoding/json"
	"github.com/fwidjaya20/demo-go-grpc-client/internal"
	"github.com/fwidjaya20/demo-go-grpc-client/internal/models"
	"github.com/fwidjaya20/demo-go-grpc-server/pkg/atm"
	"net/http"
)

func Transfer(transferService internal.ServiceInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload models.Transfer

		err := json.NewDecoder(r.Body).Decode(&payload)

		if nil != err {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		protoPayload := &atm.Transfer{
			Sender:    payload.Sender,
			Recipient: payload.Recipient,
			Amount:    payload.Amount,
		}

		result := transferService.SendMoney(protoPayload)

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(result)
	}
}