package therapy

type Day struct {
	Index         int    `json:"index"` // day 1, day 2, etc.
	Title         string `json:"title"`
	AudioFilePath string `json:"-"`
	AudioURL      string `json:"audio_url"`
}
