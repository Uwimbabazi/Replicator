package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"math/rand"
	. "github.com/pspaces/gospace"
	. "github.com/repligospaces"
)

var uri = make(map[Space]string)
var Sp = make(map[string]*Space)
var rsp Replispace = Replispace{Sp: Sp}

var debug bool
var l sync.Mutex

var writes_local int
var writes_remote int
var reads_local int
var reads_remote int
var writes_replicated int
var writestotal int
var reads_success int
var reads_insuccess int
var writes_extra int
var wg sync.WaitGroup

var s1 Space
var s2 Space
var s3 Space
var s4 Space
var s5 Space
var s6 Space
var s7 Space
var s8 Space
var s9 Space
var s10 Space
var s11 Space
var s12 Space
var s13 Space
var s14 Space
var s15 Space
var s16 Space
var s17 Space
var s18 Space
var s19 Space
var s20 Space
var s21 Space
var s22 Space
var s23 Space
var s24 Space
var s25 Space
var s26 Space
var s27 Space
var s28 Space
var s29 Space
var s30 Space
var s31 Space
var s32 Space

func main() {
	debug = false
	writes_local = 0
	writes_remote = 0
	reads_local = 0
	reads_remote = 0
	writes_replicated = Getwcount()-writes_local
	reads_success = 0
	reads_insuccess = 0
	rand.Seed(time.Now().UTC().UnixNano())

	wg.Add(32)

	s1 = NewSpace("tcp://localhost:34001/s1")
	Sp["tcp://localhost:34001/s1"] = &s1
	uri[s1] = "tcp://localhost:34001/s1"
	s2 = NewSpace("tcp://localhost:34002/s2")
	Sp["tcp://localhost:34002/s2"] = &s2
	uri[s2] = "tcp://localhost:34002/s2"
	s3 = NewSpace("tcp://localhost:34003/s3")
	Sp["tcp://localhost:34003/s3"] = &s3
	uri[s3] = "tcp://localhost:34003/s3"
	s4 = NewSpace("tcp://localhost:34004/s4")
	Sp["tcp://localhost:34004/s4"] = &s4
	uri[s4] = "tcp://localhost:34004/s4"
	s5 = NewSpace("tcp://localhost:34005/s5")
	Sp["tcp://localhost:34005/s5"] = &s5
	uri[s5] = "tcp://localhost:34005/s5"
	s6 = NewSpace("tcp://localhost:34006/s6")
	Sp["tcp://localhost:34006/s6"] = &s6
	uri[s6] = "tcp://localhost:34006/s6"
	s7 = NewSpace("tcp://localhost:34007/s7")
	Sp["tcp://localhost:34007/s7"] = &s7
	uri[s7] = "tcp://localhost:34007/s7"
	s8 = NewSpace("tcp://localhost:34008/s8")
	Sp["tcp://localhost:34008/s8"] = &s8
	uri[s8] = "tcp://localhost:34008/s8"
	s9 = NewSpace("tcp://localhost:34009/s9")
	Sp["tcp://localhost:34009/s9"] = &s9
	uri[s9] = "tcp://localhost:34009/s9"
	s10 = NewSpace("tcp://localhost:34010/s10")
	Sp["tcp://localhost:34010/s10"] = &s10
	uri[s10] = "tcp://localhost:34010/s10"
	s11 = NewSpace("tcp://localhost:34011/s11")
	Sp["tcp://localhost:34011/s11"] = &s11
	uri[s11] = "tcp://localhost:34011/s11"
	s12 = NewSpace("tcp://localhost:34012/s12")
	Sp["tcp://localhost:34012/s12"] = &s12
	uri[s12] = "tcp://localhost:34012/s12"
	s13 = NewSpace("tcp://localhost:34013/s13")
	Sp["tcp://localhost:34013/s13"] = &s13
	uri[s13] = "tcp://localhost:34013/s13"
	s14 = NewSpace("tcp://localhost:34014/s14")
	Sp["tcp://localhost:34014/s14"] = &s14
	uri[s14] = "tcp://localhost:34014/s14"
	s15 = NewSpace("tcp://localhost:34015/s15")
	Sp["tcp://localhost:34015/s15"] = &s15
	uri[s15] = "tcp://localhost:34015/s15"
	s16 = NewSpace("tcp://localhost:34016/s16")
	Sp["tcp://localhost:34016/s16"] = &s16
	uri[s16] = "tcp://localhost:34016/s16"
	s17 = NewSpace("tcp://localhost:34017/s17")
	Sp["tcp://localhost:34017/s17"] = &s17
	uri[s17] = "tcp://localhost:34017/s17"
	s18 = NewSpace("tcp://localhost:34018/s18")
	Sp["tcp://localhost:34018/s18"] = &s18
	uri[s18] = "tcp://localhost:34018/s18"
	s19 = NewSpace("tcp://localhost:34019/s19")
	Sp["tcp://localhost:34019/s19"] = &s19
	uri[s19] = "tcp://localhost:34019/s19"
	s20 = NewSpace("tcp://localhost:34020/s20")
	Sp["tcp://localhost:34020/s20"] = &s20
	uri[s20] = "tcp://localhost:34020/s20"
	s21 = NewSpace("tcp://localhost:34021/s21")
	Sp["tcp://localhost:34021/s21"] = &s21
	uri[s21] = "tcp://localhost:34021/s21"
	s22 = NewSpace("tcp://localhost:34022/s22")
	Sp["tcp://localhost:34022/s22"] = &s22
	uri[s22] = "tcp://localhost:34022/s22"
	s23 = NewSpace("tcp://localhost:34023/s23")
	Sp["tcp://localhost:34023/s23"] = &s23
	uri[s23] = "tcp://localhost:34023/s23"
	s24 = NewSpace("tcp://localhost:34024/s24")
	Sp["tcp://localhost:34024/s24"] = &s24
	uri[s24] = "tcp://localhost:34024/s24"
	s25 = NewSpace("tcp://localhost:34025/s25")
	Sp["tcp://localhost:34025/s25"] = &s25
	uri[s25] = "tcp://localhost:34025/s25"
	s26 = NewSpace("tcp://localhost:34026/s26")
	Sp["tcp://localhost:34026/s26"] = &s26
	uri[s26] = "tcp://localhost:34026/s26"
	s27 = NewSpace("tcp://localhost:34027/s27")
	Sp["tcp://localhost:34027/s27"] = &s27
	uri[s27] = "tcp://localhost:34027/s27"
	s28 = NewSpace("tcp://localhost:34028/s28")
	Sp["tcp://localhost:34028/s28"] = &s28
	uri[s28] = "tcp://localhost:34028/s28"
	s29 = NewSpace("tcp://localhost:34029/s29")
	Sp["tcp://localhost:34029/s29"] = &s29
	uri[s29] = "tcp://localhost:34029/s29"
	s30 = NewSpace("tcp://localhost:34030/s30")
	Sp["tcp://localhost:34030/s30"] = &s30
	uri[s30] = "tcp://localhost:34030/s30"
	s31 = NewSpace("tcp://localhost:34031/s31")
	Sp["tcp://localhost:34031/s31"] = &s31
	uri[s31] = "tcp://localhost:34031/s31"
	s32 = NewSpace("tcp://localhost:34032/s32")
	Sp["tcp://localhost:34032/s32"] = &s32
	uri[s32] = "tcp://localhost:34032/s32"

	go p1()
	go p2()
	go p3()
	go p4()
	go p5()
	go p6()
	go p7()
	go p8()
	go p9()
	go p10()
	go p11()
	go p12()
	go p13()
	go p14()
	go p15()
	go p16()
	go p17()
	go p18()
	go p19()
	go p20()
	go p21()
	go p22()
	go p23()
	go p24()
	go p25()
	go p26()
	go p27()
	go p28()
	go p29()
	go p30()
	go p31()
	go p32()

	wg.Wait()

	writes_replicated = Getwcount()-writes_local	// these include the normal put too

	fmt.Println("       loc w,    rem w,   repl w,    loc r,    rem r,    tot w,   succ r,   fail r")
	fmt.Printf("    %8d, %8d, %8d, %8d, %8d, %8d, %8d, %8d\n", writes_local, writes_remote, writes_replicated, reads_local, reads_remote, writestotal, reads_success, reads_insuccess)
}

