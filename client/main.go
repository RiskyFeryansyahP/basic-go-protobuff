package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/confus1on/go-protobuff/pb"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type Client struct {
	client pb.AddServiceClient
}

type Result struct {
	Result int64 `json:"result"`
}

func (client *Client) Add(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	a, err := strconv.ParseUint(vars["a"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Message" : "` + err.Error() + `"}`))
		return
	}

	b, err := strconv.ParseUint(vars["b"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Message" : "` + err.Error() + `"}`))
		return
	}

	req := &pb.Request{A: int64(a), B: int64(b)}

	if response, err := client.client.Add(context.Background(), req); err == nil {
		w.WriteHeader(http.StatusOK)
		result := &Result{Result: response.Result}
		json.NewEncoder(w).Encode(result)
	}

}

func (client *Client) Multiply(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	a, err := strconv.ParseUint(vars["a"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Message" : "` + err.Error() + `"}`))
		return
	}

	b, err := strconv.ParseUint(vars["b"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Message" : "` + err.Error() + `"}`))
		return
	}

	req := &pb.Request{A: int64(a), B: int64(b)}

	if response, err := client.client.Multiply(context.Background(), req); err == nil {
		w.WriteHeader(http.StatusOK)
		result := &Result{Result: response.Result}
		json.NewEncoder(w).Encode(result)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewAddServiceClient(conn)
	c := &Client{
		client: client,
	}

	route := mux.NewRouter()

	route.HandleFunc("/add/{a}/{b}", c.Add).Methods("GET")
	route.HandleFunc("/mult/{a}/{b}", c.Multiply).Methods("GET")

	http.Handle("/", route)
	log.Println("Server Running!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
