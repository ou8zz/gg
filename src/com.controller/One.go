package main

// 实体Model
type User struct {
	Id int 		`json:"id"`
	Name  string 	`json:"name"`
	Email string 	`json:"email"`
	Tdate string 	`json:"tdate"`

}

func One(n, y int) int {
	return 1
}
