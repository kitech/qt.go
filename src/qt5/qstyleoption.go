package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qstyleoption.h
// dst-file: /src/widgets/qstyleoption.go
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
  // proto:  void QStyleOptionComboBox::QStyleOptionComboBox(const QStyleOptionComboBox & other);
extern void* C_ZN20QStyleOptionComboBoxC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionComboBox::QStyleOptionComboBox();
extern void* C_ZN20QStyleOptionComboBoxC2Ev(); // 3
  // proto:  void QStyleOptionMenuItem::QStyleOptionMenuItem(const QStyleOptionMenuItem & other);
extern void* C_ZN20QStyleOptionMenuItemC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionMenuItem::QStyleOptionMenuItem();
extern void* C_ZN20QStyleOptionMenuItemC2Ev(); // 3
  // proto:  void QStyleHintReturnVariant::~QStyleHintReturnVariant();
extern void C_ZN23QStyleHintReturnVariantD2Ev(void* qthis); // 4
  // proto:  void QStyleHintReturnVariant::QStyleHintReturnVariant();
extern void* C_ZN23QStyleHintReturnVariantC2Ev(); // 3
  // proto:  void QStyleOptionTitleBar::QStyleOptionTitleBar();
extern void* C_ZN20QStyleOptionTitleBarC2Ev(); // 3
  // proto:  void QStyleOptionTitleBar::QStyleOptionTitleBar(const QStyleOptionTitleBar & other);
extern void* C_ZN20QStyleOptionTitleBarC2ERKS_(void* arg0); // 1
  // proto: static qreal QStyleOptionGraphicsItem::levelOfDetailFromTransform(const QTransform & worldTransform);
extern double C_ZN24QStyleOptionGraphicsItem26levelOfDetailFromTransformERK10QTransform(void* arg0); // 4
  // proto:  void QStyleOptionGraphicsItem::QStyleOptionGraphicsItem();
extern void* C_ZN24QStyleOptionGraphicsItemC2Ev(); // 3
  // proto:  void QStyleOptionGraphicsItem::QStyleOptionGraphicsItem(const QStyleOptionGraphicsItem & other);
extern void* C_ZN24QStyleOptionGraphicsItemC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOption::~QStyleOption();
extern void C_ZN12QStyleOptionD2Ev(void* qthis); // 4
  // proto:  void QStyleOption::init(const QWidget * w);
extern void C_ZN12QStyleOption4initEPK7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStyleOption::QStyleOption(const QStyleOption & other);
extern void* C_ZN12QStyleOptionC2ERKS_(void* arg0); // 3
  // proto:  void QStyleOption::QStyleOption(int version, int type);
extern void* C_ZN12QStyleOptionC2Eii(int32_t arg0, int32_t arg1); // 3
  // proto:  void QStyleOption::initFrom(const QWidget * w);
extern void C_ZN12QStyleOption8initFromEPK7QWidget(void* qthis, void* arg0); // 2
  // proto:  void QStyleOptionDockWidget::QStyleOptionDockWidget();
extern void* C_ZN22QStyleOptionDockWidgetC2Ev(); // 3
  // proto:  void QStyleOptionDockWidget::QStyleOptionDockWidget(const QStyleOptionDockWidget & other);
extern void* C_ZN22QStyleOptionDockWidgetC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionProgressBar::QStyleOptionProgressBar(const QStyleOptionProgressBar & other);
extern void* C_ZN23QStyleOptionProgressBarC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionProgressBar::QStyleOptionProgressBar();
extern void* C_ZN23QStyleOptionProgressBarC2Ev(); // 3
  // proto:  void QStyleOptionSlider::QStyleOptionSlider(const QStyleOptionSlider & other);
extern void* C_ZN18QStyleOptionSliderC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionSlider::QStyleOptionSlider();
extern void* C_ZN18QStyleOptionSliderC2Ev(); // 3
  // proto:  void QStyleOptionFrame::QStyleOptionFrame();
extern void* C_ZN17QStyleOptionFrameC2Ev(); // 3
  // proto:  void QStyleOptionFrame::QStyleOptionFrame(const QStyleOptionFrame & other);
extern void* C_ZN17QStyleOptionFrameC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionComplex::QStyleOptionComplex(int version, int type);
extern void* C_ZN19QStyleOptionComplexC2Eii(int32_t arg0, int32_t arg1); // 3
  // proto:  void QStyleOptionComplex::QStyleOptionComplex(const QStyleOptionComplex & other);
extern void* C_ZN19QStyleOptionComplexC2ERKS_(void* arg0); // 1
  // proto:  void QStyleHintReturn::QStyleHintReturn(int version, int type);
extern void* C_ZN16QStyleHintReturnC2Eii(int32_t arg0, int32_t arg1); // 3
  // proto:  void QStyleHintReturn::~QStyleHintReturn();
extern void C_ZN16QStyleHintReturnD2Ev(void* qthis); // 4
  // proto:  void QStyleOptionHeader::QStyleOptionHeader();
