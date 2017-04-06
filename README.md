# lbd - Learning By Doing [![Build Status](https://travis-ci.com/BenchR267/lbd.svg?token=WNxuwTZtUpQ6tQUcVhqK&branch=master)](https://travis-ci.com/BenchR267/lbd) [![Go Report Card](https://goreportcard.com/badge/github.com/BenchR267/lbd)](https://goreportcard.com/report/github.com/BenchR267/lbd) [![Coverage Status](https://coveralls.io/repos/github/BenchR267/lbd/badge.svg?branch=master)](https://coveralls.io/github/BenchR267/lbd?branch=master)

In the year 2017 I started to challenge myself by creating an own programming language to get more familiar with the whole topic of compilers and interpreters. Aim is not to create a new programming language that should be hyped and that does anything better than existing languages. It is my goal to learn something by creating this project.

Another challenge for me is to write one blog post a week; so creating this project is also a very important there. If you want to keep track about my thoughts during the progress of coding, check out the tag 'compiler' at my blog: [LINK](https://blog.benchr.me/tags/compiler/).

Check out all the other stuff at this [LINK](https://blog.benchr.me/).

To get a feeling for the language, here is a first draft of the MVP:
```
add = (a int, b int) -> int {
	return a + b
}

mul = (x int, y int) -> int {
	return x * y
}

a = 5
b = 4
c = add(a, b)
d = mul(a, b)
```
