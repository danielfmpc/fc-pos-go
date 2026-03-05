package main

import "context"

func main() {
	ctx := context.WithValue(context.Background(), "requestID", "12345")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	requestID := ctx.Value("requestID")
	println("Booking hotel with request ID:", requestID)
}
