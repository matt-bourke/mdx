package mdx

import (
	"testing"
)

func defaultProps(t *testing.T) []property {
	t.Helper()

	properties := make([]property, 0)
	properties = append(properties, property{Name: "class", Value: "test"})
	properties = append(properties, property{Name: "style", Value: "background-color: red"})

	return properties
}

func TestAstFragmentHtml(t *testing.T) {
	fragment := fragment{}
	fragmentHtml := fragment.Raw()
	expected := ""
	if fragmentHtml != expected {
		t.Errorf("Fragment wrong, got=%q", fragmentHtml)
	}

	fragment.Value = "Hello"
	fragmentHtml = fragment.Raw()
	expected = "Hello"
	if fragmentHtml != expected {
		t.Errorf("Fragment wrong, got=%q", fragmentHtml)
	}
}

func TestAstHeaderHtml(t *testing.T) {
	h1 := header{Level: 1, Content: []component{&fragment{Value: "Test"}}}
	headerHtml := h1.Raw()
	expected := "<h1>Test</h1>"
	if headerHtml != expected {
		t.Errorf("Header wrong, got=%q", headerHtml)
	}

	h1 = header{Level: 2, Content: []component{&fragment{Value: "Test2"}}}
	headerHtml = h1.Raw()
	expected = "<h2>Test2</h2>"
	if headerHtml != expected {
		t.Errorf("Header wrong, got=%q", headerHtml)
	}

	properties := defaultProps(t)
	h1 = header{Level: 6, Content: []component{&fragment{Value: "Test with Props"}}, Properties: properties}
	headerHtml = h1.Raw()
	expected = "<h6 class=\"test\" style=\"background-color: red\">Test with Props</h6>"
	if headerHtml != expected {
		t.Errorf("Header wrong, got=%q", headerHtml)
	}
}

func TestAstParagraphHtml(t *testing.T) {
	paragraph := paragraph{Content: []component{&fragment{Value: "Paragraph test"}}}
	paragraphHtml := paragraph.Raw()
	expected := "<p>Paragraph test</p>"
	if paragraphHtml != expected {
		t.Errorf("Paragraph wrong, got=%q", paragraphHtml)
	}
}

func TestAstCodeHtml(t *testing.T) {
	code := code{Text: "fmt.Printf(\"Hello, world!\n\")"}
	codeHtml := code.Raw()
	expected := "<code>fmt.Printf(\"Hello, world!\n\")</code>"
	if codeHtml != expected {
		t.Errorf("Code wrong, got=%q", codeHtml)
	}

	code.Properties = defaultProps(t)
	codeHtml = code.Raw()
	expected = "<code class=\"test\" style=\"background-color: red\">fmt.Printf(\"Hello, world!\n\")</code>"
	if codeHtml != expected {
		t.Errorf("Code properties wrong, got=%q", codeHtml)
	}
}

func TestAstBoldHtml(t *testing.T) {
	bold := bold{Content: []component{&fragment{Value: "stronk"}}}
	boldHtml := bold.Raw()
	expected := "<strong>stronk</strong>"
	if boldHtml != expected {
		t.Errorf("Bold wrong, got=%q", boldHtml)
	}

	bold.Properties = defaultProps(t)
	boldHtml = bold.Raw()
	expected = "<strong class=\"test\" style=\"background-color: red\">stronk</strong>"
	if boldHtml != expected {
		t.Errorf("Bold properties wrong, got=%q", boldHtml)
	}
}

func TestAstItalicHtml(t *testing.T) {
	italic := italic{Content: []component{&fragment{Value: "italian"}}}
	italicHtml := italic.Raw()
	expected := "<em>italian</em>"
	if italicHtml != expected {
		t.Errorf("Italic wrong, got=%q", italicHtml)
	}

	italic.Properties = defaultProps(t)
	italicHtml = italic.Raw()
	expected = "<em class=\"test\" style=\"background-color: red\">italian</em>"
	if italicHtml != expected {
		t.Errorf("Italic properties wrong, got=%q", italicHtml)
	}
}

func TestAstBlockQuoteHtml(t *testing.T) {
	blockquote := blockQuote{Content: []component{&fragment{Value: "quote"}}}
	blockquoteHtml := blockquote.Raw()
	expected := "<blockquote>quote</blockquote>"
	if blockquoteHtml != expected {
		t.Errorf("Blockquote wrong, got=%q", blockquoteHtml)
	}

	blockquote.Properties = defaultProps(t)
	blockquoteHtml = blockquote.Raw()
	expected = "<blockquote class=\"test\" style=\"background-color: red\">quote</blockquote>"
	if blockquoteHtml != expected {
		t.Errorf("Blockquote properties wrong, got=%q", blockquoteHtml)
	}
}

