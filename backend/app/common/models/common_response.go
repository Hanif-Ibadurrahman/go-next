package models

type CommonResponse struct {
	Status        int         `json:"status"`
	Error         bool        `json:"error"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data"`
	EncryptedData interface{} `json:"encrypted_data,omitempty"`
	Meta          interface{} `json:"meta,omitempty"`
}

type Meta struct {
	ItemCount int      `json:"item_count"`
	ItemTotal int      `json:"item_total"`
	Page      *Page    `json:"page,omitempty"`
	Cursors   *Cursors `json:"cursors,omitempty"`
}

type Page struct {
	IsCursor bool `json:"is_cursor"`
	Current  int  `json:"current"`
	Previous int  `json:"previous"`
	Next     int  `json:"next"`
	Limit    int  `json:"limit"`
	Total    int  `json:"total"`
}

type Cursors struct {
	NextCursors string `json:"next"`
	PrevCursors string `json:"previous"`
}
