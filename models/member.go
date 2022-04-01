package models

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	ID    int
	Name  string `gorm:"name"`
	Email string `gorm:"email"`
}

//create a member
func CreateMember(db *gorm.DB, Member *Member) (err error) {
	err = db.Create(Member).Error
	if err != nil {
		return err
	}
	return nil
}

//get members
func GetMembers(db *gorm.DB, Member *[]Member) (err error) {
	err = db.Find(Member).Error
	if err != nil {
		return err
	}
	return nil
}

//get member by id
func GetMember(db *gorm.DB, Member *Member, id string) (err error) {
	err = db.Where("id = ?", id).First(Member).Error
	if err != nil {
		return err
	}
	return nil
}

//update member
func UpdateMember(db *gorm.DB, Member *Member) (err error) {
	db.Save(Member)
	return nil
}

//delete member
func DeleteMember(db *gorm.DB, Member *Member, id string) (err error) {
	db.Where("id = ?", id).Delete(Member)
	return nil
}