func p1() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s16]
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s22]
	targets7 := make([]string, 1)
	targets7[0] = uri[s24]
	targets6 := make([]string, 1)
	targets6[0] = uri[s7]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s16]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(147, value), rsp, targets0)
	log("p1 w:%d", value)
	countwrites(value, 1, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(31, value), rsp, targets1)
	log("p1 w:%d", value)
	countwrites(value, 1, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(7, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(248, value), rsp, targets2)
	log("p1 w:%d", value)
	countwrites(value, 1, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(155, value), rsp, targets3)
	log("p1 w:%d", value)
	countwrites(value, 1, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(104, value), rsp, targets4)
	log("p1 w:%d", value)
	countwrites(value, 1, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(236, value), rsp, targets5)
	log("p1 w:%d", value)
	countwrites(value, 1, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(4, value), rsp, targets6)
	log("p1 w:%d", value)
	countwrites(value, 1, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(34, &value), rsp, s1)
	log("p1 r:%d", value)
	countreads(value, 1, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(117, value), rsp, targets7)
	log("p1 w:%d", value)
	countwrites(value, 1, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(197, value), rsp, targets8)
	log("p1 w:%d", value)
	countwrites(value, 1, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(140, value), rsp, targets9)
	log("p1 w:%d", value)
	countwrites(value, 1, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(137, value), rsp, targets10)
	log("p1 w:%d", value)
	countwrites(value, 1, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(106, value), rsp, targets11)
	log("p1 w:%d", value)
	countwrites(value, 1, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(88, value), rsp, targets12)
	log("p1 w:%d", value)
	countwrites(value, 1, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(44, value), rsp, targets13)
	log("p1 w:%d", value)
	countwrites(value, 1, 6)
	l.Unlock()
}

func p2() {
	targets15 := make([]string, 0)
	targets14 := make([]string, 0)
	targets13 := make([]string, 1)
	targets13[0] = uri[s11]
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(173, value), rsp, targets0)
	log("p2 w:%d", value)
	countwrites(value, 2, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(77, value), rsp, targets1)
	log("p2 w:%d", value)
	countwrites(value, 2, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(80, value), rsp, targets2)
	log("p2 w:%d", value)
	countwrites(value, 2, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(204, value), rsp, targets3)
	log("p2 w:%d", value)
	countwrites(value, 2, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(218, value), rsp, targets4)
	log("p2 w:%d", value)
	countwrites(value, 2, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(3, value), rsp, targets5)
	log("p2 w:%d", value)
	countwrites(value, 2, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(199, value), rsp, targets6)
	log("p2 w:%d", value)
	countwrites(value, 2, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(37, value), rsp, targets7)
	log("p2 w:%d", value)
	countwrites(value, 2, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(168, value), rsp, targets8)
	log("p2 w:%d", value)
	countwrites(value, 2, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(77, value), rsp, targets9)
	log("p2 w:%d", value)
	countwrites(value, 2, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(220, value), rsp, targets10)
	log("p2 w:%d", value)
	countwrites(value, 2, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(217, value), rsp, targets11)
	log("p2 w:%d", value)
	countwrites(value, 2, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(28, value), rsp, targets12)
	log("p2 w:%d", value)
	countwrites(value, 2, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(130, value), rsp, targets13)
	log("p2 w:%d", value)
	countwrites(value, 2, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(115, value), rsp, targets14)
	log("p2 w:%d", value)
	countwrites(value, 2, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(87, value), rsp, targets15)
	log("p2 w:%d", value)
	countwrites(value, 2, 11)
	l.Unlock()
}

func p3() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 2)
	targets6[1] = uri[s14]
	targets6[0] = uri[s11]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(138, value), rsp, targets0)
	log("p3 w:%d", value)
	countwrites(value, 3, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(43, value), rsp, targets1)
	log("p3 w:%d", value)
	countwrites(value, 3, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(107, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(19, value), rsp, targets2)
	log("p3 w:%d", value)
	countwrites(value, 3, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(44, value), rsp, targets3)
	log("p3 w:%d", value)
	countwrites(value, 3, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(148, value), rsp, targets4)
	log("p3 w:%d", value)
	countwrites(value, 3, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(84, value), rsp, targets5)
	log("p3 w:%d", value)
	countwrites(value, 3, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(47, value), rsp, targets6)
	log("p3 w:%d", value)
	countwrites(value, 3, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(214, value), rsp, targets7)
	log("p3 w:%d", value)
	countwrites(value, 3, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(125, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(81, value), rsp, targets8)
	log("p3 w:%d", value)
	countwrites(value, 3, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(160, value), rsp, targets9)
	log("p3 w:%d", value)
	countwrites(value, 3, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(16, value), rsp, targets10)
	log("p3 w:%d", value)
	countwrites(value, 3, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(145, value), rsp, targets11)
	log("p3 w:%d", value)
	countwrites(value, 3, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(226, value), rsp, targets12)
	log("p3 w:%d", value)
	countwrites(value, 3, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(55, &value), rsp, s3)
	log("p3 r:%d", value)
	countreads(value, 3, 7)
	l.Unlock()
}

func p4() {
	targets15 := make([]string, 0)
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s10]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 2)
	targets4[1] = uri[s28]
	targets4[0] = uri[s8]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(126, value), rsp, targets0)
	log("p4 w:%d", value)
	countwrites(value, 4, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(231, value), rsp, targets1)
	log("p4 w:%d", value)
	countwrites(value, 4, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(159, value), rsp, targets2)
	log("p4 w:%d", value)
	countwrites(value, 4, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(230, value), rsp, targets3)
	log("p4 w:%d", value)
	countwrites(value, 4, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(165, value), rsp, targets4)
	log("p4 w:%d", value)
	countwrites(value, 4, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(68, value), rsp, targets5)
	log("p4 w:%d", value)
	countwrites(value, 4, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(175, value), rsp, targets6)
	log("p4 w:%d", value)
	countwrites(value, 4, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(132, value), rsp, targets7)
	log("p4 w:%d", value)
	countwrites(value, 4, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(223, value), rsp, targets8)
	log("p4 w:%d", value)
	countwrites(value, 4, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(96, value), rsp, targets9)
	log("p4 w:%d", value)
	countwrites(value, 4, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(119, value), rsp, targets10)
	log("p4 w:%d", value)
	countwrites(value, 4, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(100, value), rsp, targets11)
	log("p4 w:%d", value)
	countwrites(value, 4, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(108, value), rsp, targets12)
	log("p4 w:%d", value)
	countwrites(value, 4, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(87, value), rsp, targets13)
	log("p4 w:%d", value)
	countwrites(value, 4, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(99, value), rsp, targets14)
	log("p4 w:%d", value)
	countwrites(value, 4, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(186, value), rsp, targets15)
	log("p4 w:%d", value)
	countwrites(value, 4, 24)
	l.Unlock()
}

func p5() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s17]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s8]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(123, value), rsp, targets0)
	log("p5 w:%d", value)
	countwrites(value, 5, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(230, value), rsp, targets1)
	log("p5 w:%d", value)
	countwrites(value, 5, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(81, value), rsp, targets2)
	log("p5 w:%d", value)
	countwrites(value, 5, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(79, value), rsp, targets3)
	log("p5 w:%d", value)
	countwrites(value, 5, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(251, value), rsp, targets4)
	log("p5 w:%d", value)
	countwrites(value, 5, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(161, value), rsp, targets5)
	log("p5 w:%d", value)
	countwrites(value, 5, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(180, value), rsp, targets6)
	log("p5 w:%d", value)
	countwrites(value, 5, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(15, value), rsp, targets7)
	log("p5 w:%d", value)
	countwrites(value, 5, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(138, value), rsp, targets8)
	log("p5 w:%d", value)
	countwrites(value, 5, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(66, value), rsp, targets9)
	log("p5 w:%d", value)
	countwrites(value, 5, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(205, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(139, value), rsp, targets10)
	log("p5 w:%d", value)
	countwrites(value, 5, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(238, value), rsp, targets11)
	log("p5 w:%d", value)
	countwrites(value, 5, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(113, value), rsp, targets12)
	log("p5 w:%d", value)
	countwrites(value, 5, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(97, &value), rsp, s5)
	log("p5 r:%d", value)
	countreads(value, 5, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(93, value), rsp, targets13)
	log("p5 w:%d", value)
	countwrites(value, 5, 12)
	l.Unlock()
}

