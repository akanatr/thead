package main

func Example_help() {
	goMain([]string{"/some/path/of/thead", "-h"})
	// Output:
	// thead [OPTIONS]  [FILEs...]
	// OPTIONS
	// 	-l, --lines <LINES>             指定された行数を各ファイルごとに出力する．
	// 	-q, --quiet                     複数ファイルを出力する際の各ファイル名を表示しない．
	// 	-h, --help                      このメッセージを出力する.
	// ARGUMENTS
	// 	FILEs...                        カウント対象を指定する．
	// 	DIRs...                         指定したディレクトリ内のファイルを入力ファイルとする．
}

// func test_thead() {
// 	goMain([]string{"thead", ".testdata/App.tsx"})
// 	// Output:
// 	// import React, { useState } from 'react';
// 	// import {
// 	// 	StyleSheet,
// 	// 	Text,
// 	// 	View,
// 	// 	TextInput,
// 	// 	Button
// 	// } from 'react-native';

// 	// interface IToDo {

// }
