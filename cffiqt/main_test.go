package ffiqt

import (
	"log"
	"regexp"
	"testing"
)

func main_test() {
	// test()
	// main_a()
	test_handy_qt()
}

func TestGen0(t *testing.T) {
	main_test()
}

func TestDynSlot0(t *testing.T) {
	dso := NewQDynSlotObject("dummysignalfromffiqtgo", 0)
	log.Println(dso)
	dso2 := NewQDynSlotObject("dummy2signalfromffiqtgo", 0)
	log.Println(dso2)
	dso.Connect_test(dso2)
	dso.Trigger_test()
}

func TestQclsName(t *testing.T) {
	reg := regexp.MustCompile(`^(qt.*)\.Q[A-Z](.+)`)
	if !reg.MatchString("qtcore.QString") {

	}
	t.Error(reg.FindAllStringSubmatch("qtcore.QString", -1))
}

func TestQSize0(t *testing.T) {
	szo := NewQSize()
	log.Println(szo)
	log.Println(szo.Height())
	log.Println(szo.RHeight())
}

func TestQSizeF0(t *testing.T) {
	szo := NewQSizeF()
	log.Println(szo)
	log.Println(szo.Height())
	log.Println(szo.RHeight())
}
