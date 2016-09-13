#!/usr/bin/env bash

BUILD_DATE=`date +%Y-%m-%d`
VERSIONFILE="lib/version.go"
VERSION="0.0.1"

rm -f $VERSIONFILE
echo "package lib" > $VERSIONFILE
echo "const (" >> $VERSIONFILE
echo "  Version = \"$VERSION\"" >> $VERSIONFILE
echo "  BuildDate = \"$BUILD_DATE\"" >> $VERSIONFILE
echo ")" >> $VERSIONFILE

go install