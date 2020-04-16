#!/bin/bash

set -xe

pushd "$(pwd)"

mkdir -p data/dic/tmp
cd data/dic/tmp

# fetch mecab dictionary
# http://taku910.github.io/mecab/
if [ ! -e data/dic/tmp/mecab-ipadic.tar.gz ]; then
  curl -sSL -o mecab-ipadic.tar.gz "https://drive.google.com/uc?export=download&id=0B4y35FiV1wh7MWVlSDBCSXZMTXM"
  tar zxf mecab-ipadic.tar.gz
fi

# fetch neologd
# https://github.com/neologd/mecab-ipadic-neologd
if [ ! -e data/dic/tmp/mecab-user-dict-seed.csv ]; then
  curl -sSL -o mecab-user-dict-seed.csv.xz https://github.com/neologd/mecab-ipadic-neologd/raw/5cadccf5c20f1cb0ab6bf967a58dd7338400dd38/seed/mecab-user-dict-seed.20200312.csv.xz
  unxz mecab-user-dict-seed.csv.xz
fi

# fetch kagome tools
# https://ikawaha.hateblo.jp/entry/2016/04/20/145940
if [ ! -e ./data/dic/tmp/ipa.dic ]; then
  git clone https://github.com/ikawaha/kagome
  go run kagome/cmd/_dictool/main.go ipa -mecab mecab-ipadic-2.7.0-20070801/ -neologd mecab-user-dict-seed.csv
fi

popd

mv data/dic/tmp/ipa.dic ./data/ipa.dic
rm -rf data/dic/
