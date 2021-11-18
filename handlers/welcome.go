package handlers

import (
	"fmt"
	"net/http"

	"github.com/vouch/vouch-proxy/pkg/cfg"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Debug("/welcome")

	loginURL := fmt.Sprintf("%s/login?%s", cfg.Cfg.DocumentRoot, r.URL.RawQuery)

	fmt.Fprint(w, "<p>Welcome to the DESI test auth application!</p>\n")
	fmt.Fprint(w, "<p>You are not currently logged in.</p>\n")
	fmt.Fprintf(w, "<p><a href=\"%s\">Login</a>\n</p>", loginURL)

	fmt.Fprintf(w, "<p>path: %s \n</p>", r.URL.EscapedPath())
	fmt.Fprintf(w, "<p>query: %s \n</p>", r.URL.RawQuery)
	fmt.Fprintf(w, "<p>login url: %s\n</p>", loginURL)
}
