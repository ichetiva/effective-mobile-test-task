package controller

// swagger:parameters createCar
type CarRequest struct {
	// license plate
	// in: body
	RegNum string `json:"regNum"`

	// brand
	// in: body
	Mark string `json:"mark"`

	// model
	// in: body
	Model string `json:"model"`

	// car owner
	// in: body
	Owner Owner `json:"owner"`
}

// swagger:model
type Owner struct {
	// first name
	// in: body
	Name string `json:"name"`

	// last name
	// in: body
	Surname string `json:"surname"`
}

// swagger:model
type CarResponse struct {
	// the id of this car
	ID uint `json:"id"`

	// license plate
	RegNum string `json:"regNum"`

	// brand
	Mark string `json:"mark"`

	// model
	Model string `json:"model"`

	// owner
	Owner OwnerResponse `json:"owner"`
}

// swagger:model
type OwnerResponse struct {
	// name
	Name string `json:"name"`

	// surname
	Surname string `json:"surname"`
}

// swagger:model
type CarsResponse []CarResponse

// swagger:model
type EmptyResponse struct{}

// swagger:parameters carFilters
type carFilters struct {
	// car id
	// in: query
	ID uint `json:"id"`

	// license plate
	// in: query
	RegNum string `json:"regNum"`

	// brand
	// in: query
	Mark string `json:"mark"`

	// model
	// in: query
	Model string `json:"model"`
}

// swagger:parameters carId
type carIDParameters struct {
	// car id
	// in: path
	ID uint `json:"id"`
}

// swagger:parameters updateCar
type updateCar struct {
	CarRequest
	carIDParameters
}
