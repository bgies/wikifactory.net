// Public Domain (-) 2015 The Wikifactory.net Website Authors.
// See the Wikifactory.net UNLICENSE file for details.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "github.com/tav/golly/fsutil"
	// "genlogo"
	"os"
	// "strconv"
	// "strings"
)

const (
	outputDirectory = "www"
	tagline         = ""
	cta			    = ""
)

var (
	header      = `<!doctype html>
<meta charset=utf-8>
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<title>Wikifactory</title>
<meta name="description" content="">`
//<script>Google analytics</script>`
)

type Person struct {
	ID       string
	Name     string
	Link     string
	GitHub   string
	LinkedIn string
	Twitter  string
	Text     string
	Image    string
}

type Project struct {
	Title    string
	Link     string
	Year     int
	Facebook string
	GitHub   string
	Twitter  string
	Text     string
	Image    string
	Position string
	Area     string
	YouTube  string
}


var Members = []*Person{
	{
		ID:       "tom",
		Name:     "Tom Salfield",
		GitHub:   "salfield",
		LinkedIn: "pub/tom-salfield/19/893/258",
		Twitter:  "tsalfield",
		Text:     `Technical architect and software developer that is passionate about employing P2P and open source technologies to solve systemic problems and bring about a more open, sustainable economy.`,
		Image:    "",
	},
	{
		ID:       "christina",
		Name:     "Christina Rebel",
		LinkedIn: "in/christinarebel",
		Twitter:  "christina_rebel",
		Text:     `Constantly building on her range of skillsets - from web development to illustration, strategic planning to video production, and more - to see social innovation projects through early stages and beyond.`,
		Image:    "",
	},
	{
		ID:       "max",
		Name:     "Maximilian Kampik",
		GitHub:   "mkampik",
		LinkedIn: "in/maximiliankampik",
		Twitter:  "mkampik",
		Text:     `Technologist and aspiring futurist that enjoys keeping up with the latest tech innovations and implementations. Has a background in politics and international relations.`,
		Image:    "",
	},
	{
		ID:       "tav",
		Name:     "tav",
		Link:     "http://tav.espians.com/",
		GitHub:   "tav",
		LinkedIn: "in/asktav",
		Twitter:  "tav",
		Text:     `Systems designer, entrepreneur and aspiring polymath. Spends his time innovating on the cutting edge of social, economic and<br> technological systems.`,
		Image:    "",
	},
	{
		ID:       "nicolai",
		Name:     "Nicolai Peitersen",
		LinkedIn: "pub/nicolai-peitersen/0/904/852",
		Twitter:  "NPeitersen",
		Text:     `A thinker, doer and entrepreneur for a range of worldwide issues, his latest book ‘The Ethical Economy’ guides a call to build the instruments, institutions, and technologies to realise the democratisation of our economies.`,
		Image:    "",
	},
}

var Projects = []*Project{
	{
		Title:    "Wikifactory Social Design Platform",
		Link:     "https://www.wikifactory.org/",
		Year:     2014,
		Facebook: "wikifactory",
		GitHub:   "tav/wikifactory",
		Twitter:  "wikifactory",
		YouTube:  "WikifactoryMovement",
		Text:     ``,
	},
	{// Wikifactory Lab @iBox Chengdu, China.
		Title:   "",
		Link:    "",
		Year:    0,
		GitHub:  "",
		Twitter: "",
		YouTube: "",
		Text:    ``,
	},
	{// Wikifactory Innovation Hub at NIMI University Chengdu, China.
		Title:   "",
		Link:    "",
		Year:    0,
		GitHub:  "",
		Twitter: "",
		YouTube: "",
		Text:    ``,
	},
	{// Wikifactory China Courses.
		Title:   "",
		Link:    "",
		Year:    0,
		GitHub:  "",
		Twitter: "",
		YouTube: "",
		Text:    ``,
	},
	{// Wikifactory Innovation Hub London.
		Title:   "",
		Link:    "",
		Year:    0,
		GitHub:  "",
		Twitter: "",
		YouTube: "",
		Text:    ``,
	},
}

var Partners = []*Project{
	{// Ethical Works
		Title:    "",
		Link:     "",
		Year:     0,
		Facebook: "",
		GitHub:   "",
		Twitter:  "",
		YouTube:  "",
		Text:     ``,
	},
	{// Nordic International Management Institute Chengdu
		Title:    "",
		Link:     "",
		Year:     0,
		Facebook: "",
		GitHub:   "",
		Twitter:  "",
		YouTube:  "",
		Text:     ``,
	},
}

func exit(args ...interface{}) {
	if len(args) == 1 {
		fmt.Printf("\n!! ERROR: %s\n\n", args...)
	} else {
		fmt.Printf("\n!! ERROR: "+args[0].(string)+"\n\n", args[1:]...)
	}
	os.Exit(1)
}

func main() {

	assetsFile, err := os.Open("assets.json")
	if err != nil {
		exit(err)
	}

	assetMap := map[string]string{}
	assetsDecoder := json.NewDecoder(assetsFile)

	err = assetsDecoder.Decode(&assetMap)
	if err != nil {
		exit(err)
	}

	getPath := func(key string) string {
		val, exists := assetMap[key]
		if !exists {
			exit("Cannot find %s in assets.json", key)
		}
		return "static/" + val
	}

	buf := &bytes.Buffer{}
	o := func(s string, args ...interface{}) {
		s += "\n"
		if len(args) > 1 {
			fmt.Fprintf(buf, s, args...)
		} else {
			buf.WriteString(s)
		}
	}

	o(header)
	o("<link rel=stylesheet href=" + getPath("style.css") + ">")
	o("<link href='http://fonts.googleapis.com/css?family=Montserrat:400,700' rel='stylesheet' type='text/css'>")
	o("<script src=http://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js></script>")
	// HEADER
	o("<header class=nav-down>")
	o("<img src=gfx/w.png>")
	o("<nav>")
	o("<ul>")
	o("<li><a></a></li>")
	o("<li><a></a></li>")
	o("<li><a></a></li>")
	o("</ul>")
	o("</nav>")
	o("</header>")
	// PARTICLES + TAGLINE
	o("<div id=particles-js><div class=wrapper><id=tagline><img src=" + "" + "></img></div>" + "</div>" + "</div>")
	o("<script src=" + getPath("site.js") + " async></script>")
	o("<script src=http://vincentgarreau.com/particles.js/particles.js></script>")
	// MAIN CONTENT
	o("<main>")
	o("<div class=wrapper>")
	o("</main")
	// FOOTER
	o("</footer>")
	// CLOSE MAIN WRAPPER
	o("</div>")//

	// GENERATE INDEX.HTML
	index, err := os.Create("www/index.html")
	if err != nil {
		exit(err)
	}

	_, err = index.Write(buf.Bytes())
	if err != nil {
		exit(err)
	}

	fmt.Println(">> Generated www/index.html")

}
