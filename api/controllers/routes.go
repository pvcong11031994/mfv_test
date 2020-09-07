package controllers

import "mfv_test/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// CRUD Route
	s.Router.HandleFunc("/api/users/{user_id}/transactions", middlewares.SetMiddlewareJSON(s.GetUserTransactionByUserId)).Methods("GET")
	s.Router.HandleFunc("/api/users/{user_id}/transactions", middlewares.SetMiddlewareJSON(s.CreateUserTransaction)).Methods("POST")
	s.Router.HandleFunc("/api/users/{user_id}/transactions", middlewares.SetMiddlewareJSON(s.UpdateNameUserByUserID)).Methods("PUT")
	s.Router.HandleFunc("/api/users/{user_id}/transactions", middlewares.SetMiddlewareJSON(s.DeleteUserByUserID)).Methods("DELETE")

}
