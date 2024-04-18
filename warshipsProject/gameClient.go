package warships

import (
	"io/ioutil"
	"log"
	"net/http"
)

func InitGame() error {
	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)
}
func Board() ([]string, error)
func Status() (*StatusResponse, error)
func Fire(coord string) (string, error)
