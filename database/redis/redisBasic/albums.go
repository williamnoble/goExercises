package main

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool
var ErrNoAlbum = errors.New("No Album Found matching provided criteria")

type Album struct {
	Title  string  `redis:"title"`
	Artist string  `redis:"artist"`
	Price  float64 `redis:"price"`
	Likes  int     `redis:"likes"`
}

func findAlbum(id string) (*Album, error) {
	conn := pool.Get() //Get a single Redis conn
	defer conn.Close()
	values, err := redis.Values(conn.Do("HGETALL", "album:"+id))
	if err != nil {
		return nil, err
	} else if len(values) == 0 {
		return nil, ErrNoAlbum
	}
	var album Album
	err = redis.ScanStruct(values, &album)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func IncrementLikes(id string) error {
	conn := pool.Get()
	defer conn.Close()

	exists, err := redis.Int(conn.Do("EXISTS", "album:"+id))
	if err != nil {
		return err
	} else if exists == 0 {
		return ErrNoAlbum
	}

	err = conn.Send("MULTI")
	if err != nil {
		return err
	}

	err = conn.Send("HINCRBY", "album:"+id, "likes", 1)
	if err != nil {
		return err
	}
	// And we do the same with the increment on our sorted set.
	err = conn.Send("ZINCRBY", "likes", 1, id)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXEC")
	if err != nil {
		return err
	}
	return nil
}
