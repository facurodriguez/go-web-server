package main

import (
	"fmt"
	"sort"
	"strconv"

	bolt "go.etcd.io/bbolt"
)

const bucketName = "StoreBucket"

func NewBoltPlayerStore(store *bolt.DB) *BoltPlayerStore {
	store.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})

	return &BoltPlayerStore{store}
}

type BoltPlayerStore struct {
	store *bolt.DB
}

func (b *BoltPlayerStore) RecordWin(name string) {
	b.store.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		score, _ := strconv.Atoi(string(bucket.Get([]byte(name))))
		score++

		err := bucket.Put([]byte(name), []byte(strconv.Itoa(score)))
		return err
	})
}

func (b *BoltPlayerStore) GetPlayerScore(name string) int {
	score := 0
	err := b.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		value := string(bucket.Get([]byte(name)))

		score, _ = strconv.Atoi(value)

		return nil
	})

	if err != nil {
		fmt.Println("Error while trying to get player score")
	}

	return score
}

func (b *BoltPlayerStore) GetLeague() []Player {
	league := make([]Player, 0)
	err := b.store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))

		return bucket.ForEach(func(k, v []byte) error {
			wins, err := strconv.Atoi(string(v))
			player := Player{
				Name: string(k),
				Wins: wins,
			}
			league = append(league, player)

			return err
		})
	})

	if err != nil {
		fmt.Println("Error while trying to get league")
	}

	sort.Slice(league, func(i, j int) bool {
		return league[i].Wins > league[j].Wins
	})

	return league
}
