package database

// custom struct

type Products struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

// product slice
var ProductList []Products


// init function
func init() {
	prd1 := Products{
		ID:          1,
		Title:       "Mango",
		Description: "Mango is green color, Mango is sweet. This is mango and mango is my favourite fruit.",
		Price:       500.78,
		ImgUrl:      "https://www.frutas-hortalizas.com/img/fruites_verdures/presentacio/18.jpg",
	}

	ProductList = append(ProductList, prd1)
}
