package server

type Product struct {
	ID    int
	Name  string
	Price float32
	Image string
}

var products []Product

func fillProducts() {
	products = make([]Product, 0)

	products = append(products, Product{
		ID:    1,
		Name:  "Апельсин",
		Price: 199,
		Image: "1.png",
	})
	products = append(products, Product{
		ID:    2,
		Name:  "Яблоко",
		Price: 119,
		Image: "2.png",
	})
	products = append(products, Product{
		ID:    3,
		Name:  "Груша",
		Price: 189,
		Image: "3.png",
	})
	products = append(products, Product{
		ID:    4,
		Name:  "Киви",
		Price: 149,
		Image: "4.png",
	})
	products = append(products, Product{
		ID:    5,
		Name:  "Гранат",
		Price: 239,
		Image: "5.png",
	})
}
