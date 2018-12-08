package qtcore

// /usr/include/qt/QtCore/qfileinfo.h
// #include <qfileinfo.h>
// #include <QtCore>

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
}

//  keep block end

//  body block begin
type QFileInfoList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QFileInfoList) Operator_equal0() *QFileInfoList {
	// QFileInfoList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QFileInfoList) Operator_equal1() *QFileInfoList {
	// QFileInfoList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QFileInfoList) Swap0() {
	// QFileInfoList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QFileInfoList) Operator_equal_equal0() bool {
	// QFileInfoList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QFileInfoList) Operator_not_equal0() bool {
	// QFileInfoList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QFileInfoList) Size0() int {
	// QFileInfoList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QFileInfoList) Detach0() {
	// QFileInfoList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QFileInfoList) DetachShared0() {
	// QFileInfoList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QFileInfoList) IsDetached0() bool {
	// QFileInfoList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QFileInfoList) SetSharable0() {
	// QFileInfoList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QFileInfoList) IsSharedWith0() bool {
	// QFileInfoList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QFileInfoList) IsEmpty0() bool {
	// QFileInfoList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QFileInfoList) Clear0() {
	// QFileInfoList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QFileInfoList) At0() *QFileInfo {
	// QFileInfoList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// const T & operator[](int)
func (this *QFileInfoList) Operator_get_index0() *QFileInfo {
	// QFileInfoList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// T & operator[](int)
func (this *QFileInfoList) Operator_get_index1() *QFileInfo {
	// QFileInfoList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// void reserve(int)
func (this *QFileInfoList) Reserve0() {
	// QFileInfoList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QFileInfoList) Append0() {
	// QFileInfoList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QFileInfoList) Append1() {
	// QFileInfoList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QFileInfoList) Prepend0() {
	// QFileInfoList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QFileInfoList) Insert0() {
	// QFileInfoList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QFileInfoList) Replace0() {
	// QFileInfoList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QFileInfoList) RemoveAt0() {
	// QFileInfoList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QFileInfoList) RemoveAll0() int {
	// QFileInfoList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QFileInfoList) RemoveOne0() bool {
	// QFileInfoList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QFileInfoList) TakeAt0() *QFileInfo {
	// QFileInfoList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// T takeFirst()
func (this *QFileInfoList) TakeFirst0() *QFileInfo {
	// QFileInfoList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// T takeLast()
func (this *QFileInfoList) TakeLast0() *QFileInfo {
	// QFileInfoList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// void move(int, int)
func (this *QFileInfoList) Move0() {
	// QFileInfoList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QFileInfoList) Swap1() {
	// QFileInfoList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QFileInfoList) IndexOf0() int {
	// QFileInfoList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QFileInfoList) LastIndexOf0() int {
	// QFileInfoList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QFileInfoList) Contains0() bool {
	// QFileInfoList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QFileInfoList) Count0() int {
	// QFileInfoList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QFileInfoList) Begin0() {
	// QFileInfoList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QFileInfoList) Begin1() {
	// QFileInfoList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QFileInfoList) Cbegin0() {
	// QFileInfoList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QFileInfoList) ConstBegin0() {
	// QFileInfoList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QFileInfoList) End0() {
	// QFileInfoList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QFileInfoList) End1() {
	// QFileInfoList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QFileInfoList) Cend0() {
	// QFileInfoList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QFileInfoList) ConstEnd0() {
	// QFileInfoList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QFileInfoList) Rbegin0() {
	// QFileInfoList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QFileInfoList) Rend0() {
	// QFileInfoList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QFileInfoList) Rbegin1() {
	// QFileInfoList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QFileInfoList) Rend1() {
	// QFileInfoList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QFileInfoList) Crbegin0() {
	// QFileInfoList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QFileInfoList) Crend0() {
	// QFileInfoList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QFileInfoList) Insert1() {
	// QFileInfoList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QFileInfoList) Erase0() {
	// QFileInfoList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QFileInfoList) Erase1() {
	// QFileInfoList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QFileInfoList) Count1() int {
	// QFileInfoList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QFileInfoList) Length0() int {
	// QFileInfoList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QFileInfoList) First0() *QFileInfo {
	// QFileInfoList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// const T & constFirst()
func (this *QFileInfoList) ConstFirst0() *QFileInfo {
	// QFileInfoList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// const T & first()
func (this *QFileInfoList) First1() *QFileInfo {
	// QFileInfoList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// T & last()
func (this *QFileInfoList) Last0() *QFileInfo {
	// QFileInfoList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// const T & last()
func (this *QFileInfoList) Last1() *QFileInfo {
	// QFileInfoList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// const T & constLast()
func (this *QFileInfoList) ConstLast0() *QFileInfo {
	// QFileInfoList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// void removeFirst()
func (this *QFileInfoList) RemoveFirst0() {
	// QFileInfoList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QFileInfoList) RemoveLast0() {
	// QFileInfoList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QFileInfoList) StartsWith0() bool {
	// QFileInfoList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QFileInfoList) EndsWith0() bool {
	// QFileInfoList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QFileInfoList) Mid0() *QFileInfoList {
	// QFileInfoList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QFileInfoList) Value0() *QFileInfo {
	// QFileInfoList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// T value(int, const T &)
func (this *QFileInfoList) Value1() *QFileInfo {
	// QFileInfoList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// void push_back(const T &)
func (this *QFileInfoList) Push_back0() {
	// QFileInfoList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QFileInfoList) Push_front0() {
	// QFileInfoList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QFileInfoList) Front0() *QFileInfo {
	// QFileInfoList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// const T & front()
func (this *QFileInfoList) Front1() *QFileInfo {
	// QFileInfoList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// T & back()
func (this *QFileInfoList) Back0() *QFileInfo {
	// QFileInfoList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// const T & back()
func (this *QFileInfoList) Back1() *QFileInfo {
	// QFileInfoList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QFileInfo{}
}

// void pop_front()
func (this *QFileInfoList) Pop_front0() {
	// QFileInfoList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QFileInfoList) Pop_back0() {
	// QFileInfoList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QFileInfoList) Empty0() bool {
	// QFileInfoList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QFileInfoList) Operator_add_equal0() *QFileInfoList {
	// QFileInfoList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QFileInfoList) Operator_add0() *QFileInfoList {
	// QFileInfoList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QFileInfoList) Operator_add_equal1() *QFileInfoList {
	// QFileInfoList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QFileInfoList) Operator_left_shift0() *QFileInfoList {
	// QFileInfoList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QFileInfoList) Operator_left_shift1() *QFileInfoList {
	// QFileInfoList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QFileInfoList) ToVector0() {
	// QFileInfoList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QFileInfoList) ToSet0() {
	// QFileInfoList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QFileInfoList) FromVector0() *QFileInfoList {
	// QFileInfoList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QFileInfoList) FromSet0() *QFileInfoList {
	// QFileInfoList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QFileInfoList) FromStdList0() *QFileInfoList {
	// QFileInfoList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QFileInfoList) ToStdList0() {
	// QFileInfoList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QFileInfoList) Detach_helper_grow0() {
	// QFileInfoList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QFileInfoList) Detach_helper0() {
	// QFileInfoList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QFileInfoList) Detach_helper1() {
	// QFileInfoList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QFileInfoList) Dealloc0() {
	// QFileInfoList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QFileInfoList) Node_construct0() {
	// QFileInfoList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QFileInfoList) Node_destruct0() {
	// QFileInfoList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QFileInfoList) Node_copy0() {
	// QFileInfoList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QFileInfoList) Node_destruct1() {
	// QFileInfoList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QFileInfoList) IsValidIterator0() bool {
	// QFileInfoList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QFileInfoList) Op_eq_impl0() bool {
	// QFileInfoList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QFileInfoList) Op_eq_impl1() bool {
	// QFileInfoList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QFileInfoList) Contains_impl0() bool {
	// QFileInfoList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QFileInfoList) Contains_impl1() bool {
	// QFileInfoList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QFileInfoList) Count_impl0() int {
	// QFileInfoList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QFileInfoList) Count_impl1() int {
	// QFileInfoList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QFileInfoList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
