package order

import (
	"context"
	"fmt"
	"insanulab/business/order"

	"github.com/go-kivik/kivik/v4"
)

//MongoDBRepository The implementation of article.Repository object
type CouchDBRepository struct {
	col *kivik.DB
}

type collection struct {
	ID      string `json:"_id"`
	OrderID string `json:"order_id"`
	NoPKS   string `json:"no_pks"`
}

type CouchDBQuery struct {
	Selector map[string]interface{} `json:"selector"`
	Limit    uint64                 `json:"limit,omitempty"`
	Skip     uint64                 `json:"skip,omitempty"`
	Fields   []string               `json:"fields,omitempty"`
	Sort     []map[string]string    `json:"sort,omitempty"`
}

//NewCouchDBRepository Generate mongo DB article repository
func NewCouchDBRepository(client *kivik.Client) *CouchDBRepository {
	return &CouchDBRepository{
		col: client.DB("feedfund_orders"),
	}
}

func (col *collection) ToArticle() order.Order {
	var order order.Order
	order.ID = col.ID
	order.OrderID = col.OrderID
	order.NoPKS = col.NoPKS

	return order
}

//FindArticleByID Find article based on given ID. Its return nil if not found
func (repo *CouchDBRepository) FetchOrders() ([]order.Order, error) {
	fmt.Println("masuk repository")
	var col collection

	orders := []order.Order{}
	query := CouchDBQuery{
		Selector: map[string]interface{}{
			"system_success": true,
		},
	}

	row, err := repo.col.Find(context.TODO(), query)
	fmt.Println("err repo: ", err)

	if err != nil {
		return orders, err
	}
	defer row.Close()

	for row.Next() {
		if err = row.ScanDoc(&col); err != nil {
			return orders, err
		}
		orders = append(orders, col.ToArticle())
	}
	fmt.Println("orders:", orders)
	return orders, nil
}
