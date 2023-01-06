package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

//type File struct {
//	*file
//}

//type file struct {
//	pfd         poll.FD // 非負整数fd(ファイルディスクリプタ) プロセス中でopenしたファイルに対して順番に割り当てられる番号
//	name        string
//	dirinfo     *dirInfo // nil unless directory being read
//	nonblock    bool     // whether we set nonblocking mode
//	stdoutOrErr bool     // whether this is stdout or stderr
//	appendMode  bool     // whether file is opened for appending
//}

// openしていない全てのファイルに対しても整数の識別子をつけて管理しており、これをinode番号
//type FD struct {
//	// System file descriptor. Immutable until Close.
//	Sysfd int
//
//	// Whether this is a streaming descriptor, as opposed to a
//	// packet-based descriptor like a UDP socket. Immutable.
//	IsStream bool
//
//	// Whether a zero byte read indicates EOF. This is false for a
//	// message based socket connection.
//	ZeroReadIsEOF bool
//
//	// contains filtered or unexported fields
//}

//func Open(name string) (*File, error) {
//	return OpenFile(name, O_RDONLY, 0)
//}

// 「サーバーからデータを受け取る」「クライアントからデータを送る」というのは、言い換えると「コネクションからデータを読み取る・書き込む」
// ファイルディスクプリタはファイルと一対一に対応していて、ファイルからデータを入力する場合は、ファイルディスクプリタを経由してデータが渡されます。逆に、ファイルへデータを出力するときも、ファイルディスクプリタを経由して行われます。
// ファイルディスクプリタはパッケージ os に定義されている構造体 File に格納されてユーザーに渡されます。ただし、Go 言語の File には低レベルな処理 (メソッド) しか用意されていません。Go 言語の場合、パッケージ bufio に高レベルな処理を行うための構造体 Reader や Writer が用意されているので、File からそれらの型を生成して入出力処理を行います。

// "everything-is-a-file philosophy" というものがあります。これは、キーボードからの入力も、プリンターへの出力も、ハードディスクやネットワークからのI/Oもありとあらゆるものを全て「OSのファイルシステム上にあるファイルへのI/Oとして捉える」という思想です。
// ネットワークからのデータ読み取り・書き込みも、OS内部的には通常のファイルI/Oと変わらないのです。そのため、ネットワークコネクションに対しても、通常ファイルと同様にfdが与えられるのです。

// システムコール = OSへのコール

// basic
func no1() {
	f1, err := os.Open("algo/inout/open.txt") //名前付きのファイルをreadonlyで開く(O_RDONLY)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("あったよ")
	}
	defer func() {
		err := f1.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	data := make([]byte, 1024)
	count, err := f1.Read(data)
	fmt.Printf("read %d bytes:\n", count)
	fmt.Println(string(data[:count]))

	// 引数として指定したファイルが既に存在している場合、中身を空にして開く
	f2, err := os.Create("algo/inout/create.txt") //umask 0666で作る(O_RDWR)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("作ったよ")
	}
	str := "これを書き込むんだー1"
	data = []byte(str)
	// Writeメソッドを使う予定のファイルオブジェクトは、書き込み権限がついたos.Create()から作ったものでなくてはなりません。
	// os.Open()で開いたファイルは読み込み専用なので、これにWriteメソッドを使うと、以下のようなエラーが出ます。
	count, err = f2.Write(data) //うわがくよ
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("書いたよ")
	}
	fmt.Printf("write %d bytes\n", count)
}

// nw i/o
func no2() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
	}

	send := "送信するもの"
	data := []byte(send)
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println(err)
	}

	//
	//// ここから読み取り
	//
	//data := make([]byte, 1024)
	//count, _ := conn.Read(data)
	//fmt.Println(string(data[:count]))
}

// バイトスライスpを用意して、そこに読み込んだ内容をいれる
// Read(p []byte) (n int, err error)
// バイトスライスpの中身を書き込む
// Write(p []byte) (n int, err error)
//
//	type Reader interface {
//	   Read(p []byte) (n int, err error)
//	}
//
//	type Writer interface {
//		Write(p []byte) (n int, err error)
//	}
//
// interface
func no3(r io.Reader) {
	data := make([]byte, 300)
	len, _ := r.Read(data)
	str := string(data[:len])

	result := strings.ReplaceAll(str, "Hello", "Guten Tag")
	fmt.Println(result)
}