func p6() {
	targets14 := make([]string, 1)
	targets14[0] = uri[s9]
	targets13 := make([]string, 1)
	targets13[0] = uri[s31]
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s9]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(217, value), rsp, targets0)
	log("p6 w:%d", value)
	countwrites(value, 6, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(75, value), rsp, targets1)
	log("p6 w:%d", value)
	countwrites(value, 6, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(126, value), rsp, targets2)
	log("p6 w:%d", value)
	countwrites(value, 6, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(142, value), rsp, targets3)
	log("p6 w:%d", value)
	countwrites(value, 6, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(171, value), rsp, targets4)
	log("p6 w:%d", value)
	countwrites(value, 6, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(96, value), rsp, targets5)
	log("p6 w:%d", value)
	countwrites(value, 6, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(114, value), rsp, targets6)
	log("p6 w:%d", value)
	countwrites(value, 6, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(68, value), rsp, targets7)
	log("p6 w:%d", value)
	countwrites(value, 6, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(220, value), rsp, targets8)
	log("p6 w:%d", value)
	countwrites(value, 6, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(246, value), rsp, targets9)
	log("p6 w:%d", value)
	countwrites(value, 6, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(218, value), rsp, targets10)
	log("p6 w:%d", value)
	countwrites(value, 6, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(19, value), rsp, targets11)
	log("p6 w:%d", value)
	countwrites(value, 6, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(139, value), rsp, targets12)
	log("p6 w:%d", value)
	countwrites(value, 6, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(209, &value), rsp, s6)
	log("p6 r:%d", value)
	countreads(value, 6, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(103, value), rsp, targets13)
	log("p6 w:%d", value)
	countwrites(value, 6, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(2, value), rsp, targets14)
	log("p6 w:%d", value)
	countwrites(value, 6, 1)
	l.Unlock()
}

func p7() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s8]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s9]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s17]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(180, value), rsp, targets0)
	log("p7 w:%d", value)
	countwrites(value, 7, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(4, &value), rsp, s7)
	log("p7 r:%d", value)
	countreads(value, 7, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(160, value), rsp, targets1)
	log("p7 w:%d", value)
	countwrites(value, 7, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(177, value), rsp, targets2)
	log("p7 w:%d", value)
	countwrites(value, 7, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(114, value), rsp, targets3)
	log("p7 w:%d", value)
	countwrites(value, 7, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(228, value), rsp, targets4)
	log("p7 w:%d", value)
	countwrites(value, 7, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(66, value), rsp, targets5)
	log("p7 w:%d", value)
	countwrites(value, 7, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(179, value), rsp, targets6)
	log("p7 w:%d", value)
	countwrites(value, 7, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(248, value), rsp, targets7)
	log("p7 w:%d", value)
	countwrites(value, 7, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(246, value), rsp, targets8)
	log("p7 w:%d", value)
	countwrites(value, 7, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(119, value), rsp, targets9)
	log("p7 w:%d", value)
	countwrites(value, 7, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(252, value), rsp, targets10)
	log("p7 w:%d", value)
	countwrites(value, 7, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(239, value), rsp, targets11)
	log("p7 w:%d", value)
	countwrites(value, 7, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(184, value), rsp, targets12)
	log("p7 w:%d", value)
	countwrites(value, 7, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(42, value), rsp, targets13)
	log("p7 w:%d", value)
	countwrites(value, 7, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(79, value), rsp, targets14)
	log("p7 w:%d", value)
	countwrites(value, 7, 10)
	l.Unlock()
}

func p8() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s23]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(67, value), rsp, targets0)
	log("p8 w:%d", value)
	countwrites(value, 8, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(65, value), rsp, targets1)
	log("p8 w:%d", value)
	countwrites(value, 8, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(70, value), rsp, targets2)
	log("p8 w:%d", value)
	countwrites(value, 8, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(71, value), rsp, targets3)
	log("p8 w:%d", value)
	countwrites(value, 8, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(166, value), rsp, targets4)
	log("p8 w:%d", value)
	countwrites(value, 8, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(18, value), rsp, targets5)
	log("p8 w:%d", value)
	countwrites(value, 8, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(199, value), rsp, targets6)
	log("p8 w:%d", value)
	countwrites(value, 8, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(48, value), rsp, targets7)
	log("p8 w:%d", value)
	countwrites(value, 8, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(98, value), rsp, targets8)
	log("p8 w:%d", value)
	countwrites(value, 8, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(239, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(253, value), rsp, targets9)
	log("p8 w:%d", value)
	countwrites(value, 8, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(165, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(127, value), rsp, targets10)
	log("p8 w:%d", value)
	countwrites(value, 8, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(235, value), rsp, targets11)
	log("p8 w:%d", value)
	countwrites(value, 8, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(123, &value), rsp, s8)
	log("p8 r:%d", value)
	countreads(value, 8, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(191, value), rsp, targets12)
	log("p8 w:%d", value)
	countwrites(value, 8, 24)
	l.Unlock()
}

func p9() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s23]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(2, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(178, value), rsp, targets0)
	log("p9 w:%d", value)
	countwrites(value, 9, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(105, value), rsp, targets1)
	log("p9 w:%d", value)
	countwrites(value, 9, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(33, value), rsp, targets2)
	log("p9 w:%d", value)
	countwrites(value, 9, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(142, value), rsp, targets3)
	log("p9 w:%d", value)
	countwrites(value, 9, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(246, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(164, value), rsp, targets4)
	log("p9 w:%d", value)
	countwrites(value, 9, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(98, value), rsp, targets5)
	log("p9 w:%d", value)
	countwrites(value, 9, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(104, value), rsp, targets6)
	log("p9 w:%d", value)
	countwrites(value, 9, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(41, value), rsp, targets7)
	log("p9 w:%d", value)
	countwrites(value, 9, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(238, value), rsp, targets8)
	log("p9 w:%d", value)
	countwrites(value, 9, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(23, value), rsp, targets9)
	log("p9 w:%d", value)
	countwrites(value, 9, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(185, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(82, value), rsp, targets10)
	log("p9 w:%d", value)
	countwrites(value, 9, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(240, value), rsp, targets11)
	log("p9 w:%d", value)
	countwrites(value, 9, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(17, &value), rsp, s9)
	log("p9 r:%d", value)
	countreads(value, 9, 3)
	l.Unlock()
}

func p10() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s24]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s11]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(3, value), rsp, targets0)
	log("p10 w:%d", value)
	countwrites(value, 10, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(57, value), rsp, targets1)
	log("p10 w:%d", value)
	countwrites(value, 10, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(112, value), rsp, targets2)
	log("p10 w:%d", value)
	countwrites(value, 10, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(82, value), rsp, targets3)
	log("p10 w:%d", value)
	countwrites(value, 10, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(186, value), rsp, targets4)
	log("p10 w:%d", value)
	countwrites(value, 10, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(64, value), rsp, targets5)
	log("p10 w:%d", value)
	countwrites(value, 10, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(210, value), rsp, targets6)
	log("p10 w:%d", value)
	countwrites(value, 10, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(8, value), rsp, targets7)
	log("p10 w:%d", value)
	countwrites(value, 10, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(132, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(238, value), rsp, targets8)
	log("p10 w:%d", value)
	countwrites(value, 10, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(173, value), rsp, targets9)
	log("p10 w:%d", value)
	countwrites(value, 10, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(61, value), rsp, targets10)
	log("p10 w:%d", value)
	countwrites(value, 10, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(117, value), rsp, targets11)
	log("p10 w:%d", value)
	countwrites(value, 10, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(174, &value), rsp, s10)
	log("p10 r:%d", value)
	countreads(value, 10, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(41, value), rsp, targets12)
	log("p10 w:%d", value)
	countwrites(value, 10, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(229, value), rsp, targets13)
	log("p10 w:%d", value)
	countwrites(value, 10, 29)
	l.Unlock()
}

