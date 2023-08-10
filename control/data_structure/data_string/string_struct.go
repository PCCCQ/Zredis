package data_string

type StringData struct {
	Key   string  `json:"key,omitempty"`
	Value any     `json:"value,omitempty"`
	TTL   float32 `json:"ttl,omitempty"`
}
