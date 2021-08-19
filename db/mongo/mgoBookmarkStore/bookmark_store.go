package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Bookmark struct {
	ID                          bson.ObjectId `bson:"_id,omitempty"`
	Name, Description, Location string
	Priority                    int
	CreatedOn                   time.Time
	Tags                        []string
}

type BookmarkStore struct {
	C *mgo.Collection
}

func (store BookmarkStore) Create(b *Bookmark) error {
	b.ID = bson.NewObjectId()
	err := store.C.Insert(b)
	return err
}

func (store BookmarkStore) Update(b Bookmark) error {
	err := store.C.Update(bson.M{"_id": b.ID},
		bson.M{"$set": bson.M{
			"name":        b.Name,
			"description": b.Description,
			"location":    b.Location,
			"priority":    b.Priority,
			"tags":        b.Tags,
		}})
	return err
}

func (store BookmarkStore) Delete(id string) error {
	err := store.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (store BookmarkStore) GetAll() []Bookmark {
	var b []Bookmark
	iter := store.C.Find(nil).Sort("prority", "-createdon").Iter()
	result := Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}

func (store BookmarkStore) GetByID(id string) (Bookmark, error) {
	var b Bookmark
	err := store.C.FindId(bson.ObjectIdHex(id)).One(&b)
	return b, err
}

func (store BookmarkStore) GetByTags(tags []string) []Bookmark {
	var b []Bookmark
	iter := store.C.Find(bson.M{"tags": bson.M{"$in": tags}}).Sort("priority", "-createdon").Iter()
	result := Bookmark{}
	for iter.Next(&result) {
		b = append(b, result)
	}
	return b
}
