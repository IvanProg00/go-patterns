package computer

type Server struct {
	Type Type
	OS   string
	CPU  int
}

func NewServer() Computer {
	return &Server{
		Type: ServerType,
		OS:   "Linux",
		CPU:  64,
	}
}

func (s *Server) GetType() string {
	return string(s.Type)
}

func (s *Server) GetCPU() int {
	return s.CPU
}

func (s *Server) GetOS() string {
	return s.OS
}
