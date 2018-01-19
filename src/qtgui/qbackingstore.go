//  header block begin
// /usr/include/qt/QtGui/qbackingstore.h
// #include <qbackingstore.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 129
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
}

//  ext block end

//  body block begin
type QBackingStore struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qbackingstore.h:62
// index:0
// void QBackingStore(class QWindow *)
func NewQBackingStore(window unsafe.Pointer) *QBackingStore {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStoreC2EP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, window)
	gopp.ErrPrint(err, rv)
	return &QBackingStore{cthis}
}

// /usr/include/qt/QtGui/qbackingstore.h:63
// index:0
// void ~QBackingStore()
func DeleteQBackingStore(*QBackingStore) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStoreD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:65
// index:0
// QWindow * window()
func (this *QBackingStore) Window() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore6windowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:67
// index:0
// QPaintDevice * paintDevice()
func (this *QBackingStore) PaintDevice() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore11paintDeviceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:69
// index:0
// void flush(const class QRegion &, class QWindow *, const class QPoint &)
func (this *QBackingStore) Flush(region unsafe.Pointer, window unsafe.Pointer, offset unsafe.Pointer) {
	// 0: (, const QRegion & region, QWindow * window, const QPoint & offset), (region, window, offset)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, region, window, offset)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:71
// index:0
// void resize(const class QSize &)
func (this *QBackingStore) Resize(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore6resizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:72
// index:0
// QSize size()
func (this *QBackingStore) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:74
// index:0
// bool scroll(const class QRegion &, int, int)
func (this *QBackingStore) Scroll(area unsafe.Pointer, dx int, dy int) {
	// 0: (, const QRegion & area, int dx, int dy), (area, &dx, &dy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore6scrollERK7QRegionii", ffiqt.FFI_TYPE_VOID, this.cthis, area, &dx, &dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:76
// index:0
// void beginPaint(const class QRegion &)
func (this *QBackingStore) BeginPaint(arg0 unsafe.Pointer) {
	// 0: (, const QRegion & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore10beginPaintERK7QRegion", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:77
// index:0
// void endPaint()
func (this *QBackingStore) EndPaint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore8endPaintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:79
// index:0
// void setStaticContents(const class QRegion &)
func (this *QBackingStore) SetStaticContents(region unsafe.Pointer) {
	// 0: (, const QRegion & region), (region)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore17setStaticContentsERK7QRegion", ffiqt.FFI_TYPE_VOID, this.cthis, region)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:80
// index:0
// QRegion staticContents()
func (this *QBackingStore) StaticContents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore14staticContentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:81
// index:0
// bool hasStaticContents()
func (this *QBackingStore) HasStaticContents() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore17hasStaticContentsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:83
// index:0
// QPlatformBackingStore * handle()
func (this *QBackingStore) Handle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore6handleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
