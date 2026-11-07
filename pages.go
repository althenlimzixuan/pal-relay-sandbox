package pages

// PageCount returns how many pages perPage items per page need for items items.
func PageCount(items, perPage int) int { return items / perPage }
