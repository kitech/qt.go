package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QStyleOptionComboBox::QStyleOptionComboBox(const QStyleOptionComboBox & other);
extern void* dector_ZN20QStyleOptionComboBoxC1ERKS_(void* arg0);
extern void _ZN20QStyleOptionComboBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionComboBox::QStyleOptionComboBox();
extern void* dector_ZN20QStyleOptionComboBoxC1Ev();
extern void _ZN20QStyleOptionComboBoxC1Ev(void* qthis);
  // proto:  void QStyleOptionComboBox::QStyleOptionComboBox(int version);
extern void* dector_ZN20QStyleOptionComboBoxC1Ei(int32_t arg0);
extern void _ZN20QStyleOptionComboBoxC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionMenuItem::QStyleOptionMenuItem(const QStyleOptionMenuItem & other);
extern void* dector_ZN20QStyleOptionMenuItemC1ERKS_(void* arg0);
extern void _ZN20QStyleOptionMenuItemC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionMenuItem::QStyleOptionMenuItem();
extern void* dector_ZN20QStyleOptionMenuItemC1Ev();
extern void _ZN20QStyleOptionMenuItemC1Ev(void* qthis);
  // proto:  void QStyleOptionMenuItem::QStyleOptionMenuItem(int version);
extern void* dector_ZN20QStyleOptionMenuItemC1Ei(int32_t arg0);
extern void _ZN20QStyleOptionMenuItemC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleHintReturnVariant::~QStyleHintReturnVariant();
extern void _ZN23QStyleHintReturnVariantD0Ev(void* qthis);
  // proto:  void QStyleHintReturnVariant::QStyleHintReturnVariant();
extern void* dector_ZN23QStyleHintReturnVariantC1Ev();
extern void _ZN23QStyleHintReturnVariantC1Ev(void* qthis);
  // proto:  void QStyleOptionTitleBar::QStyleOptionTitleBar(int version);
extern void* dector_ZN20QStyleOptionTitleBarC1Ei(int32_t arg0);
extern void _ZN20QStyleOptionTitleBarC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionTitleBar::QStyleOptionTitleBar(const QStyleOptionTitleBar & other);
extern void* dector_ZN20QStyleOptionTitleBarC1ERKS_(void* arg0);
extern void _ZN20QStyleOptionTitleBarC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionTitleBar::QStyleOptionTitleBar();
extern void* dector_ZN20QStyleOptionTitleBarC1Ev();
extern void _ZN20QStyleOptionTitleBarC1Ev(void* qthis);
  // proto:  void QStyleOptionGraphicsItem::QStyleOptionGraphicsItem();
extern void* dector_ZN24QStyleOptionGraphicsItemC1Ev();
extern void _ZN24QStyleOptionGraphicsItemC1Ev(void* qthis);
  // proto: static qreal QStyleOptionGraphicsItem::levelOfDetailFromTransform(const QTransform & worldTransform);
extern void _ZN24QStyleOptionGraphicsItem26levelOfDetailFromTransformERK10QTransform(void* arg0);
  // proto:  void QStyleOptionGraphicsItem::QStyleOptionGraphicsItem(const QStyleOptionGraphicsItem & other);
extern void* dector_ZN24QStyleOptionGraphicsItemC1ERKS_(void* arg0);
extern void _ZN24QStyleOptionGraphicsItemC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionGraphicsItem::QStyleOptionGraphicsItem(int version);
extern void* dector_ZN24QStyleOptionGraphicsItemC1Ei(int32_t arg0);
extern void _ZN24QStyleOptionGraphicsItemC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOption::~QStyleOption();
extern void _ZN12QStyleOptionD0Ev(void* qthis);
  // proto:  void QStyleOption::init(const QWidget * w);
extern void _ZN12QStyleOption4initEPK7QWidget(void* qthis, void* arg0);
  // proto:  void QStyleOption::QStyleOption(const QStyleOption & other);
extern void* dector_ZN12QStyleOptionC1ERKS_(void* arg0);
extern void _ZN12QStyleOptionC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOption::QStyleOption(int version, int type);
extern void* dector_ZN12QStyleOptionC1Eii(int32_t arg0, int32_t arg1);
extern void _ZN12QStyleOptionC1Eii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QStyleOption::initFrom(const QWidget * w);
extern void demth_ZN12QStyleOption8initFromEPK7QWidget(void* qthis, void* arg0);
  // proto:  void QStyleOptionDockWidget::QStyleOptionDockWidget();
