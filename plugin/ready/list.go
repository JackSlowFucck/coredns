// Package ready is used to signal readiness of the CoreDNS binary (it's a process wide thing).
package ready

import (
	"sort"
	"strings"
	"sync"
)

// l is structure that holds the plugins that signals readiness for this server block.
type l struct {
	sync.RWMutex
	rs    []Readiness
	names []string
}

// Append adds a new readiness to m.
func (l *l) Append(r Readiness, name string) {
	l.Lock()
	defer l.Unlock()
	l.rs = append(l.rs, r)
	l.names = append(l.names, name)
}

// Ready return true when all plugins ready, if the returned value is false the string
// contains a comma separated list of plugins that are not ready.
func (l *l) Ready() (bool, string) {
	l.RLock()
	defer l.RUnlock()
	ok := true
	s := []string{}
	for i, r := range l.rs {
		if !r.Ready() {
			ok = false
			s = append(s, l.names[i])
		}
	}
	if ok {
		return true, ""
	}
	sort.Strings(s)
	return false, strings.Join(s, ",")
}
