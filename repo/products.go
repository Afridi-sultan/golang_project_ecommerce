package repo

type Products struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

type ProductRepo interface{
	Create(p Products)(*Products, error)
	Get(productId int)(*Products, error)
	List()([]*Products,error)
	Delete(id int) error
	Update(p Products)(*Products, error)
}

type productRepo struct{
	productList []*Products
}

//constructor
func NewProductRepo()ProductRepo{
	repo := &productRepo{}
	GenerateInitProduct(repo)
	return repo
}

//Methods
func (r *productRepo) Create(p Products)(*Products, error){
	p.ID = len(r.productList)+1
	r.productList = append(r.productList, &p)
	return &p,nil
}

func (r *productRepo) Get(id int)(*Products, error){
	for i := range r.productList{
		if id == r.productList[i].ID{
			return r.productList[i],nil
		}
	}
	return nil,nil
}


func (r *productRepo) List()([]*Products,error){
	return r.productList,nil
}


func (r *productRepo) Delete(id int) error{
	var tempList []*Products
	for _,p := range r.productList{
		if p.ID != id{
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}


func (r *productRepo) Update( p Products)(*Products, error){
	for idx := range r.productList{
		if p.ID == r.productList[idx].ID{
			r.productList[idx] =  &p
			return r.productList[idx], nil
		}
	}
	return nil, nil
}

func GenerateInitProduct( r *productRepo) {
	prd1 := &Products{
		ID:          1,
		Title:       "Mango",
		Description: "Mango is green color, Mango is sweet. This is mango and mango is my favourite fruit.",
		Price:       500.78,
		ImgUrl:      "https://www.frutas-hortalizas.com/img/fruites_verdures/presentacio/18.jpg",
	}
	prd2 := &Products{
		ID:          2,
		Title:       "Orange",
		Description: "Orange is yellow color, orange is sweet. This is orange and orange is my favourite fruit.",
		Price:       400.78,
		ImgUrl:      "https://www.frutas-hortalizas.com/img/fruites_verdures/presentacio/18.jpg",
	}

	r.productList = append(r.productList, prd1, prd2)
}
