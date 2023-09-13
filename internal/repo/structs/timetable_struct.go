package structs

type Timetable struct {
	Holder string `reindex:"holder,,pk" json:"-"`
	Days   []Day  `reindex:"days" json:"days"`
}

type Day struct {
	Week    int      `reindex:"week" json:"week"`
	Weekday int      `reindex:"weekday" json:"weekday"`
	Lessons []Lesson `reindex:"lessons" json:"lessons"`
}

type Lesson struct {
	Name     string     `reindex:"name" json:"name"`
	Type     string     `reindex:"type" json:"type"`
	Time     Time       `reindex:"time" json:"time"`
	Audience []Audience `reindex:"audience" json:"audience"`
	Teachers []Teacher  `reindex:"teachers" json:"teachers"`
}

type Time struct {
	Start string `reindex:"start" json:"start"`
	End   string `reindex:"end" json:"end"`
}

type Audience struct {
	Audience string `reindex:"audience" json:"audience"`
}

type Teacher struct {
	Teacher string `reindex:"teacher" json:"teacher"`
}