func TestAstListItemHtml(t *testing.T) {
	listItem := listItem{Component: &paragraph{Content: []component{&fragment{Value: "Item #1"}}}}
	listItemHtml := listItem.Raw()
	expected := "<li><p>Item #1</p></li>"
	if listItemHtml != expected {
		t.Errorf("ListItem wrong, got=%q", listItemHtml)
	}

	listItem.Properties = defaultProps(t)
	listItemHtml = listItem.Raw()
	expected = "<li class=\"test\" style=\"background-color: red\"><p>Item #1</p></li>"
	if listItemHtml != expected {
		t.Errorf("ListItem properties wrong, got=%q", listItemHtml)
	}
}

func TestAstOrderedListHtml(t *testing.T) {
	listItem1 := listItem{Component: &paragraph{Content: []component{&fragment{Value: "Item #1"}}}}
	listItem2 := listItem{Component: &paragraph{Content: []component{&fragment{Value: "Item #2"}}}}
	listItems := []listItem{listItem1, listItem2}
	list := orderedList{ListItems: listItems, Start: 1}
	listHtml := list.Raw()
	expected := "<ol start=\"1\">\n    <li><p>Item #1</p></li>\n    <li><p>Item #2</p></li>\n</ol>"
	if listHtml != expected {
		t.Errorf("OrderedList wrong, got=%q", listHtml)
	}

	list.Properties = defaultProps(t)
	list.Start = 5
	listHtml = list.Raw()
	expected = "<ol start=\"5\" class=\"test\" style=\"background-color: red\">\n    <li><p>Item #1</p></li>\n    <li><p>Item #2</p></li>\n</ol>"
	if listHtml != expected {
		t.Errorf("OrderedList properties wrong, got=%q", listHtml)
	}
}

func TestAstUnorderedListHtml(t *testing.T) {
	listItem1 := listItem{Component: &paragraph{Content: []component{&fragment{Value: "Item #1"}}}}
	listItem2 := listItem{Component: &paragraph{Content: []component{&fragment{Value: "Item #2"}}}}
	listItems := []listItem{listItem1, listItem2}
	list := unorderedList{ListItems: listItems}
	listHtml := list.Raw()
	expected := "<ul>\n    <li><p>Item #1</p></li>\n    <li><p>Item #2</p></li>\n</ul>"
	if listHtml != expected {
		t.Errorf("UnorderedList wrong, got=%q", listHtml)
	}

	list.Properties = defaultProps(t)
	listHtml = list.Raw()
	expected = "<ul class=\"test\" style=\"background-color: red\">\n    <li><p>Item #1</p></li>\n    <li><p>Item #2</p></li>\n</ul>"
	if listHtml != expected {
		t.Errorf("UnorderedList properties wrong, got=%q", listHtml)
	}
}

func TestAstImageHtml(t *testing.T) {
	img := image{ImgUrl: "https://img.pokemondb.net/artwork/avif/regirock.avif", AltText: "Reginald"}
	imgHtml := img.Raw()
	expected := "<img src=\"https://img.pokemondb.net/artwork/avif/regirock.avif\" alt=\"Reginald\"/>"
	if imgHtml != expected {
		t.Errorf("Image wrong, got=%q", imgHtml)
	}

	img.Properties = defaultProps(t)
	imgHtml = img.Raw()
	expected = "<img class=\"test\" style=\"background-color: red\" src=\"https://img.pokemondb.net/artwork/avif/regirock.avif\" alt=\"Reginald\"/>"
}

func TestAstHorizontalRuleHtml(t *testing.T) {
	rule := horizontalRule{}
	ruleHtml := rule.Raw()
	expected := "<hr/>"
	if ruleHtml != expected {
		t.Errorf("HorizontalRule wrong, got=%q", ruleHtml)
	}

	rule.Properties = defaultProps(t)
	ruleHtml = rule.Raw()
	expected = "<hr class=\"test\" style=\"background-color: red\"/>"
	if ruleHtml != expected {
		t.Errorf("HorizontalRule properties wrong, got=%q", ruleHtml)
	}
}

