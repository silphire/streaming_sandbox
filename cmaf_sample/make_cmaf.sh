#!/bin/bash

#
# HLS と MPEG-DASH でメディアを共通化してみる
# https://qiita.com/tomopyonsama/items/d520af4a6198da5af928
#

BENTO4_ROOT=
FFMPEG_ROOT=

MEDIA_FILE=$1

# create fragmented MP4
$BENTO4_ROOT/mp4fragment $MEDIA_FILE fragmented_$MEDIA_FILE

# create manifest playlist
$BENTO4_ROOT/mp4dash --force --hls --no-split --profile=on-demand fragmented_$MEDIA_FILE

exit

