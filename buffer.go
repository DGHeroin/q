package q

type (
    Buffer interface {
        Close() error
        Read(data []byte) (n int, err error)
        Write(data []byte) (n int, err error)
        WriteString(s string) (n int, err error)
    }
)
