package encode

import (
	"os"
	"testing"
)

func BenchmarkTransferImage(b *testing.B) {

	dat, err := os.ReadFile("./file_example.txt")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		err := TransferImage(string(dat))
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkTransferImageOptimized(b *testing.B) {

	dat, err := os.ReadFile("./file_example.txt")
	if err != nil {
		b.Error(err)
	}

	for n := 0; n < b.N; n++ {
		err := TransferImageOptimized(dat)
		if err != nil {
			b.Error(err)
		}
	}
}
