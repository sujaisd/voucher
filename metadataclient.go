package voucher

import (
	"context"
	"fmt"

	"github.com/docker/distribution/reference"
	"github.com/grafeas/voucher/repository"
)

// MetadataClient is an interface that represents something that communicates
// with the Metadata server.
type MetadataClient interface {
	CanAttest() bool
	NewPayloadBody(reference.Canonical) (string, error)
	GetVulnerabilities(context.Context, reference.Canonical) ([]Vulnerability, error)
	GetBuildDetail(context.Context, reference.Canonical) (repository.BuildDetail, error)
	AddAttestationToImage(context.Context, reference.Canonical, Attestation) (SignedAttestation, error)
	GetAttestations(context.Context, reference.Canonical) ([]SignedAttestation, error)
	Close()
}

// NoMetadataError is an error that is returned when we request metadata that
// should exist but doesn't. It's a general error that will wrap more specific
// errors if desired.
type NoMetadataError struct {
	Type MetadataType
	Err  error
}

// Error returns the error value of this NoMetadataError as a string.
func (err *NoMetadataError) Error() string {
	return fmt.Sprintf("no metadata of type %s returned: %s", err.Type, err.Err)
}

// IsNoMetadataError returns true if the passed error is a NoMetadataError.
func IsNoMetadataError(err error) bool {
	_, ok := err.(*NoMetadataError)
	return ok
}
