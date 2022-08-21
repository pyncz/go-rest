package validation

type Validatable interface {
	Validate() *ValidationErrors
}