extern void* C_ZN18QStyleOptionHeaderC2Ev(); // 3
  // proto:  void QStyleOptionHeader::QStyleOptionHeader(const QStyleOptionHeader & other);
extern void* C_ZN18QStyleOptionHeaderC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionToolBox::QStyleOptionToolBox();
extern void* C_ZN19QStyleOptionToolBoxC2Ev(); // 3
  // proto:  void QStyleOptionToolBox::QStyleOptionToolBox(const QStyleOptionToolBox & other);
extern void* C_ZN19QStyleOptionToolBoxC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionFocusRect::QStyleOptionFocusRect();
extern void* C_ZN21QStyleOptionFocusRectC2Ev(); // 3
  // proto:  void QStyleOptionFocusRect::QStyleOptionFocusRect(const QStyleOptionFocusRect & other);
extern void* C_ZN21QStyleOptionFocusRectC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionGroupBox::QStyleOptionGroupBox(const QStyleOptionGroupBox & other);
extern void* C_ZN20QStyleOptionGroupBoxC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionGroupBox::QStyleOptionGroupBox();
extern void* C_ZN20QStyleOptionGroupBoxC2Ev(); // 3
  // proto:  void QStyleOptionTab::QStyleOptionTab(const QStyleOptionTab & other);
extern void* C_ZN15QStyleOptionTabC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionTab::QStyleOptionTab();
extern void* C_ZN15QStyleOptionTabC2Ev(); // 3
  // proto:  void QStyleOptionTabBarBase::QStyleOptionTabBarBase();
extern void* C_ZN22QStyleOptionTabBarBaseC2Ev(); // 3
  // proto:  void QStyleOptionTabBarBase::QStyleOptionTabBarBase(const QStyleOptionTabBarBase & other);
extern void* C_ZN22QStyleOptionTabBarBaseC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionRubberBand::QStyleOptionRubberBand();
extern void* C_ZN22QStyleOptionRubberBandC2Ev(); // 3
  // proto:  void QStyleOptionRubberBand::QStyleOptionRubberBand(const QStyleOptionRubberBand & other);
extern void* C_ZN22QStyleOptionRubberBandC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionButton::QStyleOptionButton();
extern void* C_ZN18QStyleOptionButtonC2Ev(); // 3
  // proto:  void QStyleOptionButton::QStyleOptionButton(const QStyleOptionButton & other);
extern void* C_ZN18QStyleOptionButtonC2ERKS_(void* arg0); // 1
  // proto:  void QStyleHintReturnMask::QStyleHintReturnMask();
extern void* C_ZN20QStyleHintReturnMaskC2Ev(); // 3
  // proto:  void QStyleHintReturnMask::~QStyleHintReturnMask();
extern void C_ZN20QStyleHintReturnMaskD2Ev(void* qthis); // 4
  // proto:  void QStyleOptionToolButton::QStyleOptionToolButton();
extern void* C_ZN22QStyleOptionToolButtonC2Ev(); // 3
  // proto:  void QStyleOptionToolButton::QStyleOptionToolButton(const QStyleOptionToolButton & other);
extern void* C_ZN22QStyleOptionToolButtonC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionSizeGrip::QStyleOptionSizeGrip();
extern void* C_ZN20QStyleOptionSizeGripC2Ev(); // 3
  // proto:  void QStyleOptionSizeGrip::QStyleOptionSizeGrip(const QStyleOptionSizeGrip & other);
extern void* C_ZN20QStyleOptionSizeGripC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionViewItem::QStyleOptionViewItem(const QStyleOptionViewItem & other);
extern void* C_ZN20QStyleOptionViewItemC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionViewItem::QStyleOptionViewItem();
extern void* C_ZN20QStyleOptionViewItemC2Ev(); // 3
  // proto:  void QStyleOptionSpinBox::QStyleOptionSpinBox();
extern void* C_ZN19QStyleOptionSpinBoxC2Ev(); // 3
  // proto:  void QStyleOptionSpinBox::QStyleOptionSpinBox(const QStyleOptionSpinBox & other);
extern void* C_ZN19QStyleOptionSpinBoxC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionToolBar::QStyleOptionToolBar(const QStyleOptionToolBar & other);
extern void* C_ZN19QStyleOptionToolBarC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionToolBar::QStyleOptionToolBar();
extern void* C_ZN19QStyleOptionToolBarC2Ev(); // 3
  // proto:  void QStyleOptionTabWidgetFrame::QStyleOptionTabWidgetFrame(const QStyleOptionTabWidgetFrame & other);
extern void* C_ZN26QStyleOptionTabWidgetFrameC2ERKS_(void* arg0); // 1
  // proto:  void QStyleOptionTabWidgetFrame::QStyleOptionTabWidgetFrame();
extern void* C_ZN26QStyleOptionTabWidgetFrameC2Ev(); // 3
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

// class sizeof(QStyleOptionComboBox)=1
type QStyleOptionComboBox struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionMenuItem)=1
type QStyleOptionMenuItem struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleHintReturnVariant)=24
type QStyleHintReturnVariant struct {
  /*qbase*/ QStyleHintReturn;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionTitleBar)=1
