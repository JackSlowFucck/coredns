package erratic

import "sync/atomic"

// Ready only signals readiness after the first query.
func (e *Erratic) Ready() bool {
	q := atomic.LoadUint64(&e.q)
	return q > 0
}