extern void* dector_ZN22QStyleOptionDockWidgetC1Ev();
extern void _ZN22QStyleOptionDockWidgetC1Ev(void* qthis);
  // proto:  void QStyleOptionDockWidget::QStyleOptionDockWidget(int version);
extern void* dector_ZN22QStyleOptionDockWidgetC1Ei(int32_t arg0);
extern void _ZN22QStyleOptionDockWidgetC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionDockWidget::QStyleOptionDockWidget(const QStyleOptionDockWidget & other);
extern void* dector_ZN22QStyleOptionDockWidgetC1ERKS_(void* arg0);
extern void _ZN22QStyleOptionDockWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionProgressBar::QStyleOptionProgressBar(const QStyleOptionProgressBar & other);
extern void* dector_ZN23QStyleOptionProgressBarC1ERKS_(void* arg0);
extern void _ZN23QStyleOptionProgressBarC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionProgressBar::QStyleOptionProgressBar(int version);
extern void* dector_ZN23QStyleOptionProgressBarC1Ei(int32_t arg0);
extern void _ZN23QStyleOptionProgressBarC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionProgressBar::QStyleOptionProgressBar();
extern void* dector_ZN23QStyleOptionProgressBarC1Ev();
extern void _ZN23QStyleOptionProgressBarC1Ev(void* qthis);
  // proto:  void QStyleOptionSlider::QStyleOptionSlider(const QStyleOptionSlider & other);
extern void* dector_ZN18QStyleOptionSliderC1ERKS_(void* arg0);
extern void _ZN18QStyleOptionSliderC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionSlider::QStyleOptionSlider(int version);
extern void* dector_ZN18QStyleOptionSliderC1Ei(int32_t arg0);
extern void _ZN18QStyleOptionSliderC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionSlider::QStyleOptionSlider();
extern void* dector_ZN18QStyleOptionSliderC1Ev();
extern void _ZN18QStyleOptionSliderC1Ev(void* qthis);
  // proto:  void QStyleOptionFrame::QStyleOptionFrame();
extern void* dector_ZN17QStyleOptionFrameC1Ev();
extern void _ZN17QStyleOptionFrameC1Ev(void* qthis);
  // proto:  void QStyleOptionFrame::QStyleOptionFrame(const QStyleOptionFrame & other);
extern void* dector_ZN17QStyleOptionFrameC1ERKS_(void* arg0);
extern void _ZN17QStyleOptionFrameC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionFrame::QStyleOptionFrame(int version);
extern void* dector_ZN17QStyleOptionFrameC1Ei(int32_t arg0);
extern void _ZN17QStyleOptionFrameC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionComplex::QStyleOptionComplex(int version, int type);
extern void* dector_ZN19QStyleOptionComplexC1Eii(int32_t arg0, int32_t arg1);
extern void _ZN19QStyleOptionComplexC1Eii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QStyleOptionComplex::QStyleOptionComplex(const QStyleOptionComplex & other);
extern void* dector_ZN19QStyleOptionComplexC1ERKS_(void* arg0);
extern void _ZN19QStyleOptionComplexC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleHintReturn::~QStyleHintReturn();
extern void _ZN16QStyleHintReturnD0Ev(void* qthis);
  // proto:  void QStyleHintReturn::QStyleHintReturn(int version, int type);
extern void* dector_ZN16QStyleHintReturnC1Eii(int32_t arg0, int32_t arg1);
extern void _ZN16QStyleHintReturnC1Eii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QStyleOptionHeader::QStyleOptionHeader();
extern void* dector_ZN18QStyleOptionHeaderC1Ev();
extern void _ZN18QStyleOptionHeaderC1Ev(void* qthis);
  // proto:  void QStyleOptionHeader::QStyleOptionHeader(const QStyleOptionHeader & other);
extern void* dector_ZN18QStyleOptionHeaderC1ERKS_(void* arg0);
extern void _ZN18QStyleOptionHeaderC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionHeader::QStyleOptionHeader(int version);
extern void* dector_ZN18QStyleOptionHeaderC1Ei(int32_t arg0);
extern void _ZN18QStyleOptionHeaderC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionToolBox::QStyleOptionToolBox();
extern void* dector_ZN19QStyleOptionToolBoxC1Ev();
extern void _ZN19QStyleOptionToolBoxC1Ev(void* qthis);
  // proto:  void QStyleOptionToolBox::QStyleOptionToolBox(const QStyleOptionToolBox & other);