func p11() {
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s1]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s1]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(153, value), rsp, targets0)
	log("p11 w:%d", value)
	countwrites(value, 11, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(1, value), rsp, targets1)
	log("p11 w:%d", value)
	countwrites(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(232, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(7, value), rsp, targets2)
	log("p11 w:%d", value)
	countwrites(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(9, value), rsp, targets3)
	log("p11 w:%d", value)
	countwrites(value, 11, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(71, value), rsp, targets4)
	log("p11 w:%d", value)
	countwrites(value, 11, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(112, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(177, value), rsp, targets5)
	log("p11 w:%d", value)
	countwrites(value, 11, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(81, value), rsp, targets6)
	log("p11 w:%d", value)
	countwrites(value, 11, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(213, value), rsp, targets7)
	log("p11 w:%d", value)
	countwrites(value, 11, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(7, value), rsp, targets8)
	log("p11 w:%d", value)
	countwrites(value, 11, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(236, value), rsp, targets9)
	log("p11 w:%d", value)
	countwrites(value, 11, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(47, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(101, value), rsp, targets10)
	log("p11 w:%d", value)
	countwrites(value, 11, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(130, &value), rsp, s11)
	log("p11 r:%d", value)
	countreads(value, 11, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(163, value), rsp, targets11)
	log("p11 w:%d", value)
	countwrites(value, 11, 21)
	l.Unlock()
}

func p12() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 1)
	targets10[0] = uri[s12]
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s10]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(178, value), rsp, targets0)
	log("p12 w:%d", value)
	countwrites(value, 12, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(105, value), rsp, targets1)
	log("p12 w:%d", value)
	countwrites(value, 12, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(176, value), rsp, targets2)
	log("p12 w:%d", value)
	countwrites(value, 12, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(150, &value), rsp, s12)
	log("p12 r:%d", value)
	countreads(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(19, value), rsp, targets3)
	log("p12 w:%d", value)
	countwrites(value, 12, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(135, value), rsp, targets4)
	log("p12 w:%d", value)
	countwrites(value, 12, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(174, value), rsp, targets5)
	log("p12 w:%d", value)
	countwrites(value, 12, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(238, value), rsp, targets6)
	log("p12 w:%d", value)
	countwrites(value, 12, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(235, value), rsp, targets7)
	log("p12 w:%d", value)
	countwrites(value, 12, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(201, value), rsp, targets8)
	log("p12 w:%d", value)
	countwrites(value, 12, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(119, value), rsp, targets9)
	log("p12 w:%d", value)
	countwrites(value, 12, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(150, value), rsp, targets10)
	log("p12 w:%d", value)
	countwrites(value, 12, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(22, value), rsp, targets11)
	log("p12 w:%d", value)
	countwrites(value, 12, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(116, value), rsp, targets12)
	log("p12 w:%d", value)
	countwrites(value, 12, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(59, value), rsp, targets13)
	log("p12 w:%d", value)
	countwrites(value, 12, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(68, value), rsp, targets14)
	log("p12 w:%d", value)
	countwrites(value, 12, 9)
	l.Unlock()
}

func p13() {
	targets14 := make([]string, 1)
	targets14[0] = uri[s9]
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s22]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(100, value), rsp, targets0)
	log("p13 w:%d", value)
	countwrites(value, 13, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(73, value), rsp, targets1)
	log("p13 w:%d", value)
	countwrites(value, 13, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(91, value), rsp, targets2)
	log("p13 w:%d", value)
	countwrites(value, 13, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(94, value), rsp, targets3)
	log("p13 w:%d", value)
	countwrites(value, 13, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(251, value), rsp, targets4)
	log("p13 w:%d", value)
	countwrites(value, 13, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(106, value), rsp, targets5)
	log("p13 w:%d", value)
	countwrites(value, 13, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(77, value), rsp, targets6)
	log("p13 w:%d", value)
	countwrites(value, 13, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(217, value), rsp, targets7)
	log("p13 w:%d", value)
	countwrites(value, 13, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(197, value), rsp, targets8)
	log("p13 w:%d", value)
	countwrites(value, 13, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(167, value), rsp, targets9)
	log("p13 w:%d", value)
	countwrites(value, 13, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(184, value), rsp, targets10)
	log("p13 w:%d", value)
	countwrites(value, 13, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(87, value), rsp, targets11)
	log("p13 w:%d", value)
	countwrites(value, 13, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(204, value), rsp, targets12)
	log("p13 w:%d", value)
	countwrites(value, 13, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(205, &value), rsp, s13)
	log("p13 r:%d", value)
	countreads(value, 13, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(175, value), rsp, targets13)
	log("p13 w:%d", value)
	countwrites(value, 13, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(246, value), rsp, targets14)
	log("p13 w:%d", value)
	countwrites(value, 13, 31)
	l.Unlock()
}

func p14() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s24]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s26]
	targets0 := make([]string, 2)
	targets0[1] = uri[s20]
	targets0[0] = uri[s9]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(17, value), rsp, targets0)
	log("p14 w:%d", value)
	countwrites(value, 14, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(45, value), rsp, targets1)
	log("p14 w:%d", value)
	countwrites(value, 14, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(74, value), rsp, targets2)
	log("p14 w:%d", value)
	countwrites(value, 14, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(14, value), rsp, targets3)
	log("p14 w:%d", value)
	countwrites(value, 14, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(40, value), rsp, targets4)
	log("p14 w:%d", value)
	countwrites(value, 14, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(117, value), rsp, targets5)
	log("p14 w:%d", value)
	countwrites(value, 14, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(136, value), rsp, targets6)
	log("p14 w:%d", value)
	countwrites(value, 14, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(80, value), rsp, targets7)
	log("p14 w:%d", value)
	countwrites(value, 14, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(86, value), rsp, targets8)
	log("p14 w:%d", value)
	countwrites(value, 14, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(96, value), rsp, targets9)
	log("p14 w:%d", value)
	countwrites(value, 14, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(196, value), rsp, targets10)
	log("p14 w:%d", value)
	countwrites(value, 14, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(111, value), rsp, targets11)
	log("p14 w:%d", value)
	countwrites(value, 14, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(47, &value), rsp, s14)
	log("p14 r:%d", value)
	countreads(value, 14, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(183, value), rsp, targets12)
	log("p14 w:%d", value)
	countwrites(value, 14, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(208, value), rsp, targets13)
	log("p14 w:%d", value)
	countwrites(value, 14, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(213, value), rsp, targets14)
	log("p14 w:%d", value)
	countwrites(value, 14, 27)
	l.Unlock()
}

func p15() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s12]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s24]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(53, value), rsp, targets0)
	log("p15 w:%d", value)
	countwrites(value, 15, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(74, value), rsp, targets1)
	log("p15 w:%d", value)
	countwrites(value, 15, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(63, value), rsp, targets2)
	log("p15 w:%d", value)
	countwrites(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	Put(CreateTuple(150, value), rsp, targets3)
	log("p15 w:%d", value)
	countwrites(value, 15, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(245, value), rsp, targets4)
	log("p15 w:%d", value)
	countwrites(value, 15, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(135, value), rsp, targets5)
	log("p15 w:%d", value)
	countwrites(value, 15, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(194, value), rsp, targets6)
	log("p15 w:%d", value)
	countwrites(value, 15, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(9, value), rsp, targets7)
	log("p15 w:%d", value)
	countwrites(value, 15, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(109, value), rsp, targets8)
	log("p15 w:%d", value)
	countwrites(value, 15, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(91, value), rsp, targets9)
	log("p15 w:%d", value)
	countwrites(value, 15, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(61, value), rsp, targets10)
	log("p15 w:%d", value)
	countwrites(value, 15, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(8, value), rsp, targets11)
	log("p15 w:%d", value)
	countwrites(value, 15, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(83, value), rsp, targets12)
	log("p15 w:%d", value)
	countwrites(value, 15, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(69, value), rsp, targets13)
	log("p15 w:%d", value)
	countwrites(value, 15, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(209, &value), rsp, s15)
	log("p15 r:%d", value)
	countreads(value, 15, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(11, value), rsp, targets14)
	log("p15 w:%d", value)
	countwrites(value, 15, 2)
	l.Unlock()
}

