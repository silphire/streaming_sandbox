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

# 1. HTTP Live Streamingとは

HTTP Live Streamingは、信頼性があり、かつ低コストに連続した長時間の動画をインターネットを通じて配信する手段である。HLSは受信者に、その時点で可能な最も良い品質で途切れのない再生を維持するために、現在のネットワークの状態に応じて適切なビットレートのメディアを利用することを可能にする。

本文書は本プロトコルのバージョン7について記述している。

# 2. 概観

マルチメディア表現はUniform Resource Identifier (URI) [RFC3986](http://tools.ietf.org/html/rfc3986) によってプレイリストに定義される。

プレイリストはメディアプレイリストかマスタープレイリストのいずれかである。どちらともUTF-8テキストファイルであり、URIと詳細タグを含んでいる。

メディアプレイリストは、メディアセグメントのどれが、いつ続けて再生されるかを記したリストが記載されており、マルチメディアの表現として再生される。
