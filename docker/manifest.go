package docker

import (
	"net/http"

	"github.com/docker/distribution/manifest/schema2"
	"github.com/docker/distribution/reference"
)

// RequestManifest requests an Manifest for the passed repository, with the passed tag. Returns
// a schema2.Manifest, or an error if there's an issue.
func RequestManifest(token OAuthToken, ref reference.Canonical) (schema2.Manifest, error) {
	var manifest schema2.Manifest

	request, err := http.NewRequest(http.MethodGet, GetManifestURI(ref), nil)
	if nil != err {
		return manifest, err
	}

	request.Header.Add("Accept", schema2.MediaTypeManifest)
	setBearerToken(request, token)

	err = doDockerCall(request, &manifest)

	return manifest, err
}
