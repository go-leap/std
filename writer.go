package ustd

import (
	"io"
)

type Writer struct {
	io.Writer
	On struct {
		AfterEveryNth                int
		Byte                         byte
		Do                           func(int) bool
		ButDontCountImmediateRepeats bool
	}
	num  int
	last byte
}

func (me *Writer) SuspendOnDo() {
	me.num = -1
}

func (me *Writer) RestartOnDo() {
	me.num = 0
}

func (me *Writer) Write(p []byte) (n int, err error) {
	if me.On.Do == nil || me.On.AfterEveryNth <= 0 || me.num < 0 {
		n, err = me.Writer.Write(p)
	} else {
		var nw, startfrom int
		everysingletime := (me.On.AfterEveryNth == 1)
		for i, b := range p {
			if b == me.On.Byte {
				if me.On.ButDontCountImmediateRepeats && me.last == b {
					continue
				} else if me.num++; everysingletime || me.num%me.On.AfterEveryNth == 0 {
					nw, err = me.Writer.Write(p[startfrom:i])
					if n, startfrom = n+nw, i; err != nil {
						break
					} else if inclb := me.On.Do(me.num); !inclb {
						startfrom++
					}
				}
			}
			me.last = b
		}
		if err == nil && startfrom < len(p) {
			nw, err = me.Writer.Write(p[startfrom:])
			n = n + nw
		}
	}
	return
}
