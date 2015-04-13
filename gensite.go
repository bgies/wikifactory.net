// Public Domain (-) 2015 The Wikifactory.net Website Authors.
// See the Wikifactory.net UNLICENSE file for details.

package wikifactory

import (
	"appengine"
	"appengine/datastore"
	// "appengine/mail"
	"bytes"
	"encoding/json"
	"fmt"
	// "github.com/tav/golly/fsutil"
	"image"
	"image/color"
	"image/png"
	"logo"
	"net/http"
	"os"
	// "strconv"
	// "strings"
)

const (
	outputDirectory = "www"
	tagline         = ""
	cta             = ""
	whatOne         = "We are an open innovation network of technologists, designers, architects and makers."
	whatTwo         = "We are building the systems and spaces that enable open and distributed collaboration on design and hardware projects."
)

var index []byte

var (
	header = `<!doctype html>
<meta charset=utf-8>
<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
<title>Wikifactory</title>
<meta name="description" content="">`

//<script>Google analytics</script>`
)

type Person struct {
	ID       string
	Name     string
	Role     string
	Link     string
	GitHub   string
	LinkedIn string
	Twitter  string
	Skype    string
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
	Form     bool
	CTA      string
	Button   string
	Divider  string
	Email    string
}

type Node struct {
	City  string
	Area  string
	Email string
}

type Signup struct {
	Email string
}

type Contact struct {
	Name    string
	Email   string
	Subject string
	Message string
	EmailTo string
}

