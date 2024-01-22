package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math"
    "net/http"
    "regexp"
    "strconv"
    "strings"
    "github.com/google/uuid"
    "time"

)

type Receipt struct {
    Retailer     string `json:"retailer"`
    PurchaseDate string `json:"purchaseDate"`
    PurchaseTime string `json:"purchaseTime"`
    Total        string `json:"total"`
    Items        []Item `json:"items"`
}

type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price            string `json:"price"`
}

var receiptStore = make(map[string]int)

func calculatePoints(receipt *Receipt) int {
    points := 0

    // Points for each alphanumeric character in the retailer name
    alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
    points += len(alphanumeric.FindAllString(receipt.Retailer, -1))

    // Additional points for total
    total, err := strconv.ParseFloat(receipt.Total, 64)
    if err == nil {
        if total == math.Floor(total) {
            points += 50 // 50 points if the total is a round dollar amount
        }
        if math.Mod(total, 0.25) == 0 {
            points += 25 // 25 points if the total is a multiple of 0.25
        }
    }

    // Points for each item
    for _, item := range receipt.Items {
        descriptionLength := len(strings.TrimSpace(item.ShortDescription))
        if descriptionLength%3 == 0 {
            price, err := strconv.ParseFloat(item.Price, 64)
            if err == nil {
                points += int(math.Ceil(price * 0.2))
            }
        }
    }

    // 5 points for every two items on the receipt
    points += (len(receipt.Items) / 2) * 5

    // 6 points if the day in the purchase date is odd
    purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
    if err == nil && purchaseDate.Day()%2 != 0 {
        points += 6
    }

    // 10 points if the time of purchase is after 2:00pm and before 4:00pm
    purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
    if err == nil {
        if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
            points += 10
        }
    }

    return points
}


func processReceiptHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var receipt Receipt
    err := json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    points := calculatePoints(&receipt)
    receiptID := uuid.New().String()
    receiptStore[receiptID] = points

    response := map[string]string{"id": receiptID}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func getPointsHandler(w http.ResponseWriter, r *http.Request) {
    // Assume the URL is "/receipts/{id}/points"
    // Split the URL path and expect at least 4 segments: "", "receipts", "{id}", "points"
    parts := strings.Split(r.URL.Path, "/")

    if len(parts) != 4 || parts[1] != "receipts" || parts[3] != "points" {
        http.NotFound(w, r)
        return
    }
    receiptID := parts[2] 

    points, exists := receiptStore[receiptID]
    if !exists {
        http.NotFound(w, r)
        return
    }

    json.NewEncoder(w).Encode(map[string]int{"points": points})
}


func main() {
    http.HandleFunc("/receipts/process", processReceiptHandler)
    http.HandleFunc("/receipts/", getPointsHandler)
    fmt.Println("Server started at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

