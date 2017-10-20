package main

import (
	"fmt"

	"github.com/macroblock/garbage/conio"
	"github.com/macroblock/garbage/utils"
	"github.com/macroblock/zl/core/zlog"
)

var log = zlog.Instance("test1")
var gameMap TMap
var protoMap string = " # # # # # # # # # #" +
	" #                 #" +
	" #     #           #" +
	" #                 #" +
	" #                 #" +
	" #                 #" +
	" # # # # #         #" +
	" #H                #" +
	" #       #         #" +
	" # # # # # # # # # #"

type TMap struct {
	data [mapW * mapH]TCell
	hero IObject
}
type TCell struct {
	object IObject
	ground IObject
}

func newObject(r byte, x, y int) IObject {
	switch rune(r) {
	case ' ':
		return NewObject(x, y)
	case '#':
		return NewWall(100500, x, y)
	case 'H':
		gameMap.hero = NewHero(1, x, y)
		return gameMap.hero
	}
	panic("Unknown object")
	return nil
}

func newMap() {
	for i := 0; i < mapH*mapW; i++ {
		gameMap.data[i].object = newObject(protoMap[i*2], i%mapW, i/mapW)
		gameMap.data[i].ground = newObject(protoMap[i*2+1], i%mapW, i/mapW)
		if gameMap.data[i].object == nil {
			log.Warning(true, fmt.Sprintf("Something wrong %v", i))
		}
	}
}

// func updateMap() {

// 		gameMap.data[i].object =
// 		gameMap.data[i].ground =
// }

func drawCell(scr *conio.TScreen, x, y int) {
	if gameMap.data[x+y*mapW].object == nil || gameMap.data[x+y*mapW].ground == nil {
		// log.Warning(true, fmt.Sprintf("Something wrong %v %v", x, y))
		return
	}
	scr.DrawAlignedString(mapX+x*2, mapY+y, mapW*2, gameMap.data[x+y*mapW].object.Rune()+gameMap.data[x+y*mapW].ground.Rune())
}

func catchKey() {
	// ev := evs.ReadEvent()
	conio.NewKeyboardAction("MoveUp", "w", "", func(ev conio.TKeyboardEvent) bool {
		// hero.moveUp()
		gameMap.data[gameMap.hero.GetX()+gameMap.hero.GetY()*mapW].object = NewObject(gameMap.hero.GetX(), gameMap.hero.GetY())
		gameMap.hero.MoveUp()
		gameMap.data[gameMap.hero.GetX()+gameMap.hero.GetY()*mapW].object = gameMap.hero
		log.Info("Hero move up")
		return true
	})
	conio.NewKeyboardAction("MoveDown", "s", "", func(ev conio.TKeyboardEvent) bool {
		gameMap.data[gameMap.hero.GetX()+gameMap.hero.GetY()*mapW].object = NewObject(gameMap.hero.GetX(), gameMap.hero.GetY())
		gameMap.hero.MoveDown()
		gameMap.data[gameMap.hero.GetX()+gameMap.hero.GetY()*mapW].object = gameMap.hero
		log.Info("Hero move down")
		return true
	})
	conio.NewKeyboardAction("MoveLeft", "a", "", func(ev conio.TKeyboardEvent) bool {
		gameMap.data[gameMap.hero.GetX()+gameMap.hero.GetY()*mapW].object = NewObject(gameMap.hero.GetX(), gameMap.hero.GetY())
		gameMap.hero.MoveLeft()
		gameMap.data[gameMap.hero.GetX()+gameMap.hero.GetY()*mapW].object = gameMap.hero
		log.Info("Hero move left")
		return true
	})
	conio.NewKeyboardAction("MoveRight", "d", "", func(ev conio.TKeyboardEvent) bool {
		gameMap.data[gameMap.hero.GetX()+gameMap.hero.GetY()*mapW].object = NewObject(gameMap.hero.GetX(), gameMap.hero.GetY())
		gameMap.hero.MoveRight()
		gameMap.data[gameMap.hero.GetX()+gameMap.hero.GetY()*mapW].object = gameMap.hero
		log.Info("Hero move right")
		return true
	})
	conio.NewKeyboardAction("Exit", "`", "", func(ev conio.TKeyboardEvent) bool {
		canClose = true
		return true
	})
	conio.ActionMap.Apply()
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
// func (o *gameMap) moveUp() {
// 	o.data[o.hero.GetX()+o.hero.GetY()*mapW].object = NewObject(o.hero.GetX(), o.hero.GetY())
// 	o.hero.MoveUp()
// 	o.data[o.hero.GetX()+o.hero.GetY()*mapW].object = o.hero
// 	log.Info(o.hero.GetType(), " move up")
// }

////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	err := conio.Init()
	utils.Assert(err == nil, "conio init failed")
	defer conio.Close()
	evs := conio.NewEventStream()
	utils.Assert(evs != nil, "eventStream init failed")
	defer evs.Close()
	scr := conio.NewScreen()
	utils.Assert(scr != nil, "screen init failed")
	defer scr.Close()

	initialize()
	newMap()
	catchKey()
	// log.Warning(true, "ping")
	// log.Add(zlogger.Build().Format("(~m) ~l:~s~x~e\n").Styler(zlogger.AnsiStyler).Done())
	// lion := NewLion(1)
	// wall := NewWall(1)
	// barrel := NewBarrel(1)
	// cat := NewCat(1)
	// ghost := NewGhost(1)
	log.Info("text1")
	log.Info("text2")
	for !canClose {
		draw()
		ev := evs.ReadEvent()
		//msgLog = append(msgLog, fmt.Sprintf("%s %T", ev.String(), ev))
		// log.Debug(fmt.Sprintf("%s %T", ev.String(), ev))
		log.Info(fmt.Sprintf("Hero position is: %v", gameMap.hero.PrintPos()))
		conio.HandleEvent(ev)
		//handleEvent(ev)
	}
	log.Info("text")
}