var Members = []*Person{
	{
		ID:       "tom",
		Name:     "Tom Salfield",
		Role:     "Systems Developer & Strategy",
		GitHub:   "salfield",
		LinkedIn: "https://www.linkedin.com/pub/tom-salfield/19/893/258",
		Twitter:  "tsalfield",
		Image:    "tom.jpg",
	},
	{
		ID:       "christina",
		Name:     "Christina Rebel",
		Role:     "Strategy & Business Development",
		GitHub:   "christinarebel",
		LinkedIn: "https://www.linkedin.com/in/christinarebel",
		Twitter:  "christina_rebel",
		Image:    "christina.jpg",
	},
	{
		ID:       "max",
		Name:     "Maximilian Kampik",
		Role:     "Technology & Innovation",
		GitHub:   "mkampik",
		LinkedIn: "https://www.linkedin.com/in/maximiliankampik",
		Twitter:  "mkampik",
		Image:    "max.jpg",
	},
	// {
	// 	ID:       "tav",
	// 	Name:     "tav",
	// 	Role:     "Systems Designer & Innovation",
	// 	Link:     "http://tav.espians.com/",
	// 	GitHub:   "tav",
	// 	LinkedIn: "https://www.linkedin.com/in/asktav",
	// 	Twitter:  "tav",
	// 	Image:    "tav.jpg",
	// },
	{
		ID:       "nicolai",
		Name:     "Nicolai Peitersen",
		Role:     "International Business Strategy",
		LinkedIn: "https://www.linkedin.com/pub/nicolai-peitersen/0/904/852",
		Twitter:  "NPeitersen",
		Image:    "nicolai.jpg",
	},
	{
		ID:       "jimmy",
		Name:     "Jimmy Yeh",
		Role:     "Architecture & Design",
		LinkedIn: "cn.linkedin.com/pub/jimmy-yeh/12/215/276/en",
		Twitter:  "",
		Image:    "jimmy.jpg",
	},
	{
		ID:       "jonathan",
		Name:     "Jonathan Robinson",
		Role:     "Social Innovation & Strategy",
		LinkedIn: "https://www.linkedin.com/in/jonathanrobinson1",
		Twitter:  "jon_ath_an",
		Image:    "jonathan.jpg",
	},
	{
		ID:       "ted",
		Name:     "Ted Maxwell",
		Role:     "Planning & Operations",
		LinkedIn: "https://www.linkedin.com/in/theonlyted",
		Twitter:  "theonlyted",
		Image:    "ted.jpg",
	},
	// {
	// 	ID:       "katy",
	// 	Name:     "Katy Marks",
	// 	Role:     "Architecture & Design",
	// 	LinkedIn: "https://www.linkedin.com/pub/katy-marks/55/100/22",
	// 	Image:    "katy.jpg",
	// },
	{
		ID:       "jason",
		Name:     "Jason Tang",
		Role:     "Partnerships & Operations",
		LinkedIn: "cn.linkedin.com/pub/博阳-刘/95/37b/715/zh-cn",
		Image:    "jason.jpg",
	},
	{
		ID:       "isabel",
		Name:     "Isabel Yeh",
		Role:     "Communications & Community",
		LinkedIn: "cn.linkedin.com/pub/博阳-刘/95/37b/715/zh-cn",
		Image:    "isabel.jpg",
	},
	{
		ID:       "steven",
		Name:     "Steven Lau",
		Role:     "Research & Business Support",
		LinkedIn: "cn.linkedin.com/pub/博阳-刘/95/37b/715/zh-cn",
		Image:    "steven.jpg",
	},
	{
		ID:       "lychee",
		Name:     "Lychee Li",
		Role:     "PR & Events",
		LinkedIn: "cn.linkedin.com/pub/博阳-刘/95/37b/715/zh-cn",
		Image:    "lychee.jpg",
	},
	{
		ID:       "kit",
		Name:     "Kit Harford",
		Role:     "Education",
		LinkedIn: "cn.linkedin.com/pub/博阳-刘/95/37b/715/zh-cn",
		Image:    "kit.jpg",
	},
	{
		ID:       "katherine",
		Name:     "Katherine Wang",
		Role:     "Office & Admin",
		LinkedIn: "cn.linkedin.com/pub/博阳-刘/95/37b/715/zh-cn",
		Image:    "katherine.jpg",
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
		Text:     "On wikifactory.org we are developing a social collaboration platform for open design and hardware projects. Think \"github for design and hardware\" meets \"wikipedia of things\".<br><br>Features are both inspired by successful methodologies of open source software development, as well as contextualised to the processes of designing and producing physical objects.<br><br><b>Want early access?</b><br>The more feedback we can get, the faster we can improve our features for the community. If you'd like to get early access and help shape the platform, please enter your email below.",
		Image:    "",
		Form:     true,
		CTA:      "",
		Button:   "small-btn",
		Divider:  "nodes-spacer-1.jpg",
	},
	{
		Title:    "Wikifactory @ iBox Chengdu, China",
		Link:     "",
		Status:   "Launched Q3 2014",
		GitHub:   "",
		Twitter:  "",
		YouTube:  "",
		Image:    "labchengdu.jpg",
		Text:     "The first of our replicatable Wikifactories at the heart of Chengdu’s latest creative hub, alongside art galleries, artisan workshops and coworking spaces as well as cafés and restaurants.<br><br>It houses a diverse community of technologists, designers and makers and provides access to a growing range of digital fabrication technologies, desk space and business incubation support to help launch new products. In Chengdu, there is a focus on products that have a social impact.<br><br>A Wikifactory like this one can soon be connected to other Wikifactories via our Social Design Platform. Creating a global, distributed collaboration and production network of 21st Century factories.",
		Partners: "",
		CTA:      "Start a Wikifactory in your city",
		Button:   "large-btn",
		Divider:  "nodes-spacer-2.jpg",
		Email:    "max@wikifactory.org",
	},
	{
		Title:    "Print Your Hometown",
		Link:     "",
		Status:   "Launching Q3 2015",
		GitHub:   "",
		Twitter:  "",
		YouTube:  "",
		Image:    "printingthefuture.jpg",
		Text:     "We are collaborating with education experts in China to develop and deliver a learning programme around entrepreneurship, design and 3D printing at 50 universities across 5 provinces in China.<br><br>Over 10,000 students will be introduced and given access to digital fabrication technologies in developing more sustainable, innovative products through seminars and workshops.<br><br>Participating students will be encouraged to form teams and pitch their product ideas to receive mentorship from experienced pioneers in the industry and continue prototyping.",
		Partners: "cydf.png",
		CTA:      "Partner with us",
		Button:   "small-btn",
		Divider:  "nodes-spacer-3.jpg",
		Email:    "max@wikifactory.org",
	},
	{
		Title:    "WikiHouse China - Rooftop 1.0 @ iBOX Chengdu",
		Link:     "http://www.wikihouse.cc",
		Status:   "Launching Q2 2015",
		GitHub:   "",
		Twitter:  "",
		YouTube:  "",
		Image:    "wikihouse-cn.jpg",
		Text:     "In time for summer, architects and designers are coming together to prototype the first WikiHouse in China. Launching in the rooftop space of the Wikifactory Lab to provide a relaxed environment for makers of Chengdu.<br><br>We will be launching our chapter of WikiHouse China to drive open innovation in architecture for social and environmental impact. With Sichuan being both the bamboo region of China and affected by earthquakes every year, we will explore how to develop earthquake resistant shelters with more sustainable plywood fibres to help rebuild communities.",
		Partners: "cdad.png",
		CTA:      "More about WikiHouse",
		Button:   "medium-btn",
		Divider:  "nodes-spacer-4.jpg",
	},
	// {
	// 	Title:    "Innovation Hub @ NIMI Chengdu, China",
	// 	Link:     "",
	// 	Status:   "Launching Q3 2015",
	// 	GitHub:   "",
	// 	Twitter:  "",
	// 	YouTube:  "",
	// 	Image:    "nimi.jpg",
	// 	Text:     "At NIMI we are developing a multi-purpose innovation space, hosting a public facing exhibition in the future of design and production, as well as offering professional training and digital fabrication-as-a-service.<br><br>In training young talent in digital fabrication in a range of technologies from stereolithography to laser cutting, the Innovation Hub will support local businesses and industry to adopt these in their supply chains. A co-working and fully-equipped workshop space will also be open for local maker and hardware communities.",
	// 	Partners: "nimi.png",
	// 	Divider:  "nodes-spacer-5.jpg",
	// },
}

