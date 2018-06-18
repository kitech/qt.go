package qtwidgets

// /usr/include/qt/QtWidgets/qtreewidget.h
// #include <qtreewidget.h>
// #include <QtWidgets>

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
}

//  keep block end

//  body block begin
type QTreeWidgetItemList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QTreeWidgetItemList) Operator_equal_0() *QTreeWidgetItemList {
	// QTreeWidgetItemList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QTreeWidgetItemList) Operator_equal_1() *QTreeWidgetItemList {
	// QTreeWidgetItemList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QTreeWidgetItemList) Swap_0() {
	// QTreeWidgetItemList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QTreeWidgetItemList) Operator_equal_equal_0() bool {
	// QTreeWidgetItemList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QTreeWidgetItemList) Operator_not_equal_0() bool {
	// QTreeWidgetItemList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QTreeWidgetItemList) Size_0() int {
	// QTreeWidgetItemList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QTreeWidgetItemList) Detach_0() {
	// QTreeWidgetItemList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QTreeWidgetItemList) DetachShared_0() {
	// QTreeWidgetItemList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QTreeWidgetItemList) IsDetached_0() bool {
	// QTreeWidgetItemList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QTreeWidgetItemList) SetSharable_0() {
	// QTreeWidgetItemList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QTreeWidgetItemList) IsSharedWith_0() bool {
	// QTreeWidgetItemList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QTreeWidgetItemList) IsEmpty_0() bool {
	// QTreeWidgetItemList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QTreeWidgetItemList) Clear_0() {
	// QTreeWidgetItemList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QTreeWidgetItemList) At_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// const T & operator[](int)
func (this *QTreeWidgetItemList) Operator_get_index_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// T & operator[](int)
func (this *QTreeWidgetItemList) Operator_get_index_1() *QTreeWidgetItem {
	// QTreeWidgetItemList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// void reserve(int)
func (this *QTreeWidgetItemList) Reserve_0() {
	// QTreeWidgetItemList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QTreeWidgetItemList) Append_0() {
	// QTreeWidgetItemList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QTreeWidgetItemList) Append_1() {
	// QTreeWidgetItemList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QTreeWidgetItemList) Prepend_0() {
	// QTreeWidgetItemList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QTreeWidgetItemList) Insert_0() {
	// QTreeWidgetItemList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QTreeWidgetItemList) Replace_0() {
	// QTreeWidgetItemList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QTreeWidgetItemList) RemoveAt_0() {
	// QTreeWidgetItemList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QTreeWidgetItemList) RemoveAll_0() int {
	// QTreeWidgetItemList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QTreeWidgetItemList) RemoveOne_0() bool {
	// QTreeWidgetItemList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QTreeWidgetItemList) TakeAt_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// T takeFirst()
func (this *QTreeWidgetItemList) TakeFirst_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// T takeLast()
func (this *QTreeWidgetItemList) TakeLast_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// void move(int, int)
func (this *QTreeWidgetItemList) Move_0() {
	// QTreeWidgetItemList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QTreeWidgetItemList) Swap_1() {
	// QTreeWidgetItemList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QTreeWidgetItemList) IndexOf_0() int {
	// QTreeWidgetItemList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QTreeWidgetItemList) LastIndexOf_0() int {
	// QTreeWidgetItemList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QTreeWidgetItemList) Contains_0() bool {
	// QTreeWidgetItemList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QTreeWidgetItemList) Count_0() int {
	// QTreeWidgetItemList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QTreeWidgetItemList) Begin_0() {
	// QTreeWidgetItemList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QTreeWidgetItemList) Begin_1() {
	// QTreeWidgetItemList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QTreeWidgetItemList) Cbegin_0() {
	// QTreeWidgetItemList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QTreeWidgetItemList) ConstBegin_0() {
	// QTreeWidgetItemList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QTreeWidgetItemList) End_0() {
	// QTreeWidgetItemList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QTreeWidgetItemList) End_1() {
	// QTreeWidgetItemList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QTreeWidgetItemList) Cend_0() {
	// QTreeWidgetItemList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QTreeWidgetItemList) ConstEnd_0() {
	// QTreeWidgetItemList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QTreeWidgetItemList) Rbegin_0() {
	// QTreeWidgetItemList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QTreeWidgetItemList) Rend_0() {
	// QTreeWidgetItemList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QTreeWidgetItemList) Rbegin_1() {
	// QTreeWidgetItemList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QTreeWidgetItemList) Rend_1() {
	// QTreeWidgetItemList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QTreeWidgetItemList) Crbegin_0() {
	// QTreeWidgetItemList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QTreeWidgetItemList) Crend_0() {
	// QTreeWidgetItemList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QTreeWidgetItemList) Insert_1() {
	// QTreeWidgetItemList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QTreeWidgetItemList) Erase_0() {
	// QTreeWidgetItemList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QTreeWidgetItemList) Erase_1() {
	// QTreeWidgetItemList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QTreeWidgetItemList) Count_1() int {
	// QTreeWidgetItemList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QTreeWidgetItemList) Length_0() int {
	// QTreeWidgetItemList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QTreeWidgetItemList) First_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// const T & constFirst()
func (this *QTreeWidgetItemList) ConstFirst_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// const T & first()
func (this *QTreeWidgetItemList) First_1() *QTreeWidgetItem {
	// QTreeWidgetItemList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// T & last()
func (this *QTreeWidgetItemList) Last_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// const T & last()
func (this *QTreeWidgetItemList) Last_1() *QTreeWidgetItem {
	// QTreeWidgetItemList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// const T & constLast()
func (this *QTreeWidgetItemList) ConstLast_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// void removeFirst()
func (this *QTreeWidgetItemList) RemoveFirst_0() {
	// QTreeWidgetItemList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QTreeWidgetItemList) RemoveLast_0() {
	// QTreeWidgetItemList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QTreeWidgetItemList) StartsWith_0() bool {
	// QTreeWidgetItemList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QTreeWidgetItemList) EndsWith_0() bool {
	// QTreeWidgetItemList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QTreeWidgetItemList) Mid_0() *QTreeWidgetItemList {
	// QTreeWidgetItemList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QTreeWidgetItemList) Value_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// T value(int, const T &)
func (this *QTreeWidgetItemList) Value_1() *QTreeWidgetItem {
	// QTreeWidgetItemList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// void push_back(const T &)
func (this *QTreeWidgetItemList) Push_back_0() {
	// QTreeWidgetItemList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QTreeWidgetItemList) Push_front_0() {
	// QTreeWidgetItemList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QTreeWidgetItemList) Front_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// const T & front()
func (this *QTreeWidgetItemList) Front_1() *QTreeWidgetItem {
	// QTreeWidgetItemList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// T & back()
func (this *QTreeWidgetItemList) Back_0() *QTreeWidgetItem {
	// QTreeWidgetItemList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// const T & back()
func (this *QTreeWidgetItemList) Back_1() *QTreeWidgetItem {
	// QTreeWidgetItemList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTreeWidgetItem{}
}

// void pop_front()
func (this *QTreeWidgetItemList) Pop_front_0() {
	// QTreeWidgetItemList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QTreeWidgetItemList) Pop_back_0() {
	// QTreeWidgetItemList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QTreeWidgetItemList) Empty_0() bool {
	// QTreeWidgetItemList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QTreeWidgetItemList) Operator_add_equal_0() *QTreeWidgetItemList {
	// QTreeWidgetItemList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QTreeWidgetItemList) Operator_add_0() *QTreeWidgetItemList {
	// QTreeWidgetItemList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QTreeWidgetItemList) Operator_add_equal_1() *QTreeWidgetItemList {
	// QTreeWidgetItemList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QTreeWidgetItemList) Operator_left_shift_0() *QTreeWidgetItemList {
	// QTreeWidgetItemList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QTreeWidgetItemList) Operator_left_shift_1() *QTreeWidgetItemList {
	// QTreeWidgetItemList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QTreeWidgetItemList) ToVector_0() {
	// QTreeWidgetItemList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QTreeWidgetItemList) ToSet_0() {
	// QTreeWidgetItemList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QTreeWidgetItemList) FromVector_0() *QTreeWidgetItemList {
	// QTreeWidgetItemList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QTreeWidgetItemList) FromSet_0() *QTreeWidgetItemList {
	// QTreeWidgetItemList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QTreeWidgetItemList) FromStdList_0() *QTreeWidgetItemList {
	// QTreeWidgetItemList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QTreeWidgetItemList) ToStdList_0() {
	// QTreeWidgetItemList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QTreeWidgetItemList) Detach_helper_grow_0() {
	// QTreeWidgetItemList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QTreeWidgetItemList) Detach_helper_0() {
	// QTreeWidgetItemList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QTreeWidgetItemList) Detach_helper_1() {
	// QTreeWidgetItemList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QTreeWidgetItemList) Dealloc_0() {
	// QTreeWidgetItemList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QTreeWidgetItemList) Node_construct_0() {
	// QTreeWidgetItemList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QTreeWidgetItemList) Node_destruct_0() {
	// QTreeWidgetItemList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QTreeWidgetItemList) Node_copy_0() {
	// QTreeWidgetItemList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QTreeWidgetItemList) Node_destruct_1() {
	// QTreeWidgetItemList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QTreeWidgetItemList) IsValidIterator_0() bool {
	// QTreeWidgetItemList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QTreeWidgetItemList) Op_eq_impl_0() bool {
	// QTreeWidgetItemList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QTreeWidgetItemList) Op_eq_impl_1() bool {
	// QTreeWidgetItemList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QTreeWidgetItemList) Contains_impl_0() bool {
	// QTreeWidgetItemList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QTreeWidgetItemList) Contains_impl_1() bool {
	// QTreeWidgetItemList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QTreeWidgetItemList) Count_impl_0() int {
	// QTreeWidgetItemList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QTreeWidgetItemList) Count_impl_1() int {
	// QTreeWidgetItemList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QTreeWidgetItemList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
