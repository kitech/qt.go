package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerafocus.h
// #include <qcamerafocus.h>
// #include <QtMultimedia>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  keep block begin

func init_unused_10175() {
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
		log.Println(123)
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
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end

//  body block begin
type QCameraFocusZoneList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QCameraFocusZoneList) Operator_equal0() *QCameraFocusZoneList {
	// QCameraFocusZoneList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QCameraFocusZoneList) Operator_equal1() *QCameraFocusZoneList {
	// QCameraFocusZoneList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QCameraFocusZoneList) Swap0() {
	// QCameraFocusZoneList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QCameraFocusZoneList) Operator_equal_equal0() bool {
	// QCameraFocusZoneList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QCameraFocusZoneList) Operator_not_equal0() bool {
	// QCameraFocusZoneList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QCameraFocusZoneList) Size0() int {
	// QCameraFocusZoneList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QCameraFocusZoneList) Detach0() {
	// QCameraFocusZoneList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QCameraFocusZoneList) DetachShared0() {
	// QCameraFocusZoneList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QCameraFocusZoneList) IsDetached0() bool {
	// QCameraFocusZoneList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QCameraFocusZoneList) SetSharable0() {
	// QCameraFocusZoneList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QCameraFocusZoneList) IsSharedWith0() bool {
	// QCameraFocusZoneList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QCameraFocusZoneList) IsEmpty0() bool {
	// QCameraFocusZoneList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QCameraFocusZoneList) Clear0() {
	// QCameraFocusZoneList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QCameraFocusZoneList) At0() *QCameraFocusZone {
	// QCameraFocusZoneList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// const T & operator[](int)
func (this *QCameraFocusZoneList) Operator_get_index0() *QCameraFocusZone {
	// QCameraFocusZoneList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// T & operator[](int)
func (this *QCameraFocusZoneList) Operator_get_index1() *QCameraFocusZone {
	// QCameraFocusZoneList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// void reserve(int)
func (this *QCameraFocusZoneList) Reserve0() {
	// QCameraFocusZoneList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QCameraFocusZoneList) Append0() {
	// QCameraFocusZoneList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QCameraFocusZoneList) Append1() {
	// QCameraFocusZoneList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QCameraFocusZoneList) Prepend0() {
	// QCameraFocusZoneList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QCameraFocusZoneList) Insert0() {
	// QCameraFocusZoneList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QCameraFocusZoneList) Replace0() {
	// QCameraFocusZoneList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QCameraFocusZoneList) RemoveAt0() {
	// QCameraFocusZoneList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QCameraFocusZoneList) RemoveAll0() int {
	// QCameraFocusZoneList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QCameraFocusZoneList) RemoveOne0() bool {
	// QCameraFocusZoneList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QCameraFocusZoneList) TakeAt0() *QCameraFocusZone {
	// QCameraFocusZoneList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// T takeFirst()
func (this *QCameraFocusZoneList) TakeFirst0() *QCameraFocusZone {
	// QCameraFocusZoneList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// T takeLast()
func (this *QCameraFocusZoneList) TakeLast0() *QCameraFocusZone {
	// QCameraFocusZoneList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// void move(int, int)
func (this *QCameraFocusZoneList) Move0() {
	// QCameraFocusZoneList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QCameraFocusZoneList) Swap1() {
	// QCameraFocusZoneList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QCameraFocusZoneList) IndexOf0() int {
	// QCameraFocusZoneList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QCameraFocusZoneList) LastIndexOf0() int {
	// QCameraFocusZoneList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QCameraFocusZoneList) Contains0() bool {
	// QCameraFocusZoneList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QCameraFocusZoneList) Count0() int {
	// QCameraFocusZoneList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QCameraFocusZoneList) Begin0() {
	// QCameraFocusZoneList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QCameraFocusZoneList) Begin1() {
	// QCameraFocusZoneList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QCameraFocusZoneList) Cbegin0() {
	// QCameraFocusZoneList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QCameraFocusZoneList) ConstBegin0() {
	// QCameraFocusZoneList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QCameraFocusZoneList) End0() {
	// QCameraFocusZoneList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QCameraFocusZoneList) End1() {
	// QCameraFocusZoneList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QCameraFocusZoneList) Cend0() {
	// QCameraFocusZoneList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QCameraFocusZoneList) ConstEnd0() {
	// QCameraFocusZoneList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QCameraFocusZoneList) Rbegin0() {
	// QCameraFocusZoneList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QCameraFocusZoneList) Rend0() {
	// QCameraFocusZoneList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QCameraFocusZoneList) Rbegin1() {
	// QCameraFocusZoneList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QCameraFocusZoneList) Rend1() {
	// QCameraFocusZoneList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QCameraFocusZoneList) Crbegin0() {
	// QCameraFocusZoneList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QCameraFocusZoneList) Crend0() {
	// QCameraFocusZoneList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QCameraFocusZoneList) Insert1() {
	// QCameraFocusZoneList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QCameraFocusZoneList) Erase0() {
	// QCameraFocusZoneList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QCameraFocusZoneList) Erase1() {
	// QCameraFocusZoneList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QCameraFocusZoneList) Count1() int {
	// QCameraFocusZoneList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QCameraFocusZoneList) Length0() int {
	// QCameraFocusZoneList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QCameraFocusZoneList) First0() *QCameraFocusZone {
	// QCameraFocusZoneList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// const T & constFirst()
func (this *QCameraFocusZoneList) ConstFirst0() *QCameraFocusZone {
	// QCameraFocusZoneList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// const T & first()
func (this *QCameraFocusZoneList) First1() *QCameraFocusZone {
	// QCameraFocusZoneList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// T & last()
func (this *QCameraFocusZoneList) Last0() *QCameraFocusZone {
	// QCameraFocusZoneList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// const T & last()
func (this *QCameraFocusZoneList) Last1() *QCameraFocusZone {
	// QCameraFocusZoneList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// const T & constLast()
func (this *QCameraFocusZoneList) ConstLast0() *QCameraFocusZone {
	// QCameraFocusZoneList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// void removeFirst()
func (this *QCameraFocusZoneList) RemoveFirst0() {
	// QCameraFocusZoneList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QCameraFocusZoneList) RemoveLast0() {
	// QCameraFocusZoneList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QCameraFocusZoneList) StartsWith0() bool {
	// QCameraFocusZoneList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QCameraFocusZoneList) EndsWith0() bool {
	// QCameraFocusZoneList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QCameraFocusZoneList) Mid0() *QCameraFocusZoneList {
	// QCameraFocusZoneList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QCameraFocusZoneList) Value0() *QCameraFocusZone {
	// QCameraFocusZoneList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// T value(int, const T &)
func (this *QCameraFocusZoneList) Value1() *QCameraFocusZone {
	// QCameraFocusZoneList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// void push_back(const T &)
func (this *QCameraFocusZoneList) Push_back0() {
	// QCameraFocusZoneList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QCameraFocusZoneList) Push_front0() {
	// QCameraFocusZoneList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QCameraFocusZoneList) Front0() *QCameraFocusZone {
	// QCameraFocusZoneList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// const T & front()
func (this *QCameraFocusZoneList) Front1() *QCameraFocusZone {
	// QCameraFocusZoneList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// T & back()
func (this *QCameraFocusZoneList) Back0() *QCameraFocusZone {
	// QCameraFocusZoneList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// const T & back()
func (this *QCameraFocusZoneList) Back1() *QCameraFocusZone {
	// QCameraFocusZoneList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraFocusZone{}
}

// void pop_front()
func (this *QCameraFocusZoneList) Pop_front0() {
	// QCameraFocusZoneList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QCameraFocusZoneList) Pop_back0() {
	// QCameraFocusZoneList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QCameraFocusZoneList) Empty0() bool {
	// QCameraFocusZoneList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QCameraFocusZoneList) Operator_add_equal0() *QCameraFocusZoneList {
	// QCameraFocusZoneList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QCameraFocusZoneList) Operator_add0() *QCameraFocusZoneList {
	// QCameraFocusZoneList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QCameraFocusZoneList) Operator_add_equal1() *QCameraFocusZoneList {
	// QCameraFocusZoneList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QCameraFocusZoneList) Operator_left_shift0() *QCameraFocusZoneList {
	// QCameraFocusZoneList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QCameraFocusZoneList) Operator_left_shift1() *QCameraFocusZoneList {
	// QCameraFocusZoneList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QCameraFocusZoneList) ToVector0() {
	// QCameraFocusZoneList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QCameraFocusZoneList) ToSet0() {
	// QCameraFocusZoneList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QCameraFocusZoneList) FromVector0() *QCameraFocusZoneList {
	// QCameraFocusZoneList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QCameraFocusZoneList) FromSet0() *QCameraFocusZoneList {
	// QCameraFocusZoneList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QCameraFocusZoneList) FromStdList0() *QCameraFocusZoneList {
	// QCameraFocusZoneList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QCameraFocusZoneList) ToStdList0() {
	// QCameraFocusZoneList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QCameraFocusZoneList) Detach_helper_grow0() {
	// QCameraFocusZoneList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QCameraFocusZoneList) Detach_helper0() {
	// QCameraFocusZoneList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QCameraFocusZoneList) Detach_helper1() {
	// QCameraFocusZoneList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QCameraFocusZoneList) Dealloc0() {
	// QCameraFocusZoneList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QCameraFocusZoneList) Node_construct0() {
	// QCameraFocusZoneList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QCameraFocusZoneList) Node_destruct0() {
	// QCameraFocusZoneList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QCameraFocusZoneList) Node_copy0() {
	// QCameraFocusZoneList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QCameraFocusZoneList) Node_destruct1() {
	// QCameraFocusZoneList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QCameraFocusZoneList) IsValidIterator0() bool {
	// QCameraFocusZoneList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QCameraFocusZoneList) Op_eq_impl0() bool {
	// QCameraFocusZoneList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QCameraFocusZoneList) Op_eq_impl1() bool {
	// QCameraFocusZoneList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QCameraFocusZoneList) Contains_impl0() bool {
	// QCameraFocusZoneList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QCameraFocusZoneList) Contains_impl1() bool {
	// QCameraFocusZoneList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QCameraFocusZoneList) Count_impl0() int {
	// QCameraFocusZoneList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QCameraFocusZoneList) Count_impl1() int {
	// QCameraFocusZoneList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
