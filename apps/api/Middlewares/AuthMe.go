package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RajanDhamala/puzzleSmith/Types"
	"github.com/RajanDhamala/puzzleSmith/Utils"
)

func AuthMe(w http.ResponseWriter, r *http.Request) {
	accessCookie, err := r.Cookie("accessToken")

	if err == nil && accessCookie.Value != "" {
		data, err := utils.VerifyToken(accessCookie.Value)
		if err == nil {
			writeJSON(w, http.StatusOK, map[string]string{
				"_id":      data.ID,
				"fullname": data.Fullname,
				"email":    data.Email,
			})
			return
		}
	}

	refreshCookie, err := r.Cookie("refreshToken")

	if err == nil && refreshCookie.Value != "" {
		data, err := utils.VerifyToken(refreshCookie.Value)
		if err == nil {
			newAccessToken, err := utils.CreateToken(
				types.JwtObj{
					ID:       data.ID,
					Fullname: data.Fullname,
					Email:    data.Email,
				},
				"accessToken",
				15*time.Minute,
			)
			if err != nil {
				writeJSON(w, http.StatusInternalServerError, map[string]string{
					"error": "failed to generate new access token",
				})
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "accessToken",
				Value:    newAccessToken,
				Path:     "/",
				Expires:  time.Now().Add(15 * time.Minute),
				HttpOnly: true,
				Secure:   false, // true in production with HTTPS
				SameSite: http.SameSiteLaxMode,
			})

			writeJSON(w, http.StatusOK, map[string]string{
				"_id":      data.ID,
				"fullname": data.Fullname,
				"email":    data.Email,
			})
			return
		}
	}

	writeJSON(w, http.StatusUnauthorized, map[string]string{
		"error": "No valid token, login required",
	})
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}
