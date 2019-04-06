package qtqml

// /usr/include/qt/QtQml/qqmldebug.h
// #include <qqmldebug.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 47
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QQmlDebuggingEnabler struct {
	*qtrt.CObject
}
type QQmlDebuggingEnabler_ITF interface {
	QQmlDebuggingEnabler_PTR() *QQmlDebuggingEnabler
}

func (ptr *QQmlDebuggingEnabler) QQmlDebuggingEnabler_PTR() *QQmlDebuggingEnabler { return ptr }

func (this *QQmlDebuggingEnabler) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QQmlDebuggingEnabler) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQQmlDebuggingEnablerFromPointer(cthis unsafe.Pointer) *QQmlDebuggingEnabler {
	return &QQmlDebuggingEnabler{&qtrt.CObject{cthis}}
}
func (*QQmlDebuggingEnabler) NewFromPointer(cthis unsafe.Pointer) *QQmlDebuggingEnabler {
	return NewQQmlDebuggingEnablerFromPointer(cthis)
}

/*


 */
type QQmlDebuggingEnabler__StartMode = int

//
const QQmlDebuggingEnabler__DoNotWaitForClient QQmlDebuggingEnabler__StartMode = 0

//
const QQmlDebuggingEnabler__WaitForClient QQmlDebuggingEnabler__StartMode = 1

func (this *QQmlDebuggingEnabler) StartModeItemName(val int) string {
	switch val {
	case QQmlDebuggingEnabler__DoNotWaitForClient: // 0
		return "DoNotWaitForClient"
	case QQmlDebuggingEnabler__WaitForClient: // 1
		return "WaitForClient"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QQmlDebuggingEnabler_StartModeItemName(val int) string {
	var nilthis *QQmlDebuggingEnabler
	return nilthis.StartModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11477() {
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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
