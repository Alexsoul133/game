package main

import (
	"fmt"
	"strconv"
)

type TObject struct {
	lvl        int
	dmg        int
	realobject IObject
	x, y       int
}

var _ IObject = (*TObject)(nil)

type TBarrel struct {
	TObject
	realmaterial IMaterial
}

var _ IObject = (*TBarrel)(nil)
var _ IMaterial = (*TBarrel)(nil)

type TWall struct {
	TBarrel
}

var _ IObject = (*TWall)(nil)
var _ IMaterial = (*TWall)(nil)

type TGhost struct {
	TObject
	realfighter IFighter
}

var _ IObject = (*TGhost)(nil)
var _ IFighter = (*TGhost)(nil)

type TPhantom struct {
	TGhost
}

type TCat struct {
	TObject
	realfighter  IFighter
	realmaterial IMaterial
}

var _ IObject = (*TCat)(nil)
var _ IMaterial = (*TCat)(nil)
var _ IFighter = (*TCat)(nil)

type TGoblin struct {
	TCat
}

type THero struct {
	TCat
}

type TLion struct {
	TCat
}

var _ IFighter = (*TLion)(nil)

type IObject interface {
	IsDead() bool
	GetLvl() int
	GetDmg() int
	GetType() string
	GetHp() int
	GetMaxHp() int
	GetStr() int
	Status() string
	Rune() string
	MoveUp()
	MoveDown()
	MoveLeft()
	MoveRight()
	PrintPos() string
	GetX() int
	GetY() int
}

type IFighter interface {
	Attack(IObject) (int, bool)
}

type IMaterial interface {
	RecieveDmg(int) int
}

func (o *TObject) GetType() string {
	panic("abstract call TObject.GetType()")
	return ""
}

func (o *TObject) GetHp() int {
	return o.realobject.GetMaxHp() - o.realobject.GetDmg()
}

func (o *TObject) GetLvl() int {
	return o.lvl
}

func (o *TObject) GetMaxHp() int {
	panic("abstract call TObject.GetType()")
	return 0
}

func (o *TObject) GetDmg() int {
	return o.dmg
}

func (o *TObject) GetStr() int {
	panic("abstract call TObject.GetType()")
	return 0
}

func (o *TObject) IsDead() bool {
	return o.realobject.GetHp() <= 0
}

func (a *TObject) Status() string {
	s := "alive"
	if a.realobject.IsDead() {
		s = "dead"
	}
	return fmt.Sprintf("#%v %v HP:%v/%v", a.realobject.GetType(), s, a.realobject.GetHp(), a.realobject.GetMaxHp())
}
func (o *TObject) PrintStatus() {
	log.Debug(o.realobject.Status())
}

