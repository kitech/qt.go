package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qgridlayout.h
// dst-file: /src/widgets/qgridlayout.go
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
  // proto:  void QGridLayout::setRowMinimumHeight(int row, int minSize);
extern void _ZN11QGridLayout19setRowMinimumHeightEii(void* qthis, int arg0, int arg1);
  // proto:  QLayoutItem * QGridLayout::takeAt(int index);
extern void _ZN11QGridLayout6takeAtEi(void* qthis, int arg0);
  // proto:  void QGridLayout::getItemPosition(int idx, int * row, int * column, int * rowSpan, int * columnSpan);
extern void _ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_(void* qthis, int arg0, int* arg1, int* arg2, int* arg3, int* arg4);
  // proto:  int QGridLayout::minimumHeightForWidth(int );
extern void _ZNK11QGridLayout21minimumHeightForWidthEi(void* qthis, int arg0);
  // proto:  int QGridLayout::rowMinimumHeight(int row);
extern void _ZNK11QGridLayout16rowMinimumHeightEi(void* qthis, int arg0);
  // proto:  void QGridLayout::invalidate();
extern void _ZN11QGridLayout10invalidateEv(void* qthis);
  // proto:  int QGridLayout::count();
extern void _ZNK11QGridLayout5countEv(void* qthis);
  // proto:  void QGridLayout::setColumnStretch(int column, int stretch);
extern void _ZN11QGridLayout16setColumnStretchEii(void* qthis, int arg0, int arg1);
  // proto:  int QGridLayout::spacing();
extern void _ZNK11QGridLayout7spacingEv(void* qthis);
  // proto:  int QGridLayout::rowStretch(int row);
extern void _ZNK11QGridLayout10rowStretchEi(void* qthis, int arg0);
  // proto:  QSize QGridLayout::sizeHint();
extern void _ZNK11QGridLayout8sizeHintEv(void* qthis);
  // proto:  int QGridLayout::rowCount();
extern void _ZNK11QGridLayout8rowCountEv(void* qthis);
  // proto:  void QGridLayout::setGeometry(const QRect & );
extern void _ZN11QGridLayout11setGeometryERK5QRect(void* qthis, void* arg0);
  // proto:  void QGridLayout::setVerticalSpacing(int spacing);
extern void _ZN11QGridLayout18setVerticalSpacingEi(void* qthis, int arg0);
  // proto:  void QGridLayout::setHorizontalSpacing(int spacing);
extern void _ZN11QGridLayout20setHorizontalSpacingEi(void* qthis, int arg0);
  // proto:  int QGridLayout::columnStretch(int column);
extern void _ZNK11QGridLayout13columnStretchEi(void* qthis, int arg0);
  // proto:  void QGridLayout::QGridLayout(const QGridLayout & );
extern void* dector_ZN11QGridLayoutC1ERKS_(void* arg0);
extern void _ZN11QGridLayoutC1ERKS_(void* qthis, void* arg0);
  // proto:  int QGridLayout::columnCount();
extern void _ZNK11QGridLayout11columnCountEv(void* qthis);
  // proto:  int QGridLayout::columnMinimumWidth(int column);
extern void _ZNK11QGridLayout18columnMinimumWidthEi(void* qthis, int arg0);
  // proto:  QSize QGridLayout::minimumSize();
extern void _ZNK11QGridLayout11minimumSizeEv(void* qthis);
  // proto:  bool QGridLayout::hasHeightForWidth();
extern void _ZNK11QGridLayout17hasHeightForWidthEv(void* qthis);
  // proto:  QRect QGridLayout::cellRect(int row, int column);
extern void _ZNK11QGridLayout8cellRectEii(void* qthis, int arg0, int arg1);
  // proto:  void QGridLayout::setRowStretch(int row, int stretch);
extern void _ZN11QGridLayout13setRowStretchEii(void* qthis, int arg0, int arg1);
  // proto:  QLayoutItem * QGridLayout::itemAtPosition(int row, int column);
