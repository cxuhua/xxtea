package xxtea

import (
	"log"
	"testing"
)

func TestXXTea(t *testing.T) {
	data := "1234567890"
	key := "xh07143#4$8912"
	d, err := Encode([]byte(data), []byte(key))
	if err != nil {
		t.Error(err)
		t.SkipNow()
	}
	log.Println(d, len(d))
	m, err := Decode(d, []byte(key))
	if err != nil {
		t.Error(err)
		t.SkipNow()
	}
	if string(m) != data {
		t.Error("test failed")
	}
}