type QStyleOptionTitleBar struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionGraphicsItem)=1
type QStyleOptionGraphicsItem struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOption)=1
type QStyleOption struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionDockWidget)=1
type QStyleOptionDockWidget struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionProgressBar)=1
type QStyleOptionProgressBar struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionSlider)=1
type QStyleOptionSlider struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionFrame)=1
type QStyleOptionFrame struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionComplex)=1
type QStyleOptionComplex struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleHintReturn)=8
type QStyleHintReturn struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionHeader)=1
type QStyleOptionHeader struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionToolBox)=1
type QStyleOptionToolBox struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionFocusRect)=1
type QStyleOptionFocusRect struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionGroupBox)=1
type QStyleOptionGroupBox struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionTab)=1
type QStyleOptionTab struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionTabBarBase)=1
type QStyleOptionTabBarBase struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionRubberBand)=1
type QStyleOptionRubberBand struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionButton)=1
type QStyleOptionButton struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleHintReturnMask)=16
type QStyleHintReturnMask struct {
  /*qbase*/ QStyleHintReturn;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionToolButton)=1
type QStyleOptionToolButton struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionSizeGrip)=1
type QStyleOptionSizeGrip struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionViewItem)=1
type QStyleOptionViewItem struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionSpinBox)=1
type QStyleOptionSpinBox struct {
  /*qbase*/ QStyleOptionComplex;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionToolBar)=1
type QStyleOptionToolBar struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStyleOptionTabWidgetFrame)=1
type QStyleOptionTabWidgetFrame struct {
  /*qbase*/ QStyleOption;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QStyleOptionComboBox(const class QStyleOptionComboBox &)
func NewQStyleOptionComboBox(args ...interface{}) *QStyleOptionComboBox {
  // QStyleOptionComboBox(const class QStyleOptionComboBox &)
  // QStyleOptionComboBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionComboBox{}) // "const QStyleOptionComboBox &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QStyleOptionComboBoxC1ERKS_
    // invoke: void QStyleOptionComboBox(const class QStyleOptionComboBox &)
    var arg0 = args[0].(QStyleOptionComboBox).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionComboBoxC2ERKS_(arg0)
    return &QStyleOptionComboBox{qclsinst:qthis}
  case 1:
    // invoke: _ZN20QStyleOptionComboBoxC1Ev
    // invoke: void QStyleOptionComboBox()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionComboBoxC2Ev()
    return &QStyleOptionComboBox{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionComboBox", "QStyleOptionComboBox", args)
  }

  return nil // QStyleOptionComboBox{qclsinst:qthis}
}

// QStyleOptionMenuItem(const class QStyleOptionMenuItem &)
func NewQStyleOptionMenuItem(args ...interface{}) *QStyleOptionMenuItem {
  // QStyleOptionMenuItem(const class QStyleOptionMenuItem &)
  // QStyleOptionMenuItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionMenuItem{}) // "const QStyleOptionMenuItem &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QStyleOptionMenuItemC1ERKS_
    // invoke: void QStyleOptionMenuItem(const class QStyleOptionMenuItem &)
    var arg0 = args[0].(QStyleOptionMenuItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionMenuItemC2ERKS_(arg0)
    return &QStyleOptionMenuItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN20QStyleOptionMenuItemC1Ev
    // invoke: void QStyleOptionMenuItem()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionMenuItemC2Ev()
    return &QStyleOptionMenuItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionMenuItem", "QStyleOptionMenuItem", args)
  }

  return nil // QStyleOptionMenuItem{qclsinst:qthis}
}

// ~QStyleHintReturnVariant()
func (this *QStyleHintReturnVariant) Freeqstylehintreturnvariant(args ...interface{}) () {
  // ~QStyleHintReturnVariant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QStyleHintReturnVariantD0Ev
    // invoke: void ~QStyleHintReturnVariant()
    C.C_ZN23QStyleHintReturnVariantD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHintReturnVariant", "~QStyleHintReturnVariant", args)
  }

  return
}

// QStyleHintReturnVariant()
func NewQStyleHintReturnVariant(args ...interface{}) *QStyleHintReturnVariant {
  // QStyleHintReturnVariant()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QStyleHintReturnVariantC1Ev
    // invoke: void QStyleHintReturnVariant()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QStyleHintReturnVariantC2Ev()
    return &QStyleHintReturnVariant{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleHintReturnVariant", "QStyleHintReturnVariant", args)
  }

  return nil // QStyleHintReturnVariant{qclsinst:qthis}
}

