package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qformlayout.h
// dst-file: /src/widgets/qformlayout.go
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QFormLayout::invalidate();
extern void C_ZN11QFormLayout10invalidateEv(void* qthis); // 4
  // proto:  void QFormLayout::addItem(QLayoutItem * item);
extern void C_ZN11QFormLayout7addItemEP11QLayoutItem(void* qthis, void* arg0); // 4
  // proto:  int QFormLayout::verticalSpacing();
extern void C_ZNK11QFormLayout15verticalSpacingEv(void* qthis); // 4
  // proto:  Qt::Orientations QFormLayout::expandingDirections();
extern void C_ZNK11QFormLayout19expandingDirectionsEv(void* qthis); // 4
  // proto:  void QFormLayout::QFormLayout(QWidget * parent);
extern void C_ZN11QFormLayoutC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  Qt::Alignment QFormLayout::formAlignment();
extern void C_ZNK11QFormLayout13formAlignmentEv(void* qthis); // 4
  // proto:  QWidget * QFormLayout::labelForField(QLayout * field);
extern void C_ZNK11QFormLayout13labelForFieldEP7QLayout(void* qthis, void* arg0); // 4
  // proto:  QWidget * QFormLayout::labelForField(QWidget * field);
extern void C_ZNK11QFormLayout13labelForFieldEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QFormLayout::~QFormLayout();
extern void C_ZN11QFormLayoutD2Ev(void* qthis); // 4
  // proto:  QFormLayout::FieldGrowthPolicy QFormLayout::fieldGrowthPolicy();
extern void C_ZNK11QFormLayout17fieldGrowthPolicyEv(void* qthis); // 4
  // proto:  void QFormLayout::setVerticalSpacing(int spacing);
extern void C_ZN11QFormLayout18setVerticalSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  QLayoutItem * QFormLayout::takeAt(int index);
extern void C_ZN11QFormLayout6takeAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QFormLayout::setSpacing(int );
extern void C_ZN11QFormLayout10setSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  int QFormLayout::horizontalSpacing();
extern void C_ZNK11QFormLayout17horizontalSpacingEv(void* qthis); // 4
  // proto:  void QFormLayout::setGeometry(const QRect & rect);
extern void C_ZN11QFormLayout11setGeometryERK5QRect(void* qthis, void* arg0); // 4
  // proto:  int QFormLayout::spacing();
extern void C_ZNK11QFormLayout7spacingEv(void* qthis); // 4
  // proto:  int QFormLayout::rowCount();
extern void C_ZNK11QFormLayout8rowCountEv(void* qthis); // 4
  // proto:  QFormLayout::RowWrapPolicy QFormLayout::rowWrapPolicy();
extern void C_ZNK11QFormLayout13rowWrapPolicyEv(void* qthis); // 4
  // proto:  QSize QFormLayout::sizeHint();
extern void C_ZNK11QFormLayout8sizeHintEv(void* qthis); // 4
  // proto:  int QFormLayout::heightForWidth(int width);
extern void C_ZNK11QFormLayout14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  int QFormLayout::count();
extern void C_ZNK11QFormLayout5countEv(void* qthis); // 4
  // proto:  const QMetaObject * QFormLayout::metaObject();
extern void C_ZNK11QFormLayout10metaObjectEv(void* qthis); // 4
  // proto:  bool QFormLayout::hasHeightForWidth();
extern void C_ZNK11QFormLayout17hasHeightForWidthEv(void* qthis); // 4
  // proto:  QLayoutItem * QFormLayout::itemAt(int index);
extern void C_ZNK11QFormLayout6itemAtEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QFormLayout::minimumSize();
extern void C_ZNK11QFormLayout11minimumSizeEv(void* qthis); // 4
  // proto:  void QFormLayout::setHorizontalSpacing(int spacing);
extern void C_ZN11QFormLayout20setHorizontalSpacingEi(void* qthis, int32_t arg0); // 4
  // proto:  void QFormLayout::insertRow(int row, QWidget * label, QWidget * field);