func p16() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s20]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s3]
	targets1 := make([]string, 1)
	targets1[0] = uri[s7]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(194, value), rsp, targets0)
	log("p16 w:%d", value)
	countwrites(value, 16, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(4, value), rsp, targets1)
	log("p16 w:%d", value)
	countwrites(value, 16, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(55, value), rsp, targets2)
	log("p16 w:%d", value)
	countwrites(value, 16, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(181, value), rsp, targets3)
	log("p16 w:%d", value)
	countwrites(value, 16, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(65, value), rsp, targets4)
	log("p16 w:%d", value)
	countwrites(value, 16, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(41, value), rsp, targets5)
	log("p16 w:%d", value)
	countwrites(value, 16, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(24, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(126, value), rsp, targets6)
	log("p16 w:%d", value)
	countwrites(value, 16, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(254, value), rsp, targets7)
	log("p16 w:%d", value)
	countwrites(value, 16, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(234, value), rsp, targets8)
	log("p16 w:%d", value)
	countwrites(value, 16, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(21, value), rsp, targets9)
	log("p16 w:%d", value)
	countwrites(value, 16, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(230, value), rsp, targets10)
	log("p16 w:%d", value)
	countwrites(value, 16, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(137, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(147, &value), rsp, s16)
	log("p16 r:%d", value)
	countreads(value, 16, 19)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(176, value), rsp, targets11)
	log("p16 w:%d", value)
	countwrites(value, 16, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(77, value), rsp, targets12)
	log("p16 w:%d", value)
	countwrites(value, 16, 10)
	l.Unlock()
}

func p17() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s5]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s1]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(34, value), rsp, targets0)
	log("p17 w:%d", value)
	countwrites(value, 17, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(180, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(189, value), rsp, targets1)
	log("p17 w:%d", value)
	countwrites(value, 17, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(134, value), rsp, targets2)
	log("p17 w:%d", value)
	countwrites(value, 17, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(198, &value), rsp, s17)
	log("p17 r:%d", value)
	countreads(value, 17, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(238, value), rsp, targets3)
	log("p17 w:%d", value)
	countwrites(value, 17, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(9, value), rsp, targets4)
	log("p17 w:%d", value)
	countwrites(value, 17, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(30, value), rsp, targets5)
	log("p17 w:%d", value)
	countwrites(value, 17, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(159, value), rsp, targets6)
	log("p17 w:%d", value)
	countwrites(value, 17, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(236, value), rsp, targets7)
	log("p17 w:%d", value)
	countwrites(value, 17, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(97, value), rsp, targets8)
	log("p17 w:%d", value)
	countwrites(value, 17, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(253, value), rsp, targets9)
	log("p17 w:%d", value)
	countwrites(value, 17, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(119, value), rsp, targets10)
	log("p17 w:%d", value)
	countwrites(value, 17, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(208, value), rsp, targets11)
	log("p17 w:%d", value)
	countwrites(value, 17, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(226, value), rsp, targets12)
	log("p17 w:%d", value)
	countwrites(value, 17, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(212, value), rsp, targets13)
	log("p17 w:%d", value)
	countwrites(value, 17, 27)
	l.Unlock()
}

func p18() {
	targets15 := make([]string, 0)
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s11]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(182, value), rsp, targets0)
	log("p18 w:%d", value)
	countwrites(value, 18, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(189, value), rsp, targets1)
	log("p18 w:%d", value)
	countwrites(value, 18, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(51, value), rsp, targets2)
	log("p18 w:%d", value)
	countwrites(value, 18, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(90, value), rsp, targets3)
	log("p18 w:%d", value)
	countwrites(value, 18, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(142, value), rsp, targets4)
	log("p18 w:%d", value)
	countwrites(value, 18, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(112, value), rsp, targets5)
	log("p18 w:%d", value)
	countwrites(value, 18, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(100, value), rsp, targets6)
	log("p18 w:%d", value)
	countwrites(value, 18, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(179, value), rsp, targets7)
	log("p18 w:%d", value)
	countwrites(value, 18, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(226, value), rsp, targets8)
	log("p18 w:%d", value)
	countwrites(value, 18, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(75, value), rsp, targets9)
	log("p18 w:%d", value)
	countwrites(value, 18, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(169, value), rsp, targets10)
	log("p18 w:%d", value)
	countwrites(value, 18, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(172, value), rsp, targets11)
	log("p18 w:%d", value)
	countwrites(value, 18, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(82, value), rsp, targets12)
	log("p18 w:%d", value)
	countwrites(value, 18, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(245, value), rsp, targets13)
	log("p18 w:%d", value)
	countwrites(value, 18, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(206, value), rsp, targets14)
	log("p18 w:%d", value)
	countwrites(value, 18, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(212, value), rsp, targets15)
	log("p18 w:%d", value)
	countwrites(value, 18, 27)
	l.Unlock()
}

func p19() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 1)
	targets13[0] = uri[s17]
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s1]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s14]
	targets2[0] = uri[s11]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(57, value), rsp, targets0)
	log("p19 w:%d", value)
	countwrites(value, 19, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(144, value), rsp, targets1)
	log("p19 w:%d", value)
	countwrites(value, 19, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(47, value), rsp, targets2)
	log("p19 w:%d", value)
	countwrites(value, 19, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(6, value), rsp, targets3)
	log("p19 w:%d", value)
	countwrites(value, 19, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(212, value), rsp, targets4)
	log("p19 w:%d", value)
	countwrites(value, 19, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(70, value), rsp, targets5)
	log("p19 w:%d", value)
	countwrites(value, 19, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(250, &value), rsp, s19)
	log("p19 r:%d", value)
	countreads(value, 19, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(79, value), rsp, targets6)
	log("p19 w:%d", value)
	countwrites(value, 19, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(87, value), rsp, targets7)
	log("p19 w:%d", value)
	countwrites(value, 19, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(248, value), rsp, targets8)
	log("p19 w:%d", value)
	countwrites(value, 19, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(234, value), rsp, targets9)
	log("p19 w:%d", value)
	countwrites(value, 19, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(102, value), rsp, targets10)
	log("p19 w:%d", value)
	countwrites(value, 19, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(7, value), rsp, targets11)
	log("p19 w:%d", value)
	countwrites(value, 19, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(201, value), rsp, targets12)
	log("p19 w:%d", value)
	countwrites(value, 19, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(198, value), rsp, targets13)
	log("p19 w:%d", value)
	countwrites(value, 19, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(242, value), rsp, targets14)
	log("p19 w:%d", value)
	countwrites(value, 19, 31)
	l.Unlock()
}

