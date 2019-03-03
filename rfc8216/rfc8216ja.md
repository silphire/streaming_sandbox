# HTTP Live Streaming

# 概要

本文書は、マルチメディアデータのストリームを転送するためのプロトコルについて記述している。
本文書は、ファイルのデータフォーマットと、サーバ(送信者)が取るべき挙動、ならびにクライアント(受信者)が取るべき挙動について規定している。
本文書は、本プロトコルのバージョン7について述べている。

# Status of This Memo

(省略)

# Copyright Notice

Copyright (c) 2017 IETF Trust and the persons identified as the document authors.  All rights reserved.

This document is subject to BCP 78 and the IETF Trust's Legal Provisions Relating to IETF Documents (http://trustee.ietf.org/license-info) in effect on the date of publication of this document.  Please review these documents carefully, as they describe your rights and restrictions with respect to this document.

This document may not be modified, and derivative works of it may not be created, except to format it for publication as an RFC or to translate it into languages other than English.

# 1. HTTP Live Streaming とは

HTTP Live Streamingは、信頼性があり、かつ低コストに連続した長時間の動画をインターネットを通じて配信する手段である。HLSは受信者に、その時点で可能な最も良い品質で途切れのない再生を維持するために、現在のネットワークの状態に応じて適切なビットレートのメディアを利用することを可能にする。
It supports interstitial content boundaries.
また、柔軟なメディア暗号化のフレームワークを提供する。
また、音声翻訳のような、同一のコンテンツの複数種類の演出を効率的に提供できる。
また、数多くの視聴者への配信を行うための巨大規模のHTTPキャッシュインフラとの互換性を提供している。

2009年にこのInternet-Draftが初めて投稿されてから、HTTPライブストリーミングは幅広いコンテンツ制作者、ツール業者、配信者、デバイス製造メーカーによって実装され運用されてきた。初投稿以降の8年間のうちに、メディアストリーミングの実装者による広範囲に渡るレビューや議論を通じて、本プロトコルは改良を受けてきた。

本文書の目的は、メディア転送プロトコルを記述することによってHTTPライブストリーミングの実装者間で相互運用性を促進することにある。本プロトコルを使うことによって、クライアントはサーバーからメディアの連続したストリームを受信することができる。

本文書は本プロトコルのバージョン7について記述している。

# 2. 概観

マルチメディア表現は Uniform Resource Identifier (URI) [RFC3986](http://tools.ietf.org/html/rfc3986) によってプレイリストに定義される。

プレイリストはメディアプレイリストかマスタープレイリストのいずれかである。どちらともUTF-8テキストファイルであり、URIと詳細タグを含んでいる。

メディアプレイリストは、メディアセグメントのどれが、いつ続けて再生されるかを記したリストが記載されており、マルチメディアの表現として再生される。

以下はメディアプレイリストの一例である。

```
#EXTM3U
#EXT-X-TARGETDURATION:10

#EXTINF:9.009,
http://media.example.com/first.ts
#EXTINF:9.009,
http://media.example.com/second.ts
#EXTINF:3.003,
http://media.example.com/third.ts
```

最初の行は #EXTM3U というフォーマット識別タグである。 #EXT-X-TARGETDURATION が含まれている行は全てのメディアセグメントは10秒かそれ以下の長さであることを意味している。そして、3つのメディアセグメントが宣言されている。最初とその次は9.009秒、3つめは3.003秒の長さである。

このプレイリストを再生するには、クライアントは最初にプレイリストをダウンロードしてから、その中で制限されているメディアセグメントをダウンロードして再生する。クライアントは追加されているセグメントが無いかどうかを調べるために、本文書に記述されている通りプレイリストをリロードする。データは HTTP で転送される *べき* (SHOULD) だが、一般的には、そのURIには要求に応じて信頼性のある転送を行える他のプロトコルを指定することもできる。

マスタープレイリストにはより複雑な表現を記述することができる。マスタープレイリストは同一のコンテンツの別バージョンを表す異種ストリーム (Variant Stream) の集合を提供することができる。

異種ストリームはメディアプレイリストを含む。このメディアプレイリストは特定のビットレート、特定のフォーマット、特定の解像度の映像を含むメディアを指定する。

異種ストリームはまた、「演奏」の集合を指定できる。演奏は、別バージョンのコンテンツである。例えば、他の言語で制作された音声や、異なるアングルから撮影された映像がこれに含まれる。

クライアントはネットワークの状態に適応するために、異なる異種ストリームをスイッチすべきである。クライアントはユーザー設定に基づいて演奏を選択すべきである。

本文書中に現れる "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", "OPTIONAL" キーワードは、全て大文字で書き表されている時に限り、BCP 14に記載されている通りに解釈される。

# 3. メディアセグメント

メディアプレイリストは全ての映像を成すメディアセグメントの集合を含む。メディアセグメントはURIと、任意のバイト幅にて指定される。

各々のメディアセグメントの長さはメディアプレイリストのEXTINFタグで示される。

メディアプレイリスト中の各々のセグメントは、ユニークな整数であるメディアシーケンス番号を持つ。メディアプレイリスト中の最初のセグメントのメディアシーケンス番号は０かあるいはプレイリスト(4.3.3.2章)で宣言される。その他のセグメントのメディアシーケンス番号は直前のセグメントのメディアシーケンス番号にプラス1された物である。

各々のメディアセグメントは、前回のメディアシーケンス番号がついたセグメントの終端からの続きの符号化されたビットストリームを持たなければいけないし (MUST)、タイムスタンプや連続カウンタの値は途切れなく連続していなければならない (MUST)。唯一の例外はメディアプレイリストに出てくる最初のメディアセグメントと、明示的に不連続であることを意図したメディアセグメント (4.3.2.3章) である。無表記のメディア中断はプレイバックエラーを引き起こす。

## 3.1. サポートするメディアセグメントのフォーマット

### 3.2. MPEG-2 Transport Stream

MPEG-2 Transport Streamsは ISO 13818 にて規定される。

### Fragmented MPEG-4

