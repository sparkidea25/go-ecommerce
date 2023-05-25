package models

import "time"

type User struct {
  ID uint `json:"id" gorm:"primary_key"`
  FirstName string `json:"firstname"`
  LastName string `json:"lastname"`
  
//   Title string `json:"title"`
//   Description string `json:"description"`
//   Reward int `json:"reward"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
}