func TestAstLinkHtml(t *testing.T) {
	link := link{Url: "https://google.com", Content: []component{&fragment{Value: "Google"}}}
	linkHtml := link.Raw()
	expected := "<a href=\"https://google.com\" target=_blank>Google</a>"
	if linkHtml != expected {
		t.Errorf("Link wrong, got=%q", linkHtml)
	}

	link.Properties = defaultProps(t)
	linkHtml = link.Raw()
	expected = "<a class=\"test\" style=\"background-color: red\" href=\"https://google.com\" target=_blank>Google</a>"
	if linkHtml != expected {
		t.Errorf("Link properties wrong, got=%q", linkHtml)
	}
}

func TestAstButtonHtml(t *testing.T) {
	button := button{OnClick: "handleClick", Content: []component{&paragraph{Content: []component{&fragment{Value: "Click Me"}}}}}
	buttonHtml := button.Raw()
	expected := "<button onclick=\"handleClick(this)\">\n    <p>Click Me</p>\n</button>"
	if buttonHtml != expected {
		t.Errorf("Button wrong, got=%q", buttonHtml)
	}

	button.Properties = defaultProps(t)
	buttonHtml = button.Raw()
	expected = "<button class=\"test\" style=\"background-color: red\" onclick=\"handleClick(this)\">\n    <p>Click Me</p>\n</button>"
	if buttonHtml != expected {
		t.Errorf("Button properties wrong, got=%q", buttonHtml)
	}
}

func TestAstDivHtml(t *testing.T) {
	emptyDiv := div{}
	divHtml := emptyDiv.Raw()
	expected := "<div/>"
	if divHtml != expected {
		t.Errorf("Empty Div wrong, got=%q", divHtml)
	}

	properties := defaultProps(t)
	propertyDiv := div{Properties: properties}
	divHtml = propertyDiv.Raw()
	expected = "<div class=\"test\" style=\"background-color: red\"/>"
	if divHtml != expected {
		t.Errorf("Property Div wrong, got=%q", divHtml)
	}

	p := &paragraph{Content: []component{&fragment{Value: "child"}}}
	childDiv := div{Children: []component{p}}
	divHtml = childDiv.Raw()
	expected = "<div>\n    <p>child</p>\n</div>"
	if divHtml != expected {
		t.Errorf("Child div wrong, got=%q", divHtml)
	}
}

func TestAstNavHtml(t *testing.T) {
	nav := nav{}
	navHtml := nav.Raw()
	expected := "<nav/>"
	if navHtml != expected {
		t.Errorf("Empty nav wrong, got=%q", navHtml)
	}

	nav.Properties = defaultProps(t)
	navHtml = nav.Raw()
	expected = "<nav class=\"test\" style=\"background-color: red\"/>"
	if navHtml != expected {
		t.Errorf("Nav properties wrong, got=%q", navHtml)
	}

	nav.Children = []component{&link{Url: "https://test.com", Content: []component{&fragment{Value: "Test"}}}}
	navHtml = nav.Raw()
	expected = "<nav class=\"test\" style=\"background-color: red\">\n    <a href=\"https://test.com\" target=_blank>Test</a>\n</nav>"
	if navHtml != expected {
		t.Errorf("Nav children wrong, got=%q", navHtml)
	}
}

func TestAstSpanHtml(t *testing.T) {
	span := span{}
	spanHtml := span.Raw()
	expected := "<span/>"
	if spanHtml != expected {
		t.Errorf("Span wrong, got=%q", spanHtml)
	}

	span.Content = []component{&paragraph{Content: []component{&fragment{Value: "Hello"}}}}
	spanHtml = span.Raw()
	expected = "<span><p>Hello</p></span>"
	if spanHtml != expected {
		t.Errorf("Span children wrong, got=%q", spanHtml)
	}
}

func TestAstCodeBlockHtml(t *testing.T) {
	content := `package main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("Hello, world!")\n}`
	codeBlock := codeBlock{Content: content}
	codeBlockHtml := codeBlock.Raw()
	expected := `<div class="code-block">
    <pre>package main</pre>
    <pre></pre>
    <pre>import "fmt"</pre>
    <pre></pre>
    <pre>func main() {</pre>
    <pre>    fmt.Println("Hello, world!")</pre>
    <pre>}</pre>
</div>`

	if codeBlockHtml != expected {
		t.Errorf("CodeBlock wrong\ngot=     %q\nexpected=%q", codeBlockHtml, expected)
	}
}
