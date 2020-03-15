#!/bin/bash

set -xeu

pushd "$(pwd)"

mkdir -p tmp/dic
cd tmp/dic

# fetch mecab dictionary
# http://taku910.github.io/mecab/
curl -sSL -o mecab-ipadic.tar.gz "https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7MWVlSDBCSXZMTXM"
tar zxf mecab-ipadic.tar.gz

# fetch neologd
# https://github.com/neologd/mecab-ipadic-neologd
curl -sSL -o mecab-user-dict-seed.csv.xz https://github.com/neologd/mecab-ipadic-neologd/raw/5cadccf5c20f1cb0ab6bf967a58dd7338400dd38/seed/mecab-user-dict-seed.20200312.csv.xz
unxz mecab-user-dict-seed.csv.xz

# fetch kagome tools
# https://ikawaha.hateblo.jp/entry/2016/04/20/145940
git clone https://github.com/ikawaha/kagome
go run kagome/cmd/_dictool/main.go ipa -mecab mecab-ipadic-2.7.0-20070801/ -neologd mecab-user-dict-seed.csv

popd

mv tmp/dic/ipa.dic ./

rm -rf tmp/dic
