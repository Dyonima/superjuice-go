package api

import (
	"net/http"
	"strconv"

	"github.com/KyleKincer/superjuice-go/internal/views"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "html")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Parse the query parameters.
	r.ParseForm()
	citrus := r.Form.Get("citrus")
	weight, err := strconv.ParseFloat(r.Form.Get("weight"), 64)
	if err != nil {
		http.Error(w, "weight must be a number", http.StatusBadRequest)
		return
	}

	if weight <= 0 {
		http.Error(w, "weight must be greater than 0", http.StatusBadRequest)
		return
	}

	// Calculate the recipe.
	water := weight * 10.

	// citric acid
	var citricAcid float64
	switch citrus {
	case "lime":
		citricAcid = weight * .4
	case "lemon":
		citricAcid = weight * .5
	case "orange":
		citricAcid = weight * .54
	case "grapefruit":
		citricAcid = weight * .48
	default:
		http.Error(w, "citrus must be lime, lemon, orange, or grapefruit", http.StatusBadRequest)
		return
	}

	// malic acid
	var malicAcid float64
	switch citrus {
	case "lime":
		malicAcid = weight * .2
	case "lemon":
		malicAcid = weight * .1
	case "orange":
		malicAcid = weight * .06
	case "grapefruit":
		malicAcid = weight * .12
	}

	// sugar
	var sugar float64
	switch citrus {
	case "lime", "lemon", "grapefruit":
		sugar = weight * .1
	case "orange":
		sugar = weight * .2
	}

	// salt
	var salt float64
	switch citrus {
	case "lime", "lemon", "orange":
		salt = weight * .025
	case "grapefruit":
		salt = 0
	}

	// msg
	var msg float64
	switch citrus {
	case "lime", "lemon", "orange":
		msg = 0
	case "grapefruit":
		msg = weight * .02
	}

	// ascorbic acid
	ascorbicAcid := weight * .3

	ctx := r.Context()
	views.Recipe(citricAcid, malicAcid, sugar, salt, msg, ascorbicAcid, water).Render(ctx, w)
}
