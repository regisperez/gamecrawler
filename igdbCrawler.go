package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Game struct {
	position string
	name  string
	score string
}

func Decodificar() {
	// Request the HTML page.
	res, err := http.Get("https://www.igdb.com/top-100/games")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Carrega a pagina
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Encontra a textarea da página

	texto := doc.Find("tr").Text()
	link := doc.Find("a")
	fmt.Println(link)
 	t := strings.Split(texto, "#NameScore")
	s := strings.Split(t[1], "/ 100")
	for i := 0; i < len(s)-1; i++ {

		
		position:=strconv.Itoa(i+1)
		nome := strings.TrimLeft(string(s[i]), position)
		nome =  strings.TrimRight(nome, " ")
		nota:=s[i]
		nota2 := string(nota[len(nota)-3 : len(nota)-1])
		nome = strings.TrimRight(nome,nota2)		
		var game Game
		game = Game {position,nome,nota2}
		fmt.Println(game)

		//fmt.Print (position)
		//fmt.Println (nome)
		
		
	}

}

func main() {
	Decodificar()
}
