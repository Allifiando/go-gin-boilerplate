package helper

const (
	// DefaultPage ...
	DefaultPage = 1
	// DefaultLimit ...
	DefaultLimit = 10
	// MaxLimit ...
	MaxLimit = 50
)

// PaginationVM ...
type PaginationVM struct {
	CurrentPage   int `json:"current_page"`
	LastPage      int `json:"last_page"`
	Count         int `json:"count"`
	RecordPerPage int `json:"record_per_page"`
}

// PaginationPageOffset ...
func PaginationPageOffset(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > MaxLimit {
		limit = DefaultLimit
	}

	offset := (page - 1) * limit

	return limit, offset
}

// PaginationRes ...
func PaginationRes(page, count, limit int) PaginationVM {
	lastPage := count / limit
	if count%limit > 0 {
		lastPage = lastPage + 1
	}

	pagination := PaginationVM{
		CurrentPage:   page,
		LastPage:      lastPage,
		Count:         count,
		RecordPerPage: limit,
	}
	return pagination
}