func (o *TObject) Rune() string {
	return " "
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func (o *TObject) MoveUp() {
	o.y--
}
func (o *TObject) MoveDown() {
	o.y++
}
func (o *TObject) MoveLeft() {
	o.x--
}
func (o *TObject) MoveRight() {
	o.x++
}
func (o *TObject) GetX() int {
	return o.x
}
func (o *TObject) GetY() int {
	return o.y
}

func (o *TObject) PrintPos() string {
	s := strconv.Itoa(o.realobject.GetX())
	s1 := strconv.Itoa(o.realobject.GetY())
	return s + " " + s1
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewObject(x, y int) *TObject {
	o := &TObject{x: x, y: y}
	o.realobject = o
	return o
}

////////////////////////////////////////////////////////////////////////////////////////////////////////

func Attack(a IObject, b IObject) {

	if attacker, ok := a.(IFighter); ok {
		if dmg, ok := attacker.Attack(b); ok {
			log.Info(fmt.Sprintf("  %v attacks. Inflicted %v dmg %v", a.Status(), dmg, b.Status()))
		} else {
			log.Info(fmt.Sprintf("%v can't be attacked", b.GetType()))
		}
	} else {
		log.Info(fmt.Sprintf("%v can't attack", a.GetType()))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewWall(lvl, x, y int) *TWall {
	o := &TWall{TBarrel: *NewBarrel(lvl, x, y)}
	o.lvl = lvl
	o.realobject = o
	o.realmaterial = o
	return o
}

func (o *TWall) GetMaxHp() int {
	return o.lvl * 200
}

func (o *TWall) GetType() string {
	return "  Wall"
}

func (o *TWall) RecieveDmg(i int) int {
	if !o.IsDead() {
		i = i / 3
		o.dmg += i
		return i
	}
	return 0
}

func (o *TWall) Rune() string {
	return "#"
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewBarrel(lvl, x, y int) *TBarrel {
	o := &TBarrel{TObject: *NewObject(x, y)}
	o.lvl = lvl
	o.realobject = o
	o.realmaterial = o
	return o
}
func (o *TBarrel) GetMaxHp() int {
	return o.lvl * 30
}

func (o *TBarrel) GetType() string {
	return " Barrel"
}

func (o *TBarrel) RecieveDmg(i int) int {
	if !o.IsDead() {
		o.dmg += i
		return i
	}
	return 0
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewCat(lvl, x, y int) *TCat {
	o := &TCat{TObject: *NewObject(x, y)}
	o.lvl = lvl
	o.realobject = o
	o.realfighter = o
	o.realmaterial = o
	return o
}

func (o *TCat) GetType() string {
	return "  Cat"
}
func (a *TCat) Attack(b IObject) (int, bool) {
	if mat, ok := b.(IMaterial); ok {
		o := mat.RecieveDmg(a.realobject.GetStr())
		return o, true
	}
	return 0, false
}

func (o *TCat) GetStr() int {
	return o.lvl * 2
}

func (o *TCat) GetMaxHp() int {
	return 10 + (o.lvl * 2)
}

func (o *TCat) GetLvl() int {
	return o.lvl
}

func (a *TCat) Status() string {

	return fmt.Sprintf("%v   LVL:%v   STR:%v", a.TObject.Status(), a.realobject.GetLvl(), a.realobject.GetStr())
}

func (o *TCat) RecieveDmg(i int) int {
	if !o.IsDead() {
		o.dmg += i
		return i
	}
	return 0
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewHero(lvl, x, y int) *THero {
	o := &THero{TCat: *NewCat(lvl, x, y)}
	o.realobject = o
	o.realfighter = o
	o.realmaterial = o
	return o
}

func (o *THero) GetType() string {
	return "  Hero"
}

func (o *THero) GetStr() int {
	return o.TCat.GetStr() * 4
}

func (o *THero) GetMaxHp() int {
	return o.TCat.GetMaxHp() * 6
}

func (o *THero) Rune() string {
	return "H"
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewGoblin(lvl, x, y int) *TGoblin {
	o := &TGoblin{TCat: *NewCat(lvl, x, y)}
	o.realobject = o
	o.realfighter = o
	o.realmaterial = o
	return o
}

func (o *TGoblin) GetType() string {
	return "  Goblin"
}

func (o *TGoblin) GetStr() int {
	return o.TCat.GetStr() * 3
}

func (o *TGoblin) GetMaxHp() int {
	return o.TCat.GetMaxHp() * 3
}

func (o *TGoblin) Rune() string {
	return "G"
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewLion(lvl, x, y int) *TLion {
	o := &TLion{TCat: *NewCat(lvl, x, y)}
	o.realobject = o
	o.realfighter = o
	o.realmaterial = o
	return o
}

func (o *TLion) GetType() string {
	return "  Lion"
}

func (o *TLion) GetStr() int {
	return o.TCat.GetStr() * 5
}

func (o *TLion) GetMaxHp() int {
	return o.TCat.GetMaxHp() * 5
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewGhost(lvl, x, y int) *TGhost {
	o := &TGhost{TObject: *NewObject(x, y)}
	o.lvl = lvl
	o.realobject = o
	o.realfighter = o
	return o
}

func (o *TGhost) GetType() string {
	return "  Ghost"
}

func (a *TGhost) Attack(b IObject) (int, bool) {
	if mat, ok := b.(IMaterial); ok {
		o := mat.RecieveDmg(a.realobject.GetStr())
		return o, true
	}
	return 0, false

	// o, ok := (*TCat)(a).Attack(b)

	// return o, true
}

func (o *TGhost) GetMaxHp() int {
	return 10 + (o.lvl * 2)
}

func (o *TGhost) GetStr() int {
	return o.lvl * 2
}

////////////////////////////////////////////////////////////////////////////////////////////////////////
func NewPhantom(lvl int) *TPhantom {
	o := &TPhantom{}
	o.lvl = lvl
	o.realobject = o
	o.realfighter = o
	return o
}
