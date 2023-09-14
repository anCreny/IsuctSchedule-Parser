package service

import (
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"github.com/anCreny/IsuctSchedule-Packages/structs"
	"main/internal/store"
)

func parseGroups(file store.ScheduleFile) []structs.Timetable {
	var (
		ans    = make([]structs.Timetable, 0)
		groups = make([]store.Group, 0)
	)

	for _, faculty := range file.Faculties {
		groups = append(groups, faculty.Groups...)
	}

	for _, group := range groups {
		//new empty timetable to be filled from file group
		tempTimetable := structs.Timetable{}

		tempTimetable.Holder = group.Number

		days := make([]structs.Day, 0)
		for week := 1; week < 3; week++ {
			for weekDay := 1; weekDay < 7; weekDay++ {
				tempDay := structs.Day{}

				tempDay.Weekday = weekDay
				tempDay.Week = 3 - week

				lessons := make([]structs.Lesson, 0)
				for _, lesson := range group.Lessons {
					if lesson.WeekDate.Week == week && lesson.WeekDate.Weekday == weekDay {
						tempLesson := structs.Lesson{}

						tempLesson.Name = lesson.Subject
						tempLesson.Type = lesson.Type

						tempLesson.Time.Start = lesson.Time.Start
						tempLesson.Time.End = lesson.Time.End

						teachers := make([]structs.Teacher, 0)
						for _, teacher := range lesson.Teachers {
							tempTeacher := structs.Teacher{}
							tempTeacher.Teacher = teacher.Name

							teachers = append(teachers, tempTeacher)
						}

						tempLesson.Teachers = teachers

						audiences := make([]structs.Audience, 0)
						for _, audience := range lesson.Audiences {
							tempAudience := structs.Audience{}
							tempAudience.Audience = audience.Number

							audiences = append(audiences, tempAudience)
						}

						tempLesson.Audience = audiences

						lessons = append(lessons, tempLesson)
					}
				}
				tempDay.Lessons = lessons

				days = append(days, tempDay)
			}
		}
		tempTimetable.Days = days

		ans = append(ans, tempTimetable)
	}

	logger.Log.Info().Msg("Groups successfully parsed")
	return ans
}

func parseTeachersNames(file store.ScheduleFile) structs.TeachersNames {
	teachersNames := structs.TeachersNames{}
	teachersNames.Names = make([]string, 0)

	for _, faculty := range file.Faculties {
		for _, group := range faculty.Groups {
			for _, lesson := range group.Lessons {
				for _, teacherName := range lesson.Teachers {
					if name := teacherName.Name; name != "—" {
						teachersNames.Names.Upsert(name)
					}
				}
			}
		}
	}

	logger.Log.Info().Msg("Teachers names successfully parsed")
	return teachersNames
}

func parseTeachers(file store.ScheduleFile) []structs.Timetable {
	ans := make([]structs.Timetable, 0)
	teachersNames := parseTeachersNames(file)

	for _, teacherName := range teachersNames.Names {
		tempTimetable := structs.Timetable{}

		tempTimetable.Holder = teacherName

		weekMap := make(map[int]structs.Day)
		for week := 1; week < 3; week++ {
			for weekDay := 1; weekDay < 7; weekDay++ {
				weekMap[week*10+weekDay] = structs.Day{
					Week:    week,
					Weekday: weekDay,
					Lessons: make([]structs.Lesson, 0),
				}
			}
		}

		for _, faculty := range file.Faculties {
			for _, group := range faculty.Groups {
				for _, lesson := range group.Lessons {
				nameSearch:
					for _, name := range lesson.Teachers {
						if name.Name == teacherName {
							for index, innerLesson := range weekMap[lesson.WeekDate.Week*10+lesson.WeekDate.Weekday].Lessons {
								if lesson.Subject == innerLesson.Name &&
									lesson.Type == innerLesson.Type &&
									lesson.Time.Start == innerLesson.Time.Start &&
									lesson.Time.End == innerLesson.Time.End {
									weekMap[lesson.WeekDate.Week*10+lesson.WeekDate.Weekday].Lessons[index].Teachers = append(innerLesson.Teachers, structs.Teacher{Teacher: group.Number})
									break nameSearch
								}
							}

							tempLesson := structs.Lesson{}

							tempLesson.Name = lesson.Subject
							tempLesson.Type = lesson.Type

							tempLesson.Time.Start = lesson.Time.Start
							tempLesson.Time.End = lesson.Time.End

							audiences := make([]structs.Audience, 0)
							for _, audience := range lesson.Audiences {
								audiences = append(audiences, structs.Audience{Audience: audience.Number})
							}

							tempLesson.Audience = audiences

							tempLesson.Teachers = append(make([]structs.Teacher, 0), structs.Teacher{Teacher: group.Number})

							day := weekMap[lesson.WeekDate.Week*10+lesson.WeekDate.Weekday]
							day.Lessons = append(weekMap[lesson.WeekDate.Week*10+lesson.WeekDate.Weekday].Lessons, tempLesson)
							weekMap[lesson.WeekDate.Week*10+lesson.WeekDate.Weekday] = day

							break
						}
					}
				}
			}
		}

		days := make([]structs.Day, 0)
		for _, day := range weekMap {
			days = append(days, day)
		}

		for _, day := range days {
			day.Week = 3 - day.Week
		}

		tempTimetable.Days = append(tempTimetable.Days, days...)

		ans = append(ans, tempTimetable)
	}

	logger.Log.Info().Msg("Teachers successfully parsed")
	return ans
}
