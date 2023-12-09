package favorite

import (
	"context"
	"database/sql"
	"github.com/Torebekov/shop-back/internal/models"
	"github.com/Torebekov/shop-back/modules/logger"
	"go.uber.org/zap"
)

type favorite struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB, ctx context.Context) *favorite {
	return &favorite{
		db:  db,
		ctx: ctx,
	}
}

func (r *favorite) Add(favoriteModel models.Favorite) (err error) {
	l := logger.WorkLogger.Named("repo.favorite.AddFavorite").With(zap.Any("favoriteModel", favoriteModel))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("INSERT INTO user_favorite (user_id, product_id) VALUES ($1, $2)", favoriteModel.UserID, favoriteModel.ProductID)
	if err != nil {
		l.Error("couldn't insert favorite")
		return
	}

	return
}

func (r *favorite) Remove(favoriteModel models.Favorite) (err error) {
	l := logger.WorkLogger.Named("repo.favorite.AddFavorite").With(zap.Any("favoriteModel", favoriteModel))

	if r.db == nil {
		l.Error("DB not initialized")
		return
	}

	_, err = r.db.Exec("DELETE FROM user_favorite WHERE user_id = $1 AND product_id = $2", favoriteModel.UserID, favoriteModel.ProductID)
	if err != nil {
		l.Error("couldn't delete favorite")
		return
	}

	return
}