extern void _ZNK11QGridLayout14itemAtPositionEii(void* qthis, int arg0, int arg1);
  // proto:  const QMetaObject * QGridLayout::metaObject();
extern void _ZNK11QGridLayout10metaObjectEv(void* qthis);
  // proto:  int QGridLayout::verticalSpacing();
extern void _ZNK11QGridLayout15verticalSpacingEv(void* qthis);
  // proto:  void QGridLayout::QGridLayout(QWidget * parent);
extern void* dector_ZN11QGridLayoutC1EP7QWidget(void* arg0);
extern void _ZN11QGridLayoutC1EP7QWidget(void* qthis, void* arg0);
  // proto:  int QGridLayout::horizontalSpacing();
extern void _ZNK11QGridLayout17horizontalSpacingEv(void* qthis);
  // proto:  void QGridLayout::setColumnMinimumWidth(int column, int minSize);
extern void _ZN11QGridLayout21setColumnMinimumWidthEii(void* qthis, int arg0, int arg1);
  // proto:  void QGridLayout::QGridLayout();
extern void* dector_ZN11QGridLayoutC1Ev();
extern void _ZN11QGridLayoutC1Ev(void* qthis);
  // proto:  int QGridLayout::heightForWidth(int );
extern void _ZNK11QGridLayout14heightForWidthEi(void* qthis, int arg0);
  // proto:  void QGridLayout::~QGridLayout();
extern void _ZN11QGridLayoutD0Ev(void* qthis);
  // proto:  void QGridLayout::setSpacing(int spacing);
extern void _ZN11QGridLayout10setSpacingEi(void* qthis, int arg0);
  // proto:  void QGridLayout::addWidget(QWidget * w);
extern void demth_ZN11QGridLayout9addWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  QLayoutItem * QGridLayout::itemAt(int index);
extern void _ZNK11QGridLayout6itemAtEi(void* qthis, int arg0);
  // proto:  QSize QGridLayout::maximumSize();
extern void _ZNK11QGridLayout11maximumSizeEv(void* qthis);
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

// class sizeof(QGridLayout)=1
type QGridLayout struct {
  /*qbase*/ QLayout;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QGridLayout::setRowMinimumHeight(int row, int minSize);
func (this *QGridLayout) setRowMinimumHeight(args ...interface{}) () {
  // setRowMinimumHeight(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout19setRowMinimumHeightEii
  default:
    qtrt.ErrorResolve("QGridLayout", "setRowMinimumHeight", args)
  }

}

  // proto:  QLayoutItem * QGridLayout::takeAt(int index);
func (this *QGridLayout) takeAt(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout6takeAtEi
  default:
    qtrt.ErrorResolve("QGridLayout", "takeAt", args)
  }

}

  // proto:  void QGridLayout::getItemPosition(int idx, int * row, int * column, int * rowSpan, int * columnSpan);
func (this *QGridLayout) getItemPosition(args ...interface{}) () {
  // getItemPosition(int, int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"
  vtys[0][4] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout15getItemPositionEiPiS0_S0_S0_
  default:
    qtrt.ErrorResolve("QGridLayout", "getItemPosition", args)
  }

}

  // proto:  int QGridLayout::minimumHeightForWidth(int );
func (this *QGridLayout) minimumHeightForWidth(args ...interface{}) () {
  // minimumHeightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout21minimumHeightForWidthEi
  default:
    qtrt.ErrorResolve("QGridLayout", "minimumHeightForWidth", args)
  }

}

  // proto:  int QGridLayout::rowMinimumHeight(int row);
func (this *QGridLayout) rowMinimumHeight(args ...interface{}) () {
  // rowMinimumHeight(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout16rowMinimumHeightEi
  default:
    qtrt.ErrorResolve("QGridLayout", "rowMinimumHeight", args)
  }

}

  // proto:  void QGridLayout::invalidate();
