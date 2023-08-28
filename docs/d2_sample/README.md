# tips
- マークダウンは、ファイル名はREADME.mdじゃないと自動で表示されない
- .svgは自動で表示されない  
    -> README.mdの中に画像として貼り付けることで、自動で表示されるようになる。
- ファイル名がtest.d2の場合、markdown内に以下のように書くことで、Markdown上で.svgを閲覧できる
```
<img src="test.svg" />
```
- d2ファイルは、Github Actions上で.svgに自動でビルド、コミットされる。


<img src="test.svg" />
