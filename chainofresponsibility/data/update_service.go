package data

import "fmt"

type UpdateService struct {
	Next Service
	Name string
}

func (u *UpdateService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("update data from service %q already received\n", u.Name)
		u.Next.Execute(data)

		return
	}

	fmt.Printf("update data from service %q\n", u.Name)

	data.UpdateSource = true
	u.Next.Execute(data)
}

func (u *UpdateService) SetNext(svc Service) {
	u.Next = svc
}
