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

HTTP Live Streamingは、信頼性があり、かつ低コストに連続した長時間の動画をインターネットを通じて配信する手段である。