package image

import "testing"

func CompressTest(t *testing.T) {
	var read, write = "image.jpg", "new.jpg"

	Compress(read, write)
}