extern void C_ZN11QFormLayout9insertRowEiP7QWidgetS1_(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QFormLayout::insertRow(int row, QLayout * layout);
extern void C_ZN11QFormLayout9insertRowEiP7QLayout(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QFormLayout::insertRow(int row, const QString & labelText, QLayout * field);
extern void C_ZN11QFormLayout9insertRowEiRK7QStringP7QLayout(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QFormLayout::insertRow(int row, QWidget * label, QLayout * field);
extern void C_ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QFormLayout::insertRow(int row, const QString & labelText, QWidget * field);
extern void C_ZN11QFormLayout9insertRowEiRK7QStringP7QWidget(void* qthis, int32_t arg0, void* arg1, void* arg2); // 4
  // proto:  void QFormLayout::insertRow(int row, QWidget * widget);
extern void C_ZN11QFormLayout9insertRowEiP7QWidget(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QFormLayout::addRow(QWidget * label, QLayout * field);
extern void C_ZN11QFormLayout6addRowEP7QWidgetP7QLayout(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QFormLayout::addRow(QLayout * layout);
extern void C_ZN11QFormLayout6addRowEP7QLayout(void* qthis, void* arg0); // 4
  // proto:  void QFormLayout::addRow(QWidget * widget);
extern void C_ZN11QFormLayout6addRowEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QFormLayout::addRow(const QString & labelText, QLayout * field);
extern void C_ZN11QFormLayout6addRowERK7QStringP7QLayout(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QFormLayout::addRow(const QString & labelText, QWidget * field);
extern void C_ZN11QFormLayout6addRowERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QFormLayout::addRow(QWidget * label, QWidget * field);
extern void C_ZN11QFormLayout6addRowEP7QWidgetS1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  Qt::Alignment QFormLayout::labelAlignment();
extern void C_ZNK11QFormLayout14labelAlignmentEv(void* qthis); // 4
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

// class sizeof(QFormLayout)=1
type QFormLayout struct {
  /*qbase*/ QLayout;
  qclsinst unsafe.Pointer /* *C.void */;
}

// invalidate()
func (this *QFormLayout) invalidate(args ...interface{}) () {
  // invalidate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout10invalidateEv
    // invoke: void invalidate()
    C.C_ZN11QFormLayout10invalidateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "invalidate", args)
  }

}

// addItem(class QLayoutItem *)
func (this *QFormLayout) addItem(args ...interface{}) () {
  // addItem(class QLayoutItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayoutItem{}) // "QLayoutItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout7addItemEP11QLayoutItem
    // invoke: void addItem(class QLayoutItem *)
    var arg0 = args[0].(QLayoutItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFormLayout7addItemEP11QLayoutItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "addItem", args)
  }

}

// verticalSpacing()
func (this *QFormLayout) verticalSpacing(args ...interface{}) () {
  // verticalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout15verticalSpacingEv
    // invoke: int verticalSpacing()
    C.C_ZNK11QFormLayout15verticalSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "verticalSpacing", args)
  }

}

// expandingDirections()
func (this *QFormLayout) expandingDirections(args ...interface{}) () {
  // expandingDirections()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout19expandingDirectionsEv
    // invoke: Qt::Orientations expandingDirections()
    C.C_ZNK11QFormLayout19expandingDirectionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "expandingDirections", args)
  }

}

// QFormLayout(class QWidget *)
func NewQFormLayout(args ...interface{}) QFormLayout {
  // QFormLayout(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayoutC1EP7QWidget
    // invoke: void QFormLayout(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QFormLayoutC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "QFormLayout", args)
  }

  return QFormLayout{}
}

// formAlignment()
func (this *QFormLayout) formAlignment(args ...interface{}) () {
  // formAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout13formAlignmentEv
    // invoke: Qt::Alignment formAlignment()
    C.C_ZNK11QFormLayout13formAlignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "formAlignment", args)
  }

}

// labelForField(class QLayout *)
func (this *QFormLayout) labelForField(args ...interface{}) () {
  // labelForField(class QLayout *)
  // labelForField(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout13labelForFieldEP7QLayout
    // invoke: QWidget * labelForField(class QLayout *)
    var arg0 = args[0].(QLayout).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QFormLayout13labelForFieldEP7QLayout(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK11QFormLayout13labelForFieldEP7QWidget
    // invoke: QWidget * labelForField(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK11QFormLayout13labelForFieldEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "labelForField", args)
  }

}

