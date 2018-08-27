package qtmeta

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"unsafe"
)

func Test1(t *testing.T) {
	// log.Println(t)
	log.Println(unsafe.Sizeof(int(0)))
	log.Println(unsafe.Sizeof(struct{ bool }{}))
	log.Println(unsafe.Sizeof(QByteArrayData{}))
	log.Println(unsafe.Sizeof(QByteArrayData2{}))

	msd := NewQtMetaStringData()
	msd.Append("FMocCls")
	msd.Append("")
	msd.Append("FMocCls123*")
	msd.FinalPass()
	cstr := msd.ToCs()
	log.Println(msd.Count(), cstr)

	dpstr := msd.Dump()
	log.Println(dpstr)
}

func Test2(t *testing.T) {
	mdo := NewQtMetaData()
	mdo.SetClassName("QtMocRev")
	log.Println(mdo.MetaStrDat.Count())
	mdo.SetClassName("QtMocRev")
	log.Println(mdo.MetaStrDat.Count())

	mdo.AddClassInfo("Author", "heheh-auu确uthor")
	mdo.AddClassInfo("Date", "heheh-da定te123")
	mdo.AddClassInfo("Purpose", "heheh-not的useo^o")

	var sigf1 func()
	var sigf2 func(int)
	var sigf3 func(byte)
	var sigf4 func(string)
	var sigf5 func(string) int
	var sigf6 func(unsafe.Pointer)
	var sigf7 func() string

	var sigfns = []interface{}{sigf1, sigf2, sigf3, sigf4, sigf5, sigf6, sigf7}
	_ = sigfns

	for idx, sigfn := range sigfns {
		sigfnty := reflect.TypeOf(sigfn)
		rettys := []reflect.Type{}
		argtys := []reflect.Type{}
		for i := 0; i < sigfnty.NumOut(); i++ {
			rettys = append(rettys, sigfnty.Out(i))
		}
		for i := 0; i < sigfnty.NumIn(); i++ {
			argtys = append(argtys, sigfnty.In(i))
		}
		mdo.AddSignal(fmt.Sprintf("dsigf%d", idx), rettys, argtys)
		mdo.AddSlot(fmt.Sprintf("dslotf%d", idx), rettys, argtys)
	}
	mdo.FinalPass()
	log.Println(mdo.Dump())
	log.Println(mdo.MetaStrDat.Dump())

	modat := QMetaObjectData{}
	modat.superdata = nil
	modat.stringdata = mdo.MetaStrDat.ToC()
	modat.data = mdo.ToC()

}
