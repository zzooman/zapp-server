// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: products.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (seller, title, content, price, medias, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, seller, title, content, medias, price, views, created_at
`

type CreateProductParams struct {
	Seller    string             `json:"seller"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Price     int64              `json:"price"`
	Medias    []string           `json:"medias"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Seller,
		arg.Title,
		arg.Content,
		arg.Price,
		arg.Medias,
		arg.CreatedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Seller,
		&i.Title,
		&i.Content,
		&i.Medias,
		&i.Price,
		&i.Views,
		&i.CreatedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, seller, title, content, medias, price, views, created_at FROM products WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Seller,
		&i.Title,
		&i.Content,
		&i.Medias,
		&i.Price,
		&i.Views,
		&i.CreatedAt,
	)
	return i, err
}

const getProductWithSellor = `-- name: GetProductWithSellor :one
SELECT products.id, products.seller, products.title, products.content, products.medias, products.price, products.views, products.created_at, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username WHERE products.id = $1 LIMIT 1
`

type GetProductWithSellorRow struct {
	ID        int64              `json:"id"`
	Seller    string             `json:"seller"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Price     int64              `json:"price"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     pgtype.Text        `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetProductWithSellor(ctx context.Context, id int64) (GetProductWithSellorRow, error) {
	row := q.db.QueryRow(ctx, getProductWithSellor, id)
	var i GetProductWithSellorRow
	err := row.Scan(
		&i.ID,
		&i.Seller,
		&i.Title,
		&i.Content,
		&i.Medias,
		&i.Price,
		&i.Views,
		&i.CreatedAt,
		&i.Email,
		&i.Phone,
		&i.Profile,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
SELECT id, seller, title, content, medias, price, views, created_at FROM products ORDER BY created_at DESC LIMIT $1 OFFSET $2
`

type GetProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetProducts(ctx context.Context, arg GetProductsParams) ([]Product, error) {
	rows, err := q.db.Query(ctx, getProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Seller,
			&i.Title,
			&i.Content,
			&i.Medias,
			&i.Price,
			&i.Views,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsWithSeller = `-- name: GetProductsWithSeller :many
SELECT products.id, products.seller, products.title, products.content, products.medias, products.price, products.views, products.created_at, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username ORDER BY products.created_at DESC LIMIT $1 OFFSET $2
`

type GetProductsWithSellerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetProductsWithSellerRow struct {
	ID        int64              `json:"id"`
	Seller    string             `json:"seller"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Price     int64              `json:"price"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     pgtype.Text        `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetProductsWithSeller(ctx context.Context, arg GetProductsWithSellerParams) ([]GetProductsWithSellerRow, error) {
	rows, err := q.db.Query(ctx, getProductsWithSeller, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductsWithSellerRow{}
	for rows.Next() {
		var i GetProductsWithSellerRow
		if err := rows.Scan(
			&i.ID,
			&i.Seller,
			&i.Title,
			&i.Content,
			&i.Medias,
			&i.Price,
			&i.Views,
			&i.CreatedAt,
			&i.Email,
			&i.Phone,
			&i.Profile,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsWithSellerByQuery = `-- name: GetProductsWithSellerByQuery :many
SELECT products.id, products.seller, products.title, products.content, products.medias, products.price, products.views, products.created_at, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username WHERE products.title ILIKE '%' || $1 || '%' OR products.content ILIKE '%' || $1 || '%' ORDER BY products.created_at DESC LIMIT $2 OFFSET $3
`

type GetProductsWithSellerByQueryParams struct {
	Column1 pgtype.Text `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

type GetProductsWithSellerByQueryRow struct {
	ID        int64              `json:"id"`
	Seller    string             `json:"seller"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Price     int64              `json:"price"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     pgtype.Text        `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetProductsWithSellerByQuery(ctx context.Context, arg GetProductsWithSellerByQueryParams) ([]GetProductsWithSellerByQueryRow, error) {
	rows, err := q.db.Query(ctx, getProductsWithSellerByQuery, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductsWithSellerByQueryRow{}
	for rows.Next() {
		var i GetProductsWithSellerByQueryRow
		if err := rows.Scan(
			&i.ID,
			&i.Seller,
			&i.Title,
			&i.Content,
			&i.Medias,
			&i.Price,
			&i.Views,
			&i.CreatedAt,
			&i.Email,
			&i.Phone,
			&i.Profile,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsWithSellerThatIBought = `-- name: GetProductsWithSellerThatIBought :many
SELECT products.id, products.seller, products.title, products.content, products.medias, products.price, products.views, products.created_at, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username JOIN transactions ON products.id = transactions.feed_id WHERE transactions.buyer = $1 ORDER BY products.created_at DESC LIMIT $2 OFFSET $3
`

type GetProductsWithSellerThatIBoughtParams struct {
	Buyer  string `json:"buyer"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetProductsWithSellerThatIBoughtRow struct {
	ID        int64              `json:"id"`
	Seller    string             `json:"seller"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Price     int64              `json:"price"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     pgtype.Text        `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetProductsWithSellerThatIBought(ctx context.Context, arg GetProductsWithSellerThatIBoughtParams) ([]GetProductsWithSellerThatIBoughtRow, error) {
	rows, err := q.db.Query(ctx, getProductsWithSellerThatIBought, arg.Buyer, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductsWithSellerThatIBoughtRow{}
	for rows.Next() {
		var i GetProductsWithSellerThatIBoughtRow
		if err := rows.Scan(
			&i.ID,
			&i.Seller,
			&i.Title,
			&i.Content,
			&i.Medias,
			&i.Price,
			&i.Views,
			&i.CreatedAt,
			&i.Email,
			&i.Phone,
			&i.Profile,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsWithSellerThatILiked = `-- name: GetProductsWithSellerThatILiked :many
SELECT products.id, products.seller, products.title, products.content, products.medias, products.price, products.views, products.created_at, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username JOIN like_with_feed ON products.id = like_with_feed.feed_id WHERE like_with_feed.username = $1 ORDER BY products.created_at DESC LIMIT $2 OFFSET $3
`

type GetProductsWithSellerThatILikedParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

type GetProductsWithSellerThatILikedRow struct {
	ID        int64              `json:"id"`
	Seller    string             `json:"seller"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Price     int64              `json:"price"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     pgtype.Text        `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetProductsWithSellerThatILiked(ctx context.Context, arg GetProductsWithSellerThatILikedParams) ([]GetProductsWithSellerThatILikedRow, error) {
	rows, err := q.db.Query(ctx, getProductsWithSellerThatILiked, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductsWithSellerThatILikedRow{}
	for rows.Next() {
		var i GetProductsWithSellerThatILikedRow
		if err := rows.Scan(
			&i.ID,
			&i.Seller,
			&i.Title,
			&i.Content,
			&i.Medias,
			&i.Price,
			&i.Views,
			&i.CreatedAt,
			&i.Email,
			&i.Phone,
			&i.Profile,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsWithSellerThatISold = `-- name: GetProductsWithSellerThatISold :many
SELECT products.id, products.seller, products.title, products.content, products.medias, products.price, products.views, products.created_at, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username JOIN transactions ON products.id = transactions.feed_id WHERE transactions.seller = $1 ORDER BY products.created_at DESC LIMIT $2 OFFSET $3
`

type GetProductsWithSellerThatISoldParams struct {
	Seller string `json:"seller"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetProductsWithSellerThatISoldRow struct {
	ID        int64              `json:"id"`
	Seller    string             `json:"seller"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Price     int64              `json:"price"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     pgtype.Text        `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetProductsWithSellerThatISold(ctx context.Context, arg GetProductsWithSellerThatISoldParams) ([]GetProductsWithSellerThatISoldRow, error) {
	rows, err := q.db.Query(ctx, getProductsWithSellerThatISold, arg.Seller, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetProductsWithSellerThatISoldRow{}
	for rows.Next() {
		var i GetProductsWithSellerThatISoldRow
		if err := rows.Scan(
			&i.ID,
			&i.Seller,
			&i.Title,
			&i.Content,
			&i.Medias,
			&i.Price,
			&i.Views,
			&i.CreatedAt,
			&i.Email,
			&i.Phone,
			&i.Profile,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products SET title = $2, content = $3, price = $4, medias = $5 WHERE id = $1
`

type UpdateProductParams struct {
	ID      int64    `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Price   int64    `json:"price"`
	Medias  []string `json:"medias"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.Exec(ctx, updateProduct,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.Price,
		arg.Medias,
	)
	return err
}

const viewProduct = `-- name: ViewProduct :exec
UPDATE products SET views = views + 1 WHERE id = $1
`

func (q *Queries) ViewProduct(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, viewProduct, id)
	return err
}