var Nodes = []*Node{
	{
		City:  "London",
		Area:  "Borough",
		Email: "max@wikifactory.org",
	},
	{
		City:  "Chengdu",
		Area:  "Jinjang",
		Email: "max@wikifactory.org",
	},
	{
		City:  "Beijing",
		Area:  "Wangjing",
		Email: "max@wikifactory.org",
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
	o("<li><a href='#join'><h4 class=cta>JOIN</h4></a></li>")
	o("<li><a href='#who'><h4>WHO</h4></a></li>")
	o("<li><a href='#how'><h4>HOW</h4></a></li>")
	o("<li><a href='#what'><h4>WHAT</h4></a></li>")
	o("</ul>")
	o("</nav>")
	o("</header>")
	// PARTICLES
	o("<div class=tag id=home>")
	o("<div id=particles-js><div><img src=/logo.png></img></div></div><div class=bounce><img class=arrow src=/gfx/arrow.png></img></div>")
	o("<script src=" + getPath("site.js") + " async></script>")
	o("<script src=/lib/particles.js></script>")
	o("</div>")
	// WHAT
	o("<div class=tag id=what>")
	o("<div class=intro><h1>" + whatOne + "</h1></div>")
	o("<div class=intro><h1>" + whatTwo + "</h1></div>")
	o("</div>")
	// HOW
	o("<div class='nodes tag' id=how>")
	o("<div class=nodes-text-one><h2>ONLINE<br>SOCIAL DESIGN PLATFORM</h2></div>")
	o("<div class=nodes-text-two><h2>PHYSICAL<br>COLLABORATION SPACES</h2></div>")
	o("<div class=nodes-text-three><h2>TRAINING<br>& WORKSHOPS</h2></div>")
	o("<div class=nodes-text-four><h2>PROJECT<br>INCUBATION</h2></div>")
	o("</div>")
	//// PROJECTS
	o("<div class='projects-wrapper'>")
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
		if p.Form == true {
			o("<div>")
			o("<form action=signup method=post class=form>")
			o("<input class=access type=text placeholder=Email name=email></input><br><br>")
			o("<button type=submit class=submit>Get early access</button>")
			o("</form>")
			o("<div class='thankyou'><h3>Thanks!</h3><p>We will notify you when the beta goes live.</p></div>")
			o("</div>")
		}
		if p.CTA != "" && p.Link != "" {
			o("<div class=" + p.Button + "><a href=" + p.Link + "><h3>" + p.CTA + "</h3></a></div>")
		}
		if p.CTA != "" && p.Link != "http://www.wikihouse.cc" {
			o("<div id=" + p.Email + " class=" + p.Button + "><h3>" + p.CTA + "</h3></div>")
		}
		if p.Partners != "" {
			o("<div class='partners'>")
			o("<h5>Partners</h5>")
			o("<div class='partnersimage'><img class=partnerimg src=/gfx/partners/" + p.Partners + "></div>")
			o("</div>")
		}
		if p.Divider != "" {
			o("<div class='nodes-spacer'><img src=/gfx/" + p.Divider + "></div>")
		}
		o("</div>")

	}
	for _, project := range Projects {
		renderProject(project, nil)
	}
	o("</div>")
	o("<div class=tag id=who>")
	o("<div class='team-wrapper'>")
	o("<h1 class=projects>Team</h1>")
	renderPerson := func(p *Person) {
		o("<div class=person-card>")
		o("<div>")
		if p.Image != "" {
			o("<div class=person-img><img class=avatar src=/gfx/team/" + p.Image + "></div>")
		}
		o("<h2>" + p.Name + "</h2>")
		o("<h5>" + p.Role + "</h5>")
		o("<p>" + p.Text + "</p>")
		o("</div>")
		o("<div class=person-smedia>")
		if p.Twitter != "" {
			o("<a target=_blank href=http://twitter.com/" + p.Twitter + ">" + "<img class=icon-person src=gfx/icons/twitter.png>" + "</a>")
		}
		if p.LinkedIn != "" {
			o("<a target=_blank href=" + p.LinkedIn + ">" + "<img class=icon-person src=gfx/icons/linkedin.png>" + "</a>")
		}
		if p.Skype != "" {
			o("<a target=_blank href=" + p.Skype + ">" + "<img class=icon-person src=gfx/icons/skype.png>" + "</a>")
		}
		if p.GitHub != "" {
			o("<a target=_blank href=https://github.com/" + p.GitHub + ">" + "<img class=icon-person src=gfx/icons/github.png>" + "</a>")
		}
		o("</div>")
		o("</div>")
	}
	for _, members := range Members {
		renderPerson(members)
	}
	o("</div>")
	o("</div>")
	// FOOTER
	o("<div class='footer tag' id=join>")
	//// NETWORK NODES
	o("<div class='network-nodes'>")
	o("<h2>Nodes in our network</h2>")
	renderNode := func(p *Node, next *Node) {
		o("<div class=city>")
		o("<h3>" + p.City + "</h3>")
		o("<h6>" + p.Area + "</h6>")
		o("<div id=" + p.Email + " class=contact><h5>CONTACT</h5></div>")
		o("</div>")
	}
	for _, node := range Nodes {
		renderNode(node, nil)
	}
	o("<div id='form-a-node' class=new><h3>Form a node</h3></div>")
	o("</div>")
	o("<div class=smedia>")
	o("<div class='smedia-icon'><a href='https://twitter.com/wikifactory' target='_blank'><img src=/gfx/icons/twitter-white.png></a></div>")
	o("<div class='smedia-icon'><a href='https://facebook.com/wikifactory' target='_blank'><img src=/gfx/icons/facebook-white.png></a></div>")
	o("<div class='smedia-icon'><a href='https://github.com/wikifactory' target='_blank'><img src=/gfx/icons/github-white.png></a></div>")
	o("<div class='smedia-icon'><a href='https://www.linkedin.com/company/wikifactory' target='_blank'><img src=/gfx/icons/linkedin-white.png></a></div>")
	o("</div>")
	o("<div class=coded><p>Coded with <3 in London</p></div>")
	o("</div>")
	// CONTACT FORM
	o("<div class='contactform'>")
	o("<h3>Contact us</h3>")
	o("<form action=contact method=post>")
	o("<input class=contactname type=text placeholder=Name name=name></input><br>")
	o("<input class=contactemail type=text placeholder=Email name=email></input><br>")
	o("<input class=contactsubject type=text placeholder=Subject name=subject></input><br>")
	o("<textarea class=contactmessage type=textarea placeholder=Message name=message></textarea><br>")
	o("<div class='tiny-wrapper'><div class='tiny-btn cancel'><h3>Cancel</h3></div>")
	o("<button type=submit class='tiny-submit'>Submit</button></div>")
	o("</form>")
	o("</div>")
	o("</div>")
	// WRITE TO BUF
	index = buf.Bytes()
}

// EMAIL NOTIFICATIONS
// func Notification(w http.ResponseWriter, r *http.Request) {
// 	c := appengine.NewContext(r)
// 	addr := r.FormValue("emailto")
// 	msg := &mail.Message{
// 		Sender:  "Wikifactory <team@wikifactory.org>",
// 		To:      []string{addr},
// 		Subject: "" + r.FormValue("subject"),
// 		Body:    r.FormValue("message"),
// 	}
// 	if err := mail.Send(c, msg); err != nil {
// 		c.Errorf("Couldn't send email: %v", err)
// 	}
// }
func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write(index)
}
func handleContact(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	s := Contact{
		Email:   r.FormValue("email"),
		Name:    r.FormValue("name"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
		EmailTo: r.FormValue("emailto"),
	}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "contact", nil), &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func handleSignup(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	s := Signup{
		Email: r.FormValue("email"),
	}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "signup", nil), &s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	http.HandleFunc("/signup", handleSignup)
	http.HandleFunc("/contact", handleContact)
}