func p20() {
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 2)
	targets7[1] = uri[s28]
	targets7[0] = uri[s8]
	targets6 := make([]string, 0)
	targets5 := make([]string, 1)
	targets5[0] = uri[s23]
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(171, value), rsp, targets0)
	log("p20 w:%d", value)
	countwrites(value, 20, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(11, value), rsp, targets1)
	log("p20 w:%d", value)
	countwrites(value, 20, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(49, value), rsp, targets2)
	log("p20 w:%d", value)
	countwrites(value, 20, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(162, value), rsp, targets3)
	log("p20 w:%d", value)
	countwrites(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(77, value), rsp, targets4)
	log("p20 w:%d", value)
	countwrites(value, 20, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(158, value), rsp, targets5)
	log("p20 w:%d", value)
	countwrites(value, 20, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(79, value), rsp, targets6)
	log("p20 w:%d", value)
	countwrites(value, 20, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(165, value), rsp, targets7)
	log("p20 w:%d", value)
	countwrites(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(17, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(162, value), rsp, targets8)
	log("p20 w:%d", value)
	countwrites(value, 20, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(75, value), rsp, targets9)
	log("p20 w:%d", value)
	countwrites(value, 20, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(254, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(43, value), rsp, targets10)
	log("p20 w:%d", value)
	countwrites(value, 20, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(126, value), rsp, targets11)
	log("p20 w:%d", value)
	countwrites(value, 20, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(222, value), rsp, targets12)
	log("p20 w:%d", value)
	countwrites(value, 20, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(5, &value), rsp, s20)
	log("p20 r:%d", value)
	countreads(value, 20, 1)
	l.Unlock()
}

func p21() {
	targets15 := make([]string, 0)
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s20]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(157, value), rsp, targets0)
	log("p21 w:%d", value)
	countwrites(value, 21, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(254, value), rsp, targets1)
	log("p21 w:%d", value)
	countwrites(value, 21, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(222, value), rsp, targets2)
	log("p21 w:%d", value)
	countwrites(value, 21, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(50, value), rsp, targets3)
	log("p21 w:%d", value)
	countwrites(value, 21, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(33, value), rsp, targets4)
	log("p21 w:%d", value)
	countwrites(value, 21, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(82, value), rsp, targets5)
	log("p21 w:%d", value)
	countwrites(value, 21, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(144, value), rsp, targets6)
	log("p21 w:%d", value)
	countwrites(value, 21, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(157, value), rsp, targets7)
	log("p21 w:%d", value)
	countwrites(value, 21, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(40, value), rsp, targets8)
	log("p21 w:%d", value)
	countwrites(value, 21, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(235, value), rsp, targets9)
	log("p21 w:%d", value)
	countwrites(value, 21, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(199, value), rsp, targets10)
	log("p21 w:%d", value)
	countwrites(value, 21, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(175, value), rsp, targets11)
	log("p21 w:%d", value)
	countwrites(value, 21, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(182, value), rsp, targets12)
	log("p21 w:%d", value)
	countwrites(value, 21, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(238, value), rsp, targets13)
	log("p21 w:%d", value)
	countwrites(value, 21, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	Put(CreateTuple(189, value), rsp, targets14)
	log("p21 w:%d", value)
	countwrites(value, 21, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(75, value), rsp, targets15)
	log("p21 w:%d", value)
	countwrites(value, 21, 10)
	l.Unlock()
}

func p22() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 1)
	targets12[0] = uri[s24]
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s17]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 2)
	targets2[1] = uri[s20]
	targets2[0] = uri[s9]
	targets1 := make([]string, 1)
	targets1[0] = uri[s10]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(207, value), rsp, targets0)
	log("p22 w:%d", value)
	countwrites(value, 22, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(174, value), rsp, targets1)
	log("p22 w:%d", value)
	countwrites(value, 22, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(60, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(17, value), rsp, targets2)
	log("p22 w:%d", value)
	countwrites(value, 22, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(18, value), rsp, targets3)
	log("p22 w:%d", value)
	countwrites(value, 22, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(183, value), rsp, targets4)
	log("p22 w:%d", value)
	countwrites(value, 22, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(61, value), rsp, targets5)
	log("p22 w:%d", value)
	countwrites(value, 22, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(77, value), rsp, targets6)
	log("p22 w:%d", value)
	countwrites(value, 22, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(71, value), rsp, targets7)
	log("p22 w:%d", value)
	countwrites(value, 22, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(197, &value), rsp, s22)
	log("p22 r:%d", value)
	countreads(value, 22, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(219, value), rsp, targets8)
	log("p22 w:%d", value)
	countwrites(value, 22, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(198, value), rsp, targets9)
	log("p22 w:%d", value)
	countwrites(value, 22, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(22, value), rsp, targets10)
	log("p22 w:%d", value)
	countwrites(value, 22, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(119, value), rsp, targets11)
	log("p22 w:%d", value)
	countwrites(value, 22, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(117, value), rsp, targets12)
	log("p22 w:%d", value)
	countwrites(value, 22, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(139, value), rsp, targets13)
	log("p22 w:%d", value)
	countwrites(value, 22, 18)
	l.Unlock()
}

func p23() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s16]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s31]
	targets2 := make([]string, 0)
	targets1 := make([]string, 1)
	targets1[0] = uri[s3]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(8, value), rsp, targets0)
	log("p23 w:%d", value)
	countwrites(value, 23, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(125, value), rsp, targets1)
	log("p23 w:%d", value)
	countwrites(value, 23, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(161, value), rsp, targets2)
	log("p23 w:%d", value)
	countwrites(value, 23, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(52, value), rsp, targets3)
	log("p23 w:%d", value)
	countwrites(value, 23, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(173, value), rsp, targets4)
	log("p23 w:%d", value)
	countwrites(value, 23, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(25, value), rsp, targets5)
	log("p23 w:%d", value)
	countwrites(value, 23, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(93, value), rsp, targets6)
	log("p23 w:%d", value)
	countwrites(value, 23, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(251, value), rsp, targets7)
	log("p23 w:%d", value)
	countwrites(value, 23, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(98, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(170, value), rsp, targets8)
	log("p23 w:%d", value)
	countwrites(value, 23, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(24, value), rsp, targets9)
	log("p23 w:%d", value)
	countwrites(value, 23, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(12, value), rsp, targets10)
	log("p23 w:%d", value)
	countwrites(value, 23, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(94, value), rsp, targets11)
	log("p23 w:%d", value)
	countwrites(value, 23, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	Put(CreateTuple(153, value), rsp, targets12)
	log("p23 w:%d", value)
	countwrites(value, 23, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(158, &value), rsp, s23)
	log("p23 r:%d", value)
	countreads(value, 23, 20)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(87, value), rsp, targets13)
	log("p23 w:%d", value)
	countwrites(value, 23, 11)
	l.Unlock()
}

func p24() {
	targets13 := make([]string, 1)
	targets13[0] = uri[s5]
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s1]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s31]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s26]
	targets1 := make([]string, 1)
	targets1[0] = uri[s32]
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(193, value), rsp, targets0)
	log("p24 w:%d", value)
	countwrites(value, 24, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(72, value), rsp, targets1)
	log("p24 w:%d", value)
	countwrites(value, 24, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(45, value), rsp, targets2)
	log("p24 w:%d", value)
	countwrites(value, 24, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(117, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(31, value), rsp, targets3)
	log("p24 w:%d", value)
	countwrites(value, 24, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(195, value), rsp, targets4)
	log("p24 w:%d", value)
	countwrites(value, 24, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(95, value), rsp, targets5)
	log("p24 w:%d", value)
	countwrites(value, 24, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(103, value), rsp, targets6)
	log("p24 w:%d", value)
	countwrites(value, 24, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(18, value), rsp, targets7)
	log("p24 w:%d", value)
	countwrites(value, 24, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(240, value), rsp, targets8)
	log("p24 w:%d", value)
	countwrites(value, 24, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(241, value), rsp, targets9)
	log("p24 w:%d", value)
	countwrites(value, 24, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(31, value), rsp, targets10)
	log("p24 w:%d", value)
	countwrites(value, 24, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(7, value), rsp, targets11)
	log("p24 w:%d", value)
	countwrites(value, 24, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(64, value), rsp, targets12)
	log("p24 w:%d", value)
	countwrites(value, 24, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(97, value), rsp, targets13)
	log("p24 w:%d", value)
	countwrites(value, 24, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(53, &value), rsp, s24)
	log("p24 r:%d", value)
	countreads(value, 24, 7)
	l.Unlock()
}

