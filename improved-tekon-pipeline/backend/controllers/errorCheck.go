package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

func GetErrorStatusCode(err error) int {
	switch err {
	case gorm.ErrInvalidTransaction, gorm.ErrMissingWhereClause, gorm.ErrUnsupportedRelation, gorm.ErrPrimaryKeyRequired,
		gorm.ErrModelValueRequired, gorm.ErrModelAccessibleFieldsRequired, gorm.ErrSubQueryRequired, gorm.ErrInvalidField,
		gorm.ErrEmptySlice, gorm.ErrInvalidValue, gorm.ErrInvalidValueOfLength, gorm.ErrPreloadNotAllowed,
		gorm.ErrCheckConstraintViolated:
		return http.StatusBadRequest

	case gorm.ErrRecordNotFound:
		return http.StatusNotFound

	case gorm.ErrRegistered, gorm.ErrDuplicatedKey, gorm.ErrForeignKeyViolated:
		return http.StatusConflict

	case gorm.ErrInvalidData:
		return http.StatusUnprocessableEntity

	case gorm.ErrNotImplemented:
		return http.StatusNotImplemented

	case gorm.ErrDryRunModeUnsupported:
		return http.StatusServiceUnavailable

	case gorm.ErrInvalidDB:

		return http.StatusInternalServerError

	default:
		return http.StatusInternalServerError
	}
}
