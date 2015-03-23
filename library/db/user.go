package db

import (
	"errors"

	"github.com/maxwellhealth/bongo"
	"github.com/univedo/api2go"
	"gopkg.in/mgo.v2/bson"
)

//User is a generic database user
type User struct {
	ID           bson.ObjectId `bson:"_id"`
	Username     string
	PasswordHash string `json:"-"`
	exists       bool
}

//SetIsNew satisfies the document base
func (u *User) SetIsNew(isNew bool) {
	u.exists = !isNew
}

//IsNew satisfies the document base
func (u *User) IsNew() bool {
	return !u.exists
}

//GetId Satisfy the document interface
func (u *User) GetId() bson.ObjectId {
	return u.ID
}

//GetID to satisfy api2go interface
func (u User) GetID() string {
	return u.ID.Hex()
}

//SetId satisfy the document interface
func (u *User) SetId(id bson.ObjectId) {
	u.ID = id
}

//UserSource for api2go
type UserSource struct {
	connection *bongo.Connection
}

//CreateUserSource returns a configured and connected user source
// an error on failed connection
func CreateUserSource(config *bongo.Config) (*UserSource, error) {
	connection, err := bongo.Connect(config)
	if err != nil {
		return nil, err
	}

	return &UserSource{connection: connection}, nil
}

//FindAll satisfies api2go data source interface
func (s *UserSource) FindAll(r api2go.Request) (interface{}, error) {
	var users []User
	user := User{}
	//TODO introduce paging
	resultSet := s.connection.Collection("user").Find(bson.M{})
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
	err := s.connection.Collection("user").FindById(bson.ObjectIdHex(ID), &user)

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
	resultSet := s.connection.Collection("user").Find(bson.M{"_id": bson.M{"$in": findQuery}})
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

	err := s.connection.Collection("user").Save(&user)

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

	return s.connection.Collection("user").DeleteDocument(&user)
}

//Update stores all changes on the user
func (s *UserSource) Update(obj interface{}) error {
	//create and update are the same method in a odm
	_, err := s.Create(obj)
	return err
}