// QStyleOptionTitleBar()
func NewQStyleOptionTitleBar(args ...interface{}) *QStyleOptionTitleBar {
  // QStyleOptionTitleBar()
  // QStyleOptionTitleBar(const class QStyleOptionTitleBar &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionTitleBar{}) // "const QStyleOptionTitleBar &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QStyleOptionTitleBarC1Ev
    // invoke: void QStyleOptionTitleBar()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionTitleBarC2Ev()
    return &QStyleOptionTitleBar{qclsinst:qthis}
  case 1:
    // invoke: _ZN20QStyleOptionTitleBarC1ERKS_
    // invoke: void QStyleOptionTitleBar(const class QStyleOptionTitleBar &)
    var arg0 = args[0].(QStyleOptionTitleBar).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionTitleBarC2ERKS_(arg0)
    return &QStyleOptionTitleBar{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionTitleBar", "QStyleOptionTitleBar", args)
  }

  return nil // QStyleOptionTitleBar{qclsinst:qthis}
}

// levelOfDetailFromTransform(const class QTransform &)
func (this *QStyleOptionGraphicsItem) Levelofdetailfromtransform_S(args ...interface{}) (ret interface{}) {
  // levelOfDetailFromTransform(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QStyleOptionGraphicsItem26levelOfDetailFromTransformERK10QTransform
    // invoke: qreal levelOfDetailFromTransform(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN24QStyleOptionGraphicsItem26levelOfDetailFromTransformERK10QTransform(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QStyleOptionGraphicsItem", "levelOfDetailFromTransform", args)
  }

  return
}

// QStyleOptionGraphicsItem()
func NewQStyleOptionGraphicsItem(args ...interface{}) *QStyleOptionGraphicsItem {
  // QStyleOptionGraphicsItem()
  // QStyleOptionGraphicsItem(const class QStyleOptionGraphicsItem &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionGraphicsItem{}) // "const QStyleOptionGraphicsItem &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QStyleOptionGraphicsItemC1Ev
    // invoke: void QStyleOptionGraphicsItem()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QStyleOptionGraphicsItemC2Ev()
    return &QStyleOptionGraphicsItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN24QStyleOptionGraphicsItemC1ERKS_
    // invoke: void QStyleOptionGraphicsItem(const class QStyleOptionGraphicsItem &)
    var arg0 = args[0].(QStyleOptionGraphicsItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN24QStyleOptionGraphicsItemC2ERKS_(arg0)
    return &QStyleOptionGraphicsItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionGraphicsItem", "QStyleOptionGraphicsItem", args)
  }

  return nil // QStyleOptionGraphicsItem{qclsinst:qthis}
}

// ~QStyleOption()
func (this *QStyleOption) Freeqstyleoption(args ...interface{}) () {
  // ~QStyleOption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStyleOptionD0Ev
    // invoke: void ~QStyleOption()
    C.C_ZN12QStyleOptionD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleOption", "~QStyleOption", args)
  }

  return
}

// init(const class QWidget *)
func (this *QStyleOption) Init(args ...interface{}) () {
  // init(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStyleOption4initEPK7QWidget
    // invoke: void init(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QStyleOption4initEPK7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleOption", "init", args)
  }

  return
}

// QStyleOption(const class QStyleOption &)
func NewQStyleOption(args ...interface{}) *QStyleOption {
  // QStyleOption(const class QStyleOption &)
  // QStyleOption(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOption{}) // "const QStyleOption &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStyleOptionC1ERKS_
    // invoke: void QStyleOption(const class QStyleOption &)
    var arg0 = args[0].(QStyleOption).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QStyleOptionC2ERKS_(arg0)
    return &QStyleOption{qclsinst:qthis}
  case 1:
    // invoke: _ZN12QStyleOptionC1Eii
    // invoke: void QStyleOption(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QStyleOptionC2Eii(arg0, arg1)
    return &QStyleOption{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOption", "QStyleOption", args)
  }

  return nil // QStyleOption{qclsinst:qthis}
}

// initFrom(const class QWidget *)
func (this *QStyleOption) Initfrom(args ...interface{}) () {
  // initFrom(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QStyleOption8initFromEPK7QWidget
    // invoke: void initFrom(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QStyleOption8initFromEPK7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleOption", "initFrom", args)
  }

  return
}