extern void* dector_ZN19QStyleOptionToolBoxC1ERKS_(void* arg0);
extern void _ZN19QStyleOptionToolBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionToolBox::QStyleOptionToolBox(int version);
extern void* dector_ZN19QStyleOptionToolBoxC1Ei(int32_t arg0);
extern void _ZN19QStyleOptionToolBoxC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionFocusRect::QStyleOptionFocusRect(int version);
extern void* dector_ZN21QStyleOptionFocusRectC1Ei(int32_t arg0);
extern void _ZN21QStyleOptionFocusRectC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionFocusRect::QStyleOptionFocusRect();
extern void* dector_ZN21QStyleOptionFocusRectC1Ev();
extern void _ZN21QStyleOptionFocusRectC1Ev(void* qthis);
  // proto:  void QStyleOptionFocusRect::QStyleOptionFocusRect(const QStyleOptionFocusRect & other);
extern void* dector_ZN21QStyleOptionFocusRectC1ERKS_(void* arg0);
extern void _ZN21QStyleOptionFocusRectC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionGroupBox::QStyleOptionGroupBox(int version);
extern void* dector_ZN20QStyleOptionGroupBoxC1Ei(int32_t arg0);
extern void _ZN20QStyleOptionGroupBoxC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionGroupBox::QStyleOptionGroupBox(const QStyleOptionGroupBox & other);
extern void* dector_ZN20QStyleOptionGroupBoxC1ERKS_(void* arg0);
extern void _ZN20QStyleOptionGroupBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionGroupBox::QStyleOptionGroupBox();
extern void* dector_ZN20QStyleOptionGroupBoxC1Ev();
extern void _ZN20QStyleOptionGroupBoxC1Ev(void* qthis);
  // proto:  void QStyleOptionTab::QStyleOptionTab(const QStyleOptionTab & other);
extern void* dector_ZN15QStyleOptionTabC1ERKS_(void* arg0);
extern void _ZN15QStyleOptionTabC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionTab::QStyleOptionTab(int version);
extern void* dector_ZN15QStyleOptionTabC1Ei(int32_t arg0);
extern void _ZN15QStyleOptionTabC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionTab::QStyleOptionTab();
extern void* dector_ZN15QStyleOptionTabC1Ev();
extern void _ZN15QStyleOptionTabC1Ev(void* qthis);
  // proto:  void QStyleOptionTabBarBase::QStyleOptionTabBarBase();
extern void* dector_ZN22QStyleOptionTabBarBaseC1Ev();
extern void _ZN22QStyleOptionTabBarBaseC1Ev(void* qthis);
  // proto:  void QStyleOptionTabBarBase::QStyleOptionTabBarBase(int version);
extern void* dector_ZN22QStyleOptionTabBarBaseC1Ei(int32_t arg0);
extern void _ZN22QStyleOptionTabBarBaseC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionTabBarBase::QStyleOptionTabBarBase(const QStyleOptionTabBarBase & other);
extern void* dector_ZN22QStyleOptionTabBarBaseC1ERKS_(void* arg0);
extern void _ZN22QStyleOptionTabBarBaseC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionRubberBand::QStyleOptionRubberBand(int version);
extern void* dector_ZN22QStyleOptionRubberBandC1Ei(int32_t arg0);
extern void _ZN22QStyleOptionRubberBandC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionRubberBand::QStyleOptionRubberBand();
extern void* dector_ZN22QStyleOptionRubberBandC1Ev();
extern void _ZN22QStyleOptionRubberBandC1Ev(void* qthis);
  // proto:  void QStyleOptionRubberBand::QStyleOptionRubberBand(const QStyleOptionRubberBand & other);
extern void* dector_ZN22QStyleOptionRubberBandC1ERKS_(void* arg0);
extern void _ZN22QStyleOptionRubberBandC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionButton::QStyleOptionButton(int version);
extern void* dector_ZN18QStyleOptionButtonC1Ei(int32_t arg0);
extern void _ZN18QStyleOptionButtonC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionButton::QStyleOptionButton();
extern void* dector_ZN18QStyleOptionButtonC1Ev();
extern void _ZN18QStyleOptionButtonC1Ev(void* qthis);
  // proto:  void QStyleOptionButton::QStyleOptionButton(const QStyleOptionButton & other);
