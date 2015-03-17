// Public Domain (-) 2015 The Wikifactory.net Website Authors.
// See the Wikifactory.net UNLICENSE file for details.

package wikifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "github.com/tav/golly/fsutil"
	"net/http"
	"os"
	"image"
    "image/color"
    "image/png"
    "logo"
	// "strconv"
	// "strings"
)

const (
	outputDirectory	= "www"
	tagline			= ""
	cta				= ""
	whatOne			= "A global, entrepreneurial network of strategists, technologists, designers, architects and makers." //WIP
	whatTwo 		= "Working towards a more open and empowering model of collaboration for digital fabrication projects." //WIP
)

var index []byte

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
	Role	 string
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
	Status   string
	Facebook string
	GitHub   string
	Twitter  string
	Text     string
	Image    string
	Position string
	Area     string
	YouTube  string
	Partners string
	CTA      string
	Button	 string
}

var Members = []*Person{
	{
		ID:       "tom",
		Name:     "Tom Salfield",
		Role:	  "Chief Technical Architect",
		GitHub:   "salfield",
		LinkedIn: "pub/tom-salfield/19/893/258",
		Twitter:  "tsalfield",
		Text:     `Technical architect and software developer that is passionate about employing P2P and open source technologies to solve systemic problems and bring about a more open, sustainable economy.`,
		Image:    "",
	},
	{
		ID:       "christina",
		Name:     "Christina Rebel",
		Role:	  "Business and User Experience",
		LinkedIn: "in/christinarebel",
		Twitter:  "christina_rebel",
		Text:     `Constantly building on her range of skillsets - from web development to illustration, strategic planning to video production, and more - to see social innovation projects through early stages and beyond.`,
		Image:    "",
	},
	{
		ID:       "max",
		Name:     "Maximilian Kampik",
		Role:	  "Technology and User Experience",
		GitHub:   "mkampik",
		LinkedIn: "in/maximiliankampik",
		Twitter:  "mkampik",
		Text:     `Technologist and aspiring futurist that enjoys keeping up with the latest tech innovations and implementations. Has a background in politics and international relations.`,
		Image:    "",
	},
	{
		ID:       "tav",
		Name:     "tav",
		Role:	  "Systems and Innovation",
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
		Role:	  "Business and User Experience",
		LinkedIn: "pub/nicolai-peitersen/0/904/852",
		Twitter:  "NPeitersen",
		Text:     `A thinker, doer and entrepreneur for a range of worldwide issues, his latest book ‘The Ethical Economy’ guides a call to build the instruments, institutions, and technologies to realise the democratisation of our economies.`,
		Image:    "",
	},
}

