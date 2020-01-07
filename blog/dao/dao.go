package dao

// 文章结构类型
type Article struct {
	Id int64
	Title string
	Content string
	Time string
	Type int
	Status int
}

// 留言结构类型
type Message struct {
	Id int64
	User_id int64
	Content string
	Time string
	Status int
}

// 相册结构类型
type Photos struct {
	Id int64
	Position string
	Photo string
	Content string
	Time string
	Status int
}

// 评论结构类型
type Comment struct {
	Id int64
	Whisper_id int64
	User_id int64
	Comment string
	Time string
	Status int
}

// 用户信息结构类型
type Users struct {
	Id int64
	Ip string
	Nickname string
	Time string
	Status int
}

// 微语结构类型
type Whisper struct {
	Id int64
	Content string
	Photo string
	Agree_num int64
	Comment_num int64
	Time string
	Status int
}

// 管理员结构类型
type Admin struct {
	Id int64
	User string
	Password string
	Status int
}
