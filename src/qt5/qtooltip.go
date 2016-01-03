package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qtooltip.h
// dst-file: /src/widgets/qtooltip.go
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
  // proto: static void QToolTip::setFont(const QFont & );
extern void _ZN8QToolTip7setFontERK5QFont(void* arg0);
  // proto: static QPalette QToolTip::palette();
extern void _ZN8QToolTip7paletteEv();
  // proto: static void QToolTip::hideText();
extern void demth_ZN8QToolTip8hideTextEv();
  // proto: static void QToolTip::showText(const QPoint & pos, const QString & text, QWidget * w, const QRect & rect);
extern void _ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRect(void* arg0, void* arg1, void* arg2, void* arg3);
  // proto:  void QToolTip::QToolTip();
extern void* dector_ZN8QToolTipC1Ev();
extern void _ZN8QToolTipC1Ev(void* qthis);
  // proto: static void QToolTip::showText(const QPoint & pos, const QString & text, QWidget * w, const QRect & rect, int msecShowTime);
extern void _ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidgetRK5QRecti(void* arg0, void* arg1, void* arg2, void* arg3, int32_t arg4);
  // proto: static QString QToolTip::text();
extern void _ZN8QToolTip4textEv();
  // proto: static QFont QToolTip::font();
extern void _ZN8QToolTip4fontEv();
  // proto: static void QToolTip::setPalette(const QPalette & );
extern void _ZN8QToolTip10setPaletteERK8QPalette(void* arg0);
  // proto: static void QToolTip::showText(const QPoint & pos, const QString & text, QWidget * w);
extern void _ZN8QToolTip8showTextERK6QPointRK7QStringP7QWidget(void* arg0, void* arg1, void* arg2);
  // proto: static bool QToolTip::isVisible();
extern void _ZN8QToolTip9isVisibleEv();
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

// class sizeof(QToolTip)=1
type QToolTip struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto: static void QToolTip::setFont(const QFont & );
func (this *QToolTip) setFont_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolTip", "setFont", args)
  }

}

  // proto: static QPalette QToolTip::palette();
func (this *QToolTip) palette_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolTip", "palette", args)
  }

}

  // proto: static void QToolTip::hideText();
func (this *QToolTip) hideText_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolTip", "hideText", args)
  }

}

  // proto: static void QToolTip::showText(const QPoint & pos, const QString & text, QWidget * w, const QRect & rect);
func (this *QToolTip) showText_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolTip", "showText", args)
  }

}

  // proto:  void QToolTip::QToolTip();
func NewQToolTip(args ...interface{}) QToolTip {
  return QToolTip{}
}

  // proto: static QString QToolTip::text();
func (this *QToolTip) text_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolTip", "text", args)
  }

}

  // proto: static QFont QToolTip::font();
func (this *QToolTip) font_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolTip", "font", args)
  }

}

  // proto: static void QToolTip::setPalette(const QPalette & );
func (this *QToolTip) setPalette_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolTip", "setPalette", args)
  }

}

  // proto: static bool QToolTip::isVisible();
func (this *QToolTip) isVisible_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QToolTip", "isVisible", args)
  }

}

// <= body block end

