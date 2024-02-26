# Gilded Rose in Go

[![Go](https://github.com/jivagoalves/gilded-rose-go/actions/workflows/go.yml/badge.svg)](https://github.com/jivagoalves/gilded-rose-go/actions/workflows/go.yml)

This is an attempt of implementing the [Gilded Rose](https://kata-log.rocks/gilded-rose-kata) kata in [Go](https://go.dev) using principles from [Type-Driven Design](https://fsharpforfunandprofit.com/series/designing-with-types).

## The Gilded Rose kata

Gilded Rose is a small inn whose stock control system needs to be developed. The items are constantly degrading in quality as they approach their sell by date. In summary, the behaviour of the system is:

* All items have a SellIn value which denotes the number of days we have to sell the item
* All items have a Quality value which denotes how valuable the item is
* At the end of each day our system lowers both values for every item

There are also some exceptional behaviour and invariants such as:

* Once the sell by date has passed, Quality degrades twice as fast
* The Quality of an item is never negative
* "Aged Brie" actually increases in Quality the older it gets
* The Quality of an item is never more than 50
* "Sulfuras", being a legendary item, never has to be sold or decreases in Quality
* "Backstage passes", like aged brie, increases in Quality as its SellIn value approaches;
    * Quality increases by 2 when there are 10 days or less and by 3 when there are 5 days or less but
    * Quality drops to 0 after the concert

Just for clarification, an item can never have its Quality increase above 50, however "Sulfuras" is a legendary item and as such its Quality is 80 and it never alters.

## Getting Started

Please ensure you have Go installed.

Build the project by running:
```shell
go build -v ./...
```

Run the web app with:
```shell
go test -v ./...
```

## Stock API (WIP)

You can now:

* Add a new item to the stock:
```shell
curl --location 'localhost:8080/api/items' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Orange",
    "registeredOn": "2023-02-27",
    "sellBy": "2023-03-15",
    "quality": 35
}'
```

* List all items in the stock:
```shell
curl --location 'localhost:8080/api/items'
```

* List all items in the stock as of date:
```shell
curl --location 'localhost:8080/api/items/as-of/2023-03-01'
```

* Remove items from the stock:
```shell
curl --location --request DELETE 'localhost:8080/api/items/1'
```

Please bear in mind this in under development and not all features are accessible via the web API. Stay tuned as we add them! :)

## Architecture

TODO

## Disclaimer

Please bear in mind this is by no means idiomatic Go as it is a learning experiment. Feedback is welcome!

Hope you enjoy it!
