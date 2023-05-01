package entities

type Response[T interface{}] struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

type Nullstruct struct {
}
