package main 



import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


type UserRequest struct{
	Name string  `json:"name"`
	Password string `json:"password"`
}


func main(){
	mux:= http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))

	})
	
	
	mux.HandleFunc("/user", func (w http.ResponseWriter , r *http.Request)  {
		if r.Method != http.MethodPost{
			http.Error(w , "Method not allowed", http.StatusMethodNotAllowed)
			return
		}


		var user UserRequest

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil{
			http.Error(w , "Invalid json" , http.StatusBadRequest)
			return
		}
		fmt.Fprintf(
			w, "Recived user name=%s password=%s ",
			user.Name,
			user.Password,
		)
		log.Println("server listening on :8080")
		log.Fatal(http.ListenAndServe(":8080", mux))
	})

}