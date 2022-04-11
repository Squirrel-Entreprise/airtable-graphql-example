package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"airtale/graph/generated"
	"airtale/graph/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Squirrel-Entreprise/airtable"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, name string, address string) (string, error) {
	// connexion à l'API `keyXXXXX` pour cette base `appXXXXXgit remote add origin git@github.com:Squirrel-Entreprise/airtable-graphql-example.git`
	a := airtable.New("keyXXXXX", "appXXXXXgit remote add origin git@github.com:Squirrel-Entreprise/airtable-graphql-example.git")

	orderTable := airtable.Table{
		Name:       "Orders",
		MaxRecords: "100",
		View:       "Grid view",
	}

	type payloadOrder struct {
		Fields struct {
			Name    string `json:"Name"`
			Address string `json:"Address"`
		} `json:"fields"`
	}

	newOrder := payloadOrder{}
	newOrder.Fields.Name = name
	newOrder.Fields.Address = address

	payload, err := json.Marshal(newOrder)
	if err != nil {
		fmt.Println(err)
	}

	// représentation d'une commande Airtable dans Golang
	type ordertItemAirtable struct {
		ID          string    `json:"id"`
		CreatedTime time.Time `json:"createdTime"`
		Fields      struct {
			Name    string    `json:"Name"`
			Address string    `json:"Address"`
			Status  string    `json:"Status"`
			Carts   []string  `json:"Carts"`
			Amount  float64   `json:"Amount"`
			Date    time.Time `json:"Date"`
		} `json:"fields"`
	}

	order := ordertItemAirtable{}

	if err := a.Create(orderTable, payload, &order); err != nil {
		fmt.Println(err)
	}

	return order.ID, nil
}

func (r *mutationResolver) AddToCart(ctx context.Context, orderID string, productID string, quantity int) (bool, error) {
	// connexion à l'API `keyXXXXX` pour cette base `appXXXXXgit remote add origin git@github.com:Squirrel-Entreprise/airtable-graphql-example.git`
	a := airtable.New("keyXXXXX", "appXXXXXgit remote add origin git@github.com:Squirrel-Entreprise/airtable-graphql-example.git")

	cartTable := airtable.Table{
		Name:       "Carts",
		MaxRecords: "100",
		View:       "Grid view",
	}

	type payloadCart struct {
		Fields struct {
			Product []string `json:"Product"`
			Order   []string `json:"Order"`
			Qt      int      `json:"Qt"`
		} `json:"fields"`
	}

	newCartItem := payloadCart{}
	newCartItem.Fields.Product = []string{productID}
	newCartItem.Fields.Order = []string{orderID}
	newCartItem.Fields.Qt = quantity

	payload, err := json.Marshal(newCartItem)
	if err != nil {
		fmt.Println(err)
	}

	// représentation d'une commande Airtable dans Golang
	type cartAirtable struct {
		ID          string    `json:"id"`
		CreatedTime time.Time `json:"createdTime"`
		Fields      struct {
			Product   []string  `json:"Product"`
			Order     []string  `json:"Order"`
			Qt        int       `json:"Qt"`
			Name      string    `json:"Name"`
			UnitPrice []float64 `json:"Unit price"`
			Amount    float64   `json:"Amount"`
		} `json:"fields"`
	}

	cart := cartAirtable{}

	if err := a.Create(cartTable, payload, &cart); err != nil {
		fmt.Println(err)
	}

	return true, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	// connexion à l'API `keyXXXXX` pour cette base `appXXXXXgit remote add origin git@github.com:Squirrel-Entreprise/airtable-graphql-example.git`
	a := airtable.New("keyXXXXX", "appXXXXXgit remote add origin git@github.com:Squirrel-Entreprise/airtable-graphql-example.git")

	// on sélectionne ta table Products
	productTable := airtable.Table{
		Name:       "Products",
		MaxRecords: "100",
		View:       "Grid view",
	}

	// représentation d'un produit Airtable dans Golang
	type productItemAirtable struct {
		ID          string    `json:"id"`
		CreatedTime time.Time `json:"createdTime"`
		Fields      struct {
			Name     string                `json:"Name"`
			Cover    []airtable.Attachment `json:"cover"`
			Category string                `json:"Category"`
			Price    float64               `json:"Price"`
			Carts    []string              `json:"Carts"`
		} `json:"fields"`
	}

	// représentation de a liste des produits Airtable dans Golang
	type productsListAirtable struct {
		Records []productItemAirtable `json:"records"`
		Offset  string                `json:"offset"`
	}

	products := productsListAirtable{}

	// on injecte les produits da la réponse de l'API à notre variable `products`
	if err := a.List(productTable, &products); err != nil {
		fmt.Println(err)
	}

	// on définie la sortie attendus par notre schéma
	var out []*model.Product
	for _, p := range products.Records {

		var cover string
		if len(p.Fields.Cover) > 0 {
			cover = p.Fields.Cover[0].URL
		}

		// pour chaque produit de Airtable on l'ajoute dans notre sorie
		out = append(out, &model.Product{
			ID:       p.ID,
			Name:     p.Fields.Name,
			Cover:    cover,
			Category: model.Category(p.Fields.Category),
			Price:    p.Fields.Price,
		})
	}

	// on retourne la sortie
	return out, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
