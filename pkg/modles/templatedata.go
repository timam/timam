package modles

//TemplateData holds data sent from tempates
type TemplateData struct {
	StringMap	map[string]string
	IntMap		map[string]int
	FloatMap	map[string]float32
	Data 		map[string]interface{}
}

