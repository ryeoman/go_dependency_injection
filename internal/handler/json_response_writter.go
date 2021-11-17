package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// WriteJSONResponse to help write json response for http request
func WriteJSONResponse(ctx context.Context, w http.ResponseWriter, data interface{}, err error, status int) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("content-type", "application/json")
	resp := Response{
		Status:  status,
		Message: "Success",
		Data:    data,
	}
	if err != nil {
		resp.Message = err.Error()

		// additional wired behavior logging error in writing json function
		logrus.Println(err.Error())
	}
	errWrite := json.NewEncoder(w).Encode(resp)
	if errWrite != nil {
		logrus.Println(errWrite.Error())
	}
}
