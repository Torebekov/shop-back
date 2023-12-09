package product

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/Torebekov/shop-back/modules/logger"
	"go.uber.org/zap"
)

type product struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) *product {
	return &product{
		db:  db,
		ctx: ctx,
	}
}

func (r *product) List(searchText string, categoryID uint64) (products []models.ProductDTO, err error) {
	l := logger.WorkLogger.Named("repo.product.List").With(zap.String("searchText", searchText), zap.Uint64("categoryID", categoryID))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	query := `
		SELECT
			product.id,
			product.name,
			category.name,
			image,
			price,
			CASE WHEN user_favorite.product_id IS NULL THEN FALSE ELSE TRUE END
		FROM
			product
		LEFT JOIN
			user_favorite ON user_favorite.product_id = product.ID
		JOIN
			category ON category.id = product.category_id
		WHERE
			user_id = 1`

	if searchText != "" {
		query += " AND product.name like '%" + searchText + "%'"
	}
	if categoryID != 0 {
		query += fmt.Sprintf(" AND category_id = %d", categoryID)
	}

	rows, err := r.db.Query(query)
	if err != nil {
		l.Error("couldn't make request", zap.Error(err))
		return
	}
	defer rows.Close()

	rows.Next()
	{
		var productModel models.ProductDTO
		err = rows.Scan(&productModel.ID, &productModel.Name, &productModel.CategoryName, &productModel.Image, &productModel.Price, &productModel.IsFavorite)
		if err != nil {
			l.Error("couldn't scan product", zap.Error(err))
			return
		}

		products = append(products, productModel)
	}

	if err = rows.Err(); err != nil {
		l.Error("iteration error", zap.Error(err))
		return
	}

	return
}
