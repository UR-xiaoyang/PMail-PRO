package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Jinnrry/pmail/config"
	"github.com/Jinnrry/pmail/db"
	"github.com/Jinnrry/pmail/dto/response"
	"github.com/Jinnrry/pmail/i18n"
	"github.com/Jinnrry/pmail/models"
	"github.com/Jinnrry/pmail/session"
	"github.com/Jinnrry/pmail/utils/array"
	"github.com/Jinnrry/pmail/utils/context"
	"github.com/Jinnrry/pmail/utils/errors"
	"github.com/Jinnrry/pmail/utils/password"
	log "github.com/sirupsen/logrus"
)

type loginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func Login(ctx *context.Context, w http.ResponseWriter, req *http.Request) {

	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		log.Errorf("%+v", err)
	}
	var reqData loginRequest
	err = json.Unmarshal(reqBytes, &reqData)
	if err != nil {
		log.Errorf("%+v", err)
	}

	var user models.User

	// 1. Get user by account first
	_, err = db.Instance.Where("account =? and disabled=0", reqData.Account).Get(&user)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Errorf("%+v", err)
	}

	// Prevent timing attacks: simulate work if user not found
	if user.ID == 0 {
		_, _ = password.Hash(reqData.Password)
	}

	// 2. Verify password
	if user.ID != 0 && password.Verify(reqData.Password, user.Password) {
		// Optional: Upgrade to Bcrypt if legacy
		if len(user.Password) == 32 {
			newHash, err := password.Hash(reqData.Password)
			if err == nil {
				user.Password = newHash
				db.Instance.ID(user.ID).Cols("password").Update(&user)
			}
		}

		userStr, _ := json.Marshal(user)
		session.Instance.Put(req.Context(), "user", string(userStr))

		domains := config.Instance.Domains
		domains = array.Difference(domains, []string{config.Instance.Domain})
		domains = append([]string{config.Instance.Domain}, domains...)

		response.NewSuccessResponse(map[string]any{
			"account":  user.Account,
			"name":     user.Name,
			"is_admin": user.IsAdmin,
			"domains":  domains,
		}).FPrint(w)
	} else {
		response.NewErrorResponse(response.ParamsError, i18n.GetText(ctx.Lang, "aperror"), "").FPrint(w)
	}
}

func Logout(ctx *context.Context, w http.ResponseWriter, req *http.Request) {
	session.Instance.Clear(ctx.Context)
	response.NewSuccessResponse("Success").FPrint(w)
}
