package main

import (
	"fmt"
	. "gopkg.in/metakeule/cascade.v1"
	"gopkg.in/metakeule/cascade.v1/class"
	. "gopkg.in/metakeule/goh4.v5/tag"
	// . "gopkg.in/metakeule/goh4.v5/tag/short"
	"net/http"
)

type ()

var (
	_      = fmt.Println
	layout = Cascade("/assets")
	panel  = Panel("headertext", "footertext")
	menu   = DIV(class.Site_Center,
		DIV(class.Col, class.Width_Fill,
			DIV(class.Cell,
				IconButton(class.Icon_Plus, 1),
				IconButton(class.Icon_Minus, 1),
				IconButton(class.Icon_Edit, 1),
				IconButton(class.Icon_Lock, 1),
				IconButton(class.Icon_Unlock, 1),
				IconButton(class.Icon_Copy, 1),
				IconButton(class.Icon_Trash, 1),
				IconButton(class.Icon_Move, 1),
			),
		),
	)

	footer = DIV(class.Site_Footer,
		DIV(class.Col, class.Width_Fill,
			DIV(class.Cell,
				IconButton(class.Icon_Plus, 1),
				IconButton(class.Icon_Minus, 1),
				IconButton(class.Icon_Edit, 1),
				IconButton(class.Icon_Copy, 1),
				IconButton(class.Icon_Trash, 1),
				IconButton(class.Icon_Move, 1),
			),
		),
	)
)

func init() {
	panel.Add(H2(class.Cell, "hier eine Überschrift"))
	panel.Add(P(class.Cell, "hier ein Absatz"))
	panel.AddClass(class.Col, class.Width_Fill)

	a := DIV(class.Col, class.Width_1of4, H1("here one"))
	b := DIV(class.Col, class.Width_1of4, P("here another one"))
	p := panel

	layout.Add(
		DIV(class.Site_HeaderFixture,
			DIV(class.Site_HeaderGhost, menu),
			DIV(class.Site_Header, menu),
		),

		DIV(class.Site_Body,
			DIV(class.Site_Center,
				DIV(class.Cell,
					a,
					b,
					p,
				),
			),
		),

		DIV(class.Site_Center,
			DIV(class.Site_HeaderGhost, footer),
			DIV(class.Site_FooterFixture, footer),
		),
	)
}

func main() {
	fileserver := http.FileServer(http.Dir("/home/benny/Entwicklung/gopath/src/github.com/metakeule/cascade/"))
	http.Handle("/assets/", fileserver)
	// http.Handle("/css/", fileserver)
	// http.Handle("/img/", fileserver)
	http.Handle("/", layout)
	http.ListenAndServe(":8080", nil)
}
