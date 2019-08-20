# Split string
This activity let you query a MySQL DB


## Installation

```bash
flogo install github.com/ultrafab/mysqlquery_flogo
```
Link for flogo web:
```
https://github.com/ultrafab/mysqlquery_flogo
```

## Schema
Inputs and Outputs:

```json
{
 "input":[
     {
       "name": "host",
       "type": "string"
     },
     {
       "name": "port",
       "type": "string"
     },
     {
       "name": "dbNo",
       "type": "integer"
     },
     {
       "name": "variable",
       "type": "string"
     },
     {
       "name": "index",
       "type": "string"
     }
  ],
  "output": [
    {
      "name": "value",
      "type": "string"
    },
    {
      "name": "index",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input   | Description    |
|:----------|:---------------|
| host    | the Redis address |
| port    | the Redis address |
| dbNo    | the Redis database number |
| variable    | the key of the variable to read |
| index    | the _index of the variable |

## Ouputs
| Output   | Description    |
|:----------|:---------------|
| value    | the value of the variable in Redis |
| index    | the _index of the variable |
