package checksum

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func GetFileChecksum(algo string, filepath string) (string, error) {
	return "", nil
}

func GetChecksum(algo string, str string) (string, error) {
	return "", nil
}

func GetMd5(str string) (string, error) {
	h := md5.New()
	io.WriteString(h, str)
	h.Sum(nil)

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func GetMd5File(filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	b := h.Sum(nil)[:16]
	return hex.EncodeToString(b), nil
}
