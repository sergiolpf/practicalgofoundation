package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	url := "https://go.dev"
	bid := bidOn(ctx, url)
	fmt.Println(bid)
}

func bidOn(ctx context.Context, url string) Bid {

	ch := make(chan Bid, 1)
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return defaultBid
	}
}

var defaultBid = Bid{
	AdURL: "http://adsRus.com/default",
	Price: 3,
}

func bestBid(url string) Bid {
	// Simulate work
	d := 100 * time.Millisecond
	if strings.HasPrefix(url, "https://") {
		d = 20 * time.Millisecond
	}
	time.Sleep(d)

	return Bid{
		AdURL: "http://adsRus.com/ad17",
		Price: 7,
	}
}

type Bid struct {
	AdURL string
	Price int //in Cents
}
