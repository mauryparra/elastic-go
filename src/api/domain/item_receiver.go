package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const urlItems = "https://api.mercadolibre.com/items/"

// Get obtiene toda la iformación del usuario a partir del ID
func (item *Item) Get() error {
	if item.ID == "" {
		return errors.New("Id esta vacio")
	}

	final := fmt.Sprintf("%s%s", urlItems, item.ID)
	response, err := http.Get(final)
	if err != nil {
		log.Println(err)
		return err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	if err := json.Unmarshal([]byte(data), &item); err != nil {
		return err
	}

	return nil
}