// ~QFormLayout()
func (this *QFormLayout) FreeQFormLayout(args ...interface{}) () {
  // ~QFormLayout()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayoutD0Ev
    // invoke: void ~QFormLayout()
    C.C_ZN11QFormLayoutD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "~QFormLayout", args)
  }

}

// fieldGrowthPolicy()
func (this *QFormLayout) fieldGrowthPolicy(args ...interface{}) () {
  // fieldGrowthPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout17fieldGrowthPolicyEv
    // invoke: QFormLayout::FieldGrowthPolicy fieldGrowthPolicy()
    C.C_ZNK11QFormLayout17fieldGrowthPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "fieldGrowthPolicy", args)
  }

}

// setVerticalSpacing(int)
func (this *QFormLayout) setVerticalSpacing(args ...interface{}) () {
  // setVerticalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout18setVerticalSpacingEi
    // invoke: void setVerticalSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFormLayout18setVerticalSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "setVerticalSpacing", args)
  }

}

// takeAt(int)
func (this *QFormLayout) takeAt(args ...interface{}) () {
  // takeAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout6takeAtEi
    // invoke: QLayoutItem * takeAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFormLayout6takeAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "takeAt", args)
  }

}

// setSpacing(int)
func (this *QFormLayout) setSpacing(args ...interface{}) () {
  // setSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout10setSpacingEi
    // invoke: void setSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFormLayout10setSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "setSpacing", args)
  }

}

// horizontalSpacing()
func (this *QFormLayout) horizontalSpacing(args ...interface{}) () {
  // horizontalSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout17horizontalSpacingEv
    // invoke: int horizontalSpacing()
    C.C_ZNK11QFormLayout17horizontalSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "horizontalSpacing", args)
  }

}

// setGeometry(const class QRect &)
func (this *QFormLayout) setGeometry(args ...interface{}) () {
  // setGeometry(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout11setGeometryERK5QRect
    // invoke: void setGeometry(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFormLayout11setGeometryERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "setGeometry", args)
  }

}

// spacing()
func (this *QFormLayout) spacing(args ...interface{}) () {
  // spacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout7spacingEv
    // invoke: int spacing()
    C.C_ZNK11QFormLayout7spacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "spacing", args)
  }

}

// rowCount()
func (this *QFormLayout) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout8rowCountEv
    // invoke: int rowCount()
    C.C_ZNK11QFormLayout8rowCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "rowCount", args)
  }

}

// rowWrapPolicy()
func (this *QFormLayout) rowWrapPolicy(args ...interface{}) () {
  // rowWrapPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout13rowWrapPolicyEv
    // invoke: QFormLayout::RowWrapPolicy rowWrapPolicy()
    C.C_ZNK11QFormLayout13rowWrapPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "rowWrapPolicy", args)
  }

}

// sizeHint()
func (this *QFormLayout) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout8sizeHintEv
    // invoke: QSize sizeHint()
    C.C_ZNK11QFormLayout8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "sizeHint", args)
  }

}

// heightForWidth(int)
func (this *QFormLayout) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QFormLayout14heightForWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "heightForWidth", args)
  }

}

// count()
func (this *QFormLayout) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout5countEv
    // invoke: int count()
    C.C_ZNK11QFormLayout5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "count", args)
  }

}

// metaObject()
func (this *QFormLayout) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QFormLayout10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "metaObject", args)
  }

}

// hasHeightForWidth()
func (this *QFormLayout) hasHeightForWidth(args ...interface{}) () {
  // hasHeightForWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout17hasHeightForWidthEv
    // invoke: bool hasHeightForWidth()
    C.C_ZNK11QFormLayout17hasHeightForWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "hasHeightForWidth", args)
  }

}

// itemAt(int)
func (this *QFormLayout) itemAt(args ...interface{}) () {
  // itemAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout6itemAtEi
    // invoke: QLayoutItem * itemAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK11QFormLayout6itemAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "itemAt", args)
  }

}

