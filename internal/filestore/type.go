package filestore

// GetSuffix returns a file suffix for a given mime type
func GetSuffix(kind string) (string, error) {
	switch kind {
	case "image/jpeg":
		return "jpg", nil
	case "image/png":
		return "png", nil
	default:
		return "", ErrUnknownType
	}
}
