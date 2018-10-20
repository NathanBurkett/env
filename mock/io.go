package mock

// Reader mock reader struct
type Reader struct {
	ReadErr error
}

// Read mock reader.Read method
func (r Reader) Read(p []byte) (n int, err error) {
	if r.ReadErr != nil {
		return 0, r.ReadErr
	}

	return 0, nil
}
