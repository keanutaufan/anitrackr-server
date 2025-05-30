package anime

import (
	"encoding/json"
	"fmt"
	"strings"
)

// parseAiredString converts a string like "{'from': 'YYYY-MM-DD', 'to': 'YYYY-MM-DD'}"
// into a map[string]interface{}.
func parseAiredString(airedStr string) (map[string]interface{}, error) {
	if airedStr == "" || airedStr == "None" || airedStr == "{}" {
		return make(map[string]interface{}), nil // Return empty map for empty/None/empty dict string
	}
	// Replace single quotes with double quotes for JSON compatibility
	jsonCompliantStr := strings.ReplaceAll(airedStr, "'", "\"")
	jsonCompliantStr = strings.ReplaceAll(jsonCompliantStr, "None", "null")

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonCompliantStr), &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Aired string '%s' (processed: '%s'): %w", airedStr, jsonCompliantStr, err)
	}
	return data, nil
}

// parseStringList converts a string like "['item1', 'item2', 'item3']"
// into a []string.
func parseStringList(listStr string) ([]string, error) {
	if listStr == "" || listStr == "None" || listStr == "[]" {
		return []string{}, nil // Return empty slice for empty/None/empty list string
	}

	if strings.HasPrefix(listStr, "['") && strings.HasSuffix(listStr, "']") {
		if listStr == "['']" { // Handles a list with a single empty string
			return []string{""}, nil
		}
		// Trim "['" and "']"
		// Ensure length is sufficient before slicing to prevent panic
		if len(listStr) < 4 {
			return []string{}, nil // Or an error: fmt.Errorf("invalid list format: %s", listStr)
		}
		content := listStr[2 : len(listStr)-2]
		if content == "" { // Handles cases like "['']" which becomes "" after trim, effectively an empty list if not single empty string.
			return []string{}, nil
		}
		// Split by the delimiter "', '"
		items := strings.Split(content, "', '")
		return items, nil
	}
	// Fallback for simple JSON arrays if the format ever changes e.g. ["item1", "item2"]
	if strings.HasPrefix(listStr, "[") && strings.HasSuffix(listStr, "]") {
		var items []string
		if err := json.Unmarshal([]byte(listStr), &items); err == nil {
			return items, nil
		}
	}

	return nil, fmt.Errorf("unsupported string list format: %s", listStr)
}
