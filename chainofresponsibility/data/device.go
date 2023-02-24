package data

import "fmt"

type Device struct {
	Next Service
	Name string
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("data from device %q already received\n", d.Name)
		d.Next.Execute(data)

		return
	}

	fmt.Printf("get data from device %q\n", d.Name)

	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(svc Service) {
	d.Next = svc
}