// QStyleOptionDockWidget()
func NewQStyleOptionDockWidget(args ...interface{}) *QStyleOptionDockWidget {
  // QStyleOptionDockWidget()
  // QStyleOptionDockWidget(const class QStyleOptionDockWidget &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionDockWidget{}) // "const QStyleOptionDockWidget &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QStyleOptionDockWidgetC1Ev
    // invoke: void QStyleOptionDockWidget()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QStyleOptionDockWidgetC2Ev()
    return &QStyleOptionDockWidget{qclsinst:qthis}
  case 1:
    // invoke: _ZN22QStyleOptionDockWidgetC1ERKS_
    // invoke: void QStyleOptionDockWidget(const class QStyleOptionDockWidget &)
    var arg0 = args[0].(QStyleOptionDockWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QStyleOptionDockWidgetC2ERKS_(arg0)
    return &QStyleOptionDockWidget{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionDockWidget", "QStyleOptionDockWidget", args)
  }

  return nil // QStyleOptionDockWidget{qclsinst:qthis}
}

// QStyleOptionProgressBar(const class QStyleOptionProgressBar &)
func NewQStyleOptionProgressBar(args ...interface{}) *QStyleOptionProgressBar {
  // QStyleOptionProgressBar(const class QStyleOptionProgressBar &)
  // QStyleOptionProgressBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionProgressBar{}) // "const QStyleOptionProgressBar &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN23QStyleOptionProgressBarC1ERKS_
    // invoke: void QStyleOptionProgressBar(const class QStyleOptionProgressBar &)
    var arg0 = args[0].(QStyleOptionProgressBar).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QStyleOptionProgressBarC2ERKS_(arg0)
    return &QStyleOptionProgressBar{qclsinst:qthis}
  case 1:
    // invoke: _ZN23QStyleOptionProgressBarC1Ev
    // invoke: void QStyleOptionProgressBar()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN23QStyleOptionProgressBarC2Ev()
    return &QStyleOptionProgressBar{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionProgressBar", "QStyleOptionProgressBar", args)
  }

  return nil // QStyleOptionProgressBar{qclsinst:qthis}
}

// QStyleOptionSlider(const class QStyleOptionSlider &)
func NewQStyleOptionSlider(args ...interface{}) *QStyleOptionSlider {
  // QStyleOptionSlider(const class QStyleOptionSlider &)
  // QStyleOptionSlider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionSlider{}) // "const QStyleOptionSlider &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStyleOptionSliderC1ERKS_
    // invoke: void QStyleOptionSlider(const class QStyleOptionSlider &)
    var arg0 = args[0].(QStyleOptionSlider).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QStyleOptionSliderC2ERKS_(arg0)
    return &QStyleOptionSlider{qclsinst:qthis}
  case 1:
    // invoke: _ZN18QStyleOptionSliderC1Ev
    // invoke: void QStyleOptionSlider()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QStyleOptionSliderC2Ev()
    return &QStyleOptionSlider{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionSlider", "QStyleOptionSlider", args)
  }

  return nil // QStyleOptionSlider{qclsinst:qthis}
}

// QStyleOptionFrame()
func NewQStyleOptionFrame(args ...interface{}) *QStyleOptionFrame {
  // QStyleOptionFrame()
  // QStyleOptionFrame(const class QStyleOptionFrame &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionFrame{}) // "const QStyleOptionFrame &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QStyleOptionFrameC1Ev
    // invoke: void QStyleOptionFrame()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QStyleOptionFrameC2Ev()
    return &QStyleOptionFrame{qclsinst:qthis}
  case 1:
    // invoke: _ZN17QStyleOptionFrameC1ERKS_
    // invoke: void QStyleOptionFrame(const class QStyleOptionFrame &)
    var arg0 = args[0].(QStyleOptionFrame).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QStyleOptionFrameC2ERKS_(arg0)
    return &QStyleOptionFrame{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionFrame", "QStyleOptionFrame", args)
  }

  return nil // QStyleOptionFrame{qclsinst:qthis}
}

// QStyleOptionComplex(int, int)
func NewQStyleOptionComplex(args ...interface{}) *QStyleOptionComplex {
  // QStyleOptionComplex(int, int)
  // QStyleOptionComplex(const class QStyleOptionComplex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionComplex{}) // "const QStyleOptionComplex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QStyleOptionComplexC1Eii
    // invoke: void QStyleOptionComplex(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyleOptionComplexC2Eii(arg0, arg1)
    return &QStyleOptionComplex{qclsinst:qthis}
  case 1:
    // invoke: _ZN19QStyleOptionComplexC1ERKS_
    // invoke: void QStyleOptionComplex(const class QStyleOptionComplex &)
    var arg0 = args[0].(QStyleOptionComplex).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyleOptionComplexC2ERKS_(arg0)
    return &QStyleOptionComplex{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionComplex", "QStyleOptionComplex", args)
  }

  return nil // QStyleOptionComplex{qclsinst:qthis}
}

// QStyleHintReturn(int, int)
func NewQStyleHintReturn(args ...interface{}) *QStyleHintReturn {
  // QStyleHintReturn(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QStyleHintReturnC1Eii
    // invoke: void QStyleHintReturn(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QStyleHintReturnC2Eii(arg0, arg1)
    return &QStyleHintReturn{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleHintReturn", "QStyleHintReturn", args)
  }

  return nil // QStyleHintReturn{qclsinst:qthis}
}

// ~QStyleHintReturn()
func (this *QStyleHintReturn) Freeqstylehintreturn(args ...interface{}) () {
  // ~QStyleHintReturn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QStyleHintReturnD0Ev
    // invoke: void ~QStyleHintReturn()
    C.C_ZN16QStyleHintReturnD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHintReturn", "~QStyleHintReturn", args)
  }

  return
}

