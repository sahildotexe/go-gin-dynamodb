package models

type User struct {
	UserID    string `json:"user_id" dynamodbav:"user_id"`
	Username  string `json:"username" dynamodbav:"username"`
	Password  string `json:"password" dynamodbav:"password"`
	CreatedAt int64  `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt int64  `json:"updated_at" dynamodbav:"updated_at"`
}