extern void* dector_ZN18QStyleOptionButtonC1ERKS_(void* arg0);
extern void _ZN18QStyleOptionButtonC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleHintReturnMask::QStyleHintReturnMask();
extern void* dector_ZN20QStyleHintReturnMaskC1Ev();
extern void _ZN20QStyleHintReturnMaskC1Ev(void* qthis);
  // proto:  void QStyleHintReturnMask::~QStyleHintReturnMask();
extern void _ZN20QStyleHintReturnMaskD0Ev(void* qthis);
  // proto:  void QStyleOptionToolButton::QStyleOptionToolButton(int version);
extern void* dector_ZN22QStyleOptionToolButtonC1Ei(int32_t arg0);
extern void _ZN22QStyleOptionToolButtonC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionToolButton::QStyleOptionToolButton();
extern void* dector_ZN22QStyleOptionToolButtonC1Ev();
extern void _ZN22QStyleOptionToolButtonC1Ev(void* qthis);
  // proto:  void QStyleOptionToolButton::QStyleOptionToolButton(const QStyleOptionToolButton & other);
extern void* dector_ZN22QStyleOptionToolButtonC1ERKS_(void* arg0);
extern void _ZN22QStyleOptionToolButtonC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionSizeGrip::QStyleOptionSizeGrip(int version);
extern void* dector_ZN20QStyleOptionSizeGripC1Ei(int32_t arg0);
extern void _ZN20QStyleOptionSizeGripC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionSizeGrip::QStyleOptionSizeGrip();
extern void* dector_ZN20QStyleOptionSizeGripC1Ev();
extern void _ZN20QStyleOptionSizeGripC1Ev(void* qthis);
  // proto:  void QStyleOptionSizeGrip::QStyleOptionSizeGrip(const QStyleOptionSizeGrip & other);
extern void* dector_ZN20QStyleOptionSizeGripC1ERKS_(void* arg0);
extern void _ZN20QStyleOptionSizeGripC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionViewItem::QStyleOptionViewItem(const QStyleOptionViewItem & other);
extern void* dector_ZN20QStyleOptionViewItemC1ERKS_(void* arg0);
extern void _ZN20QStyleOptionViewItemC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionViewItem::QStyleOptionViewItem(int version);
extern void* dector_ZN20QStyleOptionViewItemC1Ei(int32_t arg0);
extern void _ZN20QStyleOptionViewItemC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionViewItem::QStyleOptionViewItem();
extern void* dector_ZN20QStyleOptionViewItemC1Ev();
extern void _ZN20QStyleOptionViewItemC1Ev(void* qthis);
  // proto:  void QStyleOptionSpinBox::QStyleOptionSpinBox();
extern void* dector_ZN19QStyleOptionSpinBoxC1Ev();
extern void _ZN19QStyleOptionSpinBoxC1Ev(void* qthis);
  // proto:  void QStyleOptionSpinBox::QStyleOptionSpinBox(const QStyleOptionSpinBox & other);
extern void* dector_ZN19QStyleOptionSpinBoxC1ERKS_(void* arg0);
extern void _ZN19QStyleOptionSpinBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionSpinBox::QStyleOptionSpinBox(int version);
extern void* dector_ZN19QStyleOptionSpinBoxC1Ei(int32_t arg0);
extern void _ZN19QStyleOptionSpinBoxC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionToolBar::QStyleOptionToolBar(const QStyleOptionToolBar & other);
extern void* dector_ZN19QStyleOptionToolBarC1ERKS_(void* arg0);
extern void _ZN19QStyleOptionToolBarC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionToolBar::QStyleOptionToolBar(int version);
extern void* dector_ZN19QStyleOptionToolBarC1Ei(int32_t arg0);
extern void _ZN19QStyleOptionToolBarC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionToolBar::QStyleOptionToolBar();
extern void* dector_ZN19QStyleOptionToolBarC1Ev();
extern void _ZN19QStyleOptionToolBarC1Ev(void* qthis);
  // proto:  void QStyleOptionTabWidgetFrame::QStyleOptionTabWidgetFrame(const QStyleOptionTabWidgetFrame & other);
