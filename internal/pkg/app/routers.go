package app

func (s *APIServer) confRouter() {
	s.router.HandleFunc("/getuser", s.handler.HandlerGetUser)
	s.router.HandleFunc("/setuser", s.handler.HandlerSetUser)
	s.router.HandleFunc("/adduser", s.handler.HandlerAddUser)
	s.router.HandleFunc("/deluser", s.handler.HandlerDelUser)
	s.router.HandleFunc("/keysuser", s.handler.HandlerKeysUser)
}
