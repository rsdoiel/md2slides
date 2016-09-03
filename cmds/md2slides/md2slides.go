//
// md2slides.go - A simple command line utility that uses Markdown
// to generate a sequence of HTML5 pages that can be used for presentations.
//
// @author R. S. Doiel, <rsdoiel@gmail.com>
//
// Copyright (c) 2016, R. S. Doiel
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above copyright notice,
//   this list of conditions and the following disclaimer in the documentation
//   and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"

	// 3rd Part packages

	// My packages
	"github.com/rsdoiel/md2slides"
)

const (
	version = md2slides.Version
	license = `
%s

Copyright (c) 2016, R. S. Doiel
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`
)

var (
	showHelp          bool
	showVersion       bool
	showLicense       bool
	presentationTitle string
	cssPath           string
	jsPath            string
	templateFName     string
	templateSource    = md2slides.DefaultTemplateSource
)

func init() {
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showLicense, "l", false, "display license")
	flag.BoolVar(&showLicense, "license", false, "display license")
	flag.StringVar(&presentationTitle, "t", "", "Presentation title")
	flag.StringVar(&presentationTitle, "title", "", "Presentation title")
	flag.StringVar(&cssPath, "c", cssPath, "Specify the CSS file to use")
	flag.StringVar(&cssPath, "css", cssPath, "Specify the CSS file to use")
	flag.StringVar(&jsPath, "j", jsPath, "Specify a js file to include")
	flag.StringVar(&jsPath, "js", jsPath, "Specify a js file to include")
	flag.StringVar(&templateFName, "template", templateFName, "Specify an HTML template to use")
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
			if strings.Compare(f.Name, "template") == 0 {
				fmt.Printf("\t-%s\t%s\n", f.Name, f.Usage)
			} else if len(f.Name) > 1 {
				fmt.Printf("\t-%s, -%s\t%s\n", f.Name[0:1], f.Name, f.Usage)
			}
		})

		fmt.Printf("\n\n Version %s\n", version)
		os.Exit(0)
	}
	if showVersion == true {
		fmt.Printf(" Version %s\n", version)
		os.Exit(0)
	}
	if showLicense == true {
		fmt.Printf(license, appname)
		os.Exit(0)
	}

	if templateFName != "" {
		src, err := ioutil.ReadFile(templateFName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s %s\n", templateFName, err)
			os.Exit(1)
		}
		templateSource = string(src)
	}

	//FIXME: If it is markdown file then assign fname that value, otherwise it's a template add it to the
	// list of templates to compile.
	var fname string
	args := flag.Args()
	if len(args) > 0 {
		//fname, args = args[0], args[1:]
		fname = args[0]
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
	tmpl, err := template.New("slide").Parse(templateSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	// Build the slides
	slides := md2slides.MarkdownToSlides(fname, presentationTitle, cssPath, jsPath, src)
	// Render the slides
	for i, slide := range slides {
		err := md2slides.MakeSlideFile(tmpl, slide)
		if err == nil {
			// Note: Give some feed back when slide written successful
			fmt.Fprintf(os.Stdout, "Wrote %02d-%s.html\n", slide.CurNo, slide.FName)
		} else {
			// Note: Display an error if we have a problem
			fmt.Fprintf(os.Stderr, "Can't process slide %d, %s\n", i, err)
		}
	}
	// Render the TOC slide
	slide, err := md2slides.SlidesToTOCSlide(slides)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't create a table of contents slide", err)
		os.Exit(1)
	}
	err = md2slides.MakeTOCSlideFile(tmpl, slide)
	if err == nil {
		// Note: Give some feed back when slide written successful
		fmt.Fprintf(os.Stdout, "Wrote toc-%s.html\n", slide.FName)
	} else {
		// Note: Display an error if we have a problem
		fmt.Fprintf(os.Stderr, "Can't write toc-%s.html, %s\n", slide.FName, err)
	}
}
