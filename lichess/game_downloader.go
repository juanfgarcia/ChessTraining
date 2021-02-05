package Lichess

import(
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"github.com/clinaresl/pgnparser/pgntools"
)

func GetGames(username string) (games pgntools.PgnCollection) {
	resp, err := http.Get(fmt.Sprintf("https://lichess.org/api/games/user/%s",username))
	if err != nil{
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return pgntools.GetGamesFromString(string(body), 0, "", "", false)
}
