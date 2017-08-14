package telego

// Sticker - This object represents a sticker.
// file_id	    String		Unique identifier for this file
// width	    Integer		Sticker width
// height	    Integer		Sticker height
// thumb	    PhotoSize	Optional. Sticker thumbnail in .webp or .jpg format
// emoji	    String		Optional. Emoji associated with the sticker
// file_size	Integer		Optional. File size
type Sticker struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Emoji    string     `json:"emoji,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}
