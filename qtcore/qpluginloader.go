package qtcore

// /usr/include/qt/QtCore/qpluginloader.h
// #include <qpluginloader.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QPluginLoader struct {
	*QObject
}
type QPluginLoader_ITF interface {
	QObject_ITF
	QPluginLoader_PTR() *QPluginLoader
}

func (ptr *QPluginLoader) QPluginLoader_PTR() *QPluginLoader { return ptr }

func (this *QPluginLoader) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QPluginLoader) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQPluginLoaderFromPointer(cthis unsafe.Pointer) *QPluginLoader {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QPluginLoader{bcthis0}
}
func (*QPluginLoader) NewFromPointer(cthis unsafe.Pointer) *QPluginLoader {
	return NewQPluginLoaderFromPointer(cthis)
}

//  body block end

//  keep block begin

func init_unused_10499() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