func p25() {
	targets15 := make([]string, 1)
	targets15[0] = uri[s17]
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s11]
	targets7 := make([]string, 1)
	targets7[0] = uri[s11]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s16]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(168, value), rsp, targets0)
	log("p25 w:%d", value)
	countwrites(value, 25, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(11, value), rsp, targets1)
	log("p25 w:%d", value)
	countwrites(value, 25, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(244, value), rsp, targets2)
	log("p25 w:%d", value)
	countwrites(value, 25, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(84, value), rsp, targets3)
	log("p25 w:%d", value)
	countwrites(value, 25, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(137, value), rsp, targets4)
	log("p25 w:%d", value)
	countwrites(value, 25, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(252, value), rsp, targets5)
	log("p25 w:%d", value)
	countwrites(value, 25, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(31, value), rsp, targets6)
	log("p25 w:%d", value)
	countwrites(value, 25, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(130, value), rsp, targets7)
	log("p25 w:%d", value)
	countwrites(value, 25, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(130, value), rsp, targets8)
	log("p25 w:%d", value)
	countwrites(value, 25, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(27, value), rsp, targets9)
	log("p25 w:%d", value)
	countwrites(value, 25, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(215, value), rsp, targets10)
	log("p25 w:%d", value)
	countwrites(value, 25, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(1, value), rsp, targets11)
	log("p25 w:%d", value)
	countwrites(value, 25, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(183, value), rsp, targets12)
	log("p25 w:%d", value)
	countwrites(value, 25, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(119, value), rsp, targets13)
	log("p25 w:%d", value)
	countwrites(value, 25, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(203, value), rsp, targets14)
	log("p25 w:%d", value)
	countwrites(value, 25, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(180, value), rsp, targets15)
	log("p25 w:%d", value)
	countwrites(value, 25, 23)
	l.Unlock()
}

func p26() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 1)
	targets12[0] = uri[s11]
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(210, value), rsp, targets0)
	log("p26 w:%d", value)
	countwrites(value, 26, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(100, value), rsp, targets1)
	log("p26 w:%d", value)
	countwrites(value, 26, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(62, value), rsp, targets2)
	log("p26 w:%d", value)
	countwrites(value, 26, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(45, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(196, value), rsp, targets3)
	log("p26 w:%d", value)
	countwrites(value, 26, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(211, &value), rsp, s26)
	log("p26 r:%d", value)
	countreads(value, 26, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(116, value), rsp, targets4)
	log("p26 w:%d", value)
	countwrites(value, 26, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(200, value), rsp, targets5)
	log("p26 w:%d", value)
	countwrites(value, 26, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(244, value), rsp, targets6)
	log("p26 w:%d", value)
	countwrites(value, 26, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(253, value), rsp, targets7)
	log("p26 w:%d", value)
	countwrites(value, 26, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(10, value), rsp, targets8)
	log("p26 w:%d", value)
	countwrites(value, 26, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(193, value), rsp, targets9)
	log("p26 w:%d", value)
	countwrites(value, 26, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(104, value), rsp, targets10)
	log("p26 w:%d", value)
	countwrites(value, 26, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(51, value), rsp, targets11)
	log("p26 w:%d", value)
	countwrites(value, 26, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(130, value), rsp, targets12)
	log("p26 w:%d", value)
	countwrites(value, 26, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(65, value), rsp, targets13)
	log("p26 w:%d", value)
	countwrites(value, 26, 9)
	l.Unlock()
}

func p27() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 1)
	targets13[0] = uri[s10]
	targets12 := make([]string, 0)
	targets11 := make([]string, 1)
	targets11[0] = uri[s3]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 1)
	targets7[0] = uri[s23]
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(229, value), rsp, targets0)
	log("p27 w:%d", value)
	countwrites(value, 27, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(66, value), rsp, targets1)
	log("p27 w:%d", value)
	countwrites(value, 27, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(124, value), rsp, targets2)
	log("p27 w:%d", value)
	countwrites(value, 27, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(185, &value), rsp, s27)
	log("p27 r:%d", value)
	countreads(value, 27, 24)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(6, value), rsp, targets3)
	log("p27 w:%d", value)
	countwrites(value, 27, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(105, value), rsp, targets4)
	log("p27 w:%d", value)
	countwrites(value, 27, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(224, value), rsp, targets5)
	log("p27 w:%d", value)
	countwrites(value, 27, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(6, value), rsp, targets6)
	log("p27 w:%d", value)
	countwrites(value, 27, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(98, value), rsp, targets7)
	log("p27 w:%d", value)
	countwrites(value, 27, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(15, value), rsp, targets8)
	log("p27 w:%d", value)
	countwrites(value, 27, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(170, value), rsp, targets9)
	log("p27 w:%d", value)
	countwrites(value, 27, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(175, value), rsp, targets10)
	log("p27 w:%d", value)
	countwrites(value, 27, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(55, value), rsp, targets11)
	log("p27 w:%d", value)
	countwrites(value, 27, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(78, value), rsp, targets12)
	log("p27 w:%d", value)
	countwrites(value, 27, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(132, value), rsp, targets13)
	log("p27 w:%d", value)
	countwrites(value, 27, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(71, value), rsp, targets14)
	log("p27 w:%d", value)
	countwrites(value, 27, 9)
	l.Unlock()
}

func p28() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s31]
	targets2 := make([]string, 1)
	targets2[0] = uri[s7]
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(165, &value), rsp, s28)
	log("p28 r:%d", value)
	countreads(value, 28, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(76, value), rsp, targets0)
	log("p28 w:%d", value)
	countwrites(value, 28, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(16, value), rsp, targets1)
	log("p28 w:%d", value)
	countwrites(value, 28, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(4, value), rsp, targets2)
	log("p28 w:%d", value)
	countwrites(value, 28, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(52, value), rsp, targets3)
	log("p28 w:%d", value)
	countwrites(value, 28, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(71, value), rsp, targets4)
	log("p28 w:%d", value)
	countwrites(value, 28, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(19, value), rsp, targets5)
	log("p28 w:%d", value)
	countwrites(value, 28, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(14, value), rsp, targets6)
	log("p28 w:%d", value)
	countwrites(value, 28, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(54, value), rsp, targets7)
	log("p28 w:%d", value)
	countwrites(value, 28, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	Put(CreateTuple(208, value), rsp, targets8)
	log("p28 w:%d", value)
	countwrites(value, 28, 26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(210, value), rsp, targets9)
	log("p28 w:%d", value)
	countwrites(value, 28, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(126, value), rsp, targets10)
	log("p28 w:%d", value)
	countwrites(value, 28, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(101, value), rsp, targets11)
	log("p28 w:%d", value)
	countwrites(value, 28, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(242, value), rsp, targets12)
	log("p28 w:%d", value)
	countwrites(value, 28, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(65, value), rsp, targets13)
	log("p28 w:%d", value)
	countwrites(value, 28, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(3, value), rsp, targets14)
	log("p28 w:%d", value)
	countwrites(value, 28, 1)
	l.Unlock()
}

func p29() {
	targets15 := make([]string, 0)
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 2)
	targets12[1] = uri[s20]
	targets12[0] = uri[s9]
	targets11 := make([]string, 1)
	targets11[0] = uri[s10]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 1)
	targets2[0] = uri[s17]
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s31]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(103, value), rsp, targets0)
	log("p29 w:%d", value)
	countwrites(value, 29, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(242, value), rsp, targets1)
	log("p29 w:%d", value)
	countwrites(value, 29, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(180, value), rsp, targets2)
	log("p29 w:%d", value)
	countwrites(value, 29, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(101, value), rsp, targets3)
	log("p29 w:%d", value)
	countwrites(value, 29, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(59, value), rsp, targets4)
	log("p29 w:%d", value)
	countwrites(value, 29, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(38, value), rsp, targets5)
	log("p29 w:%d", value)
	countwrites(value, 29, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(78, value), rsp, targets6)
	log("p29 w:%d", value)
	countwrites(value, 29, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(37, value), rsp, targets7)
	log("p29 w:%d", value)
	countwrites(value, 29, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(92, value), rsp, targets8)
	log("p29 w:%d", value)
	countwrites(value, 29, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(106, value), rsp, targets9)
	log("p29 w:%d", value)
	countwrites(value, 29, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	Put(CreateTuple(9, value), rsp, targets10)
	log("p29 w:%d", value)
	countwrites(value, 29, 2)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(132, value), rsp, targets11)
	log("p29 w:%d", value)
	countwrites(value, 29, 17)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(17, value), rsp, targets12)
	log("p29 w:%d", value)
	countwrites(value, 29, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	Put(CreateTuple(177, value), rsp, targets13)
	log("p29 w:%d", value)
	countwrites(value, 29, 23)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(200, value), rsp, targets14)
	log("p29 w:%d", value)
	countwrites(value, 29, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	Put(CreateTuple(135, value), rsp, targets15)
	log("p29 w:%d", value)
	countwrites(value, 29, 17)
	l.Unlock()
}

