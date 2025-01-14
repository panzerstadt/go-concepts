package gocomposing

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	payload := []byte("hello world")

	fmt.Println("\nhash and broadcast with a bytes.NewReader")
	brokenHashAndBroadcast(bytes.NewReader(payload))

	fmt.Println("\nhash and broadcast with a NewHashReader")
	hashAndBroadcast(NewHashReader(payload))

	// hashAndBroadcast(bytes.NewReader(payload))
	// now the above no longer works with bytes.NewReader
	// cause we know we need a reader that has a hash function
	// in order to use hashAndBroadcast
	// ( broadcast() itself still just needs it to satisfy
	// bytes.NewReader to work )
}

// HashReader type extends io.Reader
type HashReader interface {
	io.Reader
	hash() string
}

// hashReader data structure extends bytes.Reader
type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	hash := sha1.Sum(h.buf.Bytes())
	return hex.EncodeToString(hash[:])
}

func hashAndBroadcast(r HashReader) error {
	hash := r.hash()
	fmt.Println(hash)
	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("string of the bytes: ", string(b))
	return nil
}

func brokenHashAndBroadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	hash := sha1.Sum(b)
	fmt.Println(hex.EncodeToString(hash[:]))
	return brokenBroadcast(r)
}

func brokenBroadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	fmt.Println("string of the bytes: ", string(b))
	return nil
}

/**
you know i just realised that maybe the issue here
is that we as programmers have become accustomed to
the idea that Read() means read only, and it doesn't
affect the object being read.

if we just think of them like generators that 'yield'
data.
or have a method called consume() to differentiate by
name from read()
*/

// this might actually read better
// now one knows that once you consume it,
// its GONE
func Consume(r io.Reader) ([]byte, error) {
	return io.ReadAll(r)
}
