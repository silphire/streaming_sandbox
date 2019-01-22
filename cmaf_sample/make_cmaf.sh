#!/bin/bash

BENTO4_ROOT=
FFMPEG_ROOT=

MEDIA_FILE=$1

# create fragmented MP4
$BENTO4_ROOT/mp4fragment $MEDIA_FILE fragmented_$MEDIA_FILE

# create manifest playlist
$BENTO4_ROOT/mp4dash fragmented_$MEDIA_FILE

exit

