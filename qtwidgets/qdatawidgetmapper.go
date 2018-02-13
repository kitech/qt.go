package qtwidgets

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h
// #include <qdatawidgetmapper.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 44
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QDataWidgetMapper struct {
	*qtcore.QObject
}
type QDataWidgetMapper_ITF interface {
	qtcore.QObject_ITF
	QDataWidgetMapper_PTR() *QDataWidgetMapper
}

func (ptr *QDataWidgetMapper) QDataWidgetMapper_PTR() *QDataWidgetMapper { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QDataWidgetMapper) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDataWidgetMapper(QObject *)
func NewQDataWidgetMapper(parent *qtcore.QObject /*777 QObject **/) *QDataWidgetMapper {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapperC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDataWidgetMapperFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDataWidgetMapper()
func DeleteQDataWidgetMapper(this *QDataWidgetMapper) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapperD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModel(QAbstractItemModel *)
func (this *QDataWidgetMapper) SetModel(model *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/) {
	var convArg0 = model.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper8setModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * model()
func (this *QDataWidgetMapper) Model() *qtcore.QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemDelegate(QAbstractItemDelegate *)
func (this *QDataWidgetMapper) SetItemDelegate(delegate *QAbstractItemDelegate /*777 QAbstractItemDelegate **/) {
	var convArg0 = delegate.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setItemDelegateEP21QAbstractItemDelegate", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemDelegate * itemDelegate()
func (this *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate /*777 QAbstractItemDelegate **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12itemDelegateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemDelegateFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRootIndex(const QModelIndex &)
func (this *QDataWidgetMapper) SetRootIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper12setRootIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:74
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex rootIndex()
func (this *QDataWidgetMapper) RootIndex() *qtcore.QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper9rootIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)
func (this *QDataWidgetMapper) SetOrientation(aOrientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), aOrientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation()
func (this *QDataWidgetMapper) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSubmitPolicy(enum QDataWidgetMapper::SubmitPolicy)
func (this *QDataWidgetMapper) SetSubmitPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setSubmitPolicyENS_12SubmitPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QDataWidgetMapper::SubmitPolicy submitPolicy()
func (this *QDataWidgetMapper) SubmitPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12submitPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addMapping(QWidget *, int)
func (this *QDataWidgetMapper) AddMapping(widget *QWidget /*777 QWidget **/, section int) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper10addMappingEP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, section)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addMapping(QWidget *, int, const QByteArray &)
func (this *QDataWidgetMapper) AddMapping_1(widget *QWidget /*777 QWidget **/, section int, propertyName *qtcore.QByteArray) {
	var convArg0 = widget.GetCthis()
	var convArg2 = propertyName.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper10addMappingEP7QWidgetiRK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, section, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeMapping(QWidget *)
func (this *QDataWidgetMapper) RemoveMapping(widget *QWidget /*777 QWidget **/) {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper13removeMappingEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int mappedSection(QWidget *)
func (this *QDataWidgetMapper) MappedSection(widget *QWidget /*777 QWidget **/) int {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper13mappedSectionEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray mappedPropertyName(QWidget *)
func (this *QDataWidgetMapper) MappedPropertyName(widget *QWidget /*777 QWidget **/) *qtcore.QByteArray /*123*/ {
	var convArg0 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper18mappedPropertyNameEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * mappedWidgetAt(int)
func (this *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper14mappedWidgetAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMapping()
func (this *QDataWidgetMapper) ClearMapping() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper12clearMappingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int currentIndex()
func (this *QDataWidgetMapper) CurrentIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QDataWidgetMapper12currentIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void revert()
func (this *QDataWidgetMapper) Revert() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper6revertEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool submit()
func (this *QDataWidgetMapper) Submit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper6submitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toFirst()
func (this *QDataWidgetMapper) ToFirst() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper7toFirstEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toLast()
func (this *QDataWidgetMapper) ToLast() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper6toLastEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toNext()
func (this *QDataWidgetMapper) ToNext() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper6toNextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void toPrevious()
func (this *QDataWidgetMapper) ToPrevious() {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper10toPreviousEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:102
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setCurrentIndex(int)
func (this *QDataWidgetMapper) SetCurrentIndex(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper15setCurrentIndexEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCurrentModelIndex(const QModelIndex &)
func (this *QDataWidgetMapper) SetCurrentModelIndex(index *qtcore.QModelIndex) {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper20setCurrentModelIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdatawidgetmapper.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void currentIndexChanged(int)
func (this *QDataWidgetMapper) CurrentIndexChanged(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QDataWidgetMapper19currentIndexChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

type QDataWidgetMapper__SubmitPolicy = int

const QDataWidgetMapper__AutoSubmit QDataWidgetMapper__SubmitPolicy = 0
const QDataWidgetMapper__ManualSubmit QDataWidgetMapper__SubmitPolicy = 1

//  body block end

//  keep block begin

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
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
