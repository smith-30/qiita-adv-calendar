# qiita-adv-calendar

Display items with many like items in order in the [qiita advent calendar](https://qiita.com/advent-calendar/2017)

## requirements

- QIITA API's token  
Please get it from qiita's account page.

- URL Shortener API's token  
Please get it from console.developers.google.com.


## how to use

```bash
$ git clone git@github.com:smith-30/qiita-adv-calendar.git
$ mv .env.sample .env
# edit .env
$ dep ensure -v
$ go run main.go
```

### License

qiita-adv-calendar is licensed under the the MIT license. Please see the LICENSE file for details.