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

* __main.go__: Main class, setup an startup of the reverse proxy.
* _config_
  * __config.go__: Helper class to load config from file.
* _jquery_
  * __factory.go__: Created a jsonquery object using specified path.
  * __jsonquery.go__: Jquery object used when querying subnodes for a specified path and json-object.
  * __jsonquery_test.go__: Testclass for jsonquery.go.
* _reverseproxy_
  * __director.go__: Redirecting host -> target (also combining query strings).
  * __handler.go__: Handling request, i.e. inspecting request, checking bandwidth limits.
  * __proxy.go__: Implementation of go reverse proxy.
  * __requestinspector.go__: Allows or denies a request based on content.
  * __responsemodifier.go__: Modifies a response using a supplied jquery object.
  * __transport.go__: Http transport object.
  * __requestinspector_test.go__: Testclass for requestinspector.go.
  * __responsemodifier_test.go__: Testclass for responsemodifier.go.
* _tokenbucket_
  * __clock.go__: Implementation of a ms-clock.
  * __metric.go__: Measurer of time-diff.
  * __tokenbucket.go__: Implementation of the token bucket algorithm to handle bandwith and burst limitations. 
  * __tokenbucket_test.go__: Testclass for tokenbucket.go.