extern void* dector_ZN26QStyleOptionTabWidgetFrameC1ERKS_(void* arg0);
extern void demth_ZN26QStyleOptionTabWidgetFrameC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyleOptionTabWidgetFrame::QStyleOptionTabWidgetFrame(int version);
extern void* dector_ZN26QStyleOptionTabWidgetFrameC1Ei(int32_t arg0);
extern void _ZN26QStyleOptionTabWidgetFrameC1Ei(void* qthis, int32_t arg0);
  // proto:  void QStyleOptionTabWidgetFrame::QStyleOptionTabWidgetFrame();
extern void* dector_ZN26QStyleOptionTabWidgetFrameC1Ev();
extern void _ZN26QStyleOptionTabWidgetFrameC1Ev(void* qthis);
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

  // proto:  void QStyleOptionComboBox::QStyleOptionComboBox(const QStyleOptionComboBox & other);
func NewQStyleOptionComboBox(args ...interface{}) QStyleOptionComboBox {
  return QStyleOptionComboBox{}
}

  // proto:  void QStyleOptionMenuItem::QStyleOptionMenuItem(const QStyleOptionMenuItem & other);
func NewQStyleOptionMenuItem(args ...interface{}) QStyleOptionMenuItem {
  return QStyleOptionMenuItem{}
}

  // proto:  void QStyleHintReturnVariant::~QStyleHintReturnVariant();
func (this *QStyleHintReturnVariant) FreeQStyleHintReturnVariant(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleHintReturnVariant", "~QStyleHintReturnVariant", args)
  }

}

  // proto:  void QStyleHintReturnVariant::QStyleHintReturnVariant();
func NewQStyleHintReturnVariant(args ...interface{}) QStyleHintReturnVariant {
  return QStyleHintReturnVariant{}
}

  // proto:  void QStyleOptionTitleBar::QStyleOptionTitleBar(int version);
func NewQStyleOptionTitleBar(args ...interface{}) QStyleOptionTitleBar {
  return QStyleOptionTitleBar{}
}

  // proto:  void QStyleOptionGraphicsItem::QStyleOptionGraphicsItem();
func NewQStyleOptionGraphicsItem(args ...interface{}) QStyleOptionGraphicsItem {
  return QStyleOptionGraphicsItem{}
}

  // proto: static qreal QStyleOptionGraphicsItem::levelOfDetailFromTransform(const QTransform & worldTransform);
func (this *QStyleOptionGraphicsItem) levelOfDetailFromTransform_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleOptionGraphicsItem", "levelOfDetailFromTransform", args)
  }

}

  // proto:  void QStyleOption::~QStyleOption();
func (this *QStyleOption) FreeQStyleOption(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleOption", "~QStyleOption", args)
  }

}

  // proto:  void QStyleOption::init(const QWidget * w);
func (this *QStyleOption) init(args ...interface{}) () {
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
    C._ZN12QStyleOption4initEPK7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleOption", "init", args)
  }

}

  // proto:  void QStyleOption::QStyleOption(const QStyleOption & other);
func NewQStyleOption(args ...interface{}) QStyleOption {
  return QStyleOption{}
}

  // proto:  void QStyleOption::initFrom(const QWidget * w);
func (this *QStyleOption) initFrom(args ...interface{}) () {
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
    C.demth_ZN12QStyleOption8initFromEPK7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyleOption", "initFrom", args)
  }

}

  // proto:  void QStyleOptionDockWidget::QStyleOptionDockWidget();
func NewQStyleOptionDockWidget(args ...interface{}) QStyleOptionDockWidget {
  return QStyleOptionDockWidget{}
}

  // proto:  void QStyleOptionProgressBar::QStyleOptionProgressBar(const QStyleOptionProgressBar & other);
func NewQStyleOptionProgressBar(args ...interface{}) QStyleOptionProgressBar {
  return QStyleOptionProgressBar{}
}

  // proto:  void QStyleOptionSlider::QStyleOptionSlider(const QStyleOptionSlider & other);
func NewQStyleOptionSlider(args ...interface{}) QStyleOptionSlider {
  return QStyleOptionSlider{}
}

  // proto:  void QStyleOptionFrame::QStyleOptionFrame();
func NewQStyleOptionFrame(args ...interface{}) QStyleOptionFrame {
  return QStyleOptionFrame{}
}

  // proto:  void QStyleOptionComplex::QStyleOptionComplex(int version, int type);
