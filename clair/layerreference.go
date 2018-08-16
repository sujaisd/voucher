package clair

import (
	"github.com/Shopify/voucher/docker"
	"github.com/coreos/clair/api/v1"
	"github.com/docker/distribution/reference"
	"github.com/opencontainers/go-digest"
)

// LayerReference is a structure containing a Layer digest, as well as the repository
// URI, to simplify loading a Layer from the server.
type LayerReference struct {
	Image   reference.Canonical // The Image's reference.
	Current digest.Digest       // The digest of the current layer.
	Parent  digest.Digest       // The digest of the parent layer.
}

// GetURI gets the URI that is described in the LayerReference.
func (ref *LayerReference) GetURI() string {
	return docker.GetBlobURI(ref.Image)
}

// GetLayer returns a layer description of the LayerReference.
func (ref *LayerReference) GetLayer() v1.Layer {
	return v1.Layer{
		Name:       string(ref.Current),
		Path:       ref.GetURI(),
		Headers:    make(map[string]string),
		ParentName: string(ref.Parent),
		Format:     "Docker",
	}
}

// AddAuthorization adds a Bearer token to the v1.Layer passed to it and
// returns a new v1.Layer.
func AddAuthorization(layer v1.Layer, token string) v1.Layer {
	layer.Headers["Authorization"] = "Bearer " + token

	return layer
}
