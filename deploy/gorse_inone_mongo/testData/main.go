package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const (
	apiKey     = "your_api_key" // 替换为你的 API Key
	apiBaseURL = "http://localhost:8088/api"
)

// User represents a user in the recommendation system
type User struct {
	UserId    string   `json:"userId"`
	Labels    []string `json:"labels"`
	Subscribe []string `json:"subscribe"`
	Comment   string   `json:"comment"`
}

// Item represents an item in the recommendation system
type Item struct {
	ItemId     string    `json:"itemId"`
	IsHidden   int       `json:"is_hidden"`
	Categories []string  `json:"categories"`
	Timestamp  time.Time `json:"timestamp"`
	Labels     []string  `json:"labels"`
	Comment    string    `json:"comment"`
}

// Feedback represents user feedback in the recommendation system
type Feedback struct {
	FeedbackType string    `json:"feedbackType"`
	UserId       string    `json:"userId"`
	ItemId       string    `json:"itemId"`
	Timestamp    time.Time `json:"timestamp"`
}

// GenerateUsers generates a list of users with random labels
func GenerateUsers(n int) []User {
	users := make([]User, n)
	labels := []string{"male", "female", "18-24", "25-34", "35-44", "tech", "fashion", "food"}

	for i := 0; i < n; i++ {
		users[i] = User{
			UserId:    fmt.Sprintf("user%d", i),
			Labels:    randomLabels(labels, 2),
			Subscribe: []string{"fish", "dog", "lion"},
			Comment:   "insect",
		}
	}
	return users
}

// GenerateItems generates a list of items with random labels and categories
func GenerateItems(n int) []Item {
	items := make([]Item, n)
	labels := []string{"technology", "gadgets", "fashion", "clothing", "food", "cooking", "sports", "health"}
	categories := []string{"electronics", "apparel", "kitchen", "outdoors", "books"}

	for i := 0; i < n; i++ {
		items[i] = Item{
			ItemId:     fmt.Sprintf("item%d", i),
			IsHidden:   0,
			Categories: randomLabels(categories, 2),
			Timestamp:  time.Now().Add(-time.Duration(rand.Intn(365*24)) * time.Hour),
			Labels:     randomLabels(labels, 2),
			Comment:    fmt.Sprintf("Comment for item %d", i),
		}
	}
	return items
}

// GenerateFeedbacks generates a list of feedbacks
func GenerateFeedbacks(nUsers, nItems int) []Feedback {
	feedbacks := make([]Feedback, nUsers*10)
	feedbackTypes := []string{"wish_list", "cart", "read"}

	for i := 0; i < nUsers*10; i++ {
		feedbacks[i] = Feedback{
			FeedbackType: feedbackTypes[rand.Intn(len(feedbackTypes))],
			UserId:       fmt.Sprintf("user%d", rand.Intn(nUsers)),
			ItemId:       fmt.Sprintf("item%d", rand.Intn(nItems)),
			Timestamp:    time.Now().Add(-time.Duration(rand.Intn(365*24)) * time.Hour),
		}
	}
	return feedbacks
}

// randomLabels generates a random selection of labels
func randomLabels(labels []string, n int) []string {
	vals := []string{}
	for i := 0; i < n; i++ {
		index := rand.Intn(len(labels))
		vals = append(vals, labels[index])
	}
	return vals
}

// batchInsert sends a batch of data to the gorse API
func batchInsert(endpoint string, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling data: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", apiBaseURL, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var responseBody bytes.Buffer
		responseBody.ReadFrom(resp.Body)
		log.Fatalf("Error response from server: %s, Status: %s", responseBody.String(), resp.Status)
	} else {
		fmt.Printf("Successfully inserted data to %s\n", endpoint)
	}
}

func main() {
	nUsers := 200
	nItems := 500

	// 生成用户数据并插入
	users := GenerateUsers(nUsers)
	for _, user := range users {
		fmt.Println(user)
		batchInsert("users", []User{user})
	}

	// 生成物品数据并插入
	items := GenerateItems(nItems)
	for _, item := range items {
		fmt.Println(item)
		batchInsert("items", []Item{item})
	}

	// 生成反馈数据并插入
	feedbacks := GenerateFeedbacks(nUsers, nItems)
	for _, item := range feedbacks {
		fmt.Println(item)
		batchInsert("feedback", []Feedback{item})
	}
}
