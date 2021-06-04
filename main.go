package main

import (
	"io/ioutil"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	gmtext "github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

func main() {
	run()
}

const astTransformerPriority = 10000

type changeLog struct {
	Description string
	Sections    []*Section
}

type Section struct {
	Version    string
	Published  string
	Added      []string
	Changed    []string
	Deprecated []string
	Fixed      []string
	Removed    []string
	Security   []string
}

func run() error {
	contents, err := ioutil.ReadFile("/Users/julie/go/src/github.com/julieqiu/changelog/CHANGELOG.md")
	if err != nil {
		return err
	}
	cl := &changeLog{}
	gdMarkdown := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithASTTransformers(util.Prioritized(cl, 1000)),
		),
	)

	reader := gmtext.NewReader(contents)
	pctx := parser.NewContext()
	gdParser := gdMarkdown.Parser()
	_ = gdParser.Parse(reader, parser.WithContext(pctx))

	/*
		var b bytes.Buffer
		gdRenderer := gdMarkdown.Renderer()
		if err := gdRenderer.Render(&b, contents, doc); err != nil {
			return err
		}
	*/
	spew.Dump(cl)
	return nil
}

func (cl *changeLog) Transform(node *ast.Document, reader text.Reader, pc parser.Context) {
	var (
		curr Section
		h3   string
	)
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		switch t := c.(type) {
		case *ast.Heading:
			switch t.Level {
			case 2:
				curr = Section{}

				h := string(c.Text(reader.Source()))
				parts := strings.SplitN(strings.TrimPrefix(h, "["), "] ", 2)
				curr.Version = parts[0]
				curr.Published = parts[1]
				cl.Sections = append(cl.Sections, &curr)
			case 3:
				h3 = string(t.Text(reader.Source()))
			}
		case *ast.List:
			for x := t.FirstChild(); x != nil; x = x.NextSibling() {
				var section []*Section
				switch h3 {
				case "Added":
					section = curr.Added
				case "Changed":
					section = curr.Changed
				case "Deprecated":
					section = curr.Deprecated
				case "Removed":
					section = curr.Removed
				case "Fixed":
					section = curr.Fixed
				case "Security":
					section = curr.Security
				}
				section = append(section, string(x.Text(reader.Source())))
			}
		}
	}
}
