package command

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func toStringSlice(a []interface{}) []string {
	b := make([]string, len(a), len(a))
	for i := range a {
		b[i] = a[i].(string)
	}
	return b
}

func toInterfaceSlice(a []string) []interface{} {
	b := make([]interface{}, len(a), len(a))
	for i := range a {
		b[i] = a[i]
	}
	return b
}

func isValid(region string) bool {
	return regions.Contains(region)
}
