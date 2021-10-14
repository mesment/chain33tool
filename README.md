# chain33tool
a simple chain33 utility for testing transaction and balance modify




Installation and build
----------------------

```
go get github.com/mesment/chain33tool
go install
```

Requirements
------------
 * `go1.16` or newer.

Usage
-----

Commands
--------

### tx  

#### subCommands
1. decode

    decode hex encoded transaction and print transaction detail
    > chain33tool tx decode -d  hexEncodedTransaction
    * -d (--data) hex encoded transaction
    * -h (--help) print usage info

2. alter 

    alter transaction info, change transaction amount or to address
    > chain33tool tx alter -d hexEncodedTransaction -v 100 -t 12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv

    * -d (--data) hex encoded transaction
    * -v (--amount) new transaction amount
    * -t (--to) new transaction receiver address
    * -h (--help) print usage info



### balance  

#### subCommands
1. get

    read account balance in chain database
    > chain33tool balance get -a  12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv -f bty.toml
    * -a (--addr) account address
    * -f (--config) chain config file
    * -h (--help) print usage info

2. set 

    modify account balance in chain database
    > chain33tool balance set -a  12qyocayNF7Lv6C9qW4avxs2E7U41fKSfv -v 100.00 -f bty.toml
    * -a (--addr) account address
    * -v (--amount) new balance to set
    * -f (--config) chain config file
    * -h (--help) print usage info

