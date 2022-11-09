package simple

import "errors"

type SimpleRepository struct {
	err bool
}

func NewSimpleRepository(IsErrors bool) *SimpleRepository {
	return &SimpleRepository{IsErrors}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.err {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}

}
