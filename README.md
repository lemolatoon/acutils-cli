# acutils-cli

## 概要

Atcoder といったコンテストに参加するときに簡単に

1. init: コンテスト用のディレクトリを作成し、vscode 用の設定を設置
2. new: 問題用のディレクトリを作成し、テンプレートファイルからコピー
3. run: 対象の問題のソースコードをコンパイルし実行する
4. clip: 対象の問題のソースコードをクリップボードにコピー

ことを可能にする。

## インストール/install

1. go環境がある場合
```
$ go install github.com/lemolatoon/acutils-cli@latest
```

2. go環境がない場合
```
curl --proto '=https' --tlsv1.2 -sSf https://lemolatoon.github.io/acutils-cli/installer.sh | sh
```

```
$ acutils-cli --version
acutils-cli version 0.2.0
$ acutils-cli --help
acutils is a Atcoder utilities CLI developed only for lemolatoon.

This application is a tool to generate template directory/files for contests 
to quickly start solving problems of the contest.

Usage:
  acutils-cli [command]

Available Commands:
  clip        Copy the source code to the clipboard.
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize contest directory
  new         create directory for the problem, and put the template source file in it.
  run         Compile and Run source code of specified problem-name

Flags:
      --config string   config file (default is $HOME/.acutils-cli/config.toml)
  -h, --help            help for acutils-cli
  -v, --version         version for acutils-cli

Use "acutils-cli [command] --help" for more information about a command.
```

お好みで、以下のaliasを追加するとよいです。
`$HOME/.bashrc`や、`$HOME/.config/fish/config.fish`など、お使いのshellに合わせて設定しましょう。
```
$ alias ac='acutils-cli'
```

## 使用例

### コンテストのディレクトリを作成

```
$ acutils-cli init abc348
Using config file: /home/lemolatoon/.acutils-cli.toml
Failed to read settings.json file: /home/lemolatoon/.vscode/settings.json
$ cd abc348
```

### 問題のディレクトリを作成

```
$ acutils-cli new a
Using config file: /home/lemolatoon/.acutils-cli.toml
Failed to read template file: /home/lemolatoon/template.cpp
```

テンプレートファイルを個別に指定したい場合は `--template` フラグを利用できます（デフォルトは `$HOME/.acutils-cli/template.cpp`。存在しない場合は内蔵テンプレートを使用します）。

```
$ acutils-cli new b --template ./my-template.cpp
```

### コーディング

実際には、vscode でやる。

```
$ nvim a/main.cpp
$ cat a/main.cpp
#include <bits/stdc++.h>
using namespace std;

int main() {
int64_t n;
cin >> n;
for (const auto i : views::iota(1, n + 1)) {
cout << (i % 3 == 0 ? 'x' : 'o');
}
cout << endl;
return 0;
}
```

### コンパイル&実行

ソースコードが変更されてない場合は、コンパイルせずに実行する。

```

$ acutils-cli run a
Using config file: /home/lemolatoon/.acutils-cli.toml
+g++-11 a/main.cpp -g -Wall -Wextra -fsanitize=undefined,address -std=c++20 -o a/a.out
+./a/a.out
7
ooxooxo
$ acutils-cli run a
Using config file: /home/lemolatoon/.acutils-cli.toml
+./a/a.out
10
ooxooxooxo

```

### 提出
クリップボードにコピーする
```
$ acutils-cli clip a
```

## 注意

これは完全に個人用です。
