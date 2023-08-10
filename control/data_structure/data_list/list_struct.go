package data_list

type ListData struct {
	Key    string  `json:"key,omitempty"`
	Value  any     `json:"value,omitempty"`
	TTL    float32 `json:"ttl,omitempty"`
	Index  int64   `json:"index,omitempty"`
	AddTab string  `json:"addTab,omitempty"`
}
