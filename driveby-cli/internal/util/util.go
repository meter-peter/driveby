package util

// PreprocessExclusiveMinMax recursively converts numeric exclusiveMinimum/exclusiveMaximum to boolean if paired with minimum/maximum, otherwise removes the field
func PreprocessExclusiveMinMax(m map[string]interface{}) {
	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			PreprocessExclusiveMinMax(val)
		case []interface{}:
			for _, item := range val {
				if sub, ok := item.(map[string]interface{}); ok {
					PreprocessExclusiveMinMax(sub)
				}
			}
		case float64:
			if k == "exclusiveMinimum" {
				if _, ok := m["minimum"]; ok {
					m[k] = true
				} else {
					delete(m, k)
				}
			}
			if k == "exclusiveMaximum" {
				if _, ok := m["maximum"]; ok {
					m[k] = true
				} else {
					delete(m, k)
				}
			}
		}
	}
}
