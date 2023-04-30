package api

import (
	m "govulnapi/models"
	"net/http"
)

// @Summary		  Update email
// @Description	Updates user email
// @Tags		    User
// @Accept	    x-www-form-urlencoded
// @Produce	    plain
// @Param		    email	formData string	true "New email"
// @Success	    200	"email updated"
// @Failure	    400	"bad request"
// @Failure	    401	"unauthorized"
// @Router			/user/email [put]
// @Security		Bearer
func (a *Api) updateEmail(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(m.User)
	newEmail := r.FormValue("email")

	err := a.db.UpdateEmail(user.Id, newEmail)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Email successfully updated!"))
	}
}

// @Summary		  Update password
// @Description	Updates user password
// @Tags		    User
// @Accept	    x-www-form-urlencoded
// @Produce	    plain
// @Param		    password formData string true "New password"
// @Success	    200	"password changed"
// @Failure	    400	"bad request"
// @Failure	    401	"unauthorized"
// @Router			/user/password [put]
// @Security		Bearer
// CWE-549: Missing Password Field Masking
func (a *Api) updatePassword(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(m.User)
	newPassword := r.FormValue("password")

	// CWE-620: Unverified Password Change
	err := a.db.UpdatePassword(user.Id, newPassword)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Password successfully updated!"))
	}
}
