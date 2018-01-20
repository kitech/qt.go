//  header block begin
// /usr/include/qt/QtGui/qbackingstore.h
// #include <qbackingstore.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 147
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
	*qtrt.CObject
}

func (this *QBackingStore) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQBackingStoreFromPointer(cthis unsafe.Pointer) *QBackingStore {
	return &QBackingStore{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qbackingstore.h:62
// index:0
// Public
// void QBackingStore(class QWindow *)
func NewQBackingStore(window unsafe.Pointer) *QBackingStore {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStoreC2EP7QWindow", ffiqt.FFI_TYPE_VOID, cthis, window)
	gopp.ErrPrint(err, rv)
	gothis := NewQBackingStoreFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qbackingstore.h:63
// index:0
// Public
// void ~QBackingStore()
func DeleteQBackingStore(*QBackingStore) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStoreD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:65
// index:0
// Public
// QWindow * window()
func (this *QBackingStore) Window() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbackingstore.h:67
// index:0
// Public
// QPaintDevice * paintDevice()
func (this *QBackingStore) PaintDevice() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore11paintDeviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbackingstore.h:69
// index:0
// Public
// void flush(const class QRegion &, class QWindow *, const class QPoint &)
func (this *QBackingStore) Flush(region *QRegion, window unsafe.Pointer, offset *qtcore.QPoint) {
	var convArg0 = region.GetCthis()
	var convArg2 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore5flushERK7QRegionP7QWindowRK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, window, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:71
// index:0
// Public
// void resize(const class QSize &)
func (this *QBackingStore) Resize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore6resizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:72
// index:0
// Public
// QSize size()
func (this *QBackingStore) Size() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore4sizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbackingstore.h:74
// index:0
// Public
// bool scroll(const class QRegion &, int, int)
func (this *QBackingStore) Scroll(area *QRegion, dx int, dy int) interface{} {
	var convArg0 = area.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore6scrollERK7QRegionii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &dx, &dy)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbackingstore.h:76
// index:0
// Public
// void beginPaint(const class QRegion &)
func (this *QBackingStore) BeginPaint(arg0 *QRegion) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore10beginPaintERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:77
// index:0
// Public
// void endPaint()
func (this *QBackingStore) EndPaint() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore8endPaintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:79
// index:0
// Public
// void setStaticContents(const class QRegion &)
func (this *QBackingStore) SetStaticContents(region *QRegion) {
	var convArg0 = region.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QBackingStore17setStaticContentsERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbackingstore.h:80
// index:0
// Public
// QRegion staticContents()
func (this *QBackingStore) StaticContents() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore14staticContentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbackingstore.h:81
// index:0
// Public
// bool hasStaticContents()
func (this *QBackingStore) HasStaticContents() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore17hasStaticContentsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qbackingstore.h:83
// index:0
// Public
// QPlatformBackingStore * handle()
func (this *QBackingStore) Handle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QBackingStore6handleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
