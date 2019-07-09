package reformat

import (
	"strings"
)

// SetMapToLower we get Uppercase field names from oracle, or any other DB depending on the DB schema, as a result we need to downcase the keys in the map
// this is specifically used when we use the json name in the struct to map to the result from the DB.. THIS_FIELD should be this_field (yes, thanks Oracle!!)
func SetMapToLower(in map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, values := range in { // range over map
		result[strings.ToLower(key)] = values
	}
	return result
}

// FlattenURLValues  normalizes map[string][]string to map[string]string.
// This should ONLY be used if you are not interested in whatever "duplicate" filters gets passed. It will pick whatever is in index 0 as the value and ignore the rest.
// example.. myurl/thisendpoit?file_id=1?file_id=2 would be represeted as map{file_id[1 2]} and the converted to map[file_id=1] and the 2 will be ignored.
func FlattenURLValues(url map[string][]string) map[string]string {
	values := make(map[string]string)
	for key, val := range url {
		values[key] = val[0]
	}
	return values
}

// NormalizeURLValues normalizes map[string][]string to map[string]string.
// This should ONLY be used if you are not interested in whatever "duplicate" filters gets passed. It will pick whatever is in index 0 as the value and ignore the rest.
// example.. myurl/thisendpoit?file_id=1?file_id=2 would be represeted as map{file_id[1 2]} and the converted to map[file_id=1] and the 2 will be ignored.
func NormalizeURLValues(url map[string][]string) map[string]string {
	values := make(map[string]string)
	for key, val := range url {
		var value string
		for _, v := range val {
			if value != "" {
				value = value + "," + v
			} else {
				value = value + v
			}
		}
		values[key] = value
	}
	return values
}
