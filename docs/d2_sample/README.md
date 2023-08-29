# tips
- stlaticaでは、main branchに向けたPR内で変更があったd2ファイルを元にしてsvgファイルの生成（ビルド、コミット）が行われる
- ソースファイル名がtest.d2の場合、同階層のディレクトリにtest.svgが生成される。
- ソースファイル名がtest.d2の場合、markdown内に以下のように書くことで、Markdown上で.svgを閲覧できる
```
<img src="test.svg" />
```

<img src="test.svg" />
