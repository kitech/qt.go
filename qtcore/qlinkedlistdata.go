package qtcore

// /usr/include/qt/QtCore/qlinkedlist.h
// #include <qlinkedlist.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QLinkedListData struct {
	*qtrt.CObject
}

func (this *QLinkedListData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QLinkedListData) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQLinkedListDataFromPointer(cthis unsafe.Pointer) *QLinkedListData {
	return &QLinkedListData{&qtrt.CObject{cthis}}
}
func (*QLinkedListData) NewFromPointer(cthis unsafe.Pointer) *QLinkedListData {
	return NewQLinkedListDataFromPointer(cthis)
}

//  body block end
