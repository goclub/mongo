package mo

func In(slice interface{}) interface{} {
	if slice == nil {
		return []string{}
	}
	return slice
}
