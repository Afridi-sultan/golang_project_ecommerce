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


//get product 

func Get(id int)*Products{
	for _,product := range ProductList{
		if id == product.ID{
			return &product
		}
	}
	return nil
}

// init function
func init() {
	prd1 := Products{
		ID:          1,
		Title:       "Mango",
		Description: "Mango is green color, Mango is sweet. This is mango and mango is my favourite fruit.",
		Price:       500.78,
		ImgUrl:      "https://www.frutas-hortalizas.com/img/fruites_verdures/presentacio/18.jpg",
	}
	prd2 := Products{
		ID:          2,
		Title:       "Orange",
		Description: "Orange is yellow color, orange is sweet. This is orange and orange is my favourite fruit.",
		Price:       400.78,
		ImgUrl:      "https://www.frutas-hortalizas.com/img/fruites_verdures/presentacio/18.jpg",
	}

	ProductList = append(ProductList, prd1, prd2)
}
