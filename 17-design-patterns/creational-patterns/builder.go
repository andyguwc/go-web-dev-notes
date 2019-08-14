/* Builder Design Pattern

Example vehicle manufacturing 
*/


package creational

// On every set step, we return the same build process, so we can chain various steps together 
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//Director in charge of construction of the objects while builder return the actual vehicle

// We will start with the Manufacturing director and the Car Builder to fulfill the first two
// acceptance criteria. In the preceding code, we are creating our Manufacturing director that
// will be in charge of the creation of every vehicle during the test. After creating the
// Manufacturing director, we created a CarBuilder that we then passed to manufacturing
// by using the SetBuilder method. Once the Manufacturing director knows what it has to
// construct now, we can call the Construct method to create the VehicleProduct using
// CarBuilder. Finally, once we have all the pieces for our car, we call the GetVehicle
// method on CarBuilder to retrieve a Car instance:

// The Manufacturing director must accept a builder and construct a vehicle using the provided builder



type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

//Product
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//A Builder of type car
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

//A Builder of type motorbike
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

//A Builder of type motorbike
type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 8
	return b
}

func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
