package utils

import (
	"github.com/pkg/errors"
)

// WrapError wraps given err with allErrs and returns a unified error
func WrapError(err, allErrs error) error {
	// if allErrs is empty, return err
	if allErrs == nil {
		return err
	}

	// if err empty return allErrs
	if err == nil {
		return allErrs
	}

	// wrap err with allErrs
	allErrs = errors.Wrap(err, allErrs.Error())
	return allErrs
}
