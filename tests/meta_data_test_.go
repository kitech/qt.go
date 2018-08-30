package main

import (
	"log"
	"os"
	"testing"
	"time"
	"unsafe"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtmeta"
	"github.com/kitech/qt.go/qtrt"
)

type WantAsQClass struct {
	*qtcore.QThread  `qt:inherit` // 要继承的类
	*qtcore.QProcess `qt:inherit`

	_ struct{} `qt:classinfo,appName=12345s`  //
	_ struct{} `qt:classinfo,appCorp=hehehhe` //

	Prop123 int    `qt:property,123` // 最后字段为默认值
	Prop456 string `qt:property,"testv123"`
	// Enum123 AType  `qt:enum` // 就怕无法支持

	_          qtmeta.Q_SIGNALS_BEGIN
	Clicked123 func(bool) `qt:signal`

	_ func(int)     `qt:slot,SlotFunc1` // 如果为小写则为私有slot, 大写为公有slot
	_ func(float32) `qt:slot,SlotFunc2`
}

func Test0(t *testing.T) {
	mdo := qtmeta.NewQtMetaData()
	mdo.SetClassName("QtMocRev")
	log.Println(mdo.MetaStrDat.Count())

	mdo.AddClassInfo("Author", "heheh-auu确uthor")
	mdo.AddClassInfo("Date", "heheh-da定te123")
	mdo.AddClassInfo("Purpose", "heheh-not的useo^o")

	emuqtc := &WantAsQClass{}
	mdo.AddVarAsSignal("Clicked123h", emuqtc.Clicked123)
	mdo.FinalPass()
	log.Println(mdo.Dump()) //
	log.Println(mdo.MetaStrDat.Dump())

	objdat := &qtmeta.QMetaObjectData{}
	objdat.Stringdata = mdo.GetMetaStrDat()
	objdat.Data = mdo.GetMetaDat()
	objdat.Superdata = qtrt.GetClassStaticMetaObjectByName("QObject")
	log.Println(objdat.GetCthis(), objdat)

	qapp := qtcore.NewQCoreApplication(len(os.Args), os.Args, 0)

	arrdat := qtcore.NewQArrayDataFromPointer(objdat.Stringdata)
	log.Println(arrdat.IsMutable(), arrdat.Data())

	dumpMetaObj := func(qmetaobj *qtcore.QMetaObject) {
		log.Println("===========")
		log.Println("classname:", qmetaobj.ClassName())
		log.Println("clsinfo count:", qmetaobj.ClassInfoCount())
		log.Println("classinfo offset:", qmetaobj.ClassInfoOffset())
		log.Println("mth cnt:", qmetaobj.MethodCount())
		log.Println("mth offset:", qmetaobj.MethodOffset())
		log.Println("prop cnt:", qmetaobj.PropertyCount())
		log.Println("prop offset:", qmetaobj.PropertyOffset())
		log.Println("enum cnt:", qmetaobj.EnumeratorCount())

		//*
		for i := 0; i < qmetaobj.MethodCount(); i++ {
			qmetamth := qmetaobj.Method(i)

			log.Println(i, qmetamth.ParameterCount(), qmetamth.MethodType(), qmetamth.Name().Data_fix(), qmetamth.MethodSignature().Data_fix())
		}
		//*/
		log.Println("===========")
	}
	tmer := qtcore.NewQTimer__()
	_ = tmer
	dumpMetaObj(tmer.MetaObject())

	qmetaobj := qtcore.NewQMetaObjectFromPointer(objdat.GetCthis())
	log.Println(qmetaobj.ClassName())
	dumpMetaObj(qmetaobj)

	objdat2 := qtmeta.NewQMetaObjectData2().SetData(objdat.Superdata, objdat.Stringdata, objdat.Data, objdat.Static_metacall)
	cobj := qtmeta.DMetaObject_new(objdat2, nil, nil)
	log.Println(cobj)
	cgoobj := newFakeQObjectFromPointer(cobj)

	qtrt.ConnectRaw(cgoobj.GetCthis(), qtrt.QSIGNAL("Clicked123h(bool)"),
		tmer.GetCthis(), qtrt.QSLOT("start()"))

	time.AfterFunc(3*time.Second, func() { qapp.Quit() })
	qapp.Exec()
}

type FakeQObject struct {
	*qtrt.CObject
}

func newFakeQObjectFromPointer(cthis unsafe.Pointer) *FakeQObject {
	this := &FakeQObject{}
	this.CObject = &qtrt.CObject{}
	this.CObject.Cthis = cthis
	return this
}

func (this *FakeQObject) GetCthis() unsafe.Pointer      { return this.Cthis }
func (this *FakeQObject) SetCthis(cthis unsafe.Pointer) { this.Cthis = cthis }

func Test1(t *testing.T) {

}

func main() {
	Test0(&testing.T{})
}
