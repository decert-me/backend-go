package utils

import "github.com/google/uuid"

// IsUUID 判断是否UUID
func IsUUID(uuidStr string) bool {
	_, err := uuid.Parse(uuidStr)
	return err == nil
}
