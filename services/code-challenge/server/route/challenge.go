package route

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type challengeResponse struct {
	Message string `json:"message"`
}

// ChallengeRoutes for the challenges path
func ChallengeRoutes(r *mux.Router) {
	r.HandleFunc("/challenges", use(getChallenges, verifyJWT)).Methods("GET")
}

func getChallenges(w http.ResponseWriter, r *http.Request) {

	encode(w, challengeResponse{"Route: '/challenges' accessed"})
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func verifyJWT(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			fmt.Printf("Request token: %v\n", s)
			http.Error(w, "Not authorized", 401)
			return
		}

		// sample token string taken from the New example
		tokenString := s[1]

		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte("msp_secret"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("Valid token: %v\n", claims)
			h.ServeHTTP(w, r)
		} else {
			fmt.Printf("Error parsing token: %v\n", err)
			http.Error(w, "Not authorized", 401)
			return
		}
	}
}
