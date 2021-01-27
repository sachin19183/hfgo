package magazine

type Address struct {
	Street     string
	City       string
	State      string
	Postalcode string
}
type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address //This is one way to declare struct within a struct
}

type Employee struct {
	Name    string
	Salary  float64
	Address //This is how you declare an embedded struct
}
