package counter

type Counter interface {
	Inc()
	Dec()
	Get() int
	Set(int)
}
