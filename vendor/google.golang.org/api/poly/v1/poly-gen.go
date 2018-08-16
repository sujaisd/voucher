// Package poly provides access to the Poly API.
//
// See https://developers.google.com/poly/
//
// Usage example:
//
//   import "google.golang.org/api/poly/v1"
//   ...
//   polyService, err := poly.New(oauthHttpClient)
package poly // import "google.golang.org/api/poly/v1"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "poly:v1"
const apiName = "poly"
const apiVersion = "v1"
const basePath = "https://poly.googleapis.com/"

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Assets = NewAssetsService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Assets *AssetsService

	Users *UsersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAssetsService(s *Service) *AssetsService {
	rs := &AssetsService{s: s}
	return rs
}

type AssetsService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	rs.Assets = NewUsersAssetsService(s)
	rs.Likedassets = NewUsersLikedassetsService(s)
	return rs
}

type UsersService struct {
	s *Service

	Assets *UsersAssetsService

	Likedassets *UsersLikedassetsService
}

func NewUsersAssetsService(s *Service) *UsersAssetsService {
	rs := &UsersAssetsService{s: s}
	return rs
}

type UsersAssetsService struct {
	s *Service
}

func NewUsersLikedassetsService(s *Service) *UsersLikedassetsService {
	rs := &UsersLikedassetsService{s: s}
	return rs
}

type UsersLikedassetsService struct {
	s *Service
}

