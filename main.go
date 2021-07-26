package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

type options struct {
	lines int
	quiet bool
	help  bool
	args  []string
}

func helpMessage(originalProgramName string) string {
	programName := filepath.Base(originalProgramName)
	return fmt.Sprintf(`%s [OPTIONS]  [FILEs...]
OPTIONS
	-l, --lines <LINES>             指定された行数を各ファイルごとに出力する．
	-q, --quiet                     複数ファイルを出力する際の各ファイル名を表示しない．
	-h, --help                      このメッセージを出力する.
ARGUMENTS
	FILEs...                        カウント対象を指定する．
	DIRs...                         指定したディレクトリ内のファイルを入力ファイルとする．
`, programName)
}

func parseArgs(args []string) (*options, error) {
	opts := &options{}
	flags := flag.NewFlagSet("thead", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args[0])) }
	flags.IntVarP(&opts.lines, "lines", "l", 0, "指定された行数を各ファイルごとに出力する．")
	flags.BoolVarP(&opts.quiet, "quiet", "q", false, "複数ファイルを出力する際の各ファイル名を表示しない．")
	flags.BoolVarP(&opts.help, "help", "h", false, "このメッセージを出力する.")

	if err := flags.Parse(args); err != nil {
		return nil, err
	}

	opts.args = flags.Args()[1:]

	return opts, nil
}

//標準出力と行数指定オプションの出力
func stdout(o *options, filename string) {

	file, err := os.Open(filename)
	//エラーならエラー出力して終了
	if err != nil {
		panic(err)
	}
	rd := bufio.NewReader(file)
	if o.lines > 0 { //-lオプションが指定されていればその行数分出力
		for i := 0; i < o.lines; i++ {
			s, err := rd.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Print(s)
		}
	} else { //行のオプション指定がなければ10行分出力
		for i := 0; i < 10; i++ {
			s, err := rd.ReadString('\n')
			if err == io.EOF {
				break
			}
			fmt.Print(s)
		}
	}
	file.Close()

}

//コマンドライン引数の複数ファイルにファイル名を一緒に出力するかどうか
func isQuietFile(filenames []string, o *options, filename string) {
	if len(filenames) > 1 && !o.quiet {
		fmt.Printf("==> " + filename + " <==\n")
		stdout(o, filename)
		fmt.Println()
	} else {
		stdout(o, filename)
	}
}

//ディレクトリ内の複数ファイルにファイル名を一緒に出力するかどうか
func isQuietDir(rootFiles []fs.FileInfo, o *options, fullPath string) {
	if len(rootFiles) > 1 && !o.quiet {
		fmt.Printf("==> " + fullPath + " <==\n")
		stdout(o, fullPath)
		fmt.Println()
	} else {
		stdout(o, fullPath)
	}
}

//コマンドライン引数にディレクトリがある場合の処理
func ifDir(o *options, dirname string) {
	rootFiles, _ := ioutil.ReadDir(dirname)
	for _, rootFile := range rootFiles {
		//ディレクトリ内のディレクトリにも再帰的に処理を行う
		if rootFile.IsDir() {
			ifDir(o, filepath.Join(dirname, rootFile.Name()))
			// continue
		}
		fullPath := filepath.Join(dirname, rootFile.Name())
		_, err := os.Open(fullPath)
		if err != nil {
			continue
		}
		//-qオプションが指定されていればファイル名は表示しない
		isQuietDir(rootFiles, o, fullPath)
	}
}

//コマンドライン引数のファイルをstdoutに渡す
func perform(o *options, filenames []string) {
	for _, filename := range filenames {
		fileInfo, _ := os.Stat(filename)
		//ディレクトリの場合の処理
		if fileInfo.IsDir() {
			ifDir(o, filename)
			break
		}
		//ファイルが複数の場合はファイル名表示
		//-qオプションが指定されていればファイル名は表示しない
		isQuietFile(filenames, o, filename)
	}
}

func goMain(args []string) int {
	//optsにオプション、errにエラー
	opts, err := parseArgs(args)
	//オプション解析にエラーがあればエラー出力
	if err != nil {
		fmt.Printf("%s", err.Error())
		fmt.Println()
		fmt.Println(helpMessage(args[0]))
	}
	//オプション解析でhelpが指定されたらhelpメッセージを出力
	if opts.help {
		fmt.Println(helpMessage(args[0]))
		return 1
	}
	//エラーがない場合，通常の処理を行う
	//opts:オプションの配列、opts.args:ファイルorディレクトリ
	perform(opts, opts.args)

	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