func (this *QGridLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout10invalidateEv
  default:
    qtrt.ErrorResolve("QGridLayout", "invalidate", args)
  }

}

  // proto:  int QGridLayout::count();
func (this *QGridLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout5countEv
  default:
    qtrt.ErrorResolve("QGridLayout", "count", args)
  }

}

  // proto:  void QGridLayout::setColumnStretch(int column, int stretch);
func (this *QGridLayout) setColumnStretch(args ...interface{}) () {
  // setColumnStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout16setColumnStretchEii
  default:
    qtrt.ErrorResolve("QGridLayout", "setColumnStretch", args)
  }

}

  // proto:  int QGridLayout::spacing();
func (this *QGridLayout) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout7spacingEv
  default:
    qtrt.ErrorResolve("QGridLayout", "spacing", args)
  }

}

  // proto:  int QGridLayout::rowStretch(int row);
func (this *QGridLayout) rowStretch(args ...interface{}) () {
  // rowStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout10rowStretchEi
  default:
    qtrt.ErrorResolve("QGridLayout", "rowStretch", args)
  }

}

  // proto:  QSize QGridLayout::sizeHint();
func (this *QGridLayout) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8sizeHintEv
  default:
    qtrt.ErrorResolve("QGridLayout", "sizeHint", args)
  }

}

  // proto:  int QGridLayout::rowCount();
func (this *QGridLayout) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8rowCountEv
  default:
    qtrt.ErrorResolve("QGridLayout", "rowCount", args)
  }

}

  // proto:  void QGridLayout::setGeometry(const QRect & );
func (this *QGridLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout11setGeometryERK5QRect
  default:
    qtrt.ErrorResolve("QGridLayout", "setGeometry", args)
  }

}

  // proto:  void QGridLayout::setVerticalSpacing(int spacing);
func (this *QGridLayout) setVerticalSpacing(args ...interface{}) () {
  // setVerticalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout18setVerticalSpacingEi
  default:
    qtrt.ErrorResolve("QGridLayout", "setVerticalSpacing", args)
  }

}

  // proto:  void QGridLayout::setHorizontalSpacing(int spacing);
func (this *QGridLayout) setHorizontalSpacing(args ...interface{}) () {
  // setHorizontalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout20setHorizontalSpacingEi
  default:
    qtrt.ErrorResolve("QGridLayout", "setHorizontalSpacing", args)
  }

}

  // proto:  int QGridLayout::columnStretch(int column);
func (this *QGridLayout) columnStretch(args ...interface{}) () {
  // columnStretch(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout13columnStretchEi
  default:
    qtrt.ErrorResolve("QGridLayout", "columnStretch", args)
  }

}

  // proto:  void QGridLayout::QGridLayout(const QGridLayout & );
func NewQGridLayout(args ...interface{}) QGridLayout {
  return QGridLayout{}
}

  // proto:  int QGridLayout::columnCount();
func (this *QGridLayout) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11columnCountEv
  default:
    qtrt.ErrorResolve("QGridLayout", "columnCount", args)
  }

}

  // proto:  int QGridLayout::columnMinimumWidth(int column);
func (this *QGridLayout) columnMinimumWidth(args ...interface{}) () {
  // columnMinimumWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout18columnMinimumWidthEi
  default:
    qtrt.ErrorResolve("QGridLayout", "columnMinimumWidth", args)
  }

}

  // proto:  QSize QGridLayout::minimumSize();
func (this *QGridLayout) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11minimumSizeEv
  default:
    qtrt.ErrorResolve("QGridLayout", "minimumSize", args)
  }

}

  // proto:  bool QGridLayout::hasHeightForWidth();
func (this *QGridLayout) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout17hasHeightForWidthEv
  default:
    qtrt.ErrorResolve("QGridLayout", "hasHeightForWidth", args)
  }

}

  // proto:  QRect QGridLayout::cellRect(int row, int column);
