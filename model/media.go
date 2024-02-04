package model

type MediaItemBasic struct {
	Id       string `json:"id"`
	MediaUrl string `json:"media_url"`
}

type MediaChildren struct {
	MediaItems []MediaItemBasic `json:"data"`
}

type MediaItem struct {
	Id        string         `json:"id"`
	Caption   string         `json:"caption"`
	MediaType string         `json:"media_type"`
	MediaUrl  string         `json:"media_url"`
	Timestamp string         `json:"timestamp"`
	Children  *MediaChildren `json:"children"`
}

type PagingCursor struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type Paging struct {
	Cursors PagingCursor `json:"cursors"`
	Next    string       `json:"next"`
}

type MediaResponse struct {
	MediaItems []MediaItem `json:"data"`
	Paging     Paging      `json:"paging"`
}
