package main

import "fmt"

type vehicle interface {
	Licence() string
	Name() string
}

type car struct {
	licenceNo string
	brand string
}

func (c *car) Name() string{
	return c.brand 
}

func (c *car) Licence() string {
	return c.licenceNo
}

func main() {
	vehicle := car{licenceNo: "T3432", brand: "Toyota"}
	brand := vehicle.Name()
	licence := vehicle.Licence()
	fmt.Println(brand, licence)
 }