package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	ID 				int
	Comments 	[]string
	Likes 		int
	Friends		[]int
}

type Response struct {
	data 	any
	err 	error
}

func main() {
	start := time.Now()
	userProfile, err := getUserProfile(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(userProfile)
	fmt.Println("fetch data took: ", time.Since(start))
}

func getUserProfile(userID int) (*UserProfile, error) {
	var (
		respCh = make(chan Response, 3)
		wg = &sync.WaitGroup{}
	)

	go getComments(userID, respCh, wg)
	go getLikes(userID, respCh, wg)
	go getFriends(userID, respCh, wg)
	wg.Add(3)
	wg.Wait()
	close(respCh)

	userProfile := &UserProfile{}

	for resp := range respCh {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []int:
			userProfile.Friends = msg
		case []string:
			userProfile.Comments = msg
		}
	}

	return userProfile, nil
}

func getComments(id int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 400)
	comments := []string{
		"Hello",
		"Whats up?",
		"Goodbye",
	}

	respCh <- Response{
		data: comments,
		err: nil,
	}

	wg.Done()
}

func getLikes(id int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 300)
	respCh <- Response{
		data: 330,
		err: nil,
	}

	wg.Done()
}

func getFriends(id int, respCh chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	friendIds := []int{11, 55, 99, 200, 403}

	respCh <- Response{
		data: friendIds,
		err: nil,
	}

	wg.Done()
}