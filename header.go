package tabular

// Header represents dataset header.
type Header struct {
	Key   string
	Title string
}

// Headers represents dataset headers.
type Headers struct {
	items []*Header
}

// Add appends a new header.
func (h *Headers) Add(key string, title string) {
	h.items = append(h.items, &Header{key, title})
}

// Get returns header on given index.
func (h *Headers) Get(idx int) (*Header, bool) {
	if ok := h.isValidIndex(idx); ok {
		return h.items[idx], true
	}
	return nil, false
}

// Items returns slice of headers.
func (h *Headers) Items() []*Header {
	return h.items
}

// Len returns header count.
func (h *Headers) Len() int {
	return len(h.items)
}

// Empty checks emptiness of headers.
func (h *Headers) Empty() bool {
	return h.Len() == 0
}

func (h *Headers) isValidIndex(idx int) bool {
	if h.Empty() {
		return false
	}
	return idx >= 0 && idx <= h.Len()
}
