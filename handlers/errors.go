package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendError(w http.ResponseWriter, title, error, details string, statusCode int) {
	bytesBody, err := json.Marshal(struct {
		Error   string `json:"error"`
		Details string `json:"details"`
		Status  int    `json:"status"`
		Title   string `json:"title"`
	}{
		Title:   title,
		Details: details,
		Status:  statusCode,
		Error:   error,
	})
	if err != nil {
		fmt.Println(err)
	}

	_, err = w.Write(bytesBody)
	if err != nil {
		fmt.Println(err)
	}
}
