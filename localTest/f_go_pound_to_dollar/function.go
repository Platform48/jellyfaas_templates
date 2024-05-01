package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

// ******************************
// Main wrapper for the function
// ******************************

func main() {
	// Setup the HTTP server to listen on the default port
	port := "8080"
	ctx := context.Background() // Create a context that lives for the lifetime of the main function

	// Register the HTTP function with context
	if err := funcframework.RegisterHTTPFunctionContext(ctx, "/", convertGBPToUSD); err != nil {
		log.Fatalf("funcframework.RegisterHTTPFunctionContext failed: %v\n", err)
	}

	log.Printf("Starting server on port %s\n", port)
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start failed: %v\n", err)
	}
}

// *******************************
// End of wrapper for the function
// *******************************

func init() {
	functions.HTTP("ConvertGBPToUSD", convertGBPToUSD)
}

const exchangeRate = 0.79

type ReturnData struct {
	USDAmount float64 `json:"usd_amount"`
}

func convertGBPToUSD(w http.ResponseWriter, r *http.Request) {

	// Get the GBP amount from the request
	gbpAmountStr := r.URL.Query().Get("gbp")
	gbpAmount, err := strconv.ParseFloat(gbpAmountStr, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid GBP amount: %v", err), http.StatusBadRequest)
		return
	}

	if gbpAmount < 0 {
		http.Error(w, "GBP amount must be positive", http.StatusBadRequest)
		return
	}

	// Calculate the USD amount
	usdAmount := gbpAmount * exchangeRate

	// Return the USD amount
	returnData := ReturnData{USDAmount: usdAmount}
	//write into the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(returnData)
}
