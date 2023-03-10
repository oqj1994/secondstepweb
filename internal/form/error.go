package form

type errors map[string][]string

func newErr() errors {
	return errors{}
}

func (e errors) Add(field, msg string) {
	e[field] = append(e[field], msg)
}

func (e errors) Get(field string) string {
	errs, ok := e[field]
	if !ok {
		return ""
	}
	return errs[0]

}
