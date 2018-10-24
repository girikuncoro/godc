package libvirt

import libvirt "github.com/libvirt/libvirt-go"

// StreamIO represents libvirt streaming
type StreamIO struct {
	Stream libvirt.Stream
}

// NewStreamIO creates new streamIO instance
func NewStreamIO(s libvirt.Stream) *StreamIO {
	return &StreamIO{Stream: s}
}

// Read reads byte stream
func (sio *StreamIO) Read(p []byte) (int, error) {
	return sio.Stream.Recv(p)
}

// Write writes byte stream
func (sio *StreamIO) Write(p []byte) (int, error) {
	return sio.Stream.Send(p)
}

// Close closes the stream connection
func (sio *StreamIO) Close() error {
	return sio.Stream.Finish()
}
