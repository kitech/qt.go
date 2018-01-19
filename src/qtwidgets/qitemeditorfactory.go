//  header block begin
// /usr/include/qt/QtWidgets/qitemeditorfactory.h
// #include <qitemeditorfactory.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QItemEditorFactory struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:98
// index:0
// inline
// void QItemEditorFactory()
func NewQItemEditorFactory() *QItemEditorFactory {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QItemEditorFactoryC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QItemEditorFactory{cthis}
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:99
// index:0
// virtual
// void ~QItemEditorFactory()
func DeleteQItemEditorFactory(*QItemEditorFactory) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QItemEditorFactoryD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:101
// index:0
// virtual
// QWidget * createEditor(int, class QWidget *)
func (this *QItemEditorFactory) CreateEditor(userType int, parent unsafe.Pointer) {
	// 0: (, int userType, QWidget * parent), (&userType, parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QItemEditorFactory12createEditorEiP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, &userType, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:102
// index:0
// virtual
// QByteArray valuePropertyName(int)
func (this *QItemEditorFactory) ValuePropertyName(userType int) {
	// 0: (, int userType), (&userType)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QItemEditorFactory17valuePropertyNameEi", ffiqt.FFI_TYPE_VOID, this.cthis, &userType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:104
// index:0
// void registerEditor(int, class QItemEditorCreatorBase *)
func (this *QItemEditorFactory) RegisterEditor(userType int, creator unsafe.Pointer) {
	// 0: (, int userType, QItemEditorCreatorBase * creator), (&userType, creator)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase", ffiqt.FFI_TYPE_VOID, this.cthis, &userType, creator)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:106
// index:0
// static
// const QItemEditorFactory * defaultFactory()
func (this *QItemEditorFactory) DefaultFactory() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QItemEditorFactory14defaultFactoryEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QItemEditorFactory_DefaultFactory() {
	// 0: (), ()
	var nilthis *QItemEditorFactory
	nilthis.DefaultFactory()
}

// /usr/include/qt/QtWidgets/qitemeditorfactory.h:107
// index:0
// static
// void setDefaultFactory(class QItemEditorFactory *)
func (this *QItemEditorFactory) SetDefaultFactory(factory unsafe.Pointer) {
	// 0: (QItemEditorFactory * factory), (factory)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QItemEditorFactory17setDefaultFactoryEPS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QItemEditorFactory_SetDefaultFactory(factory unsafe.Pointer) {
	// 0: (QItemEditorFactory * factory), (factory)
	var nilthis *QItemEditorFactory
	nilthis.SetDefaultFactory(factory)
}

//  body block end
