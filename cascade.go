package cascade

import (
	"fmt"
	"github.com/metakeule/cascade/class"
	"github.com/metakeule/goh4"
	. "github.com/metakeule/goh4/tag"
	. "github.com/metakeule/goh4/tag/short"
	"path"
)

type cascade struct {
	*DocType
	AssetsDir string
	Head      *goh4.Element
	Content   *goh4.Element
	body      *goh4.Element
}

func (ø *cascade) CssPath(file string) string {
	return path.Join(ø.AssetsDir, "css", file)
}

func (ø *cascade) JsPath(file string) string {
	return path.Join(ø.AssetsDir, "js", file)
}

func (ø *cascade) CssHref(file string, opts ...interface{}) *goh4.Element {
	return CssHref(ø.CssPath(file), opts...)
}

func (ø *cascade) AddJsFile(path string) {
	ø.body.Add(JsSrc(path))
}

func (ø *cascade) Add(opts ...interface{}) {
	ø.Content.Add(opts...)
}

func Cascade(assetsDir string) (ø *cascade) {
	ø = &cascade{AssetsDir: assetsDir}
	ø.Head = HEAD(
		CharsetUtf8(),
		ø.CssHref("cascade/production/build-full.min.css", ATTR("media", "all")),
		HTML(`<!--[if lt IE 8]><link rel="stylesheet" href="`+ø.CssPath("cascade/production/icons-ie7.min.css")+`"><![endif]-->`),
		HTML(`<!--[if lt IE 9]><script src="`+ø.JsPath(`lib/polyfills/iehtmlshiv.js`)+`"></script><![endif]-->`),
		META(ATTR("name", "viewport", "content", "width=device-width, initial-scale=1")),
	)
	ø.Content = Doc()
	//	ø.body = BODY(JsSrc(ø.JsPath("app.js")), ø.Content)
	ø.body = BODY(ø.Content)
	ø.DocType = HTML5(ø.Head, ø.body)
	return
}

type panel struct {
	outer    *goh4.Element
	inner    *goh4.Element
	collapse bool
}

func (ø *panel) String() string {
	return ø.outer.String()
}

func (ø *panel) wrap(element interface{}) *goh4.Element {
	if ø.collapse {
		return DIV(class.Body, class.CollapseSection, element)
	}
	return DIV(class.Body, element)
}

func (ø *panel) Add(elements ...interface{}) {
	opts := make([]interface{}, len(elements))
	for i, e := range elements {
		opts[i] = ø.wrap(e)
	}
	ø.inner.Add(opts...)
}

func (ø *panel) AddClass(classes ...goh4.Class) {
	ø.outer.AddClass(classes...)
}

func Panel(header, footer interface{}) *panel {
	inner := Doc()
	outer := DIV(class.Cell, class.Panel)

	if header != nil {
		outer.Add(DIV(class.Header, header))
	}
	outer.Add(inner)
	if footer != nil {
		outer.Add(DIV(class.Footer, footer))
	}
	return &panel{
		outer,
		inner,
		false,
	}
}

func CollapsingPanel(header, footer interface{}) *panel {
	inner := Doc()
	outer := DIV(class.Cell, class.Panel)

	if header != nil {
		outer.Add(DIV(class.Header, DIV(class.Header, class.CollapseTrigger,
			SPAN(class.Icon, class.IconCollapse),
			AHref("#", header),
		)))
	}
	outer.Add(inner)
	if footer != nil {
		outer.Add(DIV(class.Footer, class.CollapseSection, footer))
	}
	return &panel{
		outer,
		inner,
		true,
	}
}

type tabBlock struct {
	nav     *goh4.Element
	content *goh4.Element
	outer   *goh4.Element
	id      string
	i       int
}

func TabBlockTop(id string, outer *goh4.Element) *tabBlock {
	nav := UL(class.Nav)
	content := DIV(class.TabContent)
	outer.Add(class.TabBlock, class.TopNav, DIV(class.Tabs, nav), content)
	return &tabBlock{nav, content, outer, id, 0}
}

func TabBlockBottom(id string, outer *goh4.Element) *tabBlock {
	nav := UL(class.Nav)
	content := DIV(class.TabContent)
	outer.Add(class.TabBlock, class.BottomNav, content, DIV(class.Tabs, nav))
	return &tabBlock{nav, content, outer, id, 0}
}

func TabBlockLeft(id string, outer *goh4.Element, colNum int) *tabBlock {
	nav := UL(class.Nav)
	content := DIV(class.TabContent, class.Width_Fill)
	n_class := fmt.Sprintf("width-1of%v", colNum)
	outer.Add(
		class.TabBlock, class.LeftNav,
		DIV(class.Tabs, goh4.Class(n_class), nav),
		content,
	)
	return &tabBlock{nav, content, outer, id, 0}
}

func TabBlockRight(id string, outer *goh4.Element, colNum int) *tabBlock {
	nav := UL(class.Nav)
	c_class := fmt.Sprintf("width-%vof%v", colNum-1, colNum)
	content := DIV(class.TabContent, goh4.Class(c_class))
	outer.Add(
		class.TabBlock, class.RightNav,
		content,
		DIV(class.Tabs, class.Width_Fill, nav),
	)
	return &tabBlock{nav, content, outer, id, 0}
}

func (ø *tabBlock) String() string {
	return ø.outer.String()
}

func (ø *tabBlock) Add(name string, element *goh4.Element, active bool) {
	ø.i++
	id := fmt.Sprintf("%s_%v", ø.id, ø.i)
	li := LI(AHref("#"+id, name))
	if active {
		li.AddClass(class.Active)
	}
	ø.nav.Add(li)
	element.SetId(goh4.Id(id))
	ø.content.Add(element)
}

// size is multiplied with 16
func Icon(typ goh4.Class, size int, opts ...interface{}) *goh4.Element {
	if size < 1 {
		size = 1
	}
	if size > 4 {
		panic(fmt.Sprintf("unsupported icon size: %v (choose 1 to 4)", size*16))
	}
	size = size * 16
	s := SPAN(typ, class.Icon_Adjust, class.Icon, goh4.Class(fmt.Sprintf("icon-%v", size)))
	s.Add(opts...)
	return s
}

func IconBorder(typ goh4.Class, size int, opts ...interface{}) *goh4.Element {
	i := Icon(typ, size, opts...)
	i.AddClass(class.IconBorder)
	return i
}

func IconButton(typ goh4.Class, size int, opts ...interface{}) *goh4.Element {
	btn := BUTTON(class.IconButton, class.Button, SPAN(typ, class.Icon))
	btn.Add(opts...)
	return btn
}
