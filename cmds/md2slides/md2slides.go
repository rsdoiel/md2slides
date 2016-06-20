package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	// 3rd Part packages
	"github.com/russross/blackfriday"
)

const version = "0.0.0"

type Slide struct {
	SlideNo int
	LastNo  int
	Fname   string
	Title   string
	Content string
	Nav     string
}

var (
	showHelp          bool
	showVersion       bool
	presentationTitle string
	defaultHTML       = `<!DOCTYPE html>
<html>
<head>
   <title>{{ .Title }}</title>
   <link href="css/style.css" rel="stylesheet" />
</head>
<body>
	<section>
{{ .Content }}
	</section>
	<nav>
{{ .Nav }}
	</nav>
</body>
</html>
`
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.StringVar(&presentationTitle, "t", "", "Presentation title")
}

func makeSlide(tmpl *template.Template, slide *Slide) {
	sname := fmt.Sprintf(`%02d-%s.html`, slide.SlideNo, slide.Fname)
	fp, err := os.Create(sname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s\n", sname, err)
		os.Exit(1)
	}
	defer fp.Close()
	err = tmpl.Execute(fp, slide)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s %s", sname, err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s\n", sname)
}

func nextPrev(i, last int, fname string) string {
	var p []string
	if i > 0 {
		// Include Previous
		p = append(p,
			fmt.Sprintf(`<a href="%02d-%s.html">prev</a>`, (i-1), fname))
	}
	if i < last {
		// Include Next
		p = append(p,
			fmt.Sprintf(`<a href="%02d-%s.html">next</a>`, (i+1), fname))
	}
	return strings.Join(p, "&nbsp;")
}

func main() {
	appname := path.Base(os.Args[0])

	flag.Parse()

	if showHelp == true {
		fmt.Printf(`
 USAGE: %s [OPTIONS] [FILENAME]

 OPTIONS:

`, appname)

		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\t-%s\t(defaults to %s) %s\n", f.Name, f.Value, f.Usage)
		})

		fmt.Printf("\n\n Version %s\n", version)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Printf("\n\n Version %s\n", version)
		os.Exit(0)
	}

	var fname string
	args := flag.Args()
	if len(args) > 0 {
		fname, args = args[0], args[1:]
	}
	if fname == "" {
		fmt.Fprintf(os.Stderr, "Missing filename")
		os.Exit(1)
	}
	src, err := ioutil.ReadFile(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fname = strings.TrimSuffix(path.Base(fname), path.Ext(fname))
	tmpl, err := template.New("slide").Parse(defaultHTML)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	slides := bytes.Split(src, []byte("--\n"))

	fmt.Printf("Slide count: %d\n", len(slides))
	lastSlide := len(slides) - 1
	for i, s := range slides {
		data := blackfriday.MarkdownCommon(s)
		makeSlide(tmpl, &Slide{
			Fname:   fname,
			SlideNo: i,
			LastNo:  lastSlide,
			Title:   presentationTitle,
			Content: string(data),
			Nav:     nextPrev(i, lastSlide, fname),
		})
	}
}
