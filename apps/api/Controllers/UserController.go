package Controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/RajanDhamala/puzzleSmith/Types"

	"github.com/RajanDhamala/puzzleSmith/Utils"
	"github.com/RajanDhamala/puzzleSmith/internal/db"
)

type RegisterUserReq struct {
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func (ctrl *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var data RegisterUserReq

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "error parsing body",
		})
		return
	}

	_, err := ctrl.queries.CheckIfusrExists(r.Context(), data.Email)

	if err == nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "user already exists",
		})
		return
	}

	if err != sql.ErrNoRows {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}

	hashedPwd, err := utils.EncryptPaswrod(data.Password)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}

	id, err := ctrl.queries.RegisterUser(r.Context(), db.RegisterUserParams{
		Email:    data.Email,
		Password: hashedPwd,
		Fullname: data.FullName,
	})
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "failed to register user",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"data": map[string]any{
			"id": id,
		},
	})
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (ctrl *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user login called")

	var data LoginReq

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "error parsing body",
		})
		return
	}

	user, err := ctrl.queries.LoginUser(r.Context(), data.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "invalid credentials",
			})
			return
		}

		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "internal server error",
		})
		return
	}

	if err := utils.DecrptPassword(user.Password, data.Password); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid credentials",
		})
		return
	}

	usr := types.JwtObj{
		ID:       strconv.Itoa(int(user.ID)),
		Fullname: user.Fullname,
		Email:    user.Email,
	}

	accessToken, err := utils.CreateToken(
		usr,
		"accessToken",
		15*time.Minute,
	)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to create access token",
		})
		return
	}

	refreshToken, err := utils.CreateToken(
		usr,
		"refreshToken",
		7*24*time.Hour,
	)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "failed to create refresh token",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Path:     "/",
		Expires:  time.Now().Add(15 * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/",
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "logged in successfully",
	})
}

func (ctrl *Controller) LogoutUser(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "logged out",
	})
}
