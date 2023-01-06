package main

import (
	"context"
	"fmt"
	"time"
)

// context
// - contextパッケージで定義されているContext型は、処理の締め切り・キャンセル信号・API境界やプロセス間を横断する必要のあるリクエストスコープな値を伝達させる
// 役割は3つ
// - 処理の締め切りを伝達
// - キャンセル信号の伝播
// - リクエストスコープ値の伝達
// contextが役に立つのは、一つの処理が複数のゴールーチンをまたいで行われる場合である

// type Context interface {
//    Deadline() (deadline time.Time, ok bool)
//    Done() <-chan struct{}
//    Err() error
//    Value(key interface{}) interface{}
//}

// type Context interface {
//	Done() <-chan struct{}
// #####
//}

// contextの初期化
// context.Background()関数の返り値からは、「キャンセルされない」「deadlineも持たない」「共有する値も何も持たない」状態のcontextが得られます。いわば「context初期化のための関数」です。
// func Background() Context

// contexにキャンセル機能を追加
// context.Background()から得たまっさらなcontextをcontext.WithCancel()関数に渡すことで
// 「Doneメソッドからキャンセル有無が判断できるcontext」と「第一返り値のコンテキストをキャンセルするための関数」を得ることができます。
// func WithCancel(parent Context) (ctx Context, cancel CancelFunc)

// ctx, cancel := context.WithCancel(parentCtx)
// cancel()
//
// cancel()の実行により、ctx.Done()で得られるチャネルがcloseされる
// ctxはparentCtxとは別物なので、parentCtxはcancel()の影響を受けない

// contextのDoneメソッドを用いたキャンセル
// Doneメソッドチャネルでできるのは、あくまで「呼び出し側からキャンセルされているか否かの確認」のみ

// 自動タイムアウト機能の追加
// func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)
// context.WithDeadline関数を使うことで、指定された時刻に自動的にDoneメソッドチャネルがcloseされるcontextを作成することができます。
// WithDeadline関数から得られるcontextは、「引数として渡された親contextの設定を引き継いだ上で、Doneメソッドチャネルが第二引数で指定した時刻に自動closeされる新たなcontext」ものになります。

// contextパッケージには、2種類のエラー
// - 明示的なキャンセルとタイムアウトによるキャンセルで、後処理を変えたい(doneにはない)
// - var Canceled = errors.New("context canceled")
// - var DeadlineExceeded error = deadlineExceededError{}
// type Error interface {
//	error
//	Timeout() bool   // Is the error a timeout?
//	Temporary() bool // Is the error temporary?
//}

// contextのErrメソッド
// - contextがキャンセルされていない場合: nil
// - contextが明示的にキャンセルされていた場合: Canceledエラー
// - contextがタイムアウトしていた場合: DeadlineExceededエラー
//
//	type Context interface {
//		Err() error
//		// (以下略)
//	}

// context 値の伝達
// - WithCancel関数やWithTimeout関数を用いて、contextにキャンセル機能・タイムアウト機能を追加できたように、WithValue関数を使うことで、contextに値を追加
// - func WithValue(parent Context, key, val interface{}) Context
// - Valueメソッドによるcontext中の値抽出
// ctx := context.WithValue(parentCtx, "userID", 2)
// - interfaceValue := ctx.Value("userID") // keyが"userID"であるvalueを取り出す
// - intValue, ok := interfaceValue.(int)  // interface{}をint型にアサーション
// -----------------
// 付加: WithValue関数
// 取得: Valueメソッド
// type Context interface {
//	Value(key interface{}) interface{}
//	// (以下略)
// }
// func WithValue(parent Context, key, val interface{}) Context

func no1(ctx context.Context) {
	// ctxから、メインゴールーチン側の情報を得られる
	// (例)
	// ctx.Doneからキャンセル有無の確認
	// ctx.Deadlineで締め切り時間・締め切り有無の確認
	// ctx.Errでキャンセル理由の確認
	// ctx.Valueで値の共有
}

// 直列なゴールーチン
// 同じcontextを複数のゴールーチンに渡した場合、それらが直列の関係であろうが並列の関係であろうが同じ挙動となります。
// ゴールーチンの生死を制御するcontextが同じであるので、キャンセルタイミングも当然連動することとなります。
func no2() {
	ctx0 := context.Background()

	ctx1, _ := context.WithCancel(ctx0)
	// G1
	go func(ctx1 context.Context) {
		ctx2, cancel2 := context.WithCancel(ctx1)

		// G2-1
		go func(ctx2 context.Context) {
			// G2-2
			go func(ctx2 context.Context) {
				select {
				case <-ctx2.Done():
					fmt.Println("G2-2 canceled")
				}
			}(ctx2)

			select {
			case <-ctx2.Done():
				fmt.Println("G2-1 canceled")
			}
		}(ctx2)

		cancel2()

		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}

	}(ctx1)

	time.Sleep(time.Second)
}

// 並列なゴールーチン
// 同じcontextを複数のゴールーチンに渡した場合、それらが直列の関係であろうが並列の関係であろうが同じ挙動となります。
// ゴールーチンの生死を制御するcontextが同じであるので、キャンセルタイミングも当然連動することとなります。
func no3() {
	ctx0 := context.Background()

	ctx1, cancel1 := context.WithCancel(ctx0)
	// G1-1
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-1 canceled")
		}
	}(ctx1)

	// G1-2
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1-2 canceled")
		}
	}(ctx1)

	cancel1()

	time.Sleep(time.Second)
}

// 兄弟関係にあるcontext
// ここで、ctx1をキャンセルすると、G1のみが終了し、G2はその影響を受けることなく生きている
func no4() {
	ctx0 := context.Background()

	ctx1, cancel1 := context.WithCancel(ctx0)
	// G1
	go func(ctx1 context.Context) {
		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}
	}(ctx1)

	ctx2, _ := context.WithCancel(ctx0)
	// G2
	go func(ctx2 context.Context) {
		select {
		case <-ctx2.Done():
			fmt.Println("G2 canceled")
		}
	}(ctx2)

	cancel1()

	time.Sleep(time.Second)
}

// 親子関係にあるcontext
// ctx2のキャンセルのみを実行すると、ctx2ともつG2と、その子であるctx3を持つG3が揃って終了。
// 一方、ctx2の親であるctx1を持つG1は生きたままである
func no5() {
	ctx0 := context.Background()

	ctx1, _ := context.WithCancel(ctx0)
	// G1
	go func(ctx1 context.Context) {
		ctx2, cancel2 := context.WithCancel(ctx1)

		// G2
		go func(ctx2 context.Context) {
			ctx3, _ := context.WithCancel(ctx2)

			// G3
			go func(ctx3 context.Context) {
				select {
				case <-ctx3.Done():
					fmt.Println("G3 canceled")
				}
			}(ctx3)

			select {
			case <-ctx2.Done():
				fmt.Println("G2 canceled")
			}
		}(ctx2)

		cancel2()

		select {
		case <-ctx1.Done():
			fmt.Println("G1 canceled")
		}

	}(ctx1)

	time.Sleep(time.Second)
}

func main() {
	fmt.Println("context")

	// no1(ctx)
	// no2()
	// no3()
}
