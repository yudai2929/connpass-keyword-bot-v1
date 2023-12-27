package valueobject

type Address struct {
	Prefecture string
	City       string
}

func NewAddress(prefecture string, city string) Address {
	return Address{
		Prefecture: prefecture,
		City:       city,
	}
}
