package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	DB, _ = gorm.Open(sqlite.Open("gorm.db"))
	DB.AutoMigrate(&Book{})
	DB.AutoMigrate(&Movie{})
}

type DbEntity[T any] struct {
	Id int `json:"id" gorm:"primarykey"`
}

func (R DbEntity[T]) DbList() []T {
	var rows []T
	DB.Find(&rows)
	return rows
}

func (R DbEntity[T]) DbFind(id int) *T {
	var r T
	DB.Find(&r, map[string]interface{}{"id": id})
	return &r
}

func (R DbEntity[T]) DbInsert(t *T) {
	DB.Omit("Id").Create(t)
}

func (R *DbEntity[T]) DbUpdate(t *T, id int) {
	DB.Model(t).Where("id = ?", id).Updates(t)
	R.Id = id
}

func (R DbEntity[T]) DbDelete(id int) {
	var r T
	DB.Model(r).Where("id = ?", id).Delete(&r)
}

type Book struct {
	DbEntity[Book]
	Title  string `json:"title"`
	Author string `json:"author"`
	Date   int    `json:"date,string"`
}

type Movie struct {
	DbEntity[Movie]
	Codec    string `json:"codec"`
	Starring string `json:"starring"`
	Rating   int    `json:"rating,string"`
}

type Collectible interface {
	Book | Movie
}

type Storable[T Collectible] interface {
	DbList() []T
	DbFind(int) *T
	DbInsert(*T)
	DbUpdate(*T, int)
	DbDelete(int)
}
