package geecache

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int{
	return len(v.b)
}

func (v ByteView) ByteSlice() []byte{
	return cloneBytes(v.b)
}

func (v ByteView) String() string{
	return string(v.b)
}

func cloneBytes(b []byte) []byte{
	view := make([]byte, len(b))
	copy(view, b)
	return view
}