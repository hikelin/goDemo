package example

//Size schema
type Size struct {
	H   int32   `json:"h" bson:"h"`
	W   float32 `json:"w" bson:"w"`
	Uom string  `json:"uom" bson:"uom"`
}
