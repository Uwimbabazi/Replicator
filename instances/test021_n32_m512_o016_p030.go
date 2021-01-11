package main

import (
	"os"
	"fmt"
	"sync"
	"time"
	"math/rand"
	. "github.com/pspaces/gospace"
)

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
	writes_replicated = 0
	reads_success = 0
	reads_insuccess = 0
	rand.Seed(time.Now().UTC().UnixNano())

	wg.Add(32)

	s1 = NewSpace("tcp://localhost:34001/s1")
	s2 = NewSpace("tcp://localhost:34002/s2")
	s3 = NewSpace("tcp://localhost:34003/s3")
	s4 = NewSpace("tcp://localhost:34004/s4")
    s5 = NewSpace("tcp://localhost:34005/s5")
    s6 = NewSpace("tcp://localhost:34006/s6")
    s7 = NewSpace("tcp://localhost:34007/s7")
    s8 = NewSpace("tcp://localhost:34008/s8")
    s9 = NewSpace("tcp://localhost:34009/s9")
    s10 = NewSpace("tcp://localhost:34010/s10")
    s11 = NewSpace("tcp://localhost:34011/s11")
    s12 = NewSpace("tcp://localhost:34012/s12")
    s13 = NewSpace("tcp://localhost:34013/s13")
    s14 = NewSpace("tcp://localhost:34014/s14")
    s15 = NewSpace("tcp://localhost:34015/s15")
  	s16 = NewSpace("tcp://localhost:34016/s16")
	s17 = NewSpace("tcp://localhost:34017/s17")
	s18 = NewSpace("tcp://localhost:34018/s18")
    s19 = NewSpace("tcp://localhost:34019/s19")
    s20 = NewSpace("tcp://localhost:34020/s20")
	s21 = NewSpace("tcp://localhost:34021/s21")
	s22 = NewSpace("tcp://localhost:34022/s22")
	s23 = NewSpace("tcp://localhost:34023/s23")
	s24 = NewSpace("tcp://localhost:34024/s24")
	s25 = NewSpace("tcp://localhost:34025/s25")
	s26 = NewSpace("tcp://localhost:34026/s26")
	s27 = NewSpace("tcp://localhost:34027/s27")
	s28 = NewSpace("tcp://localhost:34028/s28")
	s29 = NewSpace("tcp://localhost:34029/s29")
	s30 = NewSpace("tcp://localhost:34030/s30")
	s31 = NewSpace("tcp://localhost:34031/s31")
	s32 = NewSpace("tcp://localhost:34032/s32")



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

	writes_replicated = 0   // these include the normal put too

	fmt.Println("       loc w,    rem w,   repl w,    loc r,    rem r,    tot w,   succ r,   fail r")
	fmt.Printf ("    %8d, %8d, %8d, %8d, %8d, %8d, %8d, %8d\n", writes_local, writes_remote, writes_replicated, reads_local, reads_remote, writestotal, reads_success, reads_insuccess)
}