// Asset: Represents and describes an asset in the Poly library. An
// asset is a 3D model
// or scene created using [Tilt
// Brush](//www.tiltbrush.com),
// [Blocks](//vr.google.com/blocks/), or any 3D program that produces a
// file
// that can be upload to Poly.
type Asset struct {
	// AuthorName: The author's publicly visible name. Use this name when
	// giving credit to the
	// author. For more information, see
	// [Licensing](/poly/discover/licensing).
	AuthorName string `json:"authorName,omitempty"`

	// CreateTime: For published assets, the time when the asset was
	// published.
	// For unpublished assets, the time when the asset was created.
	CreateTime string `json:"createTime,omitempty"`

	// Description: The human-readable description, set by the asset's
	// author.
	Description string `json:"description,omitempty"`

	// DisplayName: The human-readable name, set by the asset's author.
	DisplayName string `json:"displayName,omitempty"`

	// Formats: A list of Formats where each
	// format describes one representation of the asset.
	Formats []*Format `json:"formats,omitempty"`

	// IsCurated: Whether this asset has been curated by the Poly team.
	IsCurated bool `json:"isCurated,omitempty"`

	// License: The license under which the author has made the asset
	// available
	// for use, if any.
	//
	// Possible values:
	//   "UNKNOWN" - Unknown license value.
	//   "CREATIVE_COMMONS_BY" - Creative Commons CC-BY 3.0.
	// https://creativecommons.org/licenses/by/3.0/
	//   "ALL_RIGHTS_RESERVED" - Unlicensed: All Rights Reserved by the
	// author. Unlicensed assets are
	// **not** returned by List Assets.
	License string `json:"license,omitempty"`

	// Metadata: Application-defined opaque metadata for this asset. This
	// field is only
	// returned when querying for the signed-in user's own assets, not for
	// public
	// assets. This string is limited to 1K chars. It is up to the creator
	// of
	// the asset to define the format for this string (for example, JSON).
	Metadata string `json:"metadata,omitempty"`

	// Name: The unique identifier for the asset in the
	// form:
	// `assets/{ASSET_ID}`.
	Name string `json:"name,omitempty"`

	// PresentationParams: Hints for displaying the asset. Note that these
	// parameters are not
	// immutable; the author of an asset may change them post-publication.
	PresentationParams *PresentationParams `json:"presentationParams,omitempty"`

	// RemixInfo: The remix info for the asset.
	RemixInfo *RemixInfo `json:"remixInfo,omitempty"`

	// Thumbnail: The thumbnail image for the asset.
	Thumbnail *File `json:"thumbnail,omitempty"`

	// UpdateTime: The time when the asset was last modified. For published
	// assets, whose
	// contents are immutable, the update time changes only when
	// metadata
	// properties, such as visibility, are updated.
	UpdateTime string `json:"updateTime,omitempty"`

	// Visibility: The visibility of the asset and who can access it.
	//
	// Possible values:
	//   "VISIBILITY_UNSPECIFIED" - Unknown (and invalid) visibility.
	//   "PRIVATE" - Access to the asset and its underlying files and
	// resources is restricted to
	// the author.
	// **Authentication:** You must supply an OAuth token that corresponds
	// to the
	// author's account.
	//   "UNLISTED" - Access to the asset and its underlying files and
	// resources is available to
	// anyone with the asset's name. Unlisted assets are **not**
	// returned by List Assets.
	//   "PUBLIC" - Access to the asset and its underlying files and
	// resources is available
	// to anyone.
	Visibility string `json:"visibility,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "AuthorName") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AuthorName") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Asset) MarshalJSON() ([]byte, error) {
	type NoMethod Asset
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// AssetImportMessage: A message generated by the asset import process.
type AssetImportMessage struct {
	// Code: The code associated with this message.
	//
	// Possible values:
	//   "CODE_UNSPECIFIED" - Unknown error code.
	//   "NO_IMPORTABLE_FILE" - The asset import did not include any file
	// that we can import (i.e. an OBJ
	// file).
	//   "EMPTY_MODEL" - When generating the preview for the import, no
	// geometry was found.
	//   "OBJ_PARSE_ERROR" - A problem was encountered while parsing the OBJ
	// file. The converter makes
	// a 'best effort' attempt to continue when encountering such issues.
	// In
	// some cases the resulting preview model may still be acceptable.
	// The
	// details can be found in the parse error message.
	//   "EXPIRED" - The importer was not able to import the model before
	// the expiration time.
	//   "IMAGE_ERROR" - The importer encountered a problem reading an image
	// file.
	//   "EXTRA_FILES_WITH_ARCHIVE" - Multiple files were encountered in
	// addition to a ZIP archive. When
	// uploading an archive only one file is permitted.
	//   "DEFAULT_MATERIALS" - Default materials are used in the model. This
	// means that one or more
	// faces is using default materials either because no usemtl statement
	// was
	// specified or because the requested material was not found due to
	// a
	// missing material file or bad material name. This does not cover the
	// case
	// of missing textures.
	//   "FATAL_ERROR" - The importer encountered a fatal error and was
	// unable to import the
	// model.
	//   "INVALID_ELEMENT_TYPE" - The import includes a file of an
	// unsupported element type. The file path
	// is specified.
	Code string `json:"code,omitempty"`

	// FilePath: An optional file path. Only present for those error codes
	// that specify it.
	FilePath string `json:"filePath,omitempty"`

	// ImageError: An optional image error. Only present for
	// INVALID_IMAGE_FILE.
	ImageError *ImageError `json:"imageError,omitempty"`

	// ObjParseError: An optional OBJ parse error. Only present for
	// OBJ_PARSE_ERROR.
	ObjParseError *ObjParseError `json:"objParseError,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *AssetImportMessage) MarshalJSON() ([]byte, error) {
	type NoMethod AssetImportMessage
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// File: Represents a file in Poly, which can be a root,
// resource, or thumbnail file.
type File struct {
	// ContentType: The MIME content-type, such as `image/png`.
	// For more information, see
	// [MIME
	// types](//developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME
	// _types).
	ContentType string `json:"contentType,omitempty"`

	// RelativePath: The path of the resource file relative to the root
	// file.
	// For root or thumbnail files, this is just the filename.
	RelativePath string `json:"relativePath,omitempty"`

	// Url: The URL where the file data can be retrieved.
	Url string `json:"url,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ContentType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentType") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *File) MarshalJSON() ([]byte, error) {
	type NoMethod File
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Format: The same asset can be represented in different formats, for
// example,
// a [WaveFront .obj](//en.wikipedia.org/wiki/Wavefront_.obj_file) file
// with its
// corresponding .mtl file or a [Khronos glTF](//www.khronos.org/gltf)
// file
// with its corresponding .glb binary data. A format refers to a
// specific
// representation of an asset and contains all information needed
// to
// retrieve and describe this representation.
type Format struct {
	// FormatComplexity: Complexity stats about this representation of the
	// asset.
	FormatComplexity *FormatComplexity `json:"formatComplexity,omitempty"`

	// FormatType: A short string that identifies the format type of this
	// representation.
	// Possible values are: `FBX`, `GLTF`, `GLTF2`, `OBJ`, and `TILT`.
	FormatType string `json:"formatType,omitempty"`

	// Resources: A list of dependencies of the root element. May include,
	// but is not
	// limited to, materials, textures, and shader programs.
	Resources []*File `json:"resources,omitempty"`

	// Root: The root of the file hierarchy. This will always be
	// populated.
	// For some format_types - such as `TILT`, which are self-contained
	// -
	// this is all of the data.
	//
	// Other types - such as `OBJ` - often reference other data
	// elements.
	// These are contained in the resources field.
	Root *File `json:"root,omitempty"`

	// ForceSendFields is a list of field names (e.g. "FormatComplexity") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "FormatComplexity") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Format) MarshalJSON() ([]byte, error) {
	type NoMethod Format
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// FormatComplexity: Information on the complexity of this Format.
type FormatComplexity struct {
	// LodHint: A non-negative integer that represents the level of detail
	// (LOD) of this
	// format relative to other formats of the same asset with the
	// same
	// format_type.
	// This hint allows you to sort formats from the most-detailed (0)
	// to
	// least-detailed (integers greater than 0).
	LodHint int64 `json:"lodHint,omitempty"`

	// TriangleCount: The estimated number of triangles.
	TriangleCount int64 `json:"triangleCount,omitempty,string"`

	// ForceSendFields is a list of field names (e.g. "LodHint") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "LodHint") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *FormatComplexity) MarshalJSON() ([]byte, error) {
	type NoMethod FormatComplexity
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ImageError: A message resulting from reading an image file.
type ImageError struct {
	// Code: The type of image error encountered. Optional for older image
	// errors.
	//
	// Possible values:
	//   "CODE_UNSPECIFIED" - Unknown error code.
	//   "INVALID_IMAGE" - We were unable to read the image file.
	//   "IMAGE_TOO_BIG" - The image size is too large.
	//   "WRONG_IMAGE_TYPE" - The image data does not match the expected
	// MIME type of the image.
	Code string `json:"code,omitempty"`

	// FilePath: The file path in the import of the image that was rejected.
	FilePath string `json:"filePath,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ImageError) MarshalJSON() ([]byte, error) {
	type NoMethod ImageError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListAssetsResponse: A response message from a request to list.
type ListAssetsResponse struct {
	// Assets: A list of assets that match the criteria specified in the
	// request.
	Assets []*Asset `json:"assets,omitempty"`

	// NextPageToken: The continuation token for retrieving the next page.
	// If empty,
	// indicates that there are no more pages. To get the next page, submit
	// the
	// same request specifying this value as the
	// page_token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total number of assets in the list, without
	// pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Assets") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Assets") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListAssetsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListAssetsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListLikedAssetsResponse: A response message from a request to list.
type ListLikedAssetsResponse struct {
	// Assets: A list of assets that match the criteria specified in the
	// request.
	Assets []*Asset `json:"assets,omitempty"`

	// NextPageToken: The continuation token for retrieving the next page.
	// If empty,
	// indicates that there are no more pages. To get the next page, submit
	// the
	// same request specifying this value as the
	// page_token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total number of assets in the list, without
	// pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Assets") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Assets") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListLikedAssetsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListLikedAssetsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListUserAssetsResponse: A response message from a request to list.
type ListUserAssetsResponse struct {
	// NextPageToken: The continuation token for retrieving the next page.
	// If empty,
	// indicates that there are no more pages. To get the next page, submit
	// the
	// same request specifying this value as the
	// page_token.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// TotalSize: The total number of assets in the list, without
	// pagination.
	TotalSize int64 `json:"totalSize,omitempty"`

	// UserAssets: A list of UserAssets matching the request.
	UserAssets []*UserAsset `json:"userAssets,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "NextPageToken") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "NextPageToken") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListUserAssetsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListUserAssetsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ObjParseError: Details of an error resulting from parsing an OBJ file
type ObjParseError struct {
	// Code: The type of problem found (required).
	//
	// Possible values:
	//   "CODE_UNSPECIFIED" - Unknown error code.
	//   "INCONSISTENT_VERTEX_REFS" - Vertex references are specified in an
	// inconsistent style for a face (e.g.
	// some vertices specify texture vertices but some don't).
	//   "INVALID_COMMAND" - The command is invalid.
	//   "INVALID_NUMBER" - A invalid number was specified.
	//   "INVALID_VERTEX_REF" - An invalid vertex reference was specified.
	//   "MISSING_GEOMETRIC_VERTEX" - A vertex reference does not specify a
	// geometric vertex.
	//   "MISSING_TOKEN" - An expected token was not found.
	//   "TOO_FEW_DIMENSIONS" - The vertex specified too few dimensions for
	// its usage.
	//   "TOO_FEW_VERTICES" - The face specified too few vertices.
	//   "TOO_MANY_DIMENSIONS" - The vertex specified too many dimensions
	// for its usage.
	//   "UNSUPPORTED_COMMAND" - This command is a valid OBJ command but is
	// not supported. This error is
	// only generated for the first instance of such a command.
	//   "UNUSED_TOKENS" - This line ended with unparsed token characters.
	//   "VERTEX_NOT_FOUND" - The specified vertex was not found.
	//   "NUMBER_OUT_OF_RANGE" - The specified number was too large or small
	// for its usage.
	//   "INVALID_VALUE" - The specified parameter value was not recognized.
	//   "INVALID_TEXTURE_OPTION" - The specified texture option is not
	// valid.
	//   "TOO_MANY_PROBLEMS" - The maximum number of problems to report was
	// reached. Parsing continues,
	// but further problems will be ignored.
	//   "MISSING_FILE_NAME" - An expected file name was not specified.
	//   "FILE_NOT_FOUND" - The specified file was not found in the import.
	//   "UNKNOWN_MATERIAL" - The specified material was not found in any
	// material definition in the
	// import.
	//   "NO_MATERIAL_DEFINED" - Material parameters were specified before
	// the first material definition.
	//   "INVALID_SMOOTHING_GROUP" - The smoothing group is not valid.
	//   "MISSING_VERTEX_COLORS" - Vertex colors were specified for only
	// some vertices of a face.
	//   "FILE_SUBSTITUTION" - A missing file was found at a different file
	// path.
	//   "LINE_TOO_LONG" - A line in an OBJ or MTL file exceeded the maximum
	// line length.
	//   "INVALID_FILE_PATH" - The file path was invalid. Only relative
	// paths are supported.
	Code string `json:"code,omitempty"`

	// EndIndex: The ending character index at which the problem was found.
	EndIndex int64 `json:"endIndex,omitempty"`

	// FilePath: The file path in which the problem was found.
	FilePath string `json:"filePath,omitempty"`

	// Line: The text of the line. Note that this may be truncated if the
	// line was very
	// long. This may not include the error if it occurs after line
	// truncation.
	Line string `json:"line,omitempty"`

	// LineNumber: Line number at which the problem was found.
	LineNumber int64 `json:"lineNumber,omitempty"`

	// StartIndex: The starting character index at which the problem was
	// found.
	StartIndex int64 `json:"startIndex,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ObjParseError) MarshalJSON() ([]byte, error) {
	type NoMethod ObjParseError
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// PresentationParams: Hints for displaying the asset, based on
// information available when the asset
// was uploaded.
type PresentationParams struct {
	// BackgroundColor: A background color which could be used for
	// displaying the 3D asset in a
	// 'thumbnail' or 'palette' style view. Authors have the option to set
	// this
	// background color when publishing or editing their asset.
	//
	// This is represented as a six-digit hexademical triplet specifying
	// the
	// RGB components of the background color, e.g. #FF0000 for Red.
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// ColorSpace: The materials' diffuse/albedo color. This does not apply
	// to vertex colors
	// or texture maps.
	//
	// Possible values:
	//   "UNKNOWN" - Invalid color value.
	//   "LINEAR" - Linear color values. Default.
	//   "GAMMA" - Colors should be converted to linear by assuming gamma =
	// 2.0.
	ColorSpace string `json:"colorSpace,omitempty"`

	// OrientingRotation: A rotation that should be applied to the object
	// root to make it upright.
	// More precisely, this quaternion transforms from "object space" (the
	// space
	// in which the object is defined) to "presentation space", a
	// coordinate
	// system where +Y is up, +X is right, -Z is forward. For example,
	// if
	// the object is the Eiffel Tower, in its local coordinate system
	// the
	// object might be laid out such that the base of the tower is on the
	// YZ plane and the tip of the tower is towards positive X. In this
	// case
	// this quaternion would specify a rotation (of 90 degrees about the
	// Z
	// axis) such that in the presentation space the base of the tower
	// is
	// aligned with the XZ plane, and the tip of the tower lies towards
	// +Y.
	//
	// This rotation is unrelated to the object's pose in the web
	// preview,
	// which is just a camera position setting and is *not* reflected in
	// this
	// rotation.
	//
	// Please note: this is applicable only to the gLTF.
	OrientingRotation *Quaternion `json:"orientingRotation,omitempty"`

	// ForceSendFields is a list of field names (e.g. "BackgroundColor") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "BackgroundColor") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *PresentationParams) MarshalJSON() ([]byte, error) {
	type NoMethod PresentationParams
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Quaternion: A [Quaternion](//en.wikipedia.org/wiki/Quaternion).
// Please note: if in the
// response you see "w: 1" and nothing else this is the default value
// of
// [0, 0, 0, 1] where x,y, and z are 0.
type Quaternion struct {
	// W: The scalar component.
	W float64 `json:"w,omitempty"`

	// X: The x component.
	X float64 `json:"x,omitempty"`

	// Y: The y component.
	Y float64 `json:"y,omitempty"`

	// Z: The z component.
	Z float64 `json:"z,omitempty"`

	// ForceSendFields is a list of field names (e.g. "W") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "W") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Quaternion) MarshalJSON() ([]byte, error) {
	type NoMethod Quaternion
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

func (s *Quaternion) UnmarshalJSON(data []byte) error {
	type NoMethod Quaternion
	var s1 struct {
		W gensupport.JSONFloat64 `json:"w"`
		X gensupport.JSONFloat64 `json:"x"`
		Y gensupport.JSONFloat64 `json:"y"`
		Z gensupport.JSONFloat64 `json:"z"`
		*NoMethod
	}
	s1.NoMethod = (*NoMethod)(s)
	if err := json.Unmarshal(data, &s1); err != nil {
		return err
	}
	s.W = float64(s1.W)
	s.X = float64(s1.X)
	s.Y = float64(s1.Y)
	s.Z = float64(s1.Z)
	return nil
}

// RemixInfo: Info about the sources of this asset (i.e. assets that
// were remixed to
// create this asset).
type RemixInfo struct {
	// SourceAsset: Resource ids for the sources of this remix, of the
	// form:
	// `assets/{ASSET_ID}`
	SourceAsset []string `json:"sourceAsset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "SourceAsset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "SourceAsset") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *RemixInfo) MarshalJSON() ([]byte, error) {
	type NoMethod RemixInfo
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// StartAssetImportResponse: A response message from a request to
// startImport.
// This is returned in the response field of the Operation.
type StartAssetImportResponse struct {
	// AssetId: The id of newly created asset. If this is empty when the
	// operation is
	// complete it means the import failed. Please refer to
	// the
	// assetImportMessages field to understand what went wrong.
	AssetId string `json:"assetId,omitempty"`

	// AssetImportId: The id of the asset import.
	AssetImportId string `json:"assetImportId,omitempty"`

	// AssetImportMessages: The message from the asset import. This will
	// contain any warnings
	// (or - in the case of failure - errors) that occurred during import.
	AssetImportMessages []*AssetImportMessage `json:"assetImportMessages,omitempty"`

	// PublishUrl: The publish URL for the asset.
	PublishUrl string `json:"publishUrl,omitempty"`

	// ForceSendFields is a list of field names (e.g. "AssetId") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "AssetId") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *StartAssetImportResponse) MarshalJSON() ([]byte, error) {
	type NoMethod StartAssetImportResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// UserAsset: Data about the user's asset.
type UserAsset struct {
	// Asset: An Asset.
	Asset *Asset `json:"asset,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Asset") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Asset") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *UserAsset) MarshalJSON() ([]byte, error) {
	type NoMethod UserAsset
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "poly.assets.get":

type AssetsGetCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Get: Returns detailed information about an asset given its
// name.
// PRIVATE assets are returned only if
//  the currently authenticated user (via OAuth token) is the author of
// the asset.
func (r *AssetsService) Get(name string) *AssetsGetCall {
	c := &AssetsGetCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsGetCall) Fields(s ...googleapi.Field) *AssetsGetCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetsGetCall) IfNoneMatch(entityTag string) *AssetsGetCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsGetCall) Context(ctx context.Context) *AssetsGetCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetsGetCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetsGetCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "poly.assets.get" call.
// Exactly one of *Asset or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Asset.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *AssetsGetCall) Do(opts ...googleapi.CallOption) (*Asset, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Asset{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns detailed information about an asset given its name.\nPRIVATE assets are returned only if\n the currently authenticated user (via OAuth token) is the author of the asset.",
	//   "flatPath": "v1/assets/{assetsId}",
	//   "httpMethod": "GET",
	//   "id": "poly.assets.get",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. An asset's name in the form `assets/{ASSET_ID}`.",
	//       "location": "path",
	//       "pattern": "^assets/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "Asset"
	//   }
	// }

}

// method id "poly.assets.list":

type AssetsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists all public, remixable assets. These are assets with an
// access level of
// PUBLIC and published under the
// CC-By license.
func (r *AssetsService) List() *AssetsListCall {
	c := &AssetsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Category sets the optional parameter "category": Filter assets based
// on the specified category. Supported values are:
// `animals`, `architecture`, `art`, `food`, `nature`, `objects`,
// `people`, `scenes`,
// `technology`, and `transport`.
func (c *AssetsListCall) Category(category string) *AssetsListCall {
	c.urlParams_.Set("category", category)
	return c
}

// Curated sets the optional parameter "curated": Return only assets
// that have been curated by the Poly team.
func (c *AssetsListCall) Curated(curated bool) *AssetsListCall {
	c.urlParams_.Set("curated", fmt.Sprint(curated))
	return c
}

// Format sets the optional parameter "format": Return only assets with
// the matching format. Acceptable values are:
// `BLOCKS`, `FBX`, `GLTF`, `GLTF2`, `OBJ`, `TILT`.
func (c *AssetsListCall) Format(format string) *AssetsListCall {
	c.urlParams_.Set("format", format)
	return c
}

// Keywords sets the optional parameter "keywords": One or more search
// terms to be matched against all text that Poly has
// indexed for assets, which includes display_name,
// description, and tags. Multiple keywords should be
// separated by spaces.
func (c *AssetsListCall) Keywords(keywords string) *AssetsListCall {
	c.urlParams_.Set("keywords", keywords)
	return c
}

// MaxComplexity sets the optional parameter "maxComplexity": Returns
// assets that are of the specified complexity or less. Defaults
// to
// COMPLEX. For example, a request for
// MEDIUM assets also includes
// SIMPLE assets.
//
// Possible values:
//   "COMPLEXITY_UNSPECIFIED"
//   "COMPLEX"
//   "MEDIUM"
//   "SIMPLE"
func (c *AssetsListCall) MaxComplexity(maxComplexity string) *AssetsListCall {
	c.urlParams_.Set("maxComplexity", maxComplexity)
	return c
}

// OrderBy sets the optional parameter "orderBy": Specifies an ordering
// for assets. Acceptable values are:
// `BEST`, `NEWEST`, `OLDEST`. Defaults to `BEST`, which ranks
// assets
// based on a combination of popularity and other features.
func (c *AssetsListCall) OrderBy(orderBy string) *AssetsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of assets to be returned. This value must be between `1`
// and `100`. Defaults to `20`.
func (c *AssetsListCall) PageSize(pageSize int64) *AssetsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a
// continuation token from a previous search whose results were
// split into multiple pages. To get the next page, submit the same
// request
// specifying the value from next_page_token.
func (c *AssetsListCall) PageToken(pageToken string) *AssetsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AssetsListCall) Fields(s ...googleapi.Field) *AssetsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AssetsListCall) IfNoneMatch(entityTag string) *AssetsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AssetsListCall) Context(ctx context.Context) *AssetsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AssetsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AssetsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/assets")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "poly.assets.list" call.
// Exactly one of *ListAssetsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListAssetsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AssetsListCall) Do(opts ...googleapi.CallOption) (*ListAssetsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListAssetsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all public, remixable assets. These are assets with an access level of\nPUBLIC and published under the\nCC-By license.",
	//   "flatPath": "v1/assets",
	//   "httpMethod": "GET",
	//   "id": "poly.assets.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "category": {
	//       "description": "Filter assets based on the specified category. Supported values are:\n`animals`, `architecture`, `art`, `food`, `nature`, `objects`, `people`, `scenes`,\n`technology`, and `transport`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "curated": {
	//       "description": "Return only assets that have been curated by the Poly team.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "format": {
	//       "description": "Return only assets with the matching format. Acceptable values are:\n`BLOCKS`, `FBX`, `GLTF`, `GLTF2`, `OBJ`, `TILT`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "keywords": {
	//       "description": "One or more search terms to be matched against all text that Poly has\nindexed for assets, which includes display_name,\ndescription, and tags. Multiple keywords should be\nseparated by spaces.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "maxComplexity": {
	//       "description": "Returns assets that are of the specified complexity or less. Defaults to\nCOMPLEX. For example, a request for\nMEDIUM assets also includes\nSIMPLE assets.",
	//       "enum": [
	//         "COMPLEXITY_UNSPECIFIED",
	//         "COMPLEX",
	//         "MEDIUM",
	//         "SIMPLE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Specifies an ordering for assets. Acceptable values are:\n`BEST`, `NEWEST`, `OLDEST`. Defaults to `BEST`, which ranks assets\nbased on a combination of popularity and other features.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of assets to be returned. This value must be between `1`\nand `100`. Defaults to `20`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a continuation token from a previous search whose results were\nsplit into multiple pages. To get the next page, submit the same request\nspecifying the value from next_page_token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/assets",
	//   "response": {
	//     "$ref": "ListAssetsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *AssetsListCall) Pages(ctx context.Context, f func(*ListAssetsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "poly.users.assets.list":

type UsersAssetsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists assets authored by the given user. Only the value 'me',
// representing
// the currently-authenticated user, is supported. May include assets
// with an
// access level of PRIVATE or
// UNLISTED and assets which are
// All Rights Reserved for the
// currently-authenticated user.
func (r *UsersAssetsService) List(name string) *UsersAssetsListCall {
	c := &UsersAssetsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Format sets the optional parameter "format": Return only assets with
// the matching format. Acceptable values are:
// `BLOCKS`, `FBX`, `GLTF`, `GLTF2`, `OBJ`, and `TILT`.
func (c *UsersAssetsListCall) Format(format string) *UsersAssetsListCall {
	c.urlParams_.Set("format", format)
	return c
}

// OrderBy sets the optional parameter "orderBy": Specifies an ordering
// for assets. Acceptable values are:
// `BEST`, `NEWEST`, `OLDEST`. Defaults to `BEST`, which ranks
// assets
// based on a combination of popularity and other features.
func (c *UsersAssetsListCall) OrderBy(orderBy string) *UsersAssetsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of assets to be returned. This value must be between `1`
// and `100`. Defaults to `20`.
func (c *UsersAssetsListCall) PageSize(pageSize int64) *UsersAssetsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a
// continuation token from a previous search whose results were
// split into multiple pages. To get the next page, submit the same
// request
// specifying the value from
// next_page_token.
func (c *UsersAssetsListCall) PageToken(pageToken string) *UsersAssetsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Visibility sets the optional parameter "visibility": The visibility
// of the assets to be returned.
// Defaults to VISIBILITY_UNSPECIFIED which returns all assets.
//
// Possible values:
//   "VISIBILITY_UNSPECIFIED"
//   "PUBLISHED"
//   "PRIVATE"
func (c *UsersAssetsListCall) Visibility(visibility string) *UsersAssetsListCall {
	c.urlParams_.Set("visibility", visibility)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersAssetsListCall) Fields(s ...googleapi.Field) *UsersAssetsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UsersAssetsListCall) IfNoneMatch(entityTag string) *UsersAssetsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersAssetsListCall) Context(ctx context.Context) *UsersAssetsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersAssetsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersAssetsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/assets")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "poly.users.assets.list" call.
// Exactly one of *ListUserAssetsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListUserAssetsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UsersAssetsListCall) Do(opts ...googleapi.CallOption) (*ListUserAssetsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListUserAssetsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists assets authored by the given user. Only the value 'me', representing\nthe currently-authenticated user, is supported. May include assets with an\naccess level of PRIVATE or\nUNLISTED and assets which are\nAll Rights Reserved for the\ncurrently-authenticated user.",
	//   "flatPath": "v1/users/{usersId}/assets",
	//   "httpMethod": "GET",
	//   "id": "poly.users.assets.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "format": {
	//       "description": "Return only assets with the matching format. Acceptable values are:\n`BLOCKS`, `FBX`, `GLTF`, `GLTF2`, `OBJ`, and `TILT`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "A valid user id. Currently, only the special value 'me', representing the\ncurrently-authenticated user is supported. To use 'me', you must pass\nan OAuth token with the request.",
	//       "location": "path",
	//       "pattern": "^users/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Specifies an ordering for assets. Acceptable values are:\n`BEST`, `NEWEST`, `OLDEST`. Defaults to `BEST`, which ranks assets\nbased on a combination of popularity and other features.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of assets to be returned. This value must be between `1`\nand `100`. Defaults to `20`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a continuation token from a previous search whose results were\nsplit into multiple pages. To get the next page, submit the same request\nspecifying the value from\nnext_page_token.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "visibility": {
	//       "description": "The visibility of the assets to be returned.\nDefaults to VISIBILITY_UNSPECIFIED which returns all assets.",
	//       "enum": [
	//         "VISIBILITY_UNSPECIFIED",
	//         "PUBLISHED",
	//         "PRIVATE"
	//       ],
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/assets",
	//   "response": {
	//     "$ref": "ListUserAssetsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *UsersAssetsListCall) Pages(ctx context.Context, f func(*ListUserAssetsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "poly.users.likedassets.list":

type UsersLikedassetsListCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Lists assets that the user has liked. Only the value 'me',
// representing
// the currently-authenticated user, is supported. May include assets
// with an
// access level of UNLISTED.
func (r *UsersLikedassetsService) List(name string) *UsersLikedassetsListCall {
	c := &UsersLikedassetsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Format sets the optional parameter "format": Return only assets with
// the matching format. Acceptable values are:
// `BLOCKS`, `FBX`, `GLTF`, `GLTF2`, `OBJ`, `TILT`.
func (c *UsersLikedassetsListCall) Format(format string) *UsersLikedassetsListCall {
	c.urlParams_.Set("format", format)
	return c
}

// OrderBy sets the optional parameter "orderBy": Specifies an ordering
// for assets. Acceptable values are:
// `BEST`, `NEWEST`, `OLDEST`, 'LIKED_TIME'. Defaults to `LIKED_TIME`,
// which
// ranks assets based on how recently they were liked.
func (c *UsersLikedassetsListCall) OrderBy(orderBy string) *UsersLikedassetsListCall {
	c.urlParams_.Set("orderBy", orderBy)
	return c
}

// PageSize sets the optional parameter "pageSize": The maximum number
// of assets to be returned. This value must be between `1`
// and `100`. Defaults to `20`.
func (c *UsersLikedassetsListCall) PageSize(pageSize int64) *UsersLikedassetsListCall {
	c.urlParams_.Set("pageSize", fmt.Sprint(pageSize))
	return c
}

// PageToken sets the optional parameter "pageToken": Specifies a
// continuation token from a previous search whose results were
// split into multiple pages. To get the next page, submit the same
// request
// specifying the value from
// next_page_token.
func (c *UsersLikedassetsListCall) PageToken(pageToken string) *UsersLikedassetsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersLikedassetsListCall) Fields(s ...googleapi.Field) *UsersLikedassetsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *UsersLikedassetsListCall) IfNoneMatch(entityTag string) *UsersLikedassetsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *UsersLikedassetsListCall) Context(ctx context.Context) *UsersLikedassetsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *UsersLikedassetsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *UsersLikedassetsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}/likedassets")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "poly.users.likedassets.list" call.
// Exactly one of *ListLikedAssetsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListLikedAssetsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *UsersLikedassetsListCall) Do(opts ...googleapi.CallOption) (*ListLikedAssetsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListLikedAssetsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists assets that the user has liked. Only the value 'me', representing\nthe currently-authenticated user, is supported. May include assets with an\naccess level of UNLISTED.",
	//   "flatPath": "v1/users/{usersId}/likedassets",
	//   "httpMethod": "GET",
	//   "id": "poly.users.likedassets.list",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "format": {
	//       "description": "Return only assets with the matching format. Acceptable values are:\n`BLOCKS`, `FBX`, `GLTF`, `GLTF2`, `OBJ`, `TILT`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "A valid user id. Currently, only the special value 'me', representing the\ncurrently-authenticated user is supported. To use 'me', you must pass\nan OAuth token with the request.",
	//       "location": "path",
	//       "pattern": "^users/[^/]+$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "orderBy": {
	//       "description": "Specifies an ordering for assets. Acceptable values are:\n`BEST`, `NEWEST`, `OLDEST`, 'LIKED_TIME'. Defaults to `LIKED_TIME`, which\nranks assets based on how recently they were liked.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageSize": {
	//       "description": "The maximum number of assets to be returned. This value must be between `1`\nand `100`. Defaults to `20`.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "Specifies a continuation token from a previous search whose results were\nsplit into multiple pages. To get the next page, submit the same request\nspecifying the value from\nnext_page_token.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}/likedassets",
	//   "response": {
	//     "$ref": "ListLikedAssetsResponse"
	//   }
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *UsersLikedassetsListCall) Pages(ctx context.Context, f func(*ListLikedAssetsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}
