package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"mfv_test/api/models"
	"mfv_test/api/responses"
	"mfv_test/api/utils/formaterror"

	"github.com/gorilla/mux"
)

func (server *Server) CreateUserTransaction(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.UserTransaction{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	postCreated, err := post.SavePost(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	responses.JSON(w, http.StatusCreated, postCreated)
}

func (server *Server) GetUserTransactionByUserId(w http.ResponseWriter, r *http.Request) {

	post := models.UserTransaction{}
	vars := mux.Vars(r)

	userID, err := strconv.ParseUint(vars["user_id"], 10, 64)
	userAccountID, err := strconv.ParseUint(r.FormValue("account_id"), 10, 64)
	if userID == 0 {
		err = fmt.Errorf("user_id is not null")
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	var userTransaction *[]models.UserTransactionReponse
	if userAccountID != 0 {
		userTransaction, err = post.FindUserTransactionByUserIdUserTransactionId(server.DB, userID, userAccountID)
	} else {
		userTransaction, err = post.FindUserTransactionByUserId(server.DB, userID)
	}
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, userTransaction)
}

func (server *Server) GetTransactionByUserIdUserTransactionId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["user_id"], 10, 64)
	pidAccount, err := strconv.ParseUint(vars["account_id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	post := models.UserTransaction{}
	postReceived, err := post.FindUserTransactionByUserIdUserTransactionId(server.DB, pid, pidAccount)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, postReceived)
}
