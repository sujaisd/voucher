package voucher

import (
	"fmt"
)

// AttestationPayload is a structure that contains the Attestation data that we
// want to create an Occurrence from.
type AttestationPayload struct {
	CheckName string
	Body      string
}

// Sign takes a keyring and signs the body of the payload with it, returning that as a string.
func (payload *AttestationPayload) Sign(keyring *KeyRing) (string, string, error) {
	signer, err := keyring.GetSignerByName(payload.CheckName)
	if nil != err {
		return "", "", err
	}

	// create the hex version of the KeyID.
	keyID := fmt.Sprintf("%#x", signer.PrimaryKey.KeyId)

	signature, err := Sign(signer, payload.Body)

	return signature, keyID, err
}

// NewAttestationPayload creates a new AttestationPayload for the check with the passed name,
// with the payload as the body. The payload will then be signed by the key associated
// with the check (referenced by the checkName).
func NewAttestationPayload(checkName string, payload string) AttestationPayload {
	return AttestationPayload{
		CheckName: checkName,
		Body:      payload,
	}
}
