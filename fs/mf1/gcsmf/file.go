package gcsmf

type GcsFile struct {
	openFlag int
	fhOffset int64
	closed   bool
	resource *gcsFileResource
}
