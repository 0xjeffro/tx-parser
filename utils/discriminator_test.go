package utils

import "testing"
import "github.com/mr-tron/base58"

func TestCalculateDiscriminator(t *testing.T) {
	instructionName := "global:buy"
	decode, err := base58.Decode("AJTQ2h9DXrBgNE7LeDS9iSYNeUopL9GEb")
	if err != nil {
		t.Errorf("Error decoding base58: %v", err)
	}
	discriminator := CalculateDiscriminator(instructionName)
	if string(discriminator[:]) != string(decode[:8]) {
		t.Errorf("Discriminator for '%s' is not correct", instructionName)
	} else {
		t.Logf("Discriminator for '%s' is '%s'", instructionName, string(discriminator[:]))
	}
}
