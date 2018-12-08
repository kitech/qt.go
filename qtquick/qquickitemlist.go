package qtquick

// /usr/include/qt/QtQuick/qquickitem.h
// #include <qquickitem.h>
// #include <QtQuick>

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
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end

//  body block begin
type QQuickItemList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QQuickItemList) Operator_equal0() *QQuickItemList {
	// QQuickItemList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QQuickItemList) Operator_equal1() *QQuickItemList {
	// QQuickItemList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QQuickItemList) Swap0() {
	// QQuickItemList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QQuickItemList) Operator_equal_equal0() bool {
	// QQuickItemList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QQuickItemList) Operator_not_equal0() bool {
	// QQuickItemList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QQuickItemList) Size0() int {
	// QQuickItemList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QQuickItemList) Detach0() {
	// QQuickItemList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QQuickItemList) DetachShared0() {
	// QQuickItemList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QQuickItemList) IsDetached0() bool {
	// QQuickItemList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QQuickItemList) SetSharable0() {
	// QQuickItemList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QQuickItemList) IsSharedWith0() bool {
	// QQuickItemList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QQuickItemList) IsEmpty0() bool {
	// QQuickItemList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QQuickItemList) Clear0() {
	// QQuickItemList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QQuickItemList) At0() *QQuickItem {
	// QQuickItemList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// const T & operator[](int)
func (this *QQuickItemList) Operator_get_index0() *QQuickItem {
	// QQuickItemList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// T & operator[](int)
func (this *QQuickItemList) Operator_get_index1() *QQuickItem {
	// QQuickItemList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// void reserve(int)
func (this *QQuickItemList) Reserve0() {
	// QQuickItemList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QQuickItemList) Append0() {
	// QQuickItemList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QQuickItemList) Append1() {
	// QQuickItemList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QQuickItemList) Prepend0() {
	// QQuickItemList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QQuickItemList) Insert0() {
	// QQuickItemList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QQuickItemList) Replace0() {
	// QQuickItemList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QQuickItemList) RemoveAt0() {
	// QQuickItemList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QQuickItemList) RemoveAll0() int {
	// QQuickItemList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QQuickItemList) RemoveOne0() bool {
	// QQuickItemList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QQuickItemList) TakeAt0() *QQuickItem {
	// QQuickItemList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// T takeFirst()
func (this *QQuickItemList) TakeFirst0() *QQuickItem {
	// QQuickItemList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// T takeLast()
func (this *QQuickItemList) TakeLast0() *QQuickItem {
	// QQuickItemList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// void move(int, int)
func (this *QQuickItemList) Move0() {
	// QQuickItemList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QQuickItemList) Swap1() {
	// QQuickItemList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QQuickItemList) IndexOf0() int {
	// QQuickItemList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QQuickItemList) LastIndexOf0() int {
	// QQuickItemList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QQuickItemList) Contains0() bool {
	// QQuickItemList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QQuickItemList) Count0() int {
	// QQuickItemList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QQuickItemList) Begin0() {
	// QQuickItemList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QQuickItemList) Begin1() {
	// QQuickItemList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QQuickItemList) Cbegin0() {
	// QQuickItemList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QQuickItemList) ConstBegin0() {
	// QQuickItemList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QQuickItemList) End0() {
	// QQuickItemList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QQuickItemList) End1() {
	// QQuickItemList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QQuickItemList) Cend0() {
	// QQuickItemList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QQuickItemList) ConstEnd0() {
	// QQuickItemList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QQuickItemList) Rbegin0() {
	// QQuickItemList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QQuickItemList) Rend0() {
	// QQuickItemList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QQuickItemList) Rbegin1() {
	// QQuickItemList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QQuickItemList) Rend1() {
	// QQuickItemList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QQuickItemList) Crbegin0() {
	// QQuickItemList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QQuickItemList) Crend0() {
	// QQuickItemList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QQuickItemList) Insert1() {
	// QQuickItemList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QQuickItemList) Erase0() {
	// QQuickItemList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QQuickItemList) Erase1() {
	// QQuickItemList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QQuickItemList) Count1() int {
	// QQuickItemList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QQuickItemList) Length0() int {
	// QQuickItemList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QQuickItemList) First0() *QQuickItem {
	// QQuickItemList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// const T & constFirst()
func (this *QQuickItemList) ConstFirst0() *QQuickItem {
	// QQuickItemList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// const T & first()
func (this *QQuickItemList) First1() *QQuickItem {
	// QQuickItemList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// T & last()
func (this *QQuickItemList) Last0() *QQuickItem {
	// QQuickItemList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// const T & last()
func (this *QQuickItemList) Last1() *QQuickItem {
	// QQuickItemList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// const T & constLast()
func (this *QQuickItemList) ConstLast0() *QQuickItem {
	// QQuickItemList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// void removeFirst()
func (this *QQuickItemList) RemoveFirst0() {
	// QQuickItemList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QQuickItemList) RemoveLast0() {
	// QQuickItemList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QQuickItemList) StartsWith0() bool {
	// QQuickItemList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QQuickItemList) EndsWith0() bool {
	// QQuickItemList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QQuickItemList) Mid0() *QQuickItemList {
	// QQuickItemList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QQuickItemList) Value0() *QQuickItem {
	// QQuickItemList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// T value(int, const T &)
func (this *QQuickItemList) Value1() *QQuickItem {
	// QQuickItemList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// void push_back(const T &)
func (this *QQuickItemList) Push_back0() {
	// QQuickItemList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QQuickItemList) Push_front0() {
	// QQuickItemList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QQuickItemList) Front0() *QQuickItem {
	// QQuickItemList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// const T & front()
func (this *QQuickItemList) Front1() *QQuickItem {
	// QQuickItemList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// T & back()
func (this *QQuickItemList) Back0() *QQuickItem {
	// QQuickItemList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// const T & back()
func (this *QQuickItemList) Back1() *QQuickItem {
	// QQuickItemList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QQuickItem{}
}

// void pop_front()
func (this *QQuickItemList) Pop_front0() {
	// QQuickItemList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QQuickItemList) Pop_back0() {
	// QQuickItemList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QQuickItemList) Empty0() bool {
	// QQuickItemList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QQuickItemList) Operator_add_equal0() *QQuickItemList {
	// QQuickItemList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QQuickItemList) Operator_add0() *QQuickItemList {
	// QQuickItemList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QQuickItemList) Operator_add_equal1() *QQuickItemList {
	// QQuickItemList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QQuickItemList) Operator_left_shift0() *QQuickItemList {
	// QQuickItemList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QQuickItemList) Operator_left_shift1() *QQuickItemList {
	// QQuickItemList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QQuickItemList) ToVector0() {
	// QQuickItemList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QQuickItemList) ToSet0() {
	// QQuickItemList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QQuickItemList) FromVector0() *QQuickItemList {
	// QQuickItemList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QQuickItemList) FromSet0() *QQuickItemList {
	// QQuickItemList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QQuickItemList) FromStdList0() *QQuickItemList {
	// QQuickItemList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QQuickItemList) ToStdList0() {
	// QQuickItemList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QQuickItemList) Detach_helper_grow0() {
	// QQuickItemList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QQuickItemList) Detach_helper0() {
	// QQuickItemList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QQuickItemList) Detach_helper1() {
	// QQuickItemList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QQuickItemList) Dealloc0() {
	// QQuickItemList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QQuickItemList) Node_construct0() {
	// QQuickItemList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QQuickItemList) Node_destruct0() {
	// QQuickItemList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QQuickItemList) Node_copy0() {
	// QQuickItemList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QQuickItemList) Node_destruct1() {
	// QQuickItemList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QQuickItemList) IsValidIterator0() bool {
	// QQuickItemList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QQuickItemList) Op_eq_impl0() bool {
	// QQuickItemList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QQuickItemList) Op_eq_impl1() bool {
	// QQuickItemList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QQuickItemList) Contains_impl0() bool {
	// QQuickItemList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QQuickItemList) Contains_impl1() bool {
	// QQuickItemList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QQuickItemList) Count_impl0() int {
	// QQuickItemList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QQuickItemList) Count_impl1() int {
	// QQuickItemList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QQuickItemList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
