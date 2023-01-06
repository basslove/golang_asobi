package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// 並行処理(Concurrency)
// 並列処理(Parallelism)
// - 並行処理: ある時間の範囲において、複数のタスクを扱うこと
// = 並列処理: ある時間の点において、複数のタスクを扱うこと
// ------------------------
// - 並行処理は、複数個のスレッドを共通の期間内で実行する能力のことです。
// = 並列処理は、複数個のスレッドを同時に実行する能力のことです。
// ------------------------
// - 並行処理は、複数の処理を独立に実行できる構成のこと。
// = 並列処理は、複数の処理を同時に実行すること。
// ------------------------
// - 並行処理は、問題解決の手段としてのプログラミングパターン。
// - 並列処理は、並行処理を可能にするハードウェアの特性。
// ------------------------
// Go言語で作れるのは「コード/ソフトウェア」であり、それらの性質を指し示すのは「並行性」
// 実行時間が早くなる(かもしれない)から
// 現実世界での事象が独立性・並列性を持つから
// ------------------------
// 難易度
// - コードの実行順が予測できない
// - 競合状態を避ける必要がある(非アトミックな処理を並行して行う)(sync/atomic)
// - 共有メモリに正しくアクセスしないといけない(デッドロック)
// - CPUをたくさん積んでプログラムが早く動くかどうかは、そのプログラムで解決したい問題構造に依存する。
// - 並列処理で本当に処理を早くできるのは、解決したい問題が本質的に並列な構造を持つ場合のみである。

// スレッド導入の利点
// スレッドは、プロセスの中にある「並列可能なひとまとまりの命令処理」の単位
// 「プロセス」と「スレッド」の最も大きな違い
// - 初期化(fork), コピー(clone)
// プロセスとは別に、わざわざ「リソースを共有するプロセス」であるスレッドという概念を導入することでなんのメリットがあるのでしょうか。
// 考えられるメリットとしては2つあります。
// - メモリを節約
// - プロセス切り替えよりスレッド切り替えの方がコストが低い

// go並行処理
// - ゴールーチン
// - sync.WaitGroup
// - チャネル
func no1() {
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	fmt.Printf("Today's your lucky number is %d!\n", num)
}

// メインゴールーチンはチャネルcから値を受信するまでブロックされるので,取得前にプログラムが終了する ということはない。
// そのため、これはsync.WaitGroupを使った待ち合わせを行わなくてOK。
// どちらがいいのかは場合によるとは思いますが、複数個のゴールーチンを待つ場合にはsync.WaitGroupの方が実装が簡単
func no2(c chan<- int) {
	fmt.Println("...")

	// ランダム占い時間
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	c <- num
}

// チャネル
// - nilかどうか (例: var c chan intとしたまま、値が代入されなかったcはnilチャネル)
// - closed(=close済み)かどうか
// - バッファが空いているか / バッファに値があるか
// - 送信専用 / 受信専用だったりしないか
// 挙動
// - 値の送信
// - 値の受信
// - close操作
// チャネルについて
// - nilチャネルは常にブロックされる
// - closedなチャネルは決してブロックされることはない
// バッファなしのチャネル
// - 受信側の準備が整ってなければ、送信待ちのためにそのチャネルをブロックする
// - 送信側の準備が整ってなければ、受信待ちのためにそのチャネルをブロックする
// - 「送信側-受信側」のセットができるまではブロックされます = 送られた値は必ず受信しなくてはならない
// - チャネルというのは送受信だけではなくて実行同期のための機構でもある
func no3() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}

// dstをchannel
func no4() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	c := make(chan int)

	for _, s := range src {
		go func(s int, c chan int) {
			result := s * 2
			c <- result
		}(s, c)
	}

	for _ = range src {
		num := <-c
		dst = append(dst, num)
	}

	fmt.Println(dst)
	close(c)
}

// 排他
func no5() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	var mu sync.Mutex

	for _, s := range src {
		go func(s int) {
			result := s * 2
			mu.Lock()
			dst = append(dst, result)
			mu.Unlock()
		}(s)
	}

	time.Sleep(time.Second)
	fmt.Println(dst)
}

// 異なるゴールーチンで何かデータをやり取り・共有したい場合、とりうる手段としては主に2つあります。
// - チャネルをつかって値を送受信することでやり取りする。
// - sync.Mutex等のメモリロックを使って同じメモリを共有する
// - 実装が難しい危険なメモリ共有をするくらいなら、チャネルを使って値をやり取りした方が安全
// - 共有メモリ上のデータアクセス制御のために明示的なロックを使うよりは、Goではチャネルを使ってゴールーチン間でデータの参照結果をやり取りすることを推奨しています。
// - このやり方によって、ある時点で多くても1つのゴールーチンだけがデータにアクセスできることが保証されます。
// 「拘束」
func no6() <-chan int {
	// 1. チャネルを定義
	result := make(chan int)

	// 2. ゴールーチンを立てて
	go func() {
		defer close(result) // 4. closeするのを忘れずに

		// 3. その中で、resultチャネルに値を送る処理をする
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			result <- 1
		}
	}()

	// 5. 返り値にresultチャネルを返す
	return result
}