// QStyleOptionHeader()
func NewQStyleOptionHeader(args ...interface{}) *QStyleOptionHeader {
  // QStyleOptionHeader()
  // QStyleOptionHeader(const class QStyleOptionHeader &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionHeader{}) // "const QStyleOptionHeader &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStyleOptionHeaderC1Ev
    // invoke: void QStyleOptionHeader()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QStyleOptionHeaderC2Ev()
    return &QStyleOptionHeader{qclsinst:qthis}
  case 1:
    // invoke: _ZN18QStyleOptionHeaderC1ERKS_
    // invoke: void QStyleOptionHeader(const class QStyleOptionHeader &)
    var arg0 = args[0].(QStyleOptionHeader).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QStyleOptionHeaderC2ERKS_(arg0)
    return &QStyleOptionHeader{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionHeader", "QStyleOptionHeader", args)
  }

  return nil // QStyleOptionHeader{qclsinst:qthis}
}

// QStyleOptionToolBox()
func NewQStyleOptionToolBox(args ...interface{}) *QStyleOptionToolBox {
  // QStyleOptionToolBox()
  // QStyleOptionToolBox(const class QStyleOptionToolBox &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionToolBox{}) // "const QStyleOptionToolBox &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QStyleOptionToolBoxC1Ev
    // invoke: void QStyleOptionToolBox()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyleOptionToolBoxC2Ev()
    return &QStyleOptionToolBox{qclsinst:qthis}
  case 1:
    // invoke: _ZN19QStyleOptionToolBoxC1ERKS_
    // invoke: void QStyleOptionToolBox(const class QStyleOptionToolBox &)
    var arg0 = args[0].(QStyleOptionToolBox).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyleOptionToolBoxC2ERKS_(arg0)
    return &QStyleOptionToolBox{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionToolBox", "QStyleOptionToolBox", args)
  }

  return nil // QStyleOptionToolBox{qclsinst:qthis}
}

// QStyleOptionFocusRect()
func NewQStyleOptionFocusRect(args ...interface{}) *QStyleOptionFocusRect {
  // QStyleOptionFocusRect()
  // QStyleOptionFocusRect(const class QStyleOptionFocusRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionFocusRect{}) // "const QStyleOptionFocusRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QStyleOptionFocusRectC1Ev
    // invoke: void QStyleOptionFocusRect()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QStyleOptionFocusRectC2Ev()
    return &QStyleOptionFocusRect{qclsinst:qthis}
  case 1:
    // invoke: _ZN21QStyleOptionFocusRectC1ERKS_
    // invoke: void QStyleOptionFocusRect(const class QStyleOptionFocusRect &)
    var arg0 = args[0].(QStyleOptionFocusRect).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN21QStyleOptionFocusRectC2ERKS_(arg0)
    return &QStyleOptionFocusRect{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionFocusRect", "QStyleOptionFocusRect", args)
  }

  return nil // QStyleOptionFocusRect{qclsinst:qthis}
}

// QStyleOptionGroupBox(const class QStyleOptionGroupBox &)
func NewQStyleOptionGroupBox(args ...interface{}) *QStyleOptionGroupBox {
  // QStyleOptionGroupBox(const class QStyleOptionGroupBox &)
  // QStyleOptionGroupBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionGroupBox{}) // "const QStyleOptionGroupBox &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QStyleOptionGroupBoxC1ERKS_
    // invoke: void QStyleOptionGroupBox(const class QStyleOptionGroupBox &)
    var arg0 = args[0].(QStyleOptionGroupBox).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionGroupBoxC2ERKS_(arg0)
    return &QStyleOptionGroupBox{qclsinst:qthis}
  case 1:
    // invoke: _ZN20QStyleOptionGroupBoxC1Ev
    // invoke: void QStyleOptionGroupBox()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionGroupBoxC2Ev()
    return &QStyleOptionGroupBox{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionGroupBox", "QStyleOptionGroupBox", args)
  }

  return nil // QStyleOptionGroupBox{qclsinst:qthis}
}

// QStyleOptionTab(const class QStyleOptionTab &)
func NewQStyleOptionTab(args ...interface{}) *QStyleOptionTab {
  // QStyleOptionTab(const class QStyleOptionTab &)
  // QStyleOptionTab()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionTab{}) // "const QStyleOptionTab &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QStyleOptionTabC1ERKS_
    // invoke: void QStyleOptionTab(const class QStyleOptionTab &)
    var arg0 = args[0].(QStyleOptionTab).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QStyleOptionTabC2ERKS_(arg0)
    return &QStyleOptionTab{qclsinst:qthis}
  case 1:
    // invoke: _ZN15QStyleOptionTabC1Ev
    // invoke: void QStyleOptionTab()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QStyleOptionTabC2Ev()
    return &QStyleOptionTab{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionTab", "QStyleOptionTab", args)
  }

  return nil // QStyleOptionTab{qclsinst:qthis}
}

