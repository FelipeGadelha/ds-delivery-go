package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Price       float32
	Description string
	ImageUri    string
}

func (p *Product) Update(update Product) {
	p.Name = update.Name
	p.Price = update.Price
	p.Description = update.Description
	p.ImageUri = update.ImageUri
}
