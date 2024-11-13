package payment

type Transaction struct {
	ID     int
	Amount int
	Cloth  ClothDetail
}

type ClothDetail struct {
	ID       int
	Name     string
	Price    int
	Quantity int
}
