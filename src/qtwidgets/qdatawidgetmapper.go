package qtwidgets

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h
// #include <qdatawidgetmapper.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 43
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
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QDataWidgetMapper) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQDataWidgetMapperFromPointer(cthis unsafe.Pointer) *QDataWidgetMapper {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QDataWidgetMapper{bcthis0}
}
func (*QDataWidgetMapper) NewFromPointer(cthis unsafe.Pointer) *QDataWidgetMapper {
	return NewQDataWidgetMapperFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDataWidgetMapper) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:64
// index:0
// Public
// void QDataWidgetMapper(class QObject *)
func NewQDataWidgetMapper(parent *qtcore.QObject /*444 QObject **/) *QDataWidgetMapper {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapperC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQDataWidgetMapperFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:65
// index:0
// Public virtual
// void ~QDataWidgetMapper()
func DeleteQDataWidgetMapper(*QDataWidgetMapper) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapperD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:67
// index:0
// Public
// void setModel(class QAbstractItemModel *)
func (this *QDataWidgetMapper) SetModel(model *qtcore.QAbstractItemModel /*444 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:68
// index:0
// Public
// QAbstractItemModel * model()
func (this *QDataWidgetMapper) Model() *qtcore.QAbstractItemModel /*444 QAbstractItemModel **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper5modelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:70
// index:0
// Public
// void setItemDelegate(class QAbstractItemDelegate *)
func (this *QDataWidgetMapper) SetItemDelegate(delegate *QAbstractItemDelegate /*444 QAbstractItemDelegate **/) {
	var convArg0 = delegate.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:71
// index:0
// Public
// QAbstractItemDelegate * itemDelegate()
func (this *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate /*444 QAbstractItemDelegate **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12itemDelegateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:73
// index:0
// Public
// void setRootIndex(const class QModelIndex &)
func (this *QDataWidgetMapper) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:74
// index:0
// Public
// QModelIndex rootIndex()
func (this *QDataWidgetMapper) RootIndex() *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper9rootIndexEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:76
// index:0
// Public
// void setOrientation(Qt::Orientation)
func (this *QDataWidgetMapper) SetOrientation(aOrientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper14setOrientationEN2Qt11OrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), aOrientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:77
// index:0
// Public
// Qt::Orientation orientation()
func (this *QDataWidgetMapper) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:81
// index:0
// Public
// void setSubmitPolicy(enum QDataWidgetMapper::SubmitPolicy)
func (this *QDataWidgetMapper) SetSubmitPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setSubmitPolicyENS_12SubmitPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:82
// index:0
// Public
// QDataWidgetMapper::SubmitPolicy submitPolicy()
func (this *QDataWidgetMapper) SubmitPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12submitPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:84
// index:0
// Public
// void addMapping(class QWidget *, int)
func (this *QDataWidgetMapper) AddMapping(widget *QWidget /*444 QWidget **/, section int) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper10addMappingEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, section)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:85
// index:1
// Public
// void addMapping(class QWidget *, int, const class QByteArray &)
func (this *QDataWidgetMapper) AddMapping_1(widget *QWidget /*444 QWidget **/, section int, propertyName *qtcore.QByteArray) {
	var convArg0 = widget.GetCthis()
	var convArg2 = propertyName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, section, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:86
// index:0
// Public
// void removeMapping(class QWidget *)
func (this *QDataWidgetMapper) RemoveMapping(widget *QWidget /*444 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper13removeMappingEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:87
// index:0
// Public
// int mappedSection(class QWidget *)
func (this *QDataWidgetMapper) MappedSection(widget *QWidget /*444 QWidget **/) int {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper13mappedSectionEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:88
// index:0
// Public
// QByteArray mappedPropertyName(class QWidget *)
func (this *QDataWidgetMapper) MappedPropertyName(widget *QWidget /*444 QWidget **/) *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:89
// index:0
// Public
// QWidget * mappedWidgetAt(int)
func (this *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget /*444 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper14mappedWidgetAtEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), section)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:90
// index:0
// Public
// void clearMapping()
func (this *QDataWidgetMapper) ClearMapping() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper12clearMappingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:92
// index:0
// Public
// int currentIndex()
func (this *QDataWidgetMapper) CurrentIndex() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12currentIndexEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:95
// index:0
// Public
// void revert()
func (this *QDataWidgetMapper) Revert() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper6revertEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:96
// index:0
// Public
// bool submit()
func (this *QDataWidgetMapper) Submit() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper6submitEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:98
// index:0
// Public
// void toFirst()
func (this *QDataWidgetMapper) ToFirst() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper7toFirstEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:99
// index:0
// Public
// void toLast()
func (this *QDataWidgetMapper) ToLast() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper6toLastEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:100
// index:0
// Public
// void toNext()
func (this *QDataWidgetMapper) ToNext() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper6toNextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:101
// index:0
// Public
// void toPrevious()
func (this *QDataWidgetMapper) ToPrevious() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper10toPreviousEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:102
// index:0
// Public virtual
// void setCurrentIndex(int)
func (this *QDataWidgetMapper) SetCurrentIndex(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setCurrentIndexEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:103
// index:0
// Public
// void setCurrentModelIndex(const class QModelIndex &)
func (this *QDataWidgetMapper) SetCurrentModelIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:106
// index:0
// Public
// void currentIndexChanged(int)
func (this *QDataWidgetMapper) CurrentIndexChanged(index int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QDataWidgetMapper19currentIndexChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), index)
	gopp.ErrPrint(err, rv)
}

type QDataWidgetMapper__SubmitPolicy = int

const QDataWidgetMapper__AutoSubmit QDataWidgetMapper__SubmitPolicy = 0
const QDataWidgetMapper__ManualSubmit QDataWidgetMapper__SubmitPolicy = 1

//  body block end
