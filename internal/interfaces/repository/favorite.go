package interfaces

import "github.com/Torebekov/shop-back/internal/models"

type IFavorite interface {
	Add(favoriteModel models.Favorite) (err error)
	Remove(favoriteModel models.Favorite) (err error)
}
