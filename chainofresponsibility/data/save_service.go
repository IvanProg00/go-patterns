package data

import "fmt"

type SaveService struct {
	Next Service
}

func (s *SaveService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("data not updated")
		return
	}

	fmt.Println("data saved")
}

func (s *SaveService) SetNext(svc Service) {
	s.Next = svc
}
