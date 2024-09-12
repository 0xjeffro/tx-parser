package utils

import "crypto/sha256"

func CalculateDiscriminator(instructionName string) [8]byte {
	hash := sha256.Sum256([]byte(instructionName))
	var discriminator [8]byte
	copy(discriminator[:], hash[:8])
	return discriminator
}