// QStyleOptionTabBarBase()
func NewQStyleOptionTabBarBase(args ...interface{}) *QStyleOptionTabBarBase {
  // QStyleOptionTabBarBase()
  // QStyleOptionTabBarBase(const class QStyleOptionTabBarBase &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionTabBarBase{}) // "const QStyleOptionTabBarBase &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QStyleOptionTabBarBaseC1Ev
    // invoke: void QStyleOptionTabBarBase()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QStyleOptionTabBarBaseC2Ev()
    return &QStyleOptionTabBarBase{qclsinst:qthis}
  case 1:
    // invoke: _ZN22QStyleOptionTabBarBaseC1ERKS_
    // invoke: void QStyleOptionTabBarBase(const class QStyleOptionTabBarBase &)
    var arg0 = args[0].(QStyleOptionTabBarBase).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QStyleOptionTabBarBaseC2ERKS_(arg0)
    return &QStyleOptionTabBarBase{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionTabBarBase", "QStyleOptionTabBarBase", args)
  }

  return nil // QStyleOptionTabBarBase{qclsinst:qthis}
}

// QStyleOptionRubberBand()
func NewQStyleOptionRubberBand(args ...interface{}) *QStyleOptionRubberBand {
  // QStyleOptionRubberBand()
  // QStyleOptionRubberBand(const class QStyleOptionRubberBand &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionRubberBand{}) // "const QStyleOptionRubberBand &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QStyleOptionRubberBandC1Ev
    // invoke: void QStyleOptionRubberBand()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QStyleOptionRubberBandC2Ev()
    return &QStyleOptionRubberBand{qclsinst:qthis}
  case 1:
    // invoke: _ZN22QStyleOptionRubberBandC1ERKS_
    // invoke: void QStyleOptionRubberBand(const class QStyleOptionRubberBand &)
    var arg0 = args[0].(QStyleOptionRubberBand).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QStyleOptionRubberBandC2ERKS_(arg0)
    return &QStyleOptionRubberBand{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionRubberBand", "QStyleOptionRubberBand", args)
  }

  return nil // QStyleOptionRubberBand{qclsinst:qthis}
}

// QStyleOptionButton()
func NewQStyleOptionButton(args ...interface{}) *QStyleOptionButton {
  // QStyleOptionButton()
  // QStyleOptionButton(const class QStyleOptionButton &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionButton{}) // "const QStyleOptionButton &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStyleOptionButtonC1Ev
    // invoke: void QStyleOptionButton()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QStyleOptionButtonC2Ev()
    return &QStyleOptionButton{qclsinst:qthis}
  case 1:
    // invoke: _ZN18QStyleOptionButtonC1ERKS_
    // invoke: void QStyleOptionButton(const class QStyleOptionButton &)
    var arg0 = args[0].(QStyleOptionButton).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QStyleOptionButtonC2ERKS_(arg0)
    return &QStyleOptionButton{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionButton", "QStyleOptionButton", args)
  }

  return nil // QStyleOptionButton{qclsinst:qthis}
}

// QStyleHintReturnMask()
func NewQStyleHintReturnMask(args ...interface{}) *QStyleHintReturnMask {
  // QStyleHintReturnMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QStyleHintReturnMaskC1Ev
    // invoke: void QStyleHintReturnMask()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleHintReturnMaskC2Ev()
    return &QStyleHintReturnMask{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleHintReturnMask", "QStyleHintReturnMask", args)
  }

  return nil // QStyleHintReturnMask{qclsinst:qthis}
}

// ~QStyleHintReturnMask()
func (this *QStyleHintReturnMask) Freeqstylehintreturnmask(args ...interface{}) () {
  // ~QStyleHintReturnMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QStyleHintReturnMaskD0Ev
    // invoke: void ~QStyleHintReturnMask()
    C.C_ZN20QStyleHintReturnMaskD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStyleHintReturnMask", "~QStyleHintReturnMask", args)
  }

  return
}

// QStyleOptionToolButton()
func NewQStyleOptionToolButton(args ...interface{}) *QStyleOptionToolButton {
  // QStyleOptionToolButton()
  // QStyleOptionToolButton(const class QStyleOptionToolButton &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionToolButton{}) // "const QStyleOptionToolButton &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QStyleOptionToolButtonC1Ev
    // invoke: void QStyleOptionToolButton()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QStyleOptionToolButtonC2Ev()
    return &QStyleOptionToolButton{qclsinst:qthis}
  case 1:
    // invoke: _ZN22QStyleOptionToolButtonC1ERKS_
    // invoke: void QStyleOptionToolButton(const class QStyleOptionToolButton &)
    var arg0 = args[0].(QStyleOptionToolButton).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QStyleOptionToolButtonC2ERKS_(arg0)
    return &QStyleOptionToolButton{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionToolButton", "QStyleOptionToolButton", args)
  }

  return nil // QStyleOptionToolButton{qclsinst:qthis}
}

