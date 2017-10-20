package main

import (
	"io"

	"github.com/macroblock/garbage/conio"
	"github.com/macroblock/garbage/utils"

	"github.com/macroblock/zl/core/zlogger"
)

var (
	// log           = zlog.Instance("game")
	key           int
	ch            rune
	width, height int
	canClose      bool
	msgLog        = make([]string, 0)
)

const (
	mapX     = 2
	mapY     = 2
	mapW     = 10
	mapH     = 10
	logWidth = 100
)

var CustomWriter TCustomWriter

type x io.Writer

type TCustomWriter struct {
}

func (o *TCustomWriter) Write(p []byte) (n int, err error) {
	msgLog = append(msgLog, string(p[:]))
	return len(p), nil
}

func initialize() {
	log.Add(zlogger.Build().Format("(~m) ~l~s~x~e").Writer(&CustomWriter).Done())

	conio.Screen().Flush()
	width = conio.Screen().Width()
	height = conio.Screen().Height() - 26

	// conio.NewKeyboardAction("Exit", "`", "", func(ev conio.TKeyboardEvent) bool {
	// 	canClose = true
	// 	return true
	// })
	// conio.ActionMap.Apply()
}

func draw() {
	scr := conio.Screen()
	winFg := conio.ColorWhite
	winBg := conio.ColorBlue
	logFg := conio.ColorWhite
	logBg := conio.ColorDarkGray
	scr.Clear('░', conio.ColorWhite, conio.ColorBlack)
	scr.SelectBorder("Single")
	scr.SetColor(logFg, logBg)
	drawWindow(width-logWidth-1, 0, logWidth, height, "[ Log ]", func(x, y, w, h int) {
		for i := 0; i < utils.Min(len(msgLog), h); i++ {
			scr.DrawAlignedString(x, y+h-1-i, w, msgLog[len(msgLog)-1-i])
		}
	})

	scr.SelectBorder("Double")
	scr.SetColor(winFg, winBg)
	drawWindow(mapX-1, mapY-1, mapW*2+2, mapH+2, "[ Map ]", func(x, y, w, h int) {
		for i := 0; i < mapW*mapH; i++ {
			drawCell(scr, i%mapH, i/mapH)
			//scr.DrawAlignedString(x+offsX, y+i, w-offsX, name)
		}
	})

	scr.Flush()
}

func drawWindow(x, y, w, h int, title string, draw func(x, y, w, h int)) {
	scr := conio.Screen()
	scr.DrawBorder(x, y, w, h)
	scr.FillRect(x+1, y+1, w-2, h-2, ' ')
	scr.DrawAlignedString(x+1, y, w-2, title)
	draw(x+1, y+1, w-2, h-2)
}

// func main() {
// 	err := conio.Init()
// 	utils.Assert(err == nil, "conio init failed")
// 	defer conio.Close()
// 	evs := conio.NewEventStream()
// 	utils.Assert(evs != nil, "eventStream init failed")
// 	defer evs.Close()
// 	scr := conio.NewScreen()
// 	utils.Assert(scr != nil, "screen init failed")
// 	defer scr.Close()

// 	initialize()
// 	for !canClose {
// 		draw()
// 		ev := evs.ReadEvent()
// 		//msgLog = append(msgLog, fmt.Sprintf("%s %T", ev.String(), ev))
// 		log.Debug(fmt.Sprintf("%s %T", ev.String(), ev))
// 		conio.HandleEvent(ev)
// 		//handleEvent(ev)
// 	}
// }
