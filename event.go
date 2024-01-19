package colectivo

type Event struct {
	Bytes    []byte
	Callback func() error
}

func NewStrEvent(s string, cb func() error) {

}
