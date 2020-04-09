# Covid Cases API

The purpose of this project is to familiarize myself with some AWS Services like Lambda and API Gateway. I would not recommend using this API for production.

# Demo

Demo of the API can be found [here](https://4j222wktag.execute-api.ap-southeast-1.amazonaws.com/dev/covid).

# API Usage

## Get All Cases

**Endpoint**: [https://4j222wktag.execute-api.ap-southeast-1.amazonaws.com/dev/covid](https://4j222wktag.execute-api.ap-southeast-1.amazonaws.com/dev/covid)
```bash
$ curl https://4j222wktag.execute-api.ap-southeast-1.amazonaws.com/dev/covid
```

**Response:**
```json
[
   {
      "country":{
         "name":"US",
         "latitude":40,
         "longitude":-100
      },
      "confirmed":432438,
      "deaths":14808,
      "active":0,
      "recovered":24125,
      "lastupdate":1586417611
   },
   {
      "country":{
         "name":"Spain",
         "latitude":40.463667,
         "longitude":-3.74922
      },
      "confirmed":148220,
      "deaths":14792,
      "active":85407,
      "recovered":48021,
      "lastupdate":1586417592
   },
   {
      "country":{
         "name":"Italy",
         "latitude":41.8719,
         "longitude":12.5674
      },
      "confirmed":139422,
      "deaths":17669,
      "active":95262,
      "recovered":26491,
      "lastupdate":1586417592
   },
...
```


## Get Cases by Country

**Endpoint**: [https://4j222wktag.execute-api.ap-southeast-1.amazonaws.com/dev/covid?country=Italy](https://4j222wktag.execute-api.ap-southeast-1.amazonaws.com/dev/covid?country=Italy)
```bash
$ curl https://4j222wktag.execute-api.ap-southeast-1.amazonaws.com/dev/covid?country=Italy
```

**Response:**
```json
{
  "country":{
     "name":"Italy",
     "latitude":41.8719,
     "longitude":12.5674
  },
  "confirmed":139422,
  "deaths":17669,
  "active":95262,
  "recovered":26491,
  "lastupdate":1586417592
}
```
