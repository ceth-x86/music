package spotify

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"

	settings2 "github.com/demas/music/internal/services/settings"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

const (
	redirectURI = "http://localhost:10028/callback"
)

func login(auth spotify.Authenticator) error {
	state, err := generateRandomString(32)
	if err != nil {
		return err
	}

	ch := make(chan *oauth2.Token)

	http.Handle("/callback", &authHandler{state: state, ch: ch, auth: auth})
	go http.ListenAndServe("localhost:10028", nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	tok := <-ch

	if err := saveToken(tok); err != nil {
		return err
	}
	return nil
}

func createClient() spotify.Client {

	auth := spotify.NewAuthenticator(
		redirectURI,
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserModifyPlaybackState,
	)

	settings := settings2.InitSettings()
	auth.SetAuthInfo(settings.SpotifyClientId, settings.SpotifyClientSecret)

	token, err := readToken()
	if err != nil {
		if os.IsNotExist(err) {
			if err := login(auth); err != nil {
				log.Fatal(err)
			}

			// read token one more time
			token, err = readToken()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	return auth.NewClient(token)
}

func tokenPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(usr.HomeDir, ".music")
}

func readToken() (*oauth2.Token, error) {
	content, err := ioutil.ReadFile(tokenPath())
	if err != nil {
		return nil, err
	}

	var tok oauth2.Token
	if err := json.Unmarshal(content, &tok); err != nil {
		return nil, err
	}

	return &tok, nil
}

func saveToken(tok *oauth2.Token) error {

	f, err := os.OpenFile(tokenPath(), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(tok)
}

type authHandler struct {
	state string
	ch    chan *oauth2.Token
	auth  spotify.Authenticator
}

func (a *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tok, err := a.auth.Token(a.state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	if st := r.FormValue("state"); st != a.state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, a.state)
	}

	fmt.Fprintf(w, "Login successfully. Please return to your terminal.")

	a.ch <- tok
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
