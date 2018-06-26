# abiosapiwrapper

## Config

Configuration of the reverse proxy is specified in the conf.json file according to the example below.

```json
{
  "ip": "localhost",
  "port": "5005",
  "timePerRequest": 2000,
  "burstRequests": 5,
  "allowedQueryParameters": ["access_token", "page"],
  "routes": {
    "/series/live": {"host": "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false", "jquery": "*"},
    "/players/live": {"host": "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false", "jquery": "rosters/players/*"},
    "/teams/live": {"host": "https://api.abiosgaming.com/v2/series?starts_before=now&is_over=false", "jquery": "roster/teams/*"},
  }
}
```

__ip__: server ip  
__port__: server port  
__timePerRequest__: average time (ms) between two consecutive requests  
__burstRequests__: maximum number of burst requests in a row  
__allowedQueryParameters__: array of allowed parameters in query string  
__routes__: reverse proxy routes, specified as...  
* "end point 1": {"host": "host address 1", "jquery": "jquery string"},  
* "end point 2": {"host": "host address 2", "jquery": "jquery string"},  
* ...  
* "end point n": {"host": "host address n", "jquery": "jquery string"},  
  
jquery is specified according to node_1/node_2/.../node_n/\*, where \* gets all nodes as an array.  

## Classes
