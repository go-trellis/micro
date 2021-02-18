package codec

// import "io"

// // Takes in a connection/buffer and returns a new Codec
// type NewCodec func(io.ReadWriteCloser) Codec

// // Codec encodes/decodes various types of messages used within go-micro.
// // ReadHeader and ReadBody are called in pairs to read requests/responses
// // from the connection. Close is called when finished with the
// // connection. ReadBody may be called with a nil argument to force the
// // body to be read and discarded.
// type Codec interface {
// 	Reader
// 	Writer
// 	Close() error
// 	String() string
// }

// type Reader interface {
// 	// ReadHeader(*Message, MessageType) error
// 	ReadHeader(*Message) error
// 	ReadBody(interface{}) error
// }

// type Writer interface {
// 	Write(*Message, interface{}) error
// }

// // Message represents detailed information about
// // the communication, likely followed by the body.
// // In the case of an error, body may be nil.
// type Message struct {
// 	ID string
// 	// Type     MessageType
// 	Target   string
// 	Method   string
// 	Endpoint string
// 	Error    string

// 	// The values read from the socket
// 	Header map[string]string
// 	Body   []byte
// }
