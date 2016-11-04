package samplers

type DatadogTraceSpan struct {
	Duration int    `json:"duration"`
	Name     string `json:"name"`
	ParentID int    `json:"parent_id"`
	Resource string `json:"resource"`
	Service  string `json:"service"`
	SpanID   int    `json:"span_id"`
	Start    int    `json:"start"`
	TraceID  int    `json:"trace_id"`
}
