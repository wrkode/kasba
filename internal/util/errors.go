package util

type Errors struct {
	Errors    []error
	Fatal     bool
	HasErrors bool
}

func (e *Errors) Add(err error, fatal bool) bool {
	if err != nil {
		e.HasErrors = true
		e.Errors = append(e.Errors, err)
		if !e.Fatal {
			e.Fatal = fatal
			return true
		}
	}
	return false
}
