package utils

type SuccessResponseStruct struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type SuccessWithTokenResponseStruct struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type OffsetPaginationResponseStruct struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	TotalCount  int64       `json:"total_count"`
	Limit       int         `json:"limit"`
	NextOffset  int         `json:"next_page,omitempty"`
	PrevOffset  int         `json:"prev_page,omitempty"`
	TotalPages  int         `json:"total_pages,omitempty"`
	CurrentPage int         `json:"current_page,omitempty"`
}

type OffsetPaginationResponseStructWithDueCount struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	TotalCount  int64       `json:"total_count"`
	DueCount    int64       `json:"due_count"`
	Limit       int         `json:"limit"`
	NextOffset  int         `json:"next_page,omitempty"`
	PrevOffset  int         `json:"prev_page,omitempty"`
	TotalPages  int         `json:"total_pages,omitempty"`
	CurrentPage int         `json:"current_page,omitempty"`
}

type ErrorResponseStruct struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

func OffsetPaginationResponse(status int, message string, data interface{}, totalCount int64, limit int, nextOffset int, prevOffset int, totalPages int, currentPage int) *OffsetPaginationResponseStruct {
	return &OffsetPaginationResponseStruct{
		Status:      status,
		Message:     message,
		Data:        data,
		TotalCount:  totalCount,
		Limit:       limit,
		NextOffset:  nextOffset,
		PrevOffset:  prevOffset,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
	}
}

func OffsetPaginationResponseWithDueCount(status int, message string, data interface{}, dueCount, totalCount int64, limit int, nextOffset int, prevOffset int, totalPages int, currentPage int) *OffsetPaginationResponseStructWithDueCount {
	return &OffsetPaginationResponseStructWithDueCount{
		Status:      status,
		Message:     message,
		Data:        data,
		TotalCount:  totalCount,
		DueCount:    dueCount,
		Limit:       limit,
		NextOffset:  nextOffset,
		PrevOffset:  prevOffset,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
	}
}

func SuccessResponse(status int, message string, data interface{}) *SuccessResponseStruct {
	return &SuccessResponseStruct{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(status int, message string, err error) *ErrorResponseStruct {
	return &ErrorResponseStruct{
		Status:  status,
		Message: message,
		Error:   err,
	}
}
