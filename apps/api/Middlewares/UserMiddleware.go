package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/RajanDhamala/puzzleSmith/Types"
	"github.com/RajanDhamala/puzzleSmith/Utils"
)

type contextKey string

const userContextKey contextKey = "user"

func UserAuthenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessCookie, _ := r.Cookie("accessToken")
		refreshCookie, _ := r.Cookie("refreshToken")

		if accessCookie != nil && accessCookie.Value != "" {
			data, err := utils.VerifyToken(accessCookie.Value)
			if err == nil {
				ctx := context.WithValue(r.Context(), userContextKey, data)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			fmt.Println("Access token invalid or expired:", err)
		}

		if refreshCookie != nil && refreshCookie.Value != "" {
			data, err := utils.VerifyToken(refreshCookie.Value)
			if err == nil {
				usr := types.JwtObj{
					ID:       data.ID,
					Fullname: data.Fullname,
					Email:    data.Email,
				}

				newAccessToken, err := utils.CreateToken(
					usr,
					"accessToken",
					15*time.Minute,
				)
				if err != nil {
					writeJSON(w, http.StatusInternalServerError, map[string]string{
						"error": "Failed to generate new access token",
					})
					return
				}

				http.SetCookie(w, &http.Cookie{
					Name:     "accessToken",
					Value:    newAccessToken,
					Path:     "/",
					Expires:  time.Now().Add(15 * time.Minute),
					HttpOnly: true,
					Secure:   false, // true in production
					SameSite: http.SameSiteLaxMode,
				})

				ctx := context.WithValue(r.Context(), userContextKey, data)

				fmt.Println("Access token refreshed via refresh token")

				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			fmt.Println("Refresh token invalid or expired:", err)
		}

		writeJSON(w, http.StatusUnauthorized, map[string]string{
			"error": "No valid token, login required",
		})
	})
}
