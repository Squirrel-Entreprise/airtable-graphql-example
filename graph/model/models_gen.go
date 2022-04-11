// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// Un produit est une entité qui possède un nom, une image, une catégorie et un prix
type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// L'image du produit (url)
	Cover    string   `json:"cover"`
	Category Category `json:"category"`
	// Prix unitaire en euros
	Price float64 `json:"price"`
}

// Catégories de produit
type Category string

const (
	CategoryVegetable Category = "VEGETABLE"
	CategoryFruit     Category = "FRUIT"
)

var AllCategory = []Category{
	CategoryVegetable,
	CategoryFruit,
}

func (e Category) IsValid() bool {
	switch e {
	case CategoryVegetable, CategoryFruit:
		return true
	}
	return false
}

func (e Category) String() string {
	return string(e)
}

func (e *Category) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Category(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}

func (e Category) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}