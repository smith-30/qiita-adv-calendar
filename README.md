# qiita-adv-calendar

Display items with many like items in order in the [qiita advent calendar](https://qiita.com/advent-calendar/2017)

## requirements

- QIITA API's token  
Please get it from qiita's account page.

- URL Shortener API's token  
Please get it from console.developers.google.com.


## how to use

```bash
$ go get github.com/smith-30/qiita-adv-calendar
$ cd $GOPATH/src/github.com/smith-30/qiita-adv-calendar
$ mv .env.sample .env
# edit .env
$ vi .env
$ dep ensure -v
$ make build
# ex.
$ ./build/bin/qiita-adv -n go -c 4
```

### Todo

- [ ] implement grid proxy.

- [ ] more fast.

### License

qiita-adv-calendar is licensed under the the MIT license. Please see the LICENSE file for details.