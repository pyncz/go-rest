package utils

import "errors"

func FindByField[TModel any, TValue int | string](
	list []TModel,
	accessor func(r TModel) TValue,
	value TValue,
) (*TModel, error) {
	for _, record := range list {
		if accessor(record) == value {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}
