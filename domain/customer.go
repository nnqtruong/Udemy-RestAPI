package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, error)
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Nguyen", "Wichita", "67226", "2000-07-19", "1"},
		{"1002", "Chen", "New York", "10001", "1987-02-28", "1"},
	}
	return CustomerRepositoryStub{customers}
}
