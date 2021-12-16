package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (h *handler) createToken(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username") // retrieve username from URL
	tok, pk, err := h.CreateToken(username)

	if err != nil {
		respondErr(w, err)
		return
	}

	cookie := http.Cookie{
		Name: "token", // JWT cookie should be named "token"
		Value: tok,
		HttpOnly: true, // should NOT be accessible by client-side JavaScript
		Path: "/", // visible to all paths
	}

	http.SetCookie(w, &cookie)
	// RSA public key in plain text as the response.
	respond(w, pk, http.StatusOK)
}

func (h *handler) verifyToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token") // obtain the jwt from cookie
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized) // If the cookie is not set, return an unauthorized status
			return
		}

		w.WriteHeader(http.StatusBadRequest) // For any other type of error, return a bad request status
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	username, err := h.VerifyToken(tknStr)
	if err != nil {
		respondErr(w, err) // 500 internal error
	}

	respond(w, username, http.StatusOK)
}

func (h *handler) getReadme(w http.ResponseWriter, r *http.Request) {
	readme, err := h.GetReadme()
	if err != nil {
		respondErr(w, err)
	}

	respond(w, readme, http.StatusOK)
}

func (h *handler) getStats(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token") // obtain the jwt from cookie
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized) // If the cookie is not set, return an unauthorized status
			return
		}

		w.WriteHeader(http.StatusBadRequest) // For any other type of error, return a bad request status
		return
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	username, err := h.VerifyToken(tknStr)
	if err != nil {
		respondErr(w, err)
	}

	stats, err := h.GetStats(username) // obtain the jwt from cookie
	if err != nil {
		respondErr(w, err)
	}

	respond(w, stats, http.StatusOK)
}