// select
func no7() {
	gen1, gen2 := make(chan int), make(chan int)

	// goルーチンを立てて、gen1やgen2に送信したりする
	select {
	case num := <-gen1: // gen1から受信できるとき
		fmt.Println(num)
	case num := <-gen2: // gen2から受信できるとき
		fmt.Println(num)
	default: // どっちも受信できないとき
		fmt.Println("neither chan cannot use")
	}

	//select {
	//case num := <-gen1:  // gen1から受信できるとき
	//	fmt.Println(num)
	//case channel<-1: // channelに送信できるとき
	//	fmt.Println("write channel to 1")
	//default:  // どっちも受信できないとき
	//	fmt.Println("neither chan cannot use")
	//}
}

// ゴルーチンリーク
// ゴールーチンを稼働したまま放っておくということは、そのスタック領域をGC(ガベージコレクト)されないまま放っておくという、
// パフォーマンス的にあまりよくない事態を引き起こしていることと同義なのです。
func no8(done chan struct{}) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
	LOOP:
		for {
			select {
			case <-done:
				break LOOP
			case result <- 1:
			}
		}
	}()
	return result
}

func no9() {
	done := make(chan bool)
	go func() {
		// 重たい処理
		time.Sleep(2 * time.Second)
		done <- true
	}()

	// 処理が終わるまで待つ
	<-done
}

func no10() {
	var (
		count int
		mu    sync.Mutex
	)
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			mu.Lock()
			count++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond) // 10[ms]スリープ
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			mu.Lock()
			count++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond) // 10[ms]スリープ
		}
		done <- true
	}()

	<-done
	<-done
	fmt.Println(count)
}

func no11() {
	var (
		count int
		mu    sync.Mutex
	)

	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan bool)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mu.Lock()
			count++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond) // 10[ms]スリープ
		}
		done <- true
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mu.Lock()
			count++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond) // 10[ms]スリープ
		}
	}()

	<-done
	// 2つのゴールーチンの処理が終わるまで待つ
	wg.Wait()
	fmt.Println(count)
}

func no12() {
	root := context.Background()
	ctx1, cancel := context.WithCancel(root)
	ctx2, _ := context.WithCancel(ctx1)

	var wg sync.WaitGroup
	wg.Add(2) // 2つのゴールーチンが終わるの待つため

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("cancel goroutine1")
				return
			default:
				fmt.Println("waint goroutine1")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx2.Done():
				fmt.Println("cancel goroutine2")
				return
			default:
				fmt.Println("waint goroutine2")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}

func no13() {
	root := context.Background()
	eg, ctx := errgroup.WithContext(root)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel goroutine1")
				return nil
			default:
				fmt.Println("waint goroutine1")
				time.Sleep(500 * time.Millisecond)
			}
		}
	})

	eg.Go(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("error")
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}

// goroutine 同期
func no14(code string, in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for {
			<-in
			fmt.Print(code)
			time.Sleep(50 * time.Millisecond)
			out <- 0
		}
	}()
	return out
}

// コア数
func no15() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumGoroutine())
}

func main() {
	fmt.Println("concurparallel")

	//// no1
	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	no1()
	//}()
	//wg.Wait()
	//
	//// no2
	//c := make(chan int)
	//go no2(c)
	//num := <-c
	//fmt.Printf("Today's your lucky number is %d!\n", num)
	//close(c)
	//
	//no3()
	//no4()
	//no5()
	//no6()
	//no7()
	//
	//// no8
	//done := make(chan struct{})
	//result := no8(done)
	//for i := 0; i < 5; i++ {
	//	fmt.Println(<-result)
	//}
	//close(done)
	//
	//no9()
	//
	//no10()
	//no11()
	//no12()
	//no13()
	//
	//// no14
	//runtime.GOMAXPROCS(1)
	//ch1 := make(chan int)
	//ch2 := no14("h", ch1)
	//ch3 := no14("e", ch2)
	//ch4 := no14("y", ch3)
	//ch5 := no14("!", ch4)
	//ch6 := no14(" ", ch5)
	//for i := 0; i < 10; i++ {
	//	ch1 <- 0
	//	<-ch6
	//}

	no15()
}
