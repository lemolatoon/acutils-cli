# acutils-cli

## 概要

Atcoder といったコンテストに参加するときに簡単に

1. init: コンテスト用のディレクトリを作成し、vscode 用の設定を設置
2. new: 問題用のディレクトリを作成し、テンプレートファイルからコピー
3. run: 対象の問題のソースコードをコンパイルし実行する

ことを可能にする。

## 使用例

コンテストのディレクトリを作成

```
$ acutils-cli init abc348
Using config file: /home/lemolatoon/.acutils-cli.toml
Failed to read settings.json file: /home/lemolatoon/.vscode/settings.json
$ cd abc348
```

問題のディレクトリを作成

```
$ acutils-cli new a
Using config file: /home/lemolatoon/.acutils-cli.toml
Failed to read template file: /home/lemolatoon/template.cpp
```

コーディング

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

コンパイル&実行

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

## 注意

これは完全に個人用です。

```

```

```

```
