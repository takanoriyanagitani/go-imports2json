package imports2json

// ImportInfo holds the extracted, simplified information for a single import.
// It is designed for easy JSON serialization.
type ImportInfo struct {
	Path string `json:"path"`
	Name string `json:"name"`
}
