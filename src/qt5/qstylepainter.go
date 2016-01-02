package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qstylepainter.h
// dst-file: /src/widgets/qstylepainter.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QStylePainter::QStylePainter(QWidget * w);
extern void* dector_ZN13QStylePainterC1EP7QWidget(void* arg0);
extern void demth_ZN13QStylePainterC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QStylePainter::QStylePainter(QPaintDevice * pd, QWidget * w);
extern void* dector_ZN13QStylePainterC1EP12QPaintDeviceP7QWidget(void* arg0, void* arg1);
extern void demth_ZN13QStylePainterC1EP12QPaintDeviceP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QStylePainter::QStylePainter();
extern void* dector_ZN13QStylePainterC1Ev();
extern void demth_ZN13QStylePainterC1Ev(void* qthis);
  // proto:  bool QStylePainter::begin(QPaintDevice * pd, QWidget * w);
extern void demth_ZN13QStylePainter5beginEP12QPaintDeviceP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QStylePainter::QStylePainter(const QStylePainter & );
extern void* dector_ZN13QStylePainterC1ERKS_(void* arg0);
extern void _ZN13QStylePainterC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QStylePainter::begin(QWidget * w);
extern void demth_ZN13QStylePainter5beginEP7QWidget(void* qthis, void* arg0);
  // proto:  void QStylePainter::drawItemPixmap(const QRect & r, int flags, const QPixmap & pixmap);
extern void demth_ZN13QStylePainter14drawItemPixmapERK5QRectiRK7QPixmap(void* qthis, void* arg0, int arg1, void* arg2);
  // proto:  QStyle * QStylePainter::style();
extern void demth_ZNK13QStylePainter5styleEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QStylePainter)=1
type QStylePainter struct {
  /*qbase*/ QPainter;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QStylePainter::QStylePainter(QWidget * w);
func NewQStylePainter(args ...interface{}) QStylePainter {
  return QStylePainter{}
}

  // proto:  bool QStylePainter::begin(QPaintDevice * pd, QWidget * w);
func (this *QStylePainter) begin(args ...interface{}) () {
  // begin(class QPaintDevice *, class QWidget *)
  // begin(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStylePainter5beginEP12QPaintDeviceP7QWidget
  case 1:
    // invoke: _ZN13QStylePainter5beginEP7QWidget
  default:
    qtrt.ErrorResolve("QStylePainter", "begin", args)
  }

}

  // proto:  void QStylePainter::drawItemPixmap(const QRect & r, int flags, const QPixmap & pixmap);
func (this *QStylePainter) drawItemPixmap(args ...interface{}) () {
  // drawItemPixmap(const class QRect &, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStylePainter14drawItemPixmapERK5QRectiRK7QPixmap
  default:
    qtrt.ErrorResolve("QStylePainter", "drawItemPixmap", args)
  }

}

  // proto:  QStyle * QStylePainter::style();
func (this *QStylePainter) style(args ...interface{}) () {
  // style()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStylePainter5styleEv
  default:
    qtrt.ErrorResolve("QStylePainter", "style", args)
  }

}

// <= body block end

