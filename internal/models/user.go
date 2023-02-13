package models

type User struct {
	ID      int    `korm:"pk" json:"id"`
	Name    string `korm:"size:50" json:"name"`
	Surname string `korm:"size:50" json:"surname"`
	Email   string `korm:"size:50;iunique" json:"email"`
	Pwd     string `korm:"size:250" json:"password"`
}

type SignIn struct {
	Email    string
	Password string
}

// type Profile struct {
// 	Id        int       `korm:"pk"` // AUTO Increment ID primary key
// 	Name      string    // VARCHAR(40)
// 	Email     string    `korm:"size:50;iunique"` // VARCHAR(50) insensitive unique constraint
// 	Password  string    `korm:"size:150"`        // VARCHAR(150)
// 	IsAdmin   bool      `korm:"default:false"`   // NOT NULL CHECK is_admin IN (0,1) DEFAULT 0
// 	CreatedAt time.Time `korm:"now"`             // auto now
// 	UpdatedAt time.Time `korm:"update"`          // auto update
// 	Age       int
// 	AgeTwice  int `korm:"generated:age*2"`
// 	//Role string `korm:"default:'worker'"`
// }