func NewQStyleOptionComplex(args ...interface{}) QStyleOptionComplex {
  return QStyleOptionComplex{}
}

  // proto:  void QStyleHintReturn::~QStyleHintReturn();
func (this *QStyleHintReturn) FreeQStyleHintReturn(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleHintReturn", "~QStyleHintReturn", args)
  }

}

  // proto:  void QStyleHintReturn::QStyleHintReturn(int version, int type);
func NewQStyleHintReturn(args ...interface{}) QStyleHintReturn {
  return QStyleHintReturn{}
}

  // proto:  void QStyleOptionHeader::QStyleOptionHeader();
func NewQStyleOptionHeader(args ...interface{}) QStyleOptionHeader {
  return QStyleOptionHeader{}
}

  // proto:  void QStyleOptionToolBox::QStyleOptionToolBox();
func NewQStyleOptionToolBox(args ...interface{}) QStyleOptionToolBox {
  return QStyleOptionToolBox{}
}

  // proto:  void QStyleOptionFocusRect::QStyleOptionFocusRect(int version);
func NewQStyleOptionFocusRect(args ...interface{}) QStyleOptionFocusRect {
  return QStyleOptionFocusRect{}
}

  // proto:  void QStyleOptionGroupBox::QStyleOptionGroupBox(int version);
func NewQStyleOptionGroupBox(args ...interface{}) QStyleOptionGroupBox {
  return QStyleOptionGroupBox{}
}

  // proto:  void QStyleOptionTab::QStyleOptionTab(const QStyleOptionTab & other);
func NewQStyleOptionTab(args ...interface{}) QStyleOptionTab {
  return QStyleOptionTab{}
}

  // proto:  void QStyleOptionTabBarBase::QStyleOptionTabBarBase();
func NewQStyleOptionTabBarBase(args ...interface{}) QStyleOptionTabBarBase {
  return QStyleOptionTabBarBase{}
}

  // proto:  void QStyleOptionRubberBand::QStyleOptionRubberBand(int version);
func NewQStyleOptionRubberBand(args ...interface{}) QStyleOptionRubberBand {
  return QStyleOptionRubberBand{}
}

  // proto:  void QStyleOptionButton::QStyleOptionButton(int version);
func NewQStyleOptionButton(args ...interface{}) QStyleOptionButton {
  return QStyleOptionButton{}
}

  // proto:  void QStyleHintReturnMask::QStyleHintReturnMask();
func NewQStyleHintReturnMask(args ...interface{}) QStyleHintReturnMask {
  return QStyleHintReturnMask{}
}

  // proto:  void QStyleHintReturnMask::~QStyleHintReturnMask();
func (this *QStyleHintReturnMask) FreeQStyleHintReturnMask(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyleHintReturnMask", "~QStyleHintReturnMask", args)
  }

}

  // proto:  void QStyleOptionToolButton::QStyleOptionToolButton(int version);
func NewQStyleOptionToolButton(args ...interface{}) QStyleOptionToolButton {
  return QStyleOptionToolButton{}
}

  // proto:  void QStyleOptionSizeGrip::QStyleOptionSizeGrip(int version);
func NewQStyleOptionSizeGrip(args ...interface{}) QStyleOptionSizeGrip {
  return QStyleOptionSizeGrip{}
}

  // proto:  void QStyleOptionViewItem::QStyleOptionViewItem(const QStyleOptionViewItem & other);
func NewQStyleOptionViewItem(args ...interface{}) QStyleOptionViewItem {
  return QStyleOptionViewItem{}
}

  // proto:  void QStyleOptionSpinBox::QStyleOptionSpinBox();
func NewQStyleOptionSpinBox(args ...interface{}) QStyleOptionSpinBox {
  return QStyleOptionSpinBox{}
}

  // proto:  void QStyleOptionToolBar::QStyleOptionToolBar(const QStyleOptionToolBar & other);
func NewQStyleOptionToolBar(args ...interface{}) QStyleOptionToolBar {
  return QStyleOptionToolBar{}
}

  // proto:  void QStyleOptionTabWidgetFrame::QStyleOptionTabWidgetFrame(const QStyleOptionTabWidgetFrame & other);
func NewQStyleOptionTabWidgetFrame(args ...interface{}) QStyleOptionTabWidgetFrame {
  return QStyleOptionTabWidgetFrame{}
}

// <= body block end