func p30() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 1)
	targets13[0] = uri[s19]
	targets12 := make([]string, 1)
	targets12[0] = uri[s20]
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 1)
	targets8[0] = uri[s8]
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 1)
	targets4[0] = uri[s5]
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(93, value), rsp, targets0)
	log("p30 w:%d", value)
	countwrites(value, 30, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(64, value), rsp, targets1)
	log("p30 w:%d", value)
	countwrites(value, 30, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(251, value), rsp, targets2)
	log("p30 w:%d", value)
	countwrites(value, 30, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(173, value), rsp, targets3)
	log("p30 w:%d", value)
	countwrites(value, 30, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	Put(CreateTuple(97, value), rsp, targets4)
	log("p30 w:%d", value)
	countwrites(value, 30, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	Put(CreateTuple(245, value), rsp, targets5)
	log("p30 w:%d", value)
	countwrites(value, 30, 31)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	Put(CreateTuple(84, value), rsp, targets6)
	log("p30 w:%d", value)
	countwrites(value, 30, 11)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(89, value), rsp, targets7)
	log("p30 w:%d", value)
	countwrites(value, 30, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(233, &value), rsp, s30)
	log("p30 r:%d", value)
	countreads(value, 30, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(123, value), rsp, targets8)
	log("p30 w:%d", value)
	countwrites(value, 30, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	Put(CreateTuple(229, value), rsp, targets9)
	log("p30 w:%d", value)
	countwrites(value, 30, 29)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	Put(CreateTuple(196, value), rsp, targets10)
	log("p30 w:%d", value)
	countwrites(value, 30, 25)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	Put(CreateTuple(31, value), rsp, targets11)
	log("p30 w:%d", value)
	countwrites(value, 30, 4)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(254, value), rsp, targets12)
	log("p30 w:%d", value)
	countwrites(value, 30, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(250, value), rsp, targets13)
	log("p30 w:%d", value)
	countwrites(value, 30, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(110, value), rsp, targets14)
	log("p30 w:%d", value)
	countwrites(value, 30, 14)
	l.Unlock()
}

func p31() {
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 2)
	targets11[1] = uri[s14]
	targets11[0] = uri[s11]
	targets10 := make([]string, 0)
	targets9 := make([]string, 0)
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 0)
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 1)
	targets3[0] = uri[s9]
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 1)
	targets0[0] = uri[s26]

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 27
	Put(CreateTuple(211, value), rsp, targets0)
	log("p31 w:%d", value)
	countwrites(value, 31, 27)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(59, value), rsp, targets1)
	log("p31 w:%d", value)
	countwrites(value, 31, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(103, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 13)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	Put(CreateTuple(236, value), rsp, targets2)
	log("p31 w:%d", value)
	countwrites(value, 31, 30)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	Put(CreateTuple(2, value), rsp, targets3)
	log("p31 w:%d", value)
	countwrites(value, 31, 1)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	Put(CreateTuple(128, value), rsp, targets4)
	log("p31 w:%d", value)
	countwrites(value, 31, 16)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	Put(CreateTuple(94, value), rsp, targets5)
	log("p31 w:%d", value)
	countwrites(value, 31, 12)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(143, value), rsp, targets6)
	log("p31 w:%d", value)
	countwrites(value, 31, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(118, value), rsp, targets7)
	log("p31 w:%d", value)
	countwrites(value, 31, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(223, value), rsp, targets8)
	log("p31 w:%d", value)
	countwrites(value, 31, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	Put(CreateTuple(58, value), rsp, targets9)
	log("p31 w:%d", value)
	countwrites(value, 31, 8)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(66, value), rsp, targets10)
	log("p31 w:%d", value)
	countwrites(value, 31, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	Put(CreateTuple(47, value), rsp, targets11)
	log("p31 w:%d", value)
	countwrites(value, 31, 6)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	Put(CreateTuple(70, value), rsp, targets12)
	log("p31 w:%d", value)
	countwrites(value, 31, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(21, value), rsp, targets13)
	log("p31 w:%d", value)
	countwrites(value, 31, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(52, &value), rsp, s31)
	log("p31 r:%d", value)
	countreads(value, 31, 7)
	l.Unlock()
}

func p32() {
	targets14 := make([]string, 0)
	targets13 := make([]string, 0)
	targets12 := make([]string, 0)
	targets11 := make([]string, 0)
	targets10 := make([]string, 0)
	targets9 := make([]string, 1)
	targets9[0] = uri[s20]
	targets8 := make([]string, 0)
	targets7 := make([]string, 0)
	targets6 := make([]string, 1)
	targets6[0] = uri[s10]
	targets5 := make([]string, 0)
	targets4 := make([]string, 0)
	targets3 := make([]string, 0)
	targets2 := make([]string, 0)
	targets1 := make([]string, 0)
	targets0 := make([]string, 0)

	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	Put(CreateTuple(168, value), rsp, targets0)
	log("p32 w:%d", value)
	countwrites(value, 32, 21)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	Put(CreateTuple(18, value), rsp, targets1)
	log("p32 w:%d", value)
	countwrites(value, 32, 3)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(142, value), rsp, targets2)
	log("p32 w:%d", value)
	countwrites(value, 32, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(218, value), rsp, targets3)
	log("p32 w:%d", value)
	countwrites(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	Put(CreateTuple(50, value), rsp, targets4)
	log("p32 w:%d", value)
	countwrites(value, 32, 7)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	Put(CreateTuple(105, value), rsp, targets5)
	log("p32 w:%d", value)
	countwrites(value, 32, 14)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	Put(CreateTuple(174, value), rsp, targets6)
	log("p32 w:%d", value)
	countwrites(value, 32, 22)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	Put(CreateTuple(223, value), rsp, targets7)
	log("p32 w:%d", value)
	countwrites(value, 32, 28)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	Put(CreateTuple(114, value), rsp, targets8)
	log("p32 w:%d", value)
	countwrites(value, 32, 15)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(254, value), rsp, targets9)
	log("p32 w:%d", value)
	countwrites(value, 32, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	Put(CreateTuple(36, value), rsp, targets10)
	log("p32 w:%d", value)
	countwrites(value, 32, 5)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(78, value), rsp, targets11)
	log("p32 w:%d", value)
	countwrites(value, 32, 10)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	Put(CreateTuple(139, value), rsp, targets12)
	log("p32 w:%d", value)
	countwrites(value, 32, 18)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	Put(CreateTuple(249, value), rsp, targets13)
	log("p32 w:%d", value)
	countwrites(value, 32, 32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	QueryP(CreateTuple(72, &value), rsp, s32)
	log("p32 r:%d", value)
	countreads(value, 32, 9)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	Put(CreateTuple(80, value), rsp, targets14)
	log("p32 w:%d", value)
	countwrites(value, 32, 10)
	l.Unlock()
}

func log(format string, a ...interface{}) {
	if debug {
		fmt.Fprintf(os.Stdout, format+"\n", a...)
	}
}

func delay() {
	time.Sleep(time.Duration(rand.Int63n(75)) * time.Millisecond)
}

func countreads(value int, localspace int, targetspace int) {
	if value == -1 {
		reads_insuccess += 1
	} else {
		reads_success += 1
	}

	if localspace == targetspace {
		reads_local += 1
	} else {
		reads_remote += 1
	}
}

func countwrites(value int, localspace int, targetspace int) {
	writestotal += 1

	if localspace == targetspace {
		writes_local += 1
	} else {
		writes_remote += 1
	}
}
