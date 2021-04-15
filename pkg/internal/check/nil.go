package check

func IsNil(values ...interface{}) bool {
	for _, v := range values {
		if v == nil {
			return true
		}
	}
	return false
}
