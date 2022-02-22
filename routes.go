package main

// import(
// 	rf "golang_gauntlet_api/api/main/routeFunctions"
// )

func (s *Server) initializeRoutes() {
	// v1 := s.router.Group("/v1")
	// {
	s.router.GET("/ping", InterestingResponse)
	s.router.GET("/pong", LikingThisSoFar)
	s.router.GET("/ships", s.ShowMeTheCrew)
	// }

}
