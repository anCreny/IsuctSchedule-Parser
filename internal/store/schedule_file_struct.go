package store

type ScheduleFile struct {
	Faculties []Faculty `json:"faculties"`
}

type Faculty struct {
	Groups []Group `json:"groups"`
}

type Group struct {
	Number  string   `json:"name"`
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	Subject   string     `json:"subject"`
	Type      string     `json:"type"`
	Time      LessonTime `json:"time"`
	WeekDate  WeekDate   `json:"date"`
	Audiences []Audience `json:"audiences"`
	Teachers  []Teacher  `json:"teachers"`
}

type LessonTime struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type WeekDate struct {
	Weekday int `json:"weekday"`
	Week    int `json:"week"`
}

type Audience struct {
	Number string `json:"name"`
}

type Teacher struct {
	Name string `json:"name"`
}
