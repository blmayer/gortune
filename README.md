# gortune

![](https://travis-ci.org/blmayer/gortune.svg?branch=master)

Another clone of the classic BSD fortune, this version, written in Go,
contains all fortunes burried in the executable, so batteries are included.


## Installing

Run `go get github.com/blmayer/gortunes`


## Running

Issue `gortune` on the terminal and it will print a random fortune.

Eg.

```
> blmayer% gortune 
I didn't care as an ex-ballet dancer wrote and told me she had seen
the production and fallen in love with my legs. She said that in other
circumstances she could have lived happily with my legs but that she
only had a small flat in Holland Park.
    -- Tom Baker, in his autobiography
```


## Library

This package also deploys a library you can import in your code, add
this to your *import*:

`"github.com/blmayer/gortune/fortunes"`

At this moment the package provides the *Categories* variable, listing
all available categories of fortunes. And a map named *List*, where the
key is a category and it returns a list with all fortunes for that category,
hence it is named so.

## Comming soon

- Command line options


## Credits

All fortunes were downloaded from [www.bsdfortune.com](http://www.bsdfortune.com/).
If you know a fortune not included here please make a PR.
