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

// get all product 

func GetAllProduct()[] Products{
	return ProductList
}


//get product by id
func Get(id int)*Products{
	for i := range ProductList{
		if id == ProductList[i].ID{
			return &ProductList[i]
		}
	}
	return nil
}

// update product 
func Update(id int, data Products)*Products{
	for idx := range ProductList{
		if id == ProductList[idx].ID{
			ProductList[idx] =  data
			return &ProductList[idx]
		}
	}
	return nil
}
//delete product 
func Delete(id int)bool{
	var newList []Products
	found := false
	for i := range ProductList{
		if id != ProductList[i].ID{
			newList = append(newList, ProductList[i])
			found = true
			continue
		}
	}
	ProductList =  newList
	return found
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