func (this *QGridLayout) cellRect(args ...interface{}) () {
  // cellRect(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout8cellRectEii
  default:
    qtrt.ErrorResolve("QGridLayout", "cellRect", args)
  }

}

  // proto:  void QGridLayout::setRowStretch(int row, int stretch);
func (this *QGridLayout) setRowStretch(args ...interface{}) () {
  // setRowStretch(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout13setRowStretchEii
  default:
    qtrt.ErrorResolve("QGridLayout", "setRowStretch", args)
  }

}

  // proto:  QLayoutItem * QGridLayout::itemAtPosition(int row, int column);
func (this *QGridLayout) itemAtPosition(args ...interface{}) () {
  // itemAtPosition(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout14itemAtPositionEii
  default:
    qtrt.ErrorResolve("QGridLayout", "itemAtPosition", args)
  }

}

  // proto:  const QMetaObject * QGridLayout::metaObject();
func (this *QGridLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout10metaObjectEv
  default:
    qtrt.ErrorResolve("QGridLayout", "metaObject", args)
  }

}

  // proto:  int QGridLayout::verticalSpacing();
func (this *QGridLayout) verticalSpacing(args ...interface{}) () {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout15verticalSpacingEv
  default:
    qtrt.ErrorResolve("QGridLayout", "verticalSpacing", args)
  }

}

  // proto:  int QGridLayout::horizontalSpacing();
func (this *QGridLayout) horizontalSpacing(args ...interface{}) () {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout17horizontalSpacingEv
  default:
    qtrt.ErrorResolve("QGridLayout", "horizontalSpacing", args)
  }

}

  // proto:  void QGridLayout::setColumnMinimumWidth(int column, int minSize);
func (this *QGridLayout) setColumnMinimumWidth(args ...interface{}) () {
  // setColumnMinimumWidth(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout21setColumnMinimumWidthEii
  default:
    qtrt.ErrorResolve("QGridLayout", "setColumnMinimumWidth", args)
  }

}

  // proto:  int QGridLayout::heightForWidth(int );
func (this *QGridLayout) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout14heightForWidthEi
  default:
    qtrt.ErrorResolve("QGridLayout", "heightForWidth", args)
  }

}

  // proto:  void QGridLayout::~QGridLayout();
func (this *QGridLayout) FreeQGridLayout(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QGridLayout", "~QGridLayout", args)
  }

}

  // proto:  void QGridLayout::setSpacing(int spacing);
func (this *QGridLayout) setSpacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout10setSpacingEi
  default:
    qtrt.ErrorResolve("QGridLayout", "setSpacing", args)
  }

}

  // proto:  void QGridLayout::addWidget(QWidget * w);
func (this *QGridLayout) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *, int, int, Qt::Alignment)
  // addWidget(class QWidget *, int, int, int, int, Qt::Alignment)
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int64Ty(false) // "Qt::Alignment"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(false) // "int"
  vtys[1][5] = qtrt.Int64Ty(false) // "Qt::Alignment"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QGridLayout9addWidgetEP7QWidgetii6QFlagsIN2Qt13AlignmentFlagEE
  case 1:
    // invoke: _ZN11QGridLayout9addWidgetEP7QWidgetiiii6QFlagsIN2Qt13AlignmentFlagEE
  case 2:
    // invoke: _ZN11QGridLayout9addWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QGridLayout", "addWidget", args)
  }

}

  // proto:  QLayoutItem * QGridLayout::itemAt(int index);
func (this *QGridLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout6itemAtEi
  default:
    qtrt.ErrorResolve("QGridLayout", "itemAt", args)
  }

}

  // proto:  QSize QGridLayout::maximumSize();
func (this *QGridLayout) maximumSize(args ...interface{}) () {
  // maximumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QGridLayout11maximumSizeEv
  default:
    qtrt.ErrorResolve("QGridLayout", "maximumSize", args)
  }

}

// <= body block end

