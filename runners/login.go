package runners

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/codecyclist/kp-cli/cmd"
	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

func Login(appConfig cmd.AppConfig) {
	ctx := context.Background()
	fmt.Println("AX")

	provider, err := oidc.NewProvider(ctx, "http://paste.localhost:4200/auth/realms/koalapaste")
	if err != nil {
		log.Fatal("f:", err)
	}
	oidcConfig := &oidc.Config{
		ClientID: "kp-cli",
	}
	verifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     "kp-cli",
		ClientSecret: "72c641fb-96ca-4fe6-afec-13dd77396a83",
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:5556/auth/google/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	state := "foobar" // Don't do this in production.

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, config.AuthCodeURL(state), http.StatusFound)
	})

	http.HandleFunc("/auth/google/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "state did not match", http.StatusBadRequest)
			return
		}

		oauth2Token, err := config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
			return
		}
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		oauth2Token.AccessToken = "*REDACTED*"

		resp := struct {
			OAuth2Token   *oauth2.Token
			IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
		}{oauth2Token, new(json.RawMessage)}

		if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data, err := json.MarshalIndent(resp, "", "    ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})

	log.Printf("listening on http://%s/", "127.0.0.1:5556")
	log.Fatal(http.ListenAndServe("127.0.0.1:5556", nil))
}
