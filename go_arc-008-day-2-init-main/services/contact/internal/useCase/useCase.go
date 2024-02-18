package usecase

import "architecture_go/services/contact/internal/repository"

type UseCase struct {
    ContactUseCase ContactUseCase
}

func NewUseCase(contactRepository repository.ContactRepository) *UseCase {
    contactUseCase := NewContactUseCaseImpl(contactRepository)
    
    return &UseCase{
        ContactUseCase: contactUseCase,
    }
}
