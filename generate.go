package veneur

//go:generate protoc --go_out=ssf sample.proto

//go:generate gojson -input example.yaml -o config.go -fmt yaml -pkg veneur -name Config

//go:generate gojson -input fixtures/datadog/json/trace_span.json -o samplers/trace_span.go -pkg samplers -name DatadogTraceSpan
