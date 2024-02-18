package delivery

import (
    "net/http"

	"architecture_go/services/contact/internal/usecase"

)

type Delivery struct {
    ContactHandler ContactHandler
}

func NewDelivery(contactUseCase usecase.ContactUseCase) *Delivery {
    contactHandler := NewContactHandler(contactUseCase)
    
    return &Delivery{
        ContactHandler: contactHandler,
    }
}
