package dpt

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestWriteRead(t *testing.T) {
	p := Dpt{
		Trid:     3423,
		Eventime: 4324,
		Askbook:  []*Dpt_DptBar{&Dpt_DptBar{Price: 0.56, Qty: 1.6}, &Dpt_DptBar{Price: 2.6, Qty: 33.6}, &Dpt_DptBar{Price: 4.6, Qty: 3.6}},
		Bidbook:  []*Dpt_DptBar{},
	}
	outp, _ := proto.Marshal(&p)

	size := Size{
		Size: int32(len(outp)),
	}

	outsize, _ := proto.Marshal(&size)

	f, _ := os.Create("testfile")

	outsizepad := []byte{0, 0, 0, 0}
	paddsize := len(outsizepad)
	copy(outsizepad, outsize)

	f.Write(outsizepad)
	f.Write(outp)

	f.Close()

	insize := &Size{}
	inp := &Dpt{}

	iff, _ := os.Open("testfile")

	sizebuf := make([]byte, paddsize)
	io.ReadFull(iff, sizebuf)

	proto.Unmarshal(sizebuf, insize)

	pbuf := make([]byte, int(insize.Size))

	io.ReadFull(iff, pbuf)

	proto.Unmarshal(pbuf, inp)

	fmt.Println(*insize)
	fmt.Println(*inp)
}
