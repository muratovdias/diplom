package models

import (
	"html/template"
	"time"
)

type User struct {
	ID            int       `db:"user_id" json:"id"`
	FullName      string    `db:"full_name" json:"full_name"`
	Email         string    `db:"email" json:"email"`
	Password      string    `db:"password" json:"password"`
	Role          string    `db:"role" json:"role"`
	Token         string    `db:"token" json:"token"`
	TokenDuration time.Time `db:"token_duration" json:"token_duration"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TrainerInfo struct {
	ID         int
	FullName   string
	Email      string
	Speciality string
	Phone      string
	Adress     string
	Bio        string
	Twitter    string
	Instagram  string
	Facebook   string
	Img        template.URL
}

type Trainers struct {
	ID         int    `json:"id"`
	FullName   string `json:"fullname"`
	Speciality string `json:"speciality"`
	Img        template.URL
}

type TrainerProfile struct {
	User        User
	TrainerInfo TrainerInfo
}

type TrainerSchedule struct {
	TrainerID int
	User      User
	Schedule  map[string][]Time
}

type Time struct {
	Time string
	Flag bool
}

type MainPage struct {
	Trainers []Trainers
	User     User
}

type Training struct {
	Row  int
	Name string
	Note string
	Date string
	Role string
}

type TrainerHomePage struct {
	User      User
	Trainings []Training
}

type EditProfile struct {
	Bio        string
	Speciality string
	Facebook   string
	Twitter    string
	Instagram  string
	Ava        template.URL
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