func p1() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s5.QueryP(70,&value)
	log("p1 r:%d", value)
	countreads(value,1,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(74,&value)
	log("p1 r:%d", value)
	countreads(value,1,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(413,&value)
	log("p1 r:%d", value)
	countreads(value,1,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(2,&value)
	log("p1 r:%d", value)
	countreads(value,1,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(105,&value)
	log("p1 r:%d", value)
	countreads(value,1,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(336,value)
	log("p1 w:%d", value)
	countwrites(value,1,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(257,&value)
	log("p1 r:%d", value)
	countreads(value,1,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(197,&value)
	log("p1 r:%d", value)
	countreads(value,1,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(385,&value)
	log("p1 r:%d", value)
	countreads(value,1,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(479,&value)
	log("p1 r:%d", value)
	countreads(value,1,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(387,value)
	log("p1 w:%d", value)
	countwrites(value,1,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(415,value)
	log("p1 w:%d", value)
	countwrites(value,1,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(210,&value)
	log("p1 r:%d", value)
	countreads(value,1,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(288,&value)
	log("p1 r:%d", value)
	countreads(value,1,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(99,value)
	log("p1 w:%d", value)
	countwrites(value,1,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(40,&value)
	log("p1 r:%d", value)
	countreads(value,1,3)
	l.Unlock()
}

func p2() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s4.QueryP(63,&value)
	log("p2 r:%d", value)
	countreads(value,2,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(380,&value)
	log("p2 r:%d", value)
	countreads(value,2,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(352,&value)
	log("p2 r:%d", value)
	countreads(value,2,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(241,&value)
	log("p2 r:%d", value)
	countreads(value,2,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(245,value)
	log("p2 w:%d", value)
	countwrites(value,2,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(127,value)
	log("p2 w:%d", value)
	countwrites(value,2,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(271,value)
	log("p2 w:%d", value)
	countwrites(value,2,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(22,&value)
	log("p2 r:%d", value)
	countreads(value,2,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(394,value)
	log("p2 w:%d", value)
	countwrites(value,2,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(12,value)
	log("p2 w:%d", value)
	countwrites(value,2,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(67,&value)
	log("p2 r:%d", value)
	countreads(value,2,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(76,value)
	log("p2 w:%d", value)
	countwrites(value,2,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(125,&value)
	log("p2 r:%d", value)
	countreads(value,2,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(260,&value)
	log("p2 r:%d", value)
	countreads(value,2,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(1,&value)
	log("p2 r:%d", value)
	countreads(value,2,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(144,value)
	log("p2 w:%d", value)
	countwrites(value,2,9)
	l.Unlock()
}

func p3() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	s8.Put(114,value)
	log("p3 w:%d", value)
	countwrites(value,3,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(332,&value)
	log("p3 r:%d", value)
	countreads(value,3,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(294,value)
	log("p3 w:%d", value)
	countwrites(value,3,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(93,&value)
	log("p3 r:%d", value)
	countreads(value,3,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(342,&value)
	log("p3 r:%d", value)
	countreads(value,3,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(502,&value)
	log("p3 r:%d", value)
	countreads(value,3,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(316,&value)
	log("p3 r:%d", value)
	countreads(value,3,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(476,&value)
	log("p3 r:%d", value)
	countreads(value,3,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(105,&value)
	log("p3 r:%d", value)
	countreads(value,3,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(63,value)
	log("p3 w:%d", value)
	countwrites(value,3,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(302,value)
	log("p3 w:%d", value)
	countwrites(value,3,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(267,value)
	log("p3 w:%d", value)
	countwrites(value,3,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(308,&value)
	log("p3 r:%d", value)
	countreads(value,3,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(359,&value)
	log("p3 r:%d", value)
	countreads(value,3,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(258,value)
	log("p3 w:%d", value)
	countwrites(value,3,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p3 r:%d", value)
	countreads(value,3,21)
	l.Unlock()
}

func p4() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	s32.Put(498,value)
	log("p4 w:%d", value)
	countwrites(value,4,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(278,value)
	log("p4 w:%d", value)
	countwrites(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(111,&value)
	log("p4 r:%d", value)
	countreads(value,4,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(143,value)
	log("p4 w:%d", value)
	countwrites(value,4,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(114,&value)
	log("p4 r:%d", value)
	countreads(value,4,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(251,&value)
	log("p4 r:%d", value)
	countreads(value,4,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(468,&value)
	log("p4 r:%d", value)
	countreads(value,4,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(231,&value)
	log("p4 r:%d", value)
	countreads(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(27,&value)
	log("p4 r:%d", value)
	countreads(value,4,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(282,value)
	log("p4 w:%d", value)
	countwrites(value,4,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(419,value)
	log("p4 w:%d", value)
	countwrites(value,4,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(235,&value)
	log("p4 r:%d", value)
	countreads(value,4,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(362,value)
	log("p4 w:%d", value)
	countwrites(value,4,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(335,&value)
	log("p4 r:%d", value)
	countreads(value,4,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(420,value)
	log("p4 w:%d", value)
	countwrites(value,4,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(57,&value)
	log("p4 r:%d", value)
	countreads(value,4,4)
	l.Unlock()
}

func p5() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s23.QueryP(357,&value)
	log("p5 r:%d", value)
	countreads(value,5,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(262,&value)
	log("p5 r:%d", value)
	countreads(value,5,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(151,&value)
	log("p5 r:%d", value)
	countreads(value,5,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(27,&value)
	log("p5 r:%d", value)
	countreads(value,5,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(482,&value)
	log("p5 r:%d", value)
	countreads(value,5,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(143,value)
	log("p5 w:%d", value)
	countwrites(value,5,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(383,&value)
	log("p5 r:%d", value)
	countreads(value,5,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(125,value)
	log("p5 w:%d", value)
	countwrites(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(390,&value)
	log("p5 r:%d", value)
	countreads(value,5,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(45,value)
	log("p5 w:%d", value)
	countwrites(value,5,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(342,&value)
	log("p5 r:%d", value)
	countreads(value,5,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(27,&value)
	log("p5 r:%d", value)
	countreads(value,5,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(122,&value)
	log("p5 r:%d", value)
	countreads(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(275,&value)
	log("p5 r:%d", value)
	countreads(value,5,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(114,&value)
	log("p5 r:%d", value)
	countreads(value,5,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(130,&value)
	log("p5 r:%d", value)
	countreads(value,5,9)
	l.Unlock()
}

func p6() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s9.QueryP(139,&value)
	log("p6 r:%d", value)
	countreads(value,6,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(21,&value)
	log("p6 r:%d", value)
	countreads(value,6,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(50,value)
	log("p6 w:%d", value)
	countwrites(value,6,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(92,&value)
	log("p6 r:%d", value)
	countreads(value,6,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(406,&value)
	log("p6 r:%d", value)
	countreads(value,6,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(462,&value)
	log("p6 r:%d", value)
	countreads(value,6,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(99,&value)
	log("p6 r:%d", value)
	countreads(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(205,&value)
	log("p6 r:%d", value)
	countreads(value,6,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(345,&value)
	log("p6 r:%d", value)
	countreads(value,6,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(244,&value)
	log("p6 r:%d", value)
	countreads(value,6,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(20,&value)
	log("p6 r:%d", value)
	countreads(value,6,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p6 r:%d", value)
	countreads(value,6,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(101,&value)
	log("p6 r:%d", value)
	countreads(value,6,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p6 r:%d", value)
	countreads(value,6,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(484,&value)
	log("p6 r:%d", value)
	countreads(value,6,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(388,value)
	log("p6 w:%d", value)
	countwrites(value,6,25)
	l.Unlock()
}

func p7() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s4.QueryP(54,&value)
	log("p7 r:%d", value)
	countreads(value,7,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(307,&value)
	log("p7 r:%d", value)
	countreads(value,7,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(445,value)
	log("p7 w:%d", value)
	countwrites(value,7,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(188,&value)
	log("p7 r:%d", value)
	countreads(value,7,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(104,value)
	log("p7 w:%d", value)
	countwrites(value,7,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(382,&value)
	log("p7 r:%d", value)
	countreads(value,7,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(340,value)
	log("p7 w:%d", value)
	countwrites(value,7,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(301,value)
	log("p7 w:%d", value)
	countwrites(value,7,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(217,&value)
	log("p7 r:%d", value)
	countreads(value,7,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(101,&value)
	log("p7 r:%d", value)
	countreads(value,7,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(298,&value)
	log("p7 r:%d", value)
	countreads(value,7,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(125,&value)
	log("p7 r:%d", value)
	countreads(value,7,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(131,&value)
	log("p7 r:%d", value)
	countreads(value,7,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(204,&value)
	log("p7 r:%d", value)
	countreads(value,7,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p7 r:%d", value)
	countreads(value,7,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(230,&value)
	log("p7 r:%d", value)
	countreads(value,7,15)
	l.Unlock()
}

func p8() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s16.QueryP(252,&value)
	log("p8 r:%d", value)
	countreads(value,8,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(298,&value)
	log("p8 r:%d", value)
	countreads(value,8,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(457,&value)
	log("p8 r:%d", value)
	countreads(value,8,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(215,&value)
	log("p8 r:%d", value)
	countreads(value,8,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(234,&value)
	log("p8 r:%d", value)
	countreads(value,8,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(477,&value)
	log("p8 r:%d", value)
	countreads(value,8,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(410,&value)
	log("p8 r:%d", value)
	countreads(value,8,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(503,&value)
	log("p8 r:%d", value)
	countreads(value,8,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p8 r:%d", value)
	countreads(value,8,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(19,&value)
	log("p8 r:%d", value)
	countreads(value,8,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(242,&value)
	log("p8 r:%d", value)
	countreads(value,8,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(448,&value)
	log("p8 r:%d", value)
	countreads(value,8,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(426,&value)
	log("p8 r:%d", value)
	countreads(value,8,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(31,value)
	log("p8 w:%d", value)
	countwrites(value,8,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(321,&value)
	log("p8 r:%d", value)
	countreads(value,8,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(360,&value)
	log("p8 r:%d", value)
	countreads(value,8,23)
	l.Unlock()
}

func p9() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s12.QueryP(183,&value)
	log("p9 r:%d", value)
	countreads(value,9,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(226,&value)
	log("p9 r:%d", value)
	countreads(value,9,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(402,&value)
	log("p9 r:%d", value)
	countreads(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(311,&value)
	log("p9 r:%d", value)
	countreads(value,9,20)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(406,value)
	log("p9 w:%d", value)
	countwrites(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(417,&value)
	log("p9 r:%d", value)
	countreads(value,9,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(380,value)
	log("p9 w:%d", value)
	countwrites(value,9,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(29,&value)
	log("p9 r:%d", value)
	countreads(value,9,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(29,&value)
	log("p9 r:%d", value)
	countreads(value,9,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(401,value)
	log("p9 w:%d", value)
	countwrites(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(115,value)
	log("p9 w:%d", value)
	countwrites(value,9,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(310,&value)
	log("p9 r:%d", value)
	countreads(value,9,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(404,&value)
	log("p9 r:%d", value)
	countreads(value,9,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(427,value)
	log("p9 w:%d", value)
	countwrites(value,9,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(86,&value)
	log("p9 r:%d", value)
	countreads(value,9,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(61,&value)
	log("p9 r:%d", value)
	countreads(value,9,4)
	l.Unlock()
}


func p10() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s8.QueryP(113,&value)
	log("p10 r:%d", value)
	countreads(value,10,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(127,value)
	log("p10 w:%d", value)
	countwrites(value,10,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(36,value)
	log("p10 w:%d", value)
	countwrites(value,10,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(82,&value)
	log("p10 r:%d", value)
	countreads(value,10,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(387,&value)
	log("p10 r:%d", value)
	countreads(value,10,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(115,&value)
	log("p10 r:%d", value)
	countreads(value,10,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(213,value)
	log("p10 w:%d", value)
	countwrites(value,10,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 20
	s20.Put(308,value)
	log("p10 w:%d", value)
	countwrites(value,10,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(177,&value)
	log("p10 r:%d", value)
	countreads(value,10,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(490,value)
	log("p10 w:%d", value)
	countwrites(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(468,&value)
	log("p10 r:%d", value)
	countreads(value,10,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(141,&value)
	log("p10 r:%d", value)
	countreads(value,10,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(495,&value)
	log("p10 r:%d", value)
	countreads(value,10,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(398,&value)
	log("p10 r:%d", value)
	countreads(value,10,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(347,value)
	log("p10 w:%d", value)
	countwrites(value,10,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(359,&value)
	log("p10 r:%d", value)
	countreads(value,10,23)
	l.Unlock()
}


func p11() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s19.QueryP(289,&value)
	log("p11 r:%d", value)
	countreads(value,11,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(82,&value)
	log("p11 r:%d", value)
	countreads(value,11,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(122,&value)
	log("p11 r:%d", value)
	countreads(value,11,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(218,&value)
	log("p11 r:%d", value)
	countreads(value,11,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(297,value)
	log("p11 w:%d", value)
	countwrites(value,11,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(187,&value)
	log("p11 r:%d", value)
	countreads(value,11,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(432,&value)
	log("p11 r:%d", value)
	countreads(value,11,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(259,&value)
	log("p11 r:%d", value)
	countreads(value,11,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(98,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(108,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(184,value)
	log("p11 w:%d", value)
	countwrites(value,11,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(100,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(101,&value)
	log("p11 r:%d", value)
	countreads(value,11,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(144,&value)
	log("p11 r:%d", value)
	countreads(value,11,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(256,&value)
	log("p11 r:%d", value)
	countreads(value,11,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(286,&value)
	log("p11 r:%d", value)
	countreads(value,11,18)
	l.Unlock()
}


func p12() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(465,&value)
	log("p12 r:%d", value)
	countreads(value,12,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(188,value)
	log("p12 w:%d", value)
	countwrites(value,12,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(106,&value)
	log("p12 r:%d", value)
	countreads(value,12,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(39,&value)
	log("p12 r:%d", value)
	countreads(value,12,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(478,value)
	log("p12 w:%d", value)
	countwrites(value,12,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(284,value)
	log("p12 w:%d", value)
	countwrites(value,12,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(70,&value)
	log("p12 r:%d", value)
	countreads(value,12,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(453,&value)
	log("p12 r:%d", value)
	countreads(value,12,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(355,value)
	log("p12 w:%d", value)
	countwrites(value,12,23)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(70,value)
	log("p12 w:%d", value)
	countwrites(value,12,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(428,&value)
	log("p12 r:%d", value)
	countreads(value,12,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p12 r:%d", value)
	countreads(value,12,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(96,value)
	log("p12 w:%d", value)
	countwrites(value,12,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(279,&value)
	log("p12 r:%d", value)
	countreads(value,12,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(97,value)
	log("p12 w:%d", value)
	countwrites(value,12,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(20,value)
	log("p12 w:%d", value)
	countwrites(value,12,2)
	l.Unlock()
}


func p13() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 8
	s8.Put(127,value)
	log("p13 w:%d", value)
	countwrites(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(462,&value)
	log("p13 r:%d", value)
	countreads(value,13,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(77,&value)
	log("p13 r:%d", value)
	countreads(value,13,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(404,&value)
	log("p13 r:%d", value)
	countreads(value,13,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(68,&value)
	log("p13 r:%d", value)
	countreads(value,13,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(401,&value)
	log("p13 r:%d", value)
	countreads(value,13,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(406,&value)
	log("p13 r:%d", value)
	countreads(value,13,26)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(421,value)
	log("p13 w:%d", value)
	countwrites(value,13,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(473,value)
	log("p13 w:%d", value)
	countwrites(value,13,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(403,&value)
	log("p13 r:%d", value)
	countreads(value,13,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(127,&value)
	log("p13 r:%d", value)
	countreads(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(128,value)
	log("p13 w:%d", value)
	countwrites(value,13,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(157,&value)
	log("p13 r:%d", value)
	countreads(value,13,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p13 w:%d", value)
	countwrites(value,13,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(388,value)
	log("p13 w:%d", value)
	countwrites(value,13,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(168,value)
	log("p13 w:%d", value)
	countwrites(value,13,11)
	l.Unlock()
}

func p14() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s18.QueryP(274,&value)
	log("p14 r:%d", value)
	countreads(value,14,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(17,value)
	log("p14 w:%d", value)
	countwrites(value,14,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(501,value)
	log("p14 w:%d", value)
	countwrites(value,14,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(494,value)
	log("p14 w:%d", value)
	countwrites(value,14,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(309,&value)
	log("p14 r:%d", value)
	countreads(value,14,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(213,&value)
	log("p14 r:%d", value)
	countreads(value,14,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(402,&value)
	log("p14 r:%d", value)
	countreads(value,14,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(474,&value)
	log("p14 r:%d", value)
	countreads(value,14,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(409,&value)
	log("p14 r:%d", value)
	countreads(value,14,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(32,&value)
	log("p14 r:%d", value)
	countreads(value,14,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(260,&value)
	log("p14 r:%d", value)
	countreads(value,14,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(98,value)
	log("p14 w:%d", value)
	countwrites(value,14,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(160,value)
	log("p14 w:%d", value)
	countwrites(value,14,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(428,&value)
	log("p14 r:%d", value)
	countreads(value,14,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(36,&value)
	log("p14 r:%d", value)
	countreads(value,14,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(185,value)
	log("p14 w:%d", value)
	countwrites(value,14,12)
	l.Unlock()
}


func p15() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s8.QueryP(115,&value)
	log("p15 r:%d", value)
	countreads(value,15,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(153,&value)
	log("p15 r:%d", value)
	countreads(value,15,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(498,&value)
	log("p15 r:%d", value)
	countreads(value,15,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(431,value)
	log("p15 w:%d", value)
	countwrites(value,15,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(211,value)
	log("p15 w:%d", value)
	countwrites(value,15,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(173,&value)
	log("p15 r:%d", value)
	countreads(value,15,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(418,value)
	log("p15 w:%d", value)
	countwrites(value,15,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p15 r:%d", value)
	countreads(value,15,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(23,value)
	log("p15 w:%d", value)
	countwrites(value,15,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(194,value)
	log("p15 w:%d", value)
	countwrites(value,15,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(88,&value)
	log("p15 r:%d", value)
	countreads(value,15,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(53,&value)
	log("p15 r:%d", value)
	countreads(value,15,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(416,&value)
	log("p15 r:%d", value)
	countreads(value,15,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(131,&value)
	log("p15 r:%d", value)
	countreads(value,15,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(291,value)
	log("p15 w:%d", value)
	countwrites(value,15,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(6,&value)
	log("p15 r:%d", value)
	countreads(value,15,1)
	l.Unlock()
}


func p16() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s30.QueryP(469,&value)
	log("p16 r:%d", value)
	countreads(value,16,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(406,&value)
	log("p16 r:%d", value)
	countreads(value,16,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(343,&value)
	log("p16 r:%d", value)
	countreads(value,16,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(398,&value)
	log("p16 r:%d", value)
	countreads(value,16,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(54,value)
	log("p16 w:%d", value)
	countwrites(value,16,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(236,value)
	log("p16 w:%d", value)
	countwrites(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(236,&value)
	log("p16 r:%d", value)
	countreads(value,16,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(441,&value)
	log("p16 r:%d", value)
	countreads(value,16,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(178,&value)
	log("p16 r:%d", value)
	countreads(value,16,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(165,&value)
	log("p16 r:%d", value)
	countreads(value,16,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(10,&value)
	log("p16 r:%d", value)
	countreads(value,16,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(336,&value)
	log("p16 r:%d", value)
	countreads(value,16,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(145,&value)
	log("p16 r:%d", value)
	countreads(value,16,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(104,&value)
	log("p16 r:%d", value)
	countreads(value,16,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(348,&value)
	log("p16 r:%d", value)
	countreads(value,16,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(187,&value)
	log("p16 r:%d", value)
	countreads(value,16,12)
	l.Unlock()
}


func p17() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s29.QueryP(449,&value)
	log("p17 r:%d", value)
	countreads(value,17,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 16
	s16.Put(245,value)
	log("p17 w:%d", value)
	countwrites(value,17,16)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(471,value)
	log("p17 w:%d", value)
	countwrites(value,17,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(495,value)
	log("p17 w:%d", value)
	countwrites(value,17,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(184,&value)
	log("p17 r:%d", value)
	countreads(value,17,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(64,&value)
	log("p17 r:%d", value)
	countreads(value,17,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 2
	s2.Put(22,value)
	log("p17 w:%d", value)
	countwrites(value,17,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(271,value)
	log("p17 w:%d", value)
	countwrites(value,17,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(73,&value)
	log("p17 r:%d", value)
	countreads(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(77,value)
	log("p17 w:%d", value)
	countwrites(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(65,value)
	log("p17 w:%d", value)
	countwrites(value,17,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(477,&value)
	log("p17 r:%d", value)
	countreads(value,17,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(420,&value)
	log("p17 r:%d", value)
	countreads(value,17,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s22.QueryP(346,&value)
	log("p17 r:%d", value)
	countreads(value,17,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(282,&value)
	log("p17 r:%d", value)
	countreads(value,17,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(34,&value)
	log("p17 r:%d", value)
	countreads(value,17,3)
	l.Unlock()
}


func p18() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(202,value)
	log("p18 w:%d", value)
	countwrites(value,18,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p18 r:%d", value)
	countreads(value,18,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(496,&value)
	log("p18 r:%d", value)
	countreads(value,18,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(463,&value)
	log("p18 r:%d", value)
	countreads(value,18,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(388,&value)
	log("p18 r:%d", value)
	countreads(value,18,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(464,&value)
	log("p18 r:%d", value)
	countreads(value,18,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(432,&value)
	log("p18 r:%d", value)
	countreads(value,18,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(442,&value)
	log("p18 r:%d", value)
	countreads(value,18,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(388,&value)
	log("p18 r:%d", value)
	countreads(value,18,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(115,&value)
	log("p18 r:%d", value)
	countreads(value,18,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(104,&value)
	log("p18 r:%d", value)
	countreads(value,18,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 13
	s13.Put(194,value)
	log("p18 w:%d", value)
	countwrites(value,18,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(196,&value)
	log("p18 r:%d", value)
	countreads(value,18,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(475,value)
	log("p18 w:%d", value)
	countwrites(value,18,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(431,&value)
	log("p18 r:%d", value)
	countreads(value,18,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(124,value)
	log("p18 w:%d", value)
	countwrites(value,18,8)
	l.Unlock()
}


func p19() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	s1.Put(9,value)
	log("p19 w:%d", value)
	countwrites(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(499,&value)
	log("p19 r:%d", value)
	countreads(value,19,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(1,value)
	log("p19 w:%d", value)
	countwrites(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(5,&value)
	log("p19 r:%d", value)
	countreads(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(152,&value)
	log("p19 r:%d", value)
	countreads(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(313,&value)
	log("p19 r:%d", value)
	countreads(value,19,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(215,&value)
	log("p19 r:%d", value)
	countreads(value,19,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(380,&value)
	log("p19 r:%d", value)
	countreads(value,19,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(148,&value)
	log("p19 r:%d", value)
	countreads(value,19,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(408,&value)
	log("p19 r:%d", value)
	countreads(value,19,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(45,&value)
	log("p19 r:%d", value)
	countreads(value,19,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(173,value)
	log("p19 w:%d", value)
	countwrites(value,19,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(7,&value)
	log("p19 r:%d", value)
	countreads(value,19,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(170,&value)
	log("p19 r:%d", value)
	countreads(value,19,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(97,&value)
	log("p19 r:%d", value)
	countreads(value,19,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(297,&value)
	log("p19 r:%d", value)
	countreads(value,19,19)
	l.Unlock()
}


func p20() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p20 w:%d", value)
	countwrites(value,20,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 23
	s23.Put(367,value)
	log("p20 w:%d", value)
	countwrites(value,20,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(411,&value)
	log("p20 r:%d", value)
	countreads(value,20,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(80,&value)
	log("p20 r:%d", value)
	countreads(value,20,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 12
	s12.Put(187,value)
	log("p20 w:%d", value)
	countwrites(value,20,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(35,&value)
	log("p20 r:%d", value)
	countreads(value,20,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(10,value)
	log("p20 w:%d", value)
	countwrites(value,20,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(476,value)
	log("p20 w:%d", value)
	countwrites(value,20,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(427,value)
	log("p20 w:%d", value)
	countwrites(value,20,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(369,&value)
	log("p20 r:%d", value)
	countreads(value,20,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(266,&value)
	log("p20 r:%d", value)
	countreads(value,20,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(73,&value)
	log("p20 r:%d", value)
	countreads(value,20,5)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(52,&value)
	log("p20 r:%d", value)
	countreads(value,20,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(386,&value)
	log("p20 r:%d", value)
	countreads(value,20,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(418,&value)
	log("p20 r:%d", value)
	countreads(value,20,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(123,value)
	log("p20 w:%d", value)
	countwrites(value,20,8)
	l.Unlock()
}


func p21() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s19.QueryP(296,&value)
	log("p21 r:%d", value)
	countreads(value,21,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(262,&value)
	log("p21 r:%d", value)
	countreads(value,21,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(179,&value)
	log("p21 r:%d", value)
	countreads(value,21,12)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(392,&value)
	log("p21 r:%d", value)
	countreads(value,21,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(175,value)
	log("p21 w:%d", value)
	countwrites(value,21,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(493,&value)
	log("p21 r:%d", value)
	countreads(value,21,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(479,value)
	log("p21 w:%d", value)
	countwrites(value,21,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(323,&value)
	log("p21 r:%d", value)
	countreads(value,21,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(208,&value)
	log("p21 r:%d", value)
	countreads(value,21,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(83,value)
	log("p21 w:%d", value)
	countwrites(value,21,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(58,value)
	log("p21 w:%d", value)
	countwrites(value,21,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(197,&value)
	log("p21 r:%d", value)
	countreads(value,21,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(421,value)
	log("p21 w:%d", value)
	countwrites(value,21,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(445,&value)
	log("p21 r:%d", value)
	countreads(value,21,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(215,value)
	log("p21 w:%d", value)
	countwrites(value,21,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(401,&value)
	log("p21 r:%d", value)
	countreads(value,21,26)
	l.Unlock()
}


func p22() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s22.QueryP(352,&value)
	log("p22 r:%d", value)
	countreads(value,22,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(95,&value)
	log("p22 r:%d", value)
	countreads(value,22,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(228,&value)
	log("p22 r:%d", value)
	countreads(value,22,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(29,&value)
	log("p22 r:%d", value)
	countreads(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(58,&value)
	log("p22 r:%d", value)
	countreads(value,22,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(169,&value)
	log("p22 r:%d", value)
	countreads(value,22,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(435,&value)
	log("p22 r:%d", value)
	countreads(value,22,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(322,value)
	log("p22 w:%d", value)
	countwrites(value,22,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(221,&value)
	log("p22 r:%d", value)
	countreads(value,22,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(28,&value)
	log("p22 r:%d", value)
	countreads(value,22,2)
	l.Unlock()

	delay()
	l.Lock()
	value = 28
	s28.Put(442,value)
	log("p22 w:%d", value)
	countwrites(value,22,28)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(386,value)
	log("p22 w:%d", value)
	countwrites(value,22,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(499,&value)
	log("p22 r:%d", value)
	countreads(value,22,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(190,&value)
	log("p22 r:%d", value)
	countreads(value,22,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(38,value)
	log("p22 w:%d", value)
	countwrites(value,22,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(286,&value)
	log("p22 r:%d", value)
	countreads(value,22,18)
	l.Unlock()
}


func p23() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(332,value)
	log("p23 w:%d", value)
	countwrites(value,23,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(401,value)
	log("p23 w:%d", value)
	countwrites(value,23,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(126,&value)
	log("p23 r:%d", value)
	countreads(value,23,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(468,&value)
	log("p23 r:%d", value)
	countreads(value,23,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(390,&value)
	log("p23 r:%d", value)
	countreads(value,23,25)
	l.Unlock()

	delay()
	l.Lock()
	value = 7
	s7.Put(102,value)
	log("p23 w:%d", value)
	countwrites(value,23,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(315,&value)
	log("p23 r:%d", value)
	countreads(value,23,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(427,&value)
	log("p23 r:%d", value)
	countreads(value,23,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(85,&value)
	log("p23 r:%d", value)
	countreads(value,23,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(116,&value)
	log("p23 r:%d", value)
	countreads(value,23,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(310,&value)
	log("p23 r:%d", value)
	countreads(value,23,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(290,&value)
	log("p23 r:%d", value)
	countreads(value,23,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(277,&value)
	log("p23 r:%d", value)
	countreads(value,23,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(233,&value)
	log("p23 r:%d", value)
	countreads(value,23,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(441,&value)
	log("p23 r:%d", value)
	countreads(value,23,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(379,&value)
	log("p23 r:%d", value)
	countreads(value,23,24)
	l.Unlock()
}


func p24() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(347,value)
	log("p24 w:%d", value)
	countwrites(value,24,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(457,&value)
	log("p24 r:%d", value)
	countreads(value,24,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(150,value)
	log("p24 w:%d", value)
	countwrites(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(144,value)
	log("p24 w:%d", value)
	countwrites(value,24,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(146,&value)
	log("p24 r:%d", value)
	countreads(value,24,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(253,&value)
	log("p24 r:%d", value)
	countreads(value,24,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(108,&value)
	log("p24 r:%d", value)
	countreads(value,24,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(311,&value)
	log("p24 r:%d", value)
	countreads(value,24,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(213,&value)
	log("p24 r:%d", value)
	countreads(value,24,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(47,&value)
	log("p24 r:%d", value)
	countreads(value,24,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(509,&value)
	log("p24 r:%d", value)
	countreads(value,24,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(52,&value)
	log("p24 r:%d", value)
	countreads(value,24,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(372,&value)
	log("p24 r:%d", value)
	countreads(value,24,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(368,&value)
	log("p24 r:%d", value)
	countreads(value,24,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(253,&value)
	log("p24 r:%d", value)
	countreads(value,24,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(265,&value)
	log("p24 r:%d", value)
	countreads(value,24,17)
	l.Unlock()
}


func p25() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s3.QueryP(38,&value)
	log("p25 r:%d", value)
	countreads(value,25,3)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(488,value)
	log("p25 w:%d", value)
	countwrites(value,25,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(90,&value)
	log("p25 r:%d", value)
	countreads(value,25,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(228,&value)
	log("p25 r:%d", value)
	countreads(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(14,value)
	log("p25 w:%d", value)
	countwrites(value,25,1)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(429,value)
	log("p25 w:%d", value)
	countwrites(value,25,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(327,&value)
	log("p25 r:%d", value)
	countreads(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(238,value)
	log("p25 w:%d", value)
	countwrites(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(486,&value)
	log("p25 r:%d", value)
	countreads(value,25,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(228,&value)
	log("p25 r:%d", value)
	countreads(value,25,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(104,&value)
	log("p25 r:%d", value)
	countreads(value,25,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 21
	s21.Put(327,value)
	log("p25 w:%d", value)
	countwrites(value,25,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(50,value)
	log("p25 w:%d", value)
	countwrites(value,25,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 19
	s19.Put(290,value)
	log("p25 w:%d", value)
	countwrites(value,25,19)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(143,&value)
	log("p25 r:%d", value)
	countreads(value,25,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(410,value)
	log("p25 w:%d", value)
	countwrites(value,25,26)
	l.Unlock()
}


func p26() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 4
	s4.Put(52,value)
	log("p26 w:%d", value)
	countwrites(value,26,4)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(328,&value)
	log("p26 r:%d", value)
	countreads(value,26,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 27
	s27.Put(431,value)
	log("p26 w:%d", value)
	countwrites(value,26,27)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(214,value)
	log("p26 w:%d", value)
	countwrites(value,26,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p26 r:%d", value)
	countreads(value,26,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(11,value)
	log("p26 w:%d", value)
	countwrites(value,26,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(268,&value)
	log("p26 r:%d", value)
	countreads(value,26,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 1
	s1.Put(2,value)
	log("p26 w:%d", value)
	countwrites(value,26,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(112,&value)
	log("p26 r:%d", value)
	countreads(value,26,7)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(1,&value)
	log("p26 r:%d", value)
	countreads(value,26,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(67,&value)
	log("p26 r:%d", value)
	countreads(value,26,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 9
	s9.Put(132,value)
	log("p26 w:%d", value)
	countwrites(value,26,9)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(504,&value)
	log("p26 r:%d", value)
	countreads(value,26,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(223,value)
	log("p26 w:%d", value)
	countwrites(value,26,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(235,value)
	log("p26 w:%d", value)
	countwrites(value,26,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s26.QueryP(413,&value)
	log("p26 r:%d", value)
	countreads(value,26,26)
	l.Unlock()
}


func p27() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = -1
	s2.QueryP(30,&value)
	log("p27 r:%d", value)
	countreads(value,27,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(507,&value)
	log("p27 r:%d", value)
	countreads(value,27,32)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(66,value)
	log("p27 w:%d", value)
	countwrites(value,27,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(231,value)
	log("p27 w:%d", value)
	countwrites(value,27,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(120,&value)
	log("p27 r:%d", value)
	countreads(value,27,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p27 r:%d", value)
	countreads(value,27,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(384,&value)
	log("p27 r:%d", value)
	countreads(value,27,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 29
	s29.Put(460,value)
	log("p27 w:%d", value)
	countwrites(value,27,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(220,&value)
	log("p27 r:%d", value)
	countreads(value,27,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s16.QueryP(251,&value)
	log("p27 r:%d", value)
	countreads(value,27,16)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(48,&value)
	log("p27 r:%d", value)
	countreads(value,27,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s32.QueryP(509,&value)
	log("p27 r:%d", value)
	countreads(value,27,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(165,&value)
	log("p27 r:%d", value)
	countreads(value,27,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(156,&value)
	log("p27 r:%d", value)
	countreads(value,27,10)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(37,value)
	log("p27 w:%d", value)
	countwrites(value,27,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(375,&value)
	log("p27 r:%d", value)
	countreads(value,27,24)
	l.Unlock()
}


func p28() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 21
	s21.Put(331,value)
	log("p28 w:%d", value)
	countwrites(value,28,21)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(122,value)
	log("p28 w:%d", value)
	countwrites(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 8
	s8.Put(116,value)
	log("p28 w:%d", value)
	countwrites(value,28,8)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(219,value)
	log("p28 w:%d", value)
	countwrites(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s10.QueryP(160,&value)
	log("p28 r:%d", value)
	countreads(value,28,10)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s28.QueryP(436,&value)
	log("p28 r:%d", value)
	countreads(value,28,28)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(200,&value)
	log("p28 r:%d", value)
	countreads(value,28,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(462,&value)
	log("p28 r:%d", value)
	countreads(value,28,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 32
	s32.Put(508,value)
	log("p28 w:%d", value)
	countwrites(value,28,32)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s9.QueryP(133,&value)
	log("p28 r:%d", value)
	countreads(value,28,9)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(274,value)
	log("p28 w:%d", value)
	countwrites(value,28,18)
	l.Unlock()

	delay()
	l.Lock()
	value = 14
	s14.Put(213,value)
	log("p28 w:%d", value)
	countwrites(value,28,14)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s24.QueryP(376,&value)
	log("p28 r:%d", value)
	countreads(value,28,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(322,&value)
	log("p28 r:%d", value)
	countreads(value,28,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(87,&value)
	log("p28 r:%d", value)
	countreads(value,28,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s15.QueryP(225,&value)
	log("p28 r:%d", value)
	countreads(value,28,15)
	l.Unlock()
}


func p29() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 1
	s1.Put(13,value)
	log("p29 w:%d", value)
	countwrites(value,29,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(30,&value)
	log("p29 r:%d", value)
	countreads(value,29,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(170,&value)
	log("p29 r:%d", value)
	countreads(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(57,&value)
	log("p29 r:%d", value)
	countreads(value,29,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(468,value)
	log("p29 w:%d", value)
	countwrites(value,29,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(204,&value)
	log("p29 r:%d", value)
	countreads(value,29,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(460,&value)
	log("p29 r:%d", value)
	countreads(value,29,29)
	l.Unlock()

	delay()
	l.Lock()
	value = 6
	s6.Put(81,value)
	log("p29 w:%d", value)
	countwrites(value,29,6)
	l.Unlock()

	delay()
	l.Lock()
	value = 11
	s11.Put(176,value)
	log("p29 w:%d", value)
	countwrites(value,29,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(197,&value)
	log("p29 r:%d", value)
	countreads(value,29,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s3.QueryP(37,&value)
	log("p29 r:%d", value)
	countreads(value,29,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(366,&value)
	log("p29 r:%d", value)
	countreads(value,29,23)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(50,&value)
	log("p29 r:%d", value)
	countreads(value,29,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 30
	s30.Put(466,value)
	log("p29 w:%d", value)
	countwrites(value,29,30)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s20.QueryP(318,&value)
	log("p29 r:%d", value)
	countreads(value,29,20)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s23.QueryP(365,&value)
	log("p29 r:%d", value)
	countreads(value,29,23)
	l.Unlock()
}


func p30() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 13
	s13.Put(205,value)
	log("p30 w:%d", value)
	countwrites(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(388,&value)
	log("p30 r:%d", value)
	countreads(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s30.QueryP(480,&value)
	log("p30 r:%d", value)
	countreads(value,30,30)
	l.Unlock()

	delay()
	l.Lock()
	value = 3
	s3.Put(45,value)
	log("p30 w:%d", value)
	countwrites(value,30,3)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s21.QueryP(328,&value)
	log("p30 r:%d", value)
	countreads(value,30,21)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s5.QueryP(66,&value)
	log("p30 r:%d", value)
	countreads(value,30,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 22
	s22.Put(340,value)
	log("p30 w:%d", value)
	countwrites(value,30,22)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(484,value)
	log("p30 w:%d", value)
	countwrites(value,30,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 4
	s4.Put(61,value)
	log("p30 w:%d", value)
	countwrites(value,30,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(388,value)
	log("p30 w:%d", value)
	countwrites(value,30,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(197,&value)
	log("p30 r:%d", value)
	countreads(value,30,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 15
	s15.Put(228,value)
	log("p30 w:%d", value)
	countwrites(value,30,15)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s17.QueryP(269,&value)
	log("p30 r:%d", value)
	countreads(value,30,17)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(166,&value)
	log("p30 r:%d", value)
	countreads(value,30,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(453,&value)
	log("p30 r:%d", value)
	countreads(value,30,29)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s29.QueryP(460,&value)
	log("p30 r:%d", value)
	countreads(value,30,29)
	l.Unlock()
}


func p31() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(340,value)
	log("p31 w:%d", value)
	countwrites(value,31,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s8.QueryP(120,&value)
	log("p31 r:%d", value)
	countreads(value,31,8)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(167,&value)
	log("p31 r:%d", value)
	countreads(value,31,11)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(280,value)
	log("p31 w:%d", value)
	countwrites(value,31,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s18.QueryP(287,&value)
	log("p31 r:%d", value)
	countreads(value,31,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s4.QueryP(51,&value)
	log("p31 r:%d", value)
	countreads(value,31,4)
	l.Unlock()

	delay()
	l.Lock()
	value = 26
	s26.Put(411,value)
	log("p31 w:%d", value)
	countwrites(value,31,26)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s1.QueryP(14,&value)
	log("p31 r:%d", value)
	countreads(value,31,1)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s14.QueryP(223,&value)
	log("p31 r:%d", value)
	countreads(value,31,14)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(485,value)
	log("p31 w:%d", value)
	countwrites(value,31,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 31
	s31.Put(490,value)
	log("p31 w:%d", value)
	countwrites(value,31,31)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(384,value)
	log("p31 w:%d", value)
	countwrites(value,31,24)
	l.Unlock()

	delay()
	l.Lock()
	value = 5
	s5.Put(68,value)
	log("p31 w:%d", value)
	countwrites(value,31,5)
	l.Unlock()

	delay()
	l.Lock()
	value = 24
	s24.Put(375,value)
	log("p31 w:%d", value)
	countwrites(value,31,24)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(484,&value)
	log("p31 r:%d", value)
	countreads(value,31,31)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s31.QueryP(489,&value)
	log("p31 r:%d", value)
	countreads(value,31,31)
	l.Unlock()
}


func p32() {
	defer wg.Done()
	var value int

	delay()
	l.Lock()
	value = 22
	s22.Put(352,value)
	log("p32 w:%d", value)
	countwrites(value,32,22)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(173,&value)
	log("p32 r:%d", value)
	countreads(value,32,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s12.QueryP(183,&value)
	log("p32 r:%d", value)
	countreads(value,32,12)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(280,value)
	log("p32 w:%d", value)
	countwrites(value,32,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s19.QueryP(302,&value)
	log("p32 r:%d", value)
	countreads(value,32,19)
	l.Unlock()

	delay()
	l.Lock()
	value = 18
	s18.Put(282,value)
	log("p32 w:%d", value)
	countwrites(value,32,18)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s2.QueryP(21,&value)
	log("p32 r:%d", value)
	countreads(value,32,2)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s11.QueryP(172,&value)
	log("p32 r:%d", value)
	countreads(value,32,11)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s7.QueryP(101,&value)
	log("p32 r:%d", value)
	countreads(value,32,7)
	l.Unlock()

	delay()
	l.Lock()
	value = 17
	s17.Put(264,value)
	log("p32 w:%d", value)
	countwrites(value,32,17)
	l.Unlock()

	delay()
	l.Lock()
	value = 25
	s25.Put(396,value)
	log("p32 w:%d", value)
	countwrites(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s27.QueryP(431,&value)
	log("p32 r:%d", value)
	countreads(value,32,27)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s25.QueryP(398,&value)
	log("p32 r:%d", value)
	countreads(value,32,25)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s6.QueryP(86,&value)
	log("p32 r:%d", value)
	countreads(value,32,6)
	l.Unlock()

	delay()
	l.Lock()
	value = -1
	s13.QueryP(206,&value)
	log("p32 r:%d", value)
	countreads(value,32,13)
	l.Unlock()

	delay()
	l.Lock()
	value = 10
	s10.Put(155,value)
	log("p32 w:%d", value)
	countwrites(value,32,10)
	l.Unlock()
}


func log(format string, a ...interface{}) {
	if debug {
    	fmt.Fprintf(os.Stdout, format+"\n", a ...)
    }
}

func delay() {
	time.Sleep(time.Duration(rand.Int63n(75))*time.Millisecond)
}

func countreads(value int, localspace int, targetspace int) {
	if value == -1 {
		reads_insuccess+=1
	} else {
		reads_success+=1
	}

	if localspace == targetspace {
		reads_local+=1
	} else {
		reads_remote+=1
	}
}

func countwrites(value int, localspace int, targetspace int) {
	writestotal+=1

	if localspace == targetspace {
		writes_local+=1
	} else {
		writes_remote+=1
	}
}