// minimumSize()
func (this *QFormLayout) minimumSize(args ...interface{}) () {
  // minimumSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout11minimumSizeEv
    // invoke: QSize minimumSize()
    C.C_ZNK11QFormLayout11minimumSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "minimumSize", args)
  }

}

// setHorizontalSpacing(int)
func (this *QFormLayout) setHorizontalSpacing(args ...interface{}) () {
  // setHorizontalSpacing(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout20setHorizontalSpacingEi
    // invoke: void setHorizontalSpacing(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFormLayout20setHorizontalSpacingEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFormLayout", "setHorizontalSpacing", args)
  }

}

// insertRow(int, class QWidget *, class QWidget *)
func (this *QFormLayout) insertRow(args ...interface{}) () {
  // insertRow(int, class QWidget *, class QWidget *)
  // insertRow(int, class QLayout *)
  // insertRow(int, const class QString &, class QLayout *)
  // insertRow(int, class QWidget *, class QLayout *)
  // insertRow(int, const class QString &, class QWidget *)
  // insertRow(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][2] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[3][2] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidgetS1_
    // invoke: void insertRow(int, class QWidget *, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN11QFormLayout9insertRowEiP7QWidgetS1_(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN11QFormLayout9insertRowEiP7QLayout
    // invoke: void insertRow(int, class QLayout *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QLayout).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QFormLayout9insertRowEiP7QLayout(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN11QFormLayout9insertRowEiRK7QStringP7QLayout
    // invoke: void insertRow(int, const class QString &, class QLayout *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QLayout).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN11QFormLayout9insertRowEiRK7QStringP7QLayout(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout
    // invoke: void insertRow(int, class QWidget *, class QLayout *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QLayout).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN11QFormLayout9insertRowEiP7QWidgetP7QLayout(this.qclsinst, arg0, arg1, arg2)
  case 4:
    // invoke: _ZN11QFormLayout9insertRowEiRK7QStringP7QWidget
    // invoke: void insertRow(int, const class QString &, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN11QFormLayout9insertRowEiRK7QStringP7QWidget(this.qclsinst, arg0, arg1, arg2)
  case 5:
    // invoke: _ZN11QFormLayout9insertRowEiP7QWidget
    // invoke: void insertRow(int, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QFormLayout9insertRowEiP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFormLayout", "insertRow", args)
  }

}

// addRow(class QWidget *, class QLayout *)
func (this *QFormLayout) addRow(args ...interface{}) () {
  // addRow(class QWidget *, class QLayout *)
  // addRow(class QLayout *)
  // addRow(class QWidget *)
  // addRow(const class QString &, class QLayout *)
  // addRow(const class QString &, class QWidget *)
  // addRow(class QWidget *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = reflect.TypeOf(QLayout{}) // "QLayout *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[5][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFormLayout6addRowEP7QWidgetP7QLayout
    // invoke: void addRow(class QWidget *, class QLayout *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QLayout).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QFormLayout6addRowEP7QWidgetP7QLayout(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN11QFormLayout6addRowEP7QLayout
    // invoke: void addRow(class QLayout *)
    var arg0 = args[0].(QLayout).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFormLayout6addRowEP7QLayout(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN11QFormLayout6addRowEP7QWidget
    // invoke: void addRow(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFormLayout6addRowEP7QWidget(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN11QFormLayout6addRowERK7QStringP7QLayout
    // invoke: void addRow(const class QString &, class QLayout *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QLayout).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QFormLayout6addRowERK7QStringP7QLayout(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN11QFormLayout6addRowERK7QStringP7QWidget
    // invoke: void addRow(const class QString &, class QWidget *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QFormLayout6addRowERK7QStringP7QWidget(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN11QFormLayout6addRowEP7QWidgetS1_
    // invoke: void addRow(class QWidget *, class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QFormLayout6addRowEP7QWidgetS1_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFormLayout", "addRow", args)
  }

}

// labelAlignment()
func (this *QFormLayout) labelAlignment(args ...interface{}) () {
  // labelAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFormLayout14labelAlignmentEv
    // invoke: Qt::Alignment labelAlignment()
    C.C_ZNK11QFormLayout14labelAlignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFormLayout", "labelAlignment", args)
  }

}

// <= body block end

