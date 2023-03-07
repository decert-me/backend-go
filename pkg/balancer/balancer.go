package balancer

import "sync"

// smoothRoundrobinNode is a wrapped weighted item.
type smoothRoundrobinNode struct {
	Item            string
	Weight          int64
	CurrentWeight   int64
	EffectiveWeight int64
}

type SmoothRoundrobin struct {
	items []*smoothRoundrobinNode
	n     int64
	mu    sync.Mutex
}

// NewSmoothRoundrobin (Smooth Weighted) contains weighted items and provides methods to select a weighted item.
// It is used for the smooth weighted round-robin balancing algorithm.
// This algorithm is implemented in Nginx:
// https://github.com/phusion/nginx/commit/27e94984486058d73157038f7950a0a36ecc6e35.
//
// Algorithm is as follows: on each peer selection we increase current_weight
// of each eligible peer by its weight, select peer with greatest current_weight
// and reduce its current_weight by total number of weight points distributed
// among peers.
// In case of { 5, 1, 1 } weights this gives the following sequence of
// current_weight's: (a, a, b, a, c, a, a)
func NewSmoothRoundrobin() *SmoothRoundrobin {
	return &SmoothRoundrobin{}
}

// Add a weighted server.
func (w *SmoothRoundrobin) Add(item string, weight int64) {
	wt := weight
	weighted := &smoothRoundrobinNode{Item: item, Weight: wt, EffectiveWeight: wt}
	w.items = append(w.items, weighted)
	w.n++
}

func (w *SmoothRoundrobin) Reset() {
	w.items = w.items[:0]
	w.n = 0
}

// Next returns next selected server.
func (w *SmoothRoundrobin) Next() (best *smoothRoundrobinNode) {
	w.mu.Lock()
	defer w.mu.Unlock()
	
	if w.n == 0 {
		return nil
	}

	if w.n == 1 {
		return w.items[0]
	}

	return nextSmoothWeighted(w.items)
}

// nextSmoothWeighted selects the best node through the smooth weighted roundrobin .
func nextSmoothWeighted(items []*smoothRoundrobinNode) (best *smoothRoundrobinNode) {
	total := int64(0)

	for i := 0; i < len(items); i++ {
		w := items[i]

		w.CurrentWeight += w.EffectiveWeight
		total += w.EffectiveWeight

		//if w.EffectiveWeight < w.Weight {
		//	w.EffectiveWeight++
		//}

		if best == nil || w.CurrentWeight > best.CurrentWeight {
			best = w
		}
	}

	best.CurrentWeight -= total

	return best
}

// OnInvokeSuccess invoke success to add effectiveWeight
func (w *smoothRoundrobinNode) OnInvokeSuccess() {
	if w.EffectiveWeight < w.Weight {
		w.EffectiveWeight++
	}
}

// OnInvokeFault invoke fault to reduce effectiveWeight
func (w *smoothRoundrobinNode) OnInvokeFault() {
	w.EffectiveWeight--
}
