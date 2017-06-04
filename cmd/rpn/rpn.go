package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jroimartin/gocui"
	"github.com/pcasaretto/rpn/internal/rpn"
)

var stack *rpn.Stack

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	stack = rpn.NewStack()
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
	// scanner := bufio.NewScanner(os.Stdout)
	// // Set the split function for the scanning operation.
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	text := scanner.Text()
	// 	log.Println(stack.Contents)
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading input:", err)
	// }
}

var stackView, editorView *gocui.View

func layout(g *gocui.Gui) (err error) {
	_, maxY := g.Size()
	if stackView, err = g.SetView("stack", 0, 0, 80, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return
		}
	}
	if editorView, err = g.SetView("input", 0, maxY-20, 80, 35); err != nil {
		if err != gocui.ErrUnknownView {
			return
		}
		if _, err = g.SetCurrentView("input"); err != nil {
			return
		}
		editorView.Wrap = true
		editorView.Editor = singleLineEditor
		editorView.Editable = true
	}
	return
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

var singleLineEditor gocui.Editor = gocui.EditorFunc(func(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	if key == gocui.KeyEnter {
		line := v.Buffer()

		line = strings.Replace(line, "\x00", "", -1)
		line = strings.Replace(line, "\n", "", -1)
		line = strings.Replace(line, " ", "", -1)

		if line == "+" {
			rpn.Add(stack)
		} else {
			f, err := strconv.ParseFloat(line, 64)
			if err != nil {
				panic(err)
			}
			stack.Push(f)
		}
		updateStackView()

		v.Clear()
		v.SetCursor(0, 0)
		v.SetOrigin(0, 0)
		return
	}
	gocui.DefaultEditor.Edit(v, key, ch, mod)
})

func updateStackView() {
	stackView.Clear()
	stackView.SetCursor(0, 0)
	stackView.SetOrigin(0, 0)
	for _, f := range stack.Contents {
		fmt.Fprintln(stackView, f)
	}
}
