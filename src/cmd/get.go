package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

type ICanHazDadJoke struct {
	Id string `json:"id"`
	Joke string `json:"joke"`
	Status int `json:"status"`
}

func init() {
	rootCmd.AddCommand(getJokeCmd)
}

var getJokeCmd = &cobra.Command{
	Use:   "getJoke",
	Short: "Gets a joke from ICanHazDadJoke",
	Long:  `Gets a random joke from ICanhazDadJoke`,
	Run: func(cmd *cobra.Command, args []string) {
		res := GetJoke()
		fmt.Println(res.Joke)
	},
}

func GetJoke() ICanHazDadJoke {
	var uri = "https://icanhazdadjoke.com/"
	resp := requestData(uri)

	var jsonData ICanHazDadJoke = ICanHazDadJoke{}
	json.Unmarshal(resp, &jsonData)

	return jsonData
}

func requestData(url string) []byte {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)
	
	req.Header.Set("User-Agent", "gogetdadjoke (https://github.com/jtom38/gogetdadjoke")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}
