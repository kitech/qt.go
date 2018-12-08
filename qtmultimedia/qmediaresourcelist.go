package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaresource.h
// #include <qmediaresource.h>
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
type QMediaResourceList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QMediaResourceList) Operator_equal0() *QMediaResourceList {
	// QMediaResourceList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QMediaResourceList) Operator_equal1() *QMediaResourceList {
	// QMediaResourceList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QMediaResourceList) Swap0() {
	// QMediaResourceList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QMediaResourceList) Operator_equal_equal0() bool {
	// QMediaResourceList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QMediaResourceList) Operator_not_equal0() bool {
	// QMediaResourceList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QMediaResourceList) Size0() int {
	// QMediaResourceList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QMediaResourceList) Detach0() {
	// QMediaResourceList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QMediaResourceList) DetachShared0() {
	// QMediaResourceList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QMediaResourceList) IsDetached0() bool {
	// QMediaResourceList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QMediaResourceList) SetSharable0() {
	// QMediaResourceList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QMediaResourceList) IsSharedWith0() bool {
	// QMediaResourceList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QMediaResourceList) IsEmpty0() bool {
	// QMediaResourceList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QMediaResourceList) Clear0() {
	// QMediaResourceList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QMediaResourceList) At0() *QMediaResource {
	// QMediaResourceList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// const T & operator[](int)
func (this *QMediaResourceList) Operator_get_index0() *QMediaResource {
	// QMediaResourceList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// T & operator[](int)
func (this *QMediaResourceList) Operator_get_index1() *QMediaResource {
	// QMediaResourceList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// void reserve(int)
func (this *QMediaResourceList) Reserve0() {
	// QMediaResourceList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QMediaResourceList) Append0() {
	// QMediaResourceList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QMediaResourceList) Append1() {
	// QMediaResourceList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QMediaResourceList) Prepend0() {
	// QMediaResourceList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QMediaResourceList) Insert0() {
	// QMediaResourceList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QMediaResourceList) Replace0() {
	// QMediaResourceList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QMediaResourceList) RemoveAt0() {
	// QMediaResourceList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QMediaResourceList) RemoveAll0() int {
	// QMediaResourceList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QMediaResourceList) RemoveOne0() bool {
	// QMediaResourceList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QMediaResourceList) TakeAt0() *QMediaResource {
	// QMediaResourceList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// T takeFirst()
func (this *QMediaResourceList) TakeFirst0() *QMediaResource {
	// QMediaResourceList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// T takeLast()
func (this *QMediaResourceList) TakeLast0() *QMediaResource {
	// QMediaResourceList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// void move(int, int)
func (this *QMediaResourceList) Move0() {
	// QMediaResourceList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QMediaResourceList) Swap1() {
	// QMediaResourceList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QMediaResourceList) IndexOf0() int {
	// QMediaResourceList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QMediaResourceList) LastIndexOf0() int {
	// QMediaResourceList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QMediaResourceList) Contains0() bool {
	// QMediaResourceList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QMediaResourceList) Count0() int {
	// QMediaResourceList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QMediaResourceList) Begin0() {
	// QMediaResourceList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QMediaResourceList) Begin1() {
	// QMediaResourceList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QMediaResourceList) Cbegin0() {
	// QMediaResourceList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QMediaResourceList) ConstBegin0() {
	// QMediaResourceList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QMediaResourceList) End0() {
	// QMediaResourceList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QMediaResourceList) End1() {
	// QMediaResourceList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QMediaResourceList) Cend0() {
	// QMediaResourceList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QMediaResourceList) ConstEnd0() {
	// QMediaResourceList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QMediaResourceList) Rbegin0() {
	// QMediaResourceList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QMediaResourceList) Rend0() {
	// QMediaResourceList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QMediaResourceList) Rbegin1() {
	// QMediaResourceList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QMediaResourceList) Rend1() {
	// QMediaResourceList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QMediaResourceList) Crbegin0() {
	// QMediaResourceList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QMediaResourceList) Crend0() {
	// QMediaResourceList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QMediaResourceList) Insert1() {
	// QMediaResourceList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QMediaResourceList) Erase0() {
	// QMediaResourceList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QMediaResourceList) Erase1() {
	// QMediaResourceList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QMediaResourceList) Count1() int {
	// QMediaResourceList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QMediaResourceList) Length0() int {
	// QMediaResourceList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QMediaResourceList) First0() *QMediaResource {
	// QMediaResourceList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// const T & constFirst()
func (this *QMediaResourceList) ConstFirst0() *QMediaResource {
	// QMediaResourceList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// const T & first()
func (this *QMediaResourceList) First1() *QMediaResource {
	// QMediaResourceList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// T & last()
func (this *QMediaResourceList) Last0() *QMediaResource {
	// QMediaResourceList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// const T & last()
func (this *QMediaResourceList) Last1() *QMediaResource {
	// QMediaResourceList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// const T & constLast()
func (this *QMediaResourceList) ConstLast0() *QMediaResource {
	// QMediaResourceList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// void removeFirst()
func (this *QMediaResourceList) RemoveFirst0() {
	// QMediaResourceList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QMediaResourceList) RemoveLast0() {
	// QMediaResourceList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QMediaResourceList) StartsWith0() bool {
	// QMediaResourceList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QMediaResourceList) EndsWith0() bool {
	// QMediaResourceList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QMediaResourceList) Mid0() *QMediaResourceList {
	// QMediaResourceList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QMediaResourceList) Value0() *QMediaResource {
	// QMediaResourceList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// T value(int, const T &)
func (this *QMediaResourceList) Value1() *QMediaResource {
	// QMediaResourceList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// void push_back(const T &)
func (this *QMediaResourceList) Push_back0() {
	// QMediaResourceList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QMediaResourceList) Push_front0() {
	// QMediaResourceList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QMediaResourceList) Front0() *QMediaResource {
	// QMediaResourceList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// const T & front()
func (this *QMediaResourceList) Front1() *QMediaResource {
	// QMediaResourceList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// T & back()
func (this *QMediaResourceList) Back0() *QMediaResource {
	// QMediaResourceList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// const T & back()
func (this *QMediaResourceList) Back1() *QMediaResource {
	// QMediaResourceList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QMediaResource{}
}

// void pop_front()
func (this *QMediaResourceList) Pop_front0() {
	// QMediaResourceList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QMediaResourceList) Pop_back0() {
	// QMediaResourceList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QMediaResourceList) Empty0() bool {
	// QMediaResourceList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QMediaResourceList) Operator_add_equal0() *QMediaResourceList {
	// QMediaResourceList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QMediaResourceList) Operator_add0() *QMediaResourceList {
	// QMediaResourceList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QMediaResourceList) Operator_add_equal1() *QMediaResourceList {
	// QMediaResourceList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QMediaResourceList) Operator_left_shift0() *QMediaResourceList {
	// QMediaResourceList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QMediaResourceList) Operator_left_shift1() *QMediaResourceList {
	// QMediaResourceList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QMediaResourceList) ToVector0() {
	// QMediaResourceList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QMediaResourceList) ToSet0() {
	// QMediaResourceList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QMediaResourceList) FromVector0() *QMediaResourceList {
	// QMediaResourceList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QMediaResourceList) FromSet0() *QMediaResourceList {
	// QMediaResourceList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QMediaResourceList) FromStdList0() *QMediaResourceList {
	// QMediaResourceList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QMediaResourceList) ToStdList0() {
	// QMediaResourceList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QMediaResourceList) Detach_helper_grow0() {
	// QMediaResourceList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QMediaResourceList) Detach_helper0() {
	// QMediaResourceList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QMediaResourceList) Detach_helper1() {
	// QMediaResourceList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QMediaResourceList) Dealloc0() {
	// QMediaResourceList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QMediaResourceList) Node_construct0() {
	// QMediaResourceList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QMediaResourceList) Node_destruct0() {
	// QMediaResourceList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QMediaResourceList) Node_copy0() {
	// QMediaResourceList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QMediaResourceList) Node_destruct1() {
	// QMediaResourceList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QMediaResourceList) IsValidIterator0() bool {
	// QMediaResourceList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QMediaResourceList) Op_eq_impl0() bool {
	// QMediaResourceList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QMediaResourceList) Op_eq_impl1() bool {
	// QMediaResourceList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QMediaResourceList) Contains_impl0() bool {
	// QMediaResourceList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QMediaResourceList) Contains_impl1() bool {
	// QMediaResourceList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QMediaResourceList) Count_impl0() int {
	// QMediaResourceList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QMediaResourceList) Count_impl1() int {
	// QMediaResourceList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