var Projects = []*Project{
	{
		Title:    "Social Design Platform",
		Link:     "https://www.wikifactory.org/",
		Status:   "Beta launching Q2 2015",
		Facebook: "wikifactory",
		GitHub:   "/wikifactory",
		Twitter:  "wikifactory",
		YouTube:  "WikifactoryMovement",
		Text:     "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua et dolore manga aliqua.<br><br> Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		Image:    "",
		CTA:	  "Get early access",
		Button:	  "small-btn",
	},
	{
		Title:   "Wikifactory @iBox Chengdu, China",
		Link:    "",
		Status:  "Launched Q3 2014",
		GitHub:  "",
		Twitter: "",
		YouTube: "",
		Image:   "labchengdu.jpg",
		Text:    "Our China HQ at the heart of Chengdu’s latest creative hub, alongside art galleries, artisan workshops and coworking spaces as well as cafés and restaurants.<br><br>A space to bring together a community around digital fabrication, offering access to 3D printers, desk space and business incubation support to launch new product ideas that have a social impact.",
		Partners:"wikihouse-cn.png",
		CTA:	 "Start a Wikifactory in your city",
		Button:	 "large-btn",
	},
	{
		Title:   "Printing the Future",
		Link:    "",
		Status:  "Launching Q3 2015",
		GitHub:  "",
		Twitter: "",
		YouTube: "",
 		Image:   "printingthefuture.jpg",
		Text:    "We are collaborating with education experts in China to develop and deliver a learning programme around entrepreneurship, design and 3D printing at 10 universities across 5 provinces in China.<br><br>Over 2,000 students will be introduced and given access to digital fabrication technologies in developing more sustainable, innovative products through seminars and workshops.<br><br>Participating students will be encouraged to form teams and pitch their product ideas to receive mentorship from experienced pioneers in the industry and continue prototyping.",
 		Partners:"cydf.png",
 		CTA:	 "Partner with us",
		Button:	 "small-btn",
	},
	{
		Title:   "WikiHouse China - Rooftop 1.0 @ iBOX Chengdu",
		Link:    "",
		Status:  "Launching Q2 2015",
		GitHub:  "",
		Twitter: "",
		YouTube: "",
 		Image:   "wikihouse-cn.jpg",
		Text:    "In time for summer, architects and designers are coming together to prototype the first WikiHouse in China. Launching in the rooftop space of the Wikifactory Lab to provide a relaxed environment for makers of Chengdu.<br><br>We will be launching our chapter of WikiHouse China to drive open innovation in architecture for social and environmental impact. With Sichuan being both the bamboo region of China and affected by earthquakes every year, we will explore how to develop earthquake resistant shelters with more sustainable plywood fibres to help rebuild communities.",
 		Partners:"cdad.png",
 		CTA:	 "More about WikiHouse",
		Button:	 "medium-btn",
	},
	{
		Title:   "Innovation Hub @ NIMI Chengdu, China",
		Link:    "",
		Status:  "Launching Q3 2015",
		GitHub:  "",
		Twitter: "",
		YouTube: "",
		Image:   "nimi.jpg",
		Text:    "At NIMI we are developing a multi-purpose innovation space, hosting a public facing exhibition in the future of design and production, as well as offering professional training and digital fabrication-as-a-service.<br><br>In training young talent in digital fabrication in a range of technologies from stereolithography to laser cutting, the Innovation Hub will support local businesses and industry to adopt these in their supply chains. A co-working and fully-equipped workshop space will also be open for local maker and hardware communities.",
		Partners:"nimi.png",
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

func genSite() {
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
		return "/static/" + val
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
	// NAVBAR
	o("<header class=nav-up>")
	o("<img src=/gfx/w.png>")
	o("<nav>")
	o("<ul>")
	o("<li><h4 class=cta>JOIN</h4></li>")
	o("<li><h4>WHY</h4></li>")
	o("<li><h4>WHO</h4></li>")
	o("<li><h4>HOW</h4></li>")
	o("<li><h4>WHAT</h4></li>")
	o("</ul>")
	o("</nav>")
	o("</header>")
	// PARTICLES
	o("<div id=particles-js><div><img src=/logo.png></img></div></div><div class=bounce><img class=arrow src=/gfx/arrow.png></img></div>")
	o("<script src=" + getPath("site.js") + " async></script>")
	o("<script src=http://vincentgarreau.com/particles.js/particles.js></script>")
	// WHAT
	o("<div class=tag id=what>")
	o("<div class=intro><h1>" + whatOne + "</h1></div>")
	o("<div class=intro><h1>" + whatTwo + "</h1></div>")
	o("</div>")
	// HOW
	o("<div class='nodes tag' id= how>")
	o("<div class=nodes-text-one><h2>ONLINE<br>SOCIAL DESIGN PLATFORM</h2></div>")
	o("<div class=nodes-text-two><h2>PHYSICAL<br>COLLABORATION SPACES</h2></div>")
	o("<div class=nodes-text-three><h2>TRAINING<br>& WORKSHOPS</h2></div>")
	o("<div class=nodes-text-four><h2>PROJECT<br>INCUBATION</h2></div>")
	o("</div>")
	//// PROJECTS
	o("<div>")
	o("<h1 class=projects>Projects</h1>")
	o("<div class=device-wrapper><img class=devices src=/gfx/devices.png></div>")
	renderProject := func(p *Project, next *Project) {
		if p.Image != "" {
			o("<div class='projectimage'><img class=projectimg src=/gfx/projects/" + p.Image + "></div>")
		}
		o("<div>")
		o("<h2>" + p.Title + "</h2>")
		o("<h5>" + p.Status + "</h5>")
		o("<div class=text><p>" + p.Text + "</p></div>")
		if p.Partners != "" {
			o("<div class='partnersimage'><img class=partnerimg src=/gfx/partners/" + p.Partners + "></div>")
		}
		if p.CTA != "" {
			o("<div class=" + p.Button +"><h3>" + p.CTA + "</h3></div>")
		}
		o("</div>")
	}
	for _, project := range Projects {
		renderProject(project, nil)
	}
	o("</div>")
	// WHO
	// add renter team function
	// FOOTER
	// WRITE TO BUF
	index = buf.Bytes()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write(index)
}

func handleLogo(w http.ResponseWriter, req *http.Request) {
    clr := image.NewUniform(color.RGBA{53, 44, 37, 255})
    img, err := logo.Render(24, clr)
    if err != nil {
            fmt.Fprintf(w, "ERROR: %s\n", err)
            return
    }
    err = png.Encode(w, img)
    if err != nil {
            fmt.Fprintf(w, "ERROR: %s\n", err)
            return
    }
    w.Header().Set("Content-Type", "image/png")
}

func init() {
	logo.Setup()
	genSite()
	http.HandleFunc("/logo.png", handleLogo)
    http.HandleFunc("/", handleIndex)
}
