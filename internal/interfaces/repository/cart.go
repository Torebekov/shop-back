package interfaces

import "github.com/Torebekov/shop-back/internal/models"

type ICart interface {
	Add(cartModel models.Cart) (err error)
	Remove(cartModel models.Cart) (err error)
}
