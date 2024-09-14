package server

import "billing_service/billing/wallet"

func (s *Server) InitRouting() {
	s.InitMiddlewares()
	apiGroup := s.Echo.Group("/api")

	walletGroup := apiGroup.Group("/wallet")
	wallet.InitRouting(walletGroup)
}
