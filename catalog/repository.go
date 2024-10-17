package catalog

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	ErrNotFound = errors.New("Entity not Found")
)

type Repository interface {
	Close()
	PutProduct(ctx context.Context, p Product) error
	GetProductByID(ctx context.Context, id string) (*Product, error)
	ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error)
}

type elascticRepository struct {
	client *elastic.Client
}

type productDocument struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       string `json:"price"`
}

func NewElasticRepository(url string) (Repository, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	return &elascticRepository(client), nil
}

func (r *elascticRepository) Close() {
	r.Close()
}

func (r *elascticRepository) PutProduct(ctx context.Context, p Product) error {
	_, err := r.client.Index().
		Index("catalog").
		Type("product").
		Id(p.ID).
		BodyJson(productDocument{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})

}

func (r *elascticRepository) GetProductByID(ctx context.Context, id string) (*Product, error) {
	res, err := r.client.Get().
		Index("catalog").
		Type("product").
		Id(id).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	if !res.Found {
		return nil, ErrNotFound
	}
	p := productDocument{}

	if err := json.Unmarshal(*res.Source, &p); err != nil {
		return nil, err
	}

	return &Product{
		ID:          id,
		Name:        p.Name,
		Description: p.Description,
		Prize:       p.Price,
	}, err

}

func (r *elascticRepository) ListProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error) {
	res, err := client.Search().
		Index("catalog").
		Type("product").
		Query(elastic.NewMatchAllQuery()).
		From(int(skip)).Size(int(take)).
		Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	products := []Products{}

	for _, hit := res.Hits.Hits{
		p := productDocument{}
		if err = json.Unmarshal(*hit.Source, &p); err == nil{
			products = append(products, Products{
				ID: hit.Id,
				Name: p.Name,
				Description: p.Description
			})
		}
	}

	return products, nil
}

func (r *elascticRepository) ListProductsWithIDs(ctx context.Context, ids []string) ([]Product, error) {
	items := elastic.GetMultiItem{}
	for _, ids := range ids{
		items := append(
			items,
			elastic.NewMultiGetItem().
			Index("catalog").
			Type("product").
			Id(id)
		)
	}
	res, err := client.MultiGet().
		Add(items...).
		Do(ctx)
	if err != nil{
		log.Println(err)
		return nil, err
	} 

	products := []Product{}

	for _, doc := range res.Doc{
		p := productDocument{}
		if err = json.Unmarshal(*doc.Source, &p) err == nil{
			products = append(products, Products{
				ID: doc.Id,
				Name: p.Name,
				Description: p.Description
				Price: p.Price
			})
		}
	}
	return products, nil
}

func (r *elascticRepository) SearchProducts(ctx context.Context, query string, skip uint64, take uint64) ([]Product, error) {
	res, err := r.Client.Search().
	Index("catalog").
	Type("produc").
	Query(elastic.NewMultiMatchQuery(query, "name", "description")).
	From(int(skip)).Size(int(take)).
	Do(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	products := []Product{}
	for _, hit := range res.Hits.Hits {
		p := productDocument{}
		if err = json.Unmarshal(*hits.Source, &p); err == nil {
			products := append(products, Products{
				ID: hit.Id
				Name: p.Name,
				Description: p.Description,
				Price: p.Price,
			})
		}
	}

	return products, nil
}
