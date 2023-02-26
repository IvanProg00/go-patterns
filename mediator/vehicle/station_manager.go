package vehicle

type StationManager struct {
	Queue        []Vehicle
	PlatformFree bool
}

func NewStationManager() *StationManager {
	return &StationManager{
		PlatformFree: true,
	}
}

func (s *StationManager) CanArrive(v Vehicle) bool {
	if s.PlatformFree {
		s.PlatformFree = false
		return true
	}

	s.Queue = append(s.Queue, v)

	return false
}

func (s *StationManager) NotifyAboutGo() {
	if !s.PlatformFree {
		s.PlatformFree = true
	}

	if len(s.Queue) > 0 {
		firstTrainInQueue := s.Queue[0]
		s.Queue = s.Queue[1:]

		firstTrainInQueue.PermitArrive()
	}
}
