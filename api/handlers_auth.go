package api

import (
	"net/http"
	"time"
)

// @Summary		  User login
// @Description	Provides JWT token for existing user
// @Tags			  Auth
// @Produce   	plain
// @Param	    	email		  query		string	true	"email"
// @Param	    	password	query		string	true	"password"
// @Success   	200				"login successful"
// @Failure   	401			  "invalid credentials"
// @Router			/login [get]
func (s *Api) loginUser(w http.ResponseWriter, r *http.Request) {
	// CWE-598: Use of GET Request Method With Sensitive Query Strings
	// CWE-523: Unprotected Transport of Credentials
	var (
		email    = r.FormValue("email")
		password = r.FormValue("password")
		response string
	)

	user, err := s.db.GetUserByCredentials(email, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		response = err.Error()
	} else {
		// Potential for CWE-284, CWE-1270
		// jwtAuth := jwtauth.New("HS256", []byte(password), nil)

		// CWE-613: Insufficient Session Expiration
		// Token never expires
		_, token, _ := s.jwtAuth.Encode(map[string]interface{}{"user_id": user.Id})
		response = token

		// CWE-614: Sensitive Cookie in HTTPS Session Without 'Secure' Attribute
		// CWE-1004: Sensitive Cookie Without 'HttpOnly' Flag
		http.SetCookie(w, &http.Cookie{
			Name:    "jwt",
			Value:   token,
			Expires: time.Now().Add(time.Hour * 24 * 30),
			Path:    "/",
			// Secure:   true,
			// HttpOnly: true,
		})
	}

	// CWE-778: Insufficient Logging
	// log.Println("Issued JWT token to user id %d", user.Id)

	w.Write([]byte(response))
}

// @Summary		  User registration
// @Description	Registers a user
// @Tags			  Auth
// @Produce	    plain
// @Param		    email		  query		string	true	"email"
// @Param		    password	query		string	true	"password"
// @Success	    200				"registration successful"
// @Failure	    409		  	"email already registered or invalid parameters"
// @Router			/register [get]
func (s *Api) registerUser(w http.ResponseWriter, r *http.Request) {
	// CWE-598: Use of GET Request Method With Sensitive Query Strings
	// CWE-523: Unprotected Transport of Credentials
	var (
		email    = r.FormValue("email")
		password = r.FormValue("password")
		response string
	)

	// CWE-262: Not Using Password Aging
	if err := s.db.AddUser(email, password); err != nil {
		w.WriteHeader(http.StatusConflict)
		response = err.Error()
	} else {
		response = "User successfully registered!"
	}

	w.Write([]byte(response))
}
