package lib

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

func WriteResponse(w http.ResponseWriter, msg interface{}, statusCode int, logger *logrus.Logger) {
	res, err := json.Marshal(msg)
	if err != nil {
		logger.Errorf("could not marshal response: %s", err)
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		logger.Errorf("could not send response: %s", err)
	}
	return
}
