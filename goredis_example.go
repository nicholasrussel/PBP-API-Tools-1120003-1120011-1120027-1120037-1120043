package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

/*
	Redis (Remote Dictionary Server)
	Adalah penimpanan nilai didalam memori
	Bisa dibilang kegunaanya seperti
	Database/Cache/Session/Penyimpanan Temp Data Lainnya.
*/

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func getDataRedis(con *redis.Client, key string) {
	fmt.Println("Mencoba Get Data")
	val, err := con.Get(context.Background(), key).Result()
	if err != nil {
		fmt.Println("Data Dengan Key '", key, "' Kosong")
		return
	}

	fmt.Println("GET data untuk key '", key, "' : ", val)
	fmt.Println("========================================")
}

func TestGoRedis() {

	//Port Disesuaikan
	var redisHost = "localhost:425"
	var redisPassword = ""

	con := newRedisClient(redisHost, redisPassword)

	fmt.Println("Inisialisasi Redis")
	fmt.Println("========================================")

	//Contoh Misal Ingin Menyimpan Data Level Di Sebuah Game.

	key := "level"
	data := 1
	exp := time.Duration(3) * time.Second

	ctx := context.Background()

	//Set Data
	fmt.Println("Mencoba Set Level 0")
	err := con.Set(ctx, key, data, exp).Err()
	if err != nil {
		fmt.Printf("Gagal SET data. error: %v", err)
		return
	}

	fmt.Println("SET data untuk key '", key, "' Berhasil.")
	fmt.Println("========================================")

	//Get Data
	getDataRedis(con, key)

	//Icrement Data
	fmt.Println("Mencoba Increment Level")
	err = con.Incr(ctx, key).Err()
	if err != nil {
		fmt.Printf("Gagal Increment data. error: %v", err)
		return
	}
	fmt.Println("Increment data untuk key '", key, "' Berhasil.")
	fmt.Println("========================================")

	//Get Data
	getDataRedis(con, key)

	//Delete Data
	err = con.Del(ctx, key).Err()
	if err != nil {
		fmt.Printf("Gagal Delete key. error: %v", err)
		return
	}
	fmt.Println("Delete key '", key, "' Berhasil.")
	fmt.Println("========================================")

	//Get Data
	getDataRedis(con, key)

	//Function Lainnya
	// err = con.IncrBy(ctx, "angka", 10).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// err = con.Decr(ctx, "angka").Err()
	// if err != nil {
	// 	panic(err)
	// }

	// err = con.DecrBy(ctx, "angka", 4).Err()
	// if err != nil {
	// 	panic(err)
	// }
}
