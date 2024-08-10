package statuses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yatter-backend-go/app/domain/auth"

	vo "yatter-backend-go/app/domain/value-object"
)

// Request body for `POST /v1/statuses`
type AddRequest struct {
	Content string
}

// Handle request for `POST /v1/statuses`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	ctx := r.Context()

	account_info := auth.AccountOf(r.Context()) // 認証情報を取得する
	fmt.Println(account_info)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accountID, err := vo.NewAccountID(account_info.ID.Value())
	if err != nil {
		http.Error(w, "Invalid account ID", http.StatusBadRequest)
		return
	}

	dto, err := h.statusUsecase.Create(ctx, *accountID, req.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto.Status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	panic(fmt.Sprintf("Must Implement Status Creation And Check Acount Info %v", account_info))
}
