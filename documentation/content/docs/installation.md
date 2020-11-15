---
title: "Installation"
weight: 5
---


# Installation
You can install spectools on your local machine. For working with specs only, this is enough. 
We recommend [furoBEC](/docs/tools/BEC/) if you have to generate more then "just" the specs. furoBEC is a
docker image which has nearly all dependencies already installed to generate all the additional things you may need to.

## Installation with brew
    brew tap theNorstroem/tap
    brew install spectools

## Installation with go
    GO111MODULE=on go get github.com/theNorstroem/spectools@v1.xx.xx

## Installation from sources
    git clone git@github.com:theNorstroem/spectools.git
    go install