// bufio bufioパッケージはbuffered I/Oをやるためのも
// func NewWriter(w io.Writer) *Writer
// func NewReader(rd io.Reader) *Reader
// len(p)が内部バッファの空きより大きい場合(=pの中身を一旦全部bufに書き込むだけの余裕がない場合)
// bufが先頭から空いているなら、pの中身を直接メモリに書き込む(=bufを使わない)
// bufの空きが先頭からじゃないなら、
// bufに入るだけデータを埋める
// bufの中身をメモリに書き込む[2]
// pの中でbufに書き込み済みのところを切る
// 実際にデータをメモリに書き込むのは、内部バッファbufの中身がいっぱいになったときのみ
// fun1byteごと書き込んでいる場合、bufioの有無で110倍ものパフォーマンス差が生まれる
// 1byteごと読み込んでいる処理の場合、bufio使用なし/ありでそれぞれ10157577ns/29841nsと、約340倍ものパフォーマンスの差が出る
// bufio.Scanner トークンごとに読む
// - 単語ごと(=スペース区切り)に読み取りたい
// - 行ごと(=改行文字区切り)に読み取りたい
func no4() {
	// bf := bufio.NewReader(f)
	f, _ := os.Open("algo/inout/open.txt")
	defer f.Close()

	// スキャナを用意(トークンはデフォルトの行のまま)
	sc := bufio.NewScanner(f)

	// EOFにあたるまでスキャンを繰り返す
	for sc.Scan() {
		line := sc.Text() // スキャンした内容を文字列で取得
		fmt.Println(line)
	}
}

// fmt
// Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
// Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
// os.NewFile関数: 第二引数にとった名前のファイルを、第一引数にとったfd番号でos.File型にする関数
// syscall.Stdin: syscallパッケージ内でvar Stdin = 0と定義された変数
// syscall.Stdout: syscallパッケージ内でvar Stdout = 1と定義された変数
// 標準入力: ファイル/dev/stdinをfd0番で開いたもの
// 標準出力: ファイル/dev/stdoutをfd1番で開いたもの
// - 標準入力・出力に割り当てるfd番号を0と1にするのは一種の慣例です。 また、標準エラー出力は慣例的にfd2番になります。
func no5(a ...interface{}) {
	//fmt.Fprintln(os.Stdout, a)
	fmt.Fscan(os.Stdin, a...)
}

//	type Buffer struct {
//		buf      []byte
//		// (略)
//	}
//
// bytes.Bufferを用意
// (bytes.Bufferは初期化の必要がありません)
func no6() {
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Println(b.String())
}

// stringsパッケージは、文字列を置換したり辞書順でどっちが先か比べたり だけではない
// bytes.Buffer型と同じく、文字列型をパッケージ独自型でカバーすることで、io.Readerに代入できるようにした型も定義されている
func no7() {
	var b bytes.Buffer
	b.Write([]byte("World"))

	plain := make([]byte, 10)
	b.Read(plain)

	fmt.Println("buffer: ", b.String())
	fmt.Println("output:", string(plain))
}

func no8() {
	str := "あ〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜〜"
	rd := strings.NewReader(str)

	row := make([]byte, 100)
	rd.Read(row)
	fmt.Println(string(row))
}

// 入力されたデータをそのまま画面へ出力するコマンド echo
func no9() {
	buff := make([]byte, 256)
	for {
		c, _ := os.Stdin.Read(buff)
		if c == 0 {
			break
		}
		os.Stdout.Write(buff[:c])
	}
}

// Writer はバッファが満杯になるまでデータを溜め込みます。
// 実際にファイルに書き込むのはバッファが満杯になってからです。
// 途中でバッファのデータをファイルに出力したい場合はメソッド Flush を使います。
// c が '\n' と等しい場合は Flush でバッファのデータを出力します。これで 1 行ごとにデータを出力することができます。
// 最後に、Flush でバッファ内のデータを出力します。
func no10() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		c, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		w.WriteByte(c)
		if c == '\n' {
			w.Flush()
		}
	}
	w.Flush()
}

// 行単位の入出力
// ReadBytes と ReadString は引数 delim で指定した区切り文字までデータを読み込みます。
// ReadBytes は読み込んだデータをスライスに格納して返します。ReadString は文字列にして返します。どちらのメソッドも区切り文字は捨てずにスライス (または文字列) に格納する
func no11() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		w.WriteString(s)
		w.Flush()
	}
}

// bufio.Scanner
// func (s *Scanner) Scan() bool
// func (s *Scanner) Bytes() []byte
// func (s *Scanner) Text() string
func no12() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	for s.Scan() {
		w.WriteString(s.Text())
		w.WriteString("\n")
		w.Flush()
	}
}

// ファイルのアクセス方法
// 標準入出力を使わずにファイルにアクセスする場合、次の 3 つの操作が基本
// アクセスするファイルをオープンする
// 入出力関数（メソッド）を使ってファイルを読み書きする。
// ファイルをクローズする。
func no13() {
	input, err := os.Open("algo/inout/open.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	output, _ := os.Create("algo/inout/testout.txt")
	buff := make([]byte, 256)
	for {
		c, _ := input.Read(buff)
		if c == 0 {
			break
		}
		output.Write(buff[:c])
	}
	input.Close()
	output.Close()
}

// コマンドライン引数の取得
func no14() {
	for _, name := range os.Args[1:] {
		file, err := os.Open(name)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		s := bufio.NewScanner(file)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		file.Close()
	}
}

func main() {
	fmt.Println("inout")

	// no1()
	// no2()
	// no3()
	// no4()
	// no5()
	// no6()
	// no7()
	// no8()
	// no9()
	// no10()
	// no11()
	// no12()
	// no13()
	// no14()
}
