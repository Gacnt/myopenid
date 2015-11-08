# MyOpenID 

A [Yohcop-Openid](https://github.com/yohcop/openid-go) Nonce/Discovery storage replacement other than using the in memory storage that is default

# Installation

Download the package


`go get github.com/Gacnt/myopenid`


After package is downloaded, import it into your program by


`import "github.com/Gacnt/myopenid"`

# Usage
Then declare an `init` function and your 2 global variables in your package like so:

```
var discoveryCache *myopenid.MysqlDiscoveryCache
var nonceStore *myopenid.MysqlNonceStore

// Make sure you include the ?parseTime=true on the end of your connection.
func init() {
        discoveryCache, nonceStore = myopenid.DbConnection("[your connection string]/database?parseTime=true")
}
```

Simply just use your variables in the `openid.Verify` function as you would the normal ones that are provided. You will also need to setup 2 tables in your database you can either import the SQL from the provided file.

# Table

Required tables are dumped into the `myopenid.sql` file simply import it into your database or do it manually following the layout of the file provided.  
