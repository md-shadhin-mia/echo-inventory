package models

import "github.com/jinzhu/gorm"

type BaseModel struct{
	gorm.Model
}

type Category struct{
	BaseModel
	Name string 
}

type ProductType struct{
	BaseModel
	Name string
}

type Product struct{
	BaseModel
	Name string
	CategoryID uint
	Category Category `gorm:"foreignKey:CategoryID"`
	ProductTypeID uint
	ProductType ProductType `gorm:"foreignKey:ProductTypeID"`
}

type Pricing struct {
    BaseModel
    ProductID uint
	Product Product `gorm:"foreignKey:ProductID"`
    PricePerUnit float64
    IsSalePrice bool
    SalePrice float64
}