package main

import (
	"math/rand"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	ds "github.com/bwmarrin/discordgo"
)

// Prefix
const prefix string = ".go"

func main() {
	// TODO: create a .env for the discord token
	session, err := ds.New("Bot ")

	if err != nil {
		log.Fatal(err)
	}

	// FIX: ?
	const redirectPostGophersLatam string = "Si tienes dudas sobre como arrancar con Go, puedes consultar la opinion de la comunidad en el siguiente [link](https://gophers-latam.github.io/posts/2024/04/explorando-frameworks-en-go-una-gu%C3%ADa-para-principiantes/)"

	session.AddHandler(func(s *ds.Session, m *ds.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] != prefix {
			return
		}

		// Hello world test
		if args[1] == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
		}

		// NOTE: maybe this could be better(?)
		if args[1] == "framework" {
			s.ChannelMessageSend(m.ChannelID, redirectPostGophersLatam)
		}

		// meme 😂
		if args[1] == "facts" {
			facts := []string{
				"1. Los miembros de Gophers LATAM pueden escribir una api Go en microsoft paint y compilarla en excel escrita en Ruby",
				"2. Cuando el compilador encuentra un error en el código de Gophers LATAM, el compilador se disculpa",
				"3. Gophers LATAM puede dividir por cero, pero el compilador asustado intenta multiplicar por el infinito",
				"4. Gophers LATAM puede arrojar una excepción mas lejos que nadie y en menor tiempo",
				"5. Cuando Gophers LATAM presiona [ctr-alt-del] es el resto mundo el que se reinicia",
				"6. Gophers LATAM no necesita recolector de basura. Solo mira a los objetos fijamente y los objetos se destruyen a si mismos muertos de miedo",
				"7. Gophers LATAM no necesita compilar su código Go. Le basta escribir en Javascript y se traduce el mismo a binario",
				"8. Los miembros de Gophers LATAM no tienen tecla [CONTROL] en su teclado, ellos siempre están en control",
				"9. Gophers LATAM puede detectar el siguiente número en una secuencia aleatoria",
				"10. Gophers LATAM puede ejecutar un loop infinito en 3 segundos",
				"11. Gophers LATAM no puede producir un Null Pointer Exception, si Gophers LATAM apunta a Null, un objeto se materializa instantaneamente",
				"12. Gophers LATAM puede hacer control-z con lápiz y papel",
				"14. Gophers LATAM te hace updates a la base de datos con el buscaminas",
				"15. Los arrays de Gophers LATAM son de tamaño infinito porque Gophers LATAM no tiene límites.",
				"16. Gophers LATAM terminó World of Warcraft",
				"17. Solo hay 10 clases de personas, los que son parte de Gophers LATAM y los que no",
			}

			selection := rand.Intn(len(facts))

			author := ds.MessageEmbedAuthor{
				Name: "El Programador Pobre",
			}			
			embed := ds.MessageEmbed{
				Title: facts[selection],
				Author: &author,
			}


			// s.ChannelMessageSend(m.ChannelID, facts[selection])
			s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		}

	})

	// NOTE: checking what this does
	session.Identify.Intents = ds.IntentsAllWithoutPrivileged

	err = session.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	// PERF: can be done better!
	fmt.Println("Bot online!")

	// NOTE: is this 100% necessary?
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
