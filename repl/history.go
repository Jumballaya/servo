package repl

type History struct {
	Current int
	History []string
}

func NewHistory() *History {
	return &History{Current: 0, History: []string{""}}
}

func (h *History) Back() {
	length := len(h.History)
	if h.Current < 0 || h.Current > length-1 {
		h.Current = 0
		return
	}

	h.Current--
}

func (h *History) Forward() {
	length := len(h.History)
	if h.Current < 0 {
		h.Current = 0
		return
	}
	if h.Current > length-1 {
		h.Current = length - 1
		return
	}

	h.Current += 1
}

func (h *History) CurrentInput() string {
	return h.History[h.Current]
}

func (h *History) Insert(command string) {
	h.History[h.Current] = command
	h.Current = len(h.History) + 1
	h.History = append(h.History, "")
}
