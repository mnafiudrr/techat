package types

// ThumbType represents thumbnail or thumb details.
type ThumbType struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     *int   `json:"file_size,omitempty"`
	Width        *int   `json:"width,omitempty"`
	Height       *int   `json:"height,omitempty"`
}

// AnimationType represents animation file details.
type AnimationType struct {
	MimeType     string    `json:"mime_type"`
	Duration     int       `json:"duration"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	Thumb        ThumbType `json:"thumb"`
	Thumbnail    ThumbType `json:"thumbnail"`
	FileID       string    `json:"file_id"`
	FileUniqueID string    `json:"file_unique_id"`
	FileSize     int       `json:"file_size"`
}

// AudioType represents audio file details.
type AudioType struct {
	Duration     int     `json:"duration"`
	Filename     *string `json:"filename,omitempty"`
	MimeType     *string `json:"mime_type,omitempty"`
	Title        *string `json:"title,omitempty"`
	Performer    *string `json:"performer,omitempty"`
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	FileSize     *int    `json:"file_size,omitempty"`
}

// ChatType represents chat details.
type ChatType struct {
	ID                          int64   `json:"id"`
	FirstName                   *string `json:"first_name,omitempty"`
	LastName                    *string `json:"last_name,omitempty"`
	Username                    *string `json:"username,omitempty"`
	Title                       *string `json:"title,omitempty"`
	Type                        string  `json:"type"`
	AllMembersAreAdministrators *bool   `json:"all_members_are_administrators,omitempty"`
}

// DocumentType represents document file details.
type DocumentType struct {
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	Thumb        *ThumbType `json:"thumb,omitempty"`
	Thumbnail    *ThumbType `json:"thumbnail,omitempty"`
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	FileSize     *int       `json:"file_size,omitempty"`
}

// PhotoType represents photo file details.
type PhotoType struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     *int   `json:"file_size,omitempty"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

// UserType represents user details.
type UserType struct {
	ID           int64   `json:"id"`
	IsBot        bool    `json:"is_bot"`
	FirstName    string  `json:"first_name"`
	LastName     *string `json:"last_name,omitempty"`
	Username     *string `json:"username,omitempty"`
	LanguageCode *string `json:"language_code,omitempty"`
}

// WebhookRequestType represents the incoming webhook request from Telegram.
type WebhookRequestType struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int      `json:"message_id"`
		From      UserType `json:"from"`
		Chat      ChatType `json:"chat"`
		Date      int      `json:"date"`
		Text      *string  `json:"text,omitempty"`
		Entities  *[]struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities,omitempty"`
		Document        *DocumentType       `json:"document,omitempty"`
		Animation       *AnimationType      `json:"animation,omitempty"`
		Photo           *[]PhotoType        `json:"photo,omitempty"`
		HasMediaSpoiler *bool               `json:"has_media_spoiler,omitempty"`
		ReplyToMessage  *WebhookRequestType `json:"reply_to_message,omitempty"`
	} `json:"message"`
}
