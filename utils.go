package kickgo

import "encoding/json"

// This function is made not to repeat the same sequence of code each times
func bytesToTPtr[T any](data []byte, t *T) (*T, error) {
	err := json.Unmarshal(data, t)
	return t, err
}
