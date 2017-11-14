package models

// Supplier ...
type Supplier struct {
	// BaseModel
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Actions []Action `json:"actions"`
}
