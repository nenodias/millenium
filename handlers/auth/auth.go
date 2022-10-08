package auth

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/nenodias/millenium/config"
	"github.com/nenodias/millenium/core/domain/auth"
	"github.com/nenodias/millenium/core/domain/utils"
	"github.com/rs/zerolog/log"
)

func Middleware(handler http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authRequest := r.Header.Get("Authorization")
		authRequest = strings.ReplaceAll(authRequest, "Bearer ", "")
		claims, err := auth.Verify(authRequest)
		if err != nil {
			http.Error(w, "please sign-in", http.StatusUnauthorized)
			return
		} else {
			contextWithUser := context.WithValue(r.Context(), auth.AUTH_KEY, claims)
			requestWithUser := r.WithContext(contextWithUser)
			handler(w, requestWithUser)
		}
	}
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	username := config.GetEnv("USER_DEFAULT", "admin")
	password := config.GetEnv("PASS_DEFAULT", "123456")
	bytes := username + ":" + password
	basicAuth := base64.RawStdEncoding.EncodeToString([]byte(bytes))
	basicAuthRequest := r.Header.Get("Authorization")
	basicAuthRequest = strings.ReplaceAll(basicAuthRequest, "Basic ", "")
	if basicAuthRequest == basicAuth {
		token, err := auth.GenerateJWT()
		if err != nil {
			log.Error().Msg(err.Error())
			w.WriteHeader(500)
		} else {
			tokenJson := auth.Token{
				Token: token,
			}
			utils.WriteJson(tokenJson, w, 200, 500)
		}
	} else {
		w.WriteHeader(401)
	}
}
