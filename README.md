# status-server
Basic server to check if an OpenBazaar node can accept incoming connections

### Usage

```curl
curl --request POST \
   --url http://localhost:80/ \
   --header 'Content-Type: application/json' \
   --data '[
 	    "/ip4/104.236.157.44/tcp/4001",
 	    "/ip4/139.59.174.197/tcp/4001/"
 	 ]'
```
