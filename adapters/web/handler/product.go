package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/iugmali/fullcycle-products/adapters/web/dto"
	"github.com/iugmali/fullcycle-products/application"
	"github.com/urfave/negroni"
	"net/http"
	"strconv"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func MakeProductHandler(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product", n.With(negroni.Wrap(getAllProducts(service)))).Methods("GET", "OPTIONS")
	r.Handle("/product/{id}", n.With(negroni.Wrap(getProduct(service)))).Methods("GET", "OPTIONS")
	r.Handle("/product", n.With(negroni.Wrap(createProduct(service)))).Methods("POST")
	r.Handle("/product/{id}/enable", n.With(negroni.Wrap(enableProduct(service)))).Methods("PATCH", "OPTIONS")
	r.Handle("/product/{id}/disable", n.With(negroni.Wrap(disableProduct(service)))).Methods("PATCH", "OPTIONS")
	r.Handle("/product/{id}/setprice/{price}", n.With(negroni.Wrap(setProductPrice(service)))).Methods("PATCH", "OPTIONS")
}

func getAllProducts(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		products, err := service.GetAll()
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Products not found")
		} else {
			respondWithJSON(w, http.StatusOK, products)
		}
	})
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		product, err := service.Get(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Product not found")
		} else {
			respondWithJSON(w, http.StatusOK, product)
		}
	})
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var productDto dto.Product
		err := json.NewDecoder(r.Body).Decode(&productDto)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		product, err := service.Create(productDto.Name, productDto.Price)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithJSON(w, http.StatusCreated, product)
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		product, err := service.Get(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Product not found")
			return
		}
		product, err = service.Enable(product)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"success": "Product enabled"})
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		product, err := service.Get(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Product not found")
			return
		}
		product, err = service.Disable(product)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"success": "Product disabled"})
	})
}

func setProductPrice(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		product, err := service.Get(vars["id"])
		if err != nil {
			respondWithError(w, http.StatusNotFound, "Product not found")
			return
		}
		price, err := strconv.ParseInt(vars["price"], 10, 64)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid price")
			return
		}
		product, err = service.SetPrice(product, price)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, map[string]string{"success": fmt.Sprintf("Product price has been set to %d", price)})
	})
}
