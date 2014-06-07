package get_blog_stats

var Hosts = []string {"letter.boxes.cf-app.com", "pandora.boxes.cf-app.com", "juke.boxes.cf-app.com",
	"punk.music.cf-app.com", "rock.music.cf-app.com", "dubstep.music.cf-app.com",
	"neon.atom.cf-app.com", "xenon.atom.cf-app.com", "helium.atom.cf-app.com"}

var country = map[string]string{
	"usa": "com",
	"germany": "de",
	"russia": "ru",
	"canada": "ca",
	"australia": "au",
	"uk": "co.uk",
}

var blog = "http://golang-basic.blogspot."

var Gobasic_urls = []string {
	"/2014/06/step-by-step-guide-to-ssh-using-go.html?view=sidebar",
	"/2014/06/curl-in-golang-go-curl.html?view=sidebar",
	"/2014/06/golang-closures-running-as-go-routines.html?view=sidebar",
	"/2014/06/basic-go-golang-difference-in-cc-golang.html?view=sidebar",
	"/2014/06/basic-golang-why-and-what-part-5.html?view=sidebar",
	"/2014/05/golang-recursive-function-game-tower-of.html?view=sidebar",
	"/2014/05/basic-go-golang-difference-in-cc-go.html?view=sidebar",
	"/2014/05/basic-golang-why-and-what-part-4-naming.html?view=sidebar",
	"/2014/05/tick-tack-toe-tic-tac-toe-tick-tat-toe.html?view=sidebar",
	"/2014/05/basic-golang-why-and-what-part-4-syntax.html?view=sidebar",
	"/2014/05/basic-golang-why-and-what-part-2_27.html?view=sidebar",
	"/2014/05/basic-golang-why-and-what-part-2.html?view=sidebar",
	"/2014/05/why-go-dont-use-space-indentation-as.html?view=sidebar",
	"/2014/05/basic-golang-why-and-what-part-1.html?view=sidebar",
	"/2014/05/guess-number-game-in-go.html?view=sidebar",
}
