package util

// PreprocessNullableAnyOf recursively converts OpenAPI 3.1.0 nullable anyOf patterns
// like {"anyOf": [{"type": "string"}, {"type": "null"}]} into the 3.0.x equivalent
// {"type": "string", "nullable": true}. This allows kin-openapi (which rejects
// "type": "null") to load 3.1.0 specs. Only 2-member anyOf arrays where one member
// is {"type": "null"} are converted; 3+ member anyOf arrays are left untouched.
func PreprocessNullableAnyOf(m map[string]interface{}) {
	for k, v := range m {
		switch val := v.(type) {
		case map[string]interface{}:
			PreprocessNullableAnyOf(val)
		case []interface{}:
			if k == "anyOf" && len(val) == 2 {
				// Find which member (if any) is {"type": "null"}
				nullIdx := -1
				for i, item := range val {
					obj, ok := item.(map[string]interface{})
					if !ok {
						continue
					}
					if t, ok := obj["type"].(string); ok && t == "null" {
						nullIdx = i
					}
				}
				if nullIdx >= 0 {
					concreteIdx := 1 - nullIdx
					concrete, ok := val[concreteIdx].(map[string]interface{})
					if ok {
						// Merge concrete type properties into parent, replace anyOf
						delete(m, "anyOf")
						for ck, cv := range concrete {
							m[ck] = cv
						}
						m["nullable"] = true
						// Recurse into the now-modified parent
						PreprocessNullableAnyOf(m)
						return
					}
				}
			}
			// Recurse into array elements
			for _, item := range val {
				if sub, ok := item.(map[string]interface{}); ok {
					PreprocessNullableAnyOf(sub)
				}
			}
		}
	}
}

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
