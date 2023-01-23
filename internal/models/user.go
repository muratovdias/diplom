package models

type Trainer struct {
	ID       int    `korm:"pk" json:"id"`
	Name     string `korm:"size:50" json:"name"`
	Username string `korm:"size:50;iunique" json:"username"`
	Email    string `korm:"size:50;iunique" json:"email"`
	Password string `json:"password"`
}

type Client struct {
	ID       int    `korm:"pk"`
	Name     string `korm:"size:50"`
	Username string `korm:"size:50;iunique"`
	Email    string `korm:"size:50;iunique"`
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
