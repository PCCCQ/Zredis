package data_set

type SetData struct {
	Key   string  `json:"key,omitempty"`
	Value any     `json:"value,omitempty"`
	TTL   float32 `json:"ttl,omitempty"`
}
