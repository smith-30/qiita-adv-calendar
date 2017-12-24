# qiita-adv-calendar

Display items with many like items in order in the [qiita advent calendar](https://qiita.com/advent-calendar/2017)

## architecture

```
                     +-------------+
        +------------+ aggregateCh <--------------------------------------------------------------+
        |            +------+------+                                                              |
        |                   |                                    +-----------+                    |
        |                   |                             +->   +----------+ |  +--+              |
 +------v------+      +-----+------+   +------------+     |    +---------+ +-+     |       +------+------+
 | []*GridData +------+ Aggregater +---+ Dispatcher +-------> +--------+ +-+    +--------> | aggregateCh |
 +------+------+      +-----+------+   +-----+------+     |   | Fetcher+-+         |       +-------------+
        |                   |                |            +-> +---^----+        +--+                                      +-----------+
        |                   |             +--+----+               |                    +-------------+                    |           |
    +---+----+              |  +----------> queue |               +--------------------+  Req / Res  +--------------------> qiita api |
    |        |              |  |          +-------+                                    +-------------+                    |           |
    |  Sort  |       +------+--+----+                                                                                     +-----------+
    |        |       | gridUpdateCh <--------------------------------------------------------------------------+
    +---+----+       +--------------+                                                  +------+                |
        |                                              +----------+                +-> | Grid | +-----------+  |
        |                                    +------>  |          +----------------+   +------+             |  |
Finish  ^                                    |       +-+--------+ |                                         |  |
   +----+---+                                +---->  |          +---+  +------+   +-----------------------+ |  |
   |        |              Start             |     +-+--------+ | | +--> Grid +-+ +                       | |  |
   | Output |        +----------------+      |     |          | +-+    +------+ |                         | |  |
   |        |        | GridAggregater +------+---> | Calendar +-+       +-------+                        +v-v--+--------+
   +--------+        +----------------+            |          +----+               +-------------------> | gridUpdateCh |
                                                   +----------+    |  +------+     |                     +--------------+
                                                                   +> | Grid +-+ +-+
                                                                      +------+ |
                                                                       +-------+

```

## requirements

- QIITA API's token  
Please get it from qiita's account page.

- URL Shortener API's token  
Please get it from console.developers.google.com.
`It does not cause an error even if there is no token.`

## how to use

```bash
$ go get github.com/smith-30/qiita-adv-calendar
$ cd $GOPATH/src/github.com/smith-30/qiita-adv-calendar
$ mv .env.sample .env
# edit .env
$ vi .env
$ dep ensure -v
$ make build
$ ./build/bin/qiita-adv -n <calendar-name> -c <calendar-count>
# ex.
$ ./build/bin/qiita-adv -n go -c 4
$ ./build/bin/qiita-adv -n vue -c 4
```

### Todo

- [ ] implement grid proxy.

- [ ] more fast.

### License

qiita-adv-calendar is licensed under the the MIT license. Please see the LICENSE file for details.