// QStyleOptionSizeGrip()
func NewQStyleOptionSizeGrip(args ...interface{}) *QStyleOptionSizeGrip {
  // QStyleOptionSizeGrip()
  // QStyleOptionSizeGrip(const class QStyleOptionSizeGrip &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionSizeGrip{}) // "const QStyleOptionSizeGrip &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QStyleOptionSizeGripC1Ev
    // invoke: void QStyleOptionSizeGrip()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionSizeGripC2Ev()
    return &QStyleOptionSizeGrip{qclsinst:qthis}
  case 1:
    // invoke: _ZN20QStyleOptionSizeGripC1ERKS_
    // invoke: void QStyleOptionSizeGrip(const class QStyleOptionSizeGrip &)
    var arg0 = args[0].(QStyleOptionSizeGrip).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionSizeGripC2ERKS_(arg0)
    return &QStyleOptionSizeGrip{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionSizeGrip", "QStyleOptionSizeGrip", args)
  }

  return nil // QStyleOptionSizeGrip{qclsinst:qthis}
}

// QStyleOptionViewItem(const class QStyleOptionViewItem &)
func NewQStyleOptionViewItem(args ...interface{}) *QStyleOptionViewItem {
  // QStyleOptionViewItem(const class QStyleOptionViewItem &)
  // QStyleOptionViewItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionViewItem{}) // "const QStyleOptionViewItem &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QStyleOptionViewItemC1ERKS_
    // invoke: void QStyleOptionViewItem(const class QStyleOptionViewItem &)
    var arg0 = args[0].(QStyleOptionViewItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionViewItemC2ERKS_(arg0)
    return &QStyleOptionViewItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN20QStyleOptionViewItemC1Ev
    // invoke: void QStyleOptionViewItem()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN20QStyleOptionViewItemC2Ev()
    return &QStyleOptionViewItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionViewItem", "QStyleOptionViewItem", args)
  }

  return nil // QStyleOptionViewItem{qclsinst:qthis}
}

// QStyleOptionSpinBox()
func NewQStyleOptionSpinBox(args ...interface{}) *QStyleOptionSpinBox {
  // QStyleOptionSpinBox()
  // QStyleOptionSpinBox(const class QStyleOptionSpinBox &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStyleOptionSpinBox{}) // "const QStyleOptionSpinBox &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QStyleOptionSpinBoxC1Ev
    // invoke: void QStyleOptionSpinBox()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyleOptionSpinBoxC2Ev()
    return &QStyleOptionSpinBox{qclsinst:qthis}
  case 1:
    // invoke: _ZN19QStyleOptionSpinBoxC1ERKS_
    // invoke: void QStyleOptionSpinBox(const class QStyleOptionSpinBox &)
    var arg0 = args[0].(QStyleOptionSpinBox).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyleOptionSpinBoxC2ERKS_(arg0)
    return &QStyleOptionSpinBox{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionSpinBox", "QStyleOptionSpinBox", args)
  }

  return nil // QStyleOptionSpinBox{qclsinst:qthis}
}

// QStyleOptionToolBar(const class QStyleOptionToolBar &)
func NewQStyleOptionToolBar(args ...interface{}) *QStyleOptionToolBar {
  // QStyleOptionToolBar(const class QStyleOptionToolBar &)
  // QStyleOptionToolBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionToolBar{}) // "const QStyleOptionToolBar &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QStyleOptionToolBarC1ERKS_
    // invoke: void QStyleOptionToolBar(const class QStyleOptionToolBar &)
    var arg0 = args[0].(QStyleOptionToolBar).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyleOptionToolBarC2ERKS_(arg0)
    return &QStyleOptionToolBar{qclsinst:qthis}
  case 1:
    // invoke: _ZN19QStyleOptionToolBarC1Ev
    // invoke: void QStyleOptionToolBar()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QStyleOptionToolBarC2Ev()
    return &QStyleOptionToolBar{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionToolBar", "QStyleOptionToolBar", args)
  }

  return nil // QStyleOptionToolBar{qclsinst:qthis}
}

// QStyleOptionTabWidgetFrame(const class QStyleOptionTabWidgetFrame &)
func NewQStyleOptionTabWidgetFrame(args ...interface{}) *QStyleOptionTabWidgetFrame {
  // QStyleOptionTabWidgetFrame(const class QStyleOptionTabWidgetFrame &)
  // QStyleOptionTabWidgetFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyleOptionTabWidgetFrame{}) // "const QStyleOptionTabWidgetFrame &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QStyleOptionTabWidgetFrameC1ERKS_
    // invoke: void QStyleOptionTabWidgetFrame(const class QStyleOptionTabWidgetFrame &)
    var arg0 = args[0].(QStyleOptionTabWidgetFrame).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QStyleOptionTabWidgetFrameC2ERKS_(arg0)
    return &QStyleOptionTabWidgetFrame{qclsinst:qthis}
  case 1:
    // invoke: _ZN26QStyleOptionTabWidgetFrameC1Ev
    // invoke: void QStyleOptionTabWidgetFrame()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QStyleOptionTabWidgetFrameC2Ev()
    return &QStyleOptionTabWidgetFrame{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStyleOptionTabWidgetFrame", "QStyleOptionTabWidgetFrame", args)
  }

  return nil // QStyleOptionTabWidgetFrame{qclsinst:qthis}
}

// <= body block end

