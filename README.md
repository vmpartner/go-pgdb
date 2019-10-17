# Postgres connect module
Connection module to Postgres database
   
Features: 
- Safe reuse connection on New()
- Automatic check and reconnect
- Comfortable configuration
- Stable fixed release versions

## Usage example

```golang

import "github.com/vmpartner/go-pgdb/v6"

config := LoadConfig()
pgdb.User = config.Key("user").String()
pgdb.Pass = config.Key("pass").String()
pgdb.Host = config.Key("host").String()
pgdb.Port = config.Key("port").String()
pgdb.Name = config.Key("name").String()
pgdb.Debug, _ = config.Key("debug").Bool()
pgdb.PingEachMinute, _ = config.Key("ping_each_minute").Int()
pgdb.MaxIdleConns, _ = config.Key("max_idle_conns").Int()
pgdb.MaxOpenConns, _ = config.Key("max_open_conns").Int()
var err error
DB, err = pgdb.New()
if err != nil {
    return err
}

```
