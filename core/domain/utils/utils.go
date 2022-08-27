package utils

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	DEFAULT_PAGE = 0
	DEFAULT_SIZE = 20
)

func StringToInt(text string, fallback int) int {
	value, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return fallback
	}
	return int(value)
}

func StringToInt64(text string, fallback int64) int64 {
	value, err := strconv.ParseInt(text, 10, 32)
	if err != nil {
		return fallback
	}
	return value
}

func StringToDate(text string, fallback time.Time) time.Time {
	return StringToDateFormat(text, "2006-01-02", fallback)
}

func StringToDateFormat(text string, format string, fallback time.Time) time.Time {
	value, err := time.Parse(format, text)
	if err != nil {
		return fallback
	}
	return value
}

func StringNormalized(raw string, fallback string) string {
	if strings.TrimSpace(raw) == "" {
		return strings.ToUpper(fallback)
	}
	return strings.ToUpper(raw)
}

func WriteJson(data interface{}, w http.ResponseWriter, statusOk int, statusError int) {
	w.Header().Add("Content-type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error().Msg(err.Error())
		w.WriteHeader(statusError)
	}
}
