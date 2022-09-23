package main

// MaxFlow make sure that there is can not be flow bigger than MaxFlow
const MaxFlow = 1_000_000_000

// Edge : edge from a to b with capacity cap and "current" flow flow
type Edge struct {
	a    int
	b    int
	cap  int
	flow int
}

// DinicContext graph with some useful entities for Dinic algorithm
type DinicContext struct {
	s   int     // source
	t   int     // sink
	n   int     // total vertexes
	d   []int   // depth of vertex in bfs
	ptr []int   // pointer in dfs
	q   []int   // vertex by layer depth in bfs
	e   []Edge  // list of edges
	g   [][]int // adjacency list, indexes are refer to edges
}

func bfs(ctx *DinicContext) bool {
	qh, qt := 0, 0
	ctx.q[0] = ctx.s
	qt += 1
	for i := 0; i < len(ctx.d); i++ {
		ctx.d[i] = -1
	}
	ctx.d[ctx.s] = 0
	for qh < qt && ctx.d[ctx.t] == -1 {
		v := ctx.q[qh]
		qh += 1
		for i := 0; i < len(ctx.g[v]); i++ {
			id := ctx.g[v][i]
			to := ctx.e[id].b
			if ctx.d[to] == -1 && ctx.e[id].flow < ctx.e[id].cap {
				ctx.q[qt] = to
				qt += 1
				ctx.d[to] = ctx.d[v] + 1
			}
		}
	}
	return ctx.d[ctx.t] != -1
}

func dfs(ctx *DinicContext, v int, flow int) int {
	if flow == 0 {
		return 0
	}
	if v == ctx.t {
		return flow
	}
	for ; ctx.ptr[v] < len(ctx.g[v]); ctx.ptr[v] += 1 {
		id := ctx.g[v][ctx.ptr[v]]
		to := ctx.e[id].b
		if ctx.d[to] != ctx.d[v]+1 {
			continue
		}
		pushable := dfs(ctx, to, min(flow, ctx.e[id].cap-ctx.e[id].flow))
		if pushable != 0 {
			ctx.e[id].flow += pushable
			ctx.e[id^1].flow -= pushable
			return pushable
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dinic(ctx *DinicContext) int {
	flow := 0
	for {
		if !bfs(ctx) {
			break
		}
		for i := 0; i < len(ctx.ptr); i++ {
			ctx.ptr[i] = 0
		}
		for pushed := dfs(ctx, ctx.s, MaxFlow); pushed > 0; pushed = dfs(ctx, ctx.s, MaxFlow) {
			flow += pushed
		}
	}
	return flow
}

func GetMaxFlow(input *Input) *Output {
	dinicContext := prepareData(input)
	flow := 0
	for adding := dinic(dinicContext); adding > 0; adding = dinic(dinicContext) {
		flow += adding
	}
	return &Output{flow: flow}
}

//======================================

// Input Parsed input
type Input struct {
}

// Output single int with resulting value of maximum flow
type Output struct {
	flow int
}

// Parsing input, different in every problem
func parseInput() *Input {
	return nil
}

// Dinic context preparing based on input, different in every problem
func prepareData(input *Input) *DinicContext {
	return nil
}
