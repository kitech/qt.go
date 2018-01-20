//  header block begin
// /usr/include/qt/QtWidgets/qdatawidgetmapper.h
// #include <qdatawidgetmapper.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
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
type QDataWidgetMapper struct {
	*qtcore.QObject
}

func (this *QDataWidgetMapper) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDataWidgetMapper) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:64
// index:0
// void QDataWidgetMapper(class QObject *)
func NewQDataWidgetMapper(parent unsafe.Pointer) *QDataWidgetMapper {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapperC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataWidgetMapperFromPointer(cthis)
	return gothis
}
func NewQDataWidgetMapperFromPointer(cthis unsafe.Pointer) *QDataWidgetMapper {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDataWidgetMapper{bcthis0}
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:65
// index:0
// virtual
// void ~QDataWidgetMapper()
func DeleteQDataWidgetMapper(*QDataWidgetMapper) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapperD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:67
// index:0
// void setModel(class QAbstractItemModel *)
func (this *QDataWidgetMapper) SetModel(model unsafe.Pointer) {
	// 0: (, model QAbstractItemModel *), (model)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_VOID, this.GetCthis(), model)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:68
// index:0
// QAbstractItemModel * model()
func (this *QDataWidgetMapper) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper5modelEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:70
// index:0
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QDataWidgetMapper) SetItemDelegate(delegate unsafe.Pointer) {
	// 0: (, delegate QAbstractItemDelegate *), (delegate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_VOID, this.GetCthis(), delegate)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:71
// index:0
// QAbstractItemDelegate * itemDelegate()
func (this *QDataWidgetMapper) ItemDelegate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12itemDelegateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:73
// index:0
// void setRootIndex(const class QModelIndex &)
func (this *QDataWidgetMapper) SetRootIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:74
// index:0
// QModelIndex rootIndex()
func (this *QDataWidgetMapper) RootIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper9rootIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:76
// index:0
// void setOrientation(Qt::Orientation)
func (this *QDataWidgetMapper) SetOrientation(aOrientation int) {
	// 0: (, aOrientation Qt::Orientation), (&aOrientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &aOrientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:77
// index:0
// Qt::Orientation orientation()
func (this *QDataWidgetMapper) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:81
// index:0
// void setSubmitPolicy(enum QDataWidgetMapper::SubmitPolicy)
func (this *QDataWidgetMapper) SetSubmitPolicy(policy int) {
	// 0: (, policy QDataWidgetMapper::SubmitPolicy), (&policy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setSubmitPolicyENS_12SubmitPolicyE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:82
// index:0
// QDataWidgetMapper::SubmitPolicy submitPolicy()
func (this *QDataWidgetMapper) SubmitPolicy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12submitPolicyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:84
// index:0
// void addMapping(class QWidget *, int)
func (this *QDataWidgetMapper) AddMapping(widget unsafe.Pointer, section int) {
	// 0: (, widget QWidget *, section int), (widget, &section)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper10addMappingEP7QWidgeti", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, &section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:85
// index:1
// void addMapping(class QWidget *, int, const class QByteArray &)
func (this *QDataWidgetMapper) AddMapping_1(widget unsafe.Pointer, section int, propertyName unsafe.Pointer) {
	// 1: (, widget QWidget *, section int, propertyName const QByteArray &), (widget, &section, propertyName)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget, &section, propertyName)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:86
// index:0
// void removeMapping(class QWidget *)
func (this *QDataWidgetMapper) RemoveMapping(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper13removeMappingEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:87
// index:0
// int mappedSection(class QWidget *)
func (this *QDataWidgetMapper) MappedSection(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper13mappedSectionEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:88
// index:0
// QByteArray mappedPropertyName(class QWidget *)
func (this *QDataWidgetMapper) MappedPropertyName(widget unsafe.Pointer) {
	// 0: (, widget QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:89
// index:0
// QWidget * mappedWidgetAt(int)
func (this *QDataWidgetMapper) MappedWidgetAt(section int) {
	// 0: (, section int), (&section)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper14mappedWidgetAtEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:90
// index:0
// void clearMapping()
func (this *QDataWidgetMapper) ClearMapping() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper12clearMappingEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:92
// index:0
// int currentIndex()
func (this *QDataWidgetMapper) CurrentIndex() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12currentIndexEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:95
// index:0
// void revert()
func (this *QDataWidgetMapper) Revert() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper6revertEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:96
// index:0
// bool submit()
func (this *QDataWidgetMapper) Submit() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper6submitEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:98
// index:0
// void toFirst()
func (this *QDataWidgetMapper) ToFirst() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper7toFirstEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:99
// index:0
// void toLast()
func (this *QDataWidgetMapper) ToLast() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper6toLastEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:100
// index:0
// void toNext()
func (this *QDataWidgetMapper) ToNext() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper6toNextEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:101
// index:0
// void toPrevious()
func (this *QDataWidgetMapper) ToPrevious() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper10toPreviousEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:102
// index:0
// virtual
// void setCurrentIndex(int)
func (this *QDataWidgetMapper) SetCurrentIndex(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setCurrentIndexEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:103
// index:0
// void setCurrentModelIndex(const class QModelIndex &)
func (this *QDataWidgetMapper) SetCurrentModelIndex(index unsafe.Pointer) {
	// 0: (, index const QModelIndex &), (index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex", ffiqt.FFI_TYPE_VOID, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:106
// index:0
// void currentIndexChanged(int)
func (this *QDataWidgetMapper) CurrentIndexChanged(index int) {
	// 0: (, index int), (&index)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper19currentIndexChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &index)
	gopp.ErrPrint(err, rv)
}

//  body block end
