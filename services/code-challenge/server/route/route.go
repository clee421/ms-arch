package route

import "github.com/gorilla/mux"

// Handler will take care of all the routing
func Handler(r *mux.Router) {

	r.Use(mux.CORSMethodMiddleware(r))

	// All routes will be under /api
	s := r.PathPrefix("/api").Subrouter()

	// Routes for Authentication
	AuthRoutes(s)

	// Routes for Challenges
	ChallengeRoutes(s)

	// The catch all route - KEEP THIS LAST
	DefaultRoute(r)
}
