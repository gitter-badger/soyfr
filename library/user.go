package library

import (
	"errors"

	"github.com/maxwellhealth/bongo"
	"github.com/univedo/api2go"
	"labix.org/v2/mgo/bson"
)

//User is a generic database user
type User struct {
	ID       bson.ObjectId `bson:"_id"`
	Username string
}

//GetId Satisfy the document interface
func (u *User) GetId() bson.ObjectId {
	return u.ID
}

//SetId satisfy the document interface
func (u *User) SetId(id bson.ObjectId) {
	u.ID = id
}

//UserSource for api2go
type UserSource struct {
	Connection *bongo.Connection
}

//FindAll satisfies api2go data source interface
func (s *UserSource) FindAll(r api2go.Request) (interface{}, error) {
	var users []User
	user := User{}
	//TODO introduce paging
	resultSet := s.Connection.Collection("user").Find(bson.M{})
	if resultSet.Error != nil {
		return users, resultSet.Error
	}

	for resultSet.Next(&user) {
		users = append(users, user)
	}

	return users, nil
}

//FindOne satisfies api2go data source interface
func (s *UserSource) FindOne(ID string, r api2go.Request) (interface{}, error) {
	user := User{}
	err := s.Connection.Collection("user").FindById(bson.ObjectIdHex(ID), &user)

	return user, err
}

//FindMultiple satifies api2go data source interface
func (s *UserSource) FindMultiple(IDs []string, r api2go.Request) (interface{}, error) {
	var users []User
	user := User{}

	var findQuery []bson.ObjectId

	for _, s := range IDs {
		findQuery = append(findQuery, bson.ObjectIdHex(s))
	}

	//TODO introduce paging
	resultSet := s.Connection.Collection("user").Find(bson.M{"_id": bson.M{"$in": findQuery}})
	if resultSet.Error != nil {
		return users, resultSet.Error
	}

	for resultSet.Next(&user) {
		users = append(users, user)
	}

	return users, nil
}

//Create satisfies api2go create interface
func (s *UserSource) Create(obj interface{}) (string, error) {
	user, ok := obj.(User)
	if !ok {
		return "", errors.New("Invalid instance given")
	}

	err := s.Connection.Collection("user").Save(&user)

	if err != nil {
		return "", err
	}

	return user.GetId().Hex(), nil
}

//Delete deletes the instance
func (s *UserSource) Delete(id string) error {
	obj, err := s.FindOne(id, api2go.Request{})
	if err != nil {
		return err
	}

	user, ok := obj.(User)
	if !ok {
		return errors.New("Invalid instance given")
	}

	return s.Connection.Collection("user").Delete(&user)
}

//Update stores all changes on the user
func (s *UserSource) Update(obj interface{}) error {
	return errors.New("not implemented")
}
