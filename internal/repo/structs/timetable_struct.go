package structs

type Timetable struct {
	Holder string `reindex:"holder,,pk"`
	Days   []Day  `reindex:"days"`
}

type Day struct {
	Week    int      `reindex:"week"`
	Weekday int      `reindex:"weekday"`
	Lessons []Lesson `reindex:"lessons"`
}

type Lesson struct {
	Name     string     `reindex:"name"`
	Type     string     `reindex:"type"`
	Time     Time       `reindex:"time"`
	Audience []Audience `reindex:"audience"`
	Teachers []Teacher  `reindex:"teachers"`
}

type Time struct {
	Start string `reindex:"start"`
	End   string `reindex:"end"`
}

type Audience struct {
	Audience string `reindex:"audience"`
}

type Teacher struct {
	Teacher string `reindex:"teacher"`
}
