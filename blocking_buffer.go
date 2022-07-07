package q

import (
    "bytes"
    "io"
    "sync"
)

type blockingBuffer struct {
    buf  bytes.Buffer
    cond *sync.Cond
}

func (b *blockingBuffer) Close() error {
    b.cond.Broadcast()
    return nil
}

func NewBlockingBuffer() Buffer {
    m := sync.Mutex{}
    return &blockingBuffer{
        cond: sync.NewCond(&m),
        buf:  bytes.Buffer{},
    }
}

func (b *blockingBuffer) Write(data []byte) (n int, err error) {
    n, err = b.buf.Write(data)
    b.cond.Broadcast()
    return
}

func (b *blockingBuffer) Read(data []byte) (n int, err error) {
    n, err = b.buf.Read(data)
    if err == io.EOF {
        b.cond.L.Lock()
        b.cond.Wait()
        b.cond.L.Unlock()
        n, err = b.buf.Read(data)
    }
    return
}

func (b *blockingBuffer) WriteString(s string) (ln int, err error) {
    return b.Write([]byte(s))
}
