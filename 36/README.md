
`docker pull daocloud.io/elasticsearch`

`docker network create esnet`

`docker run -d -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" --name 01-elasticsearch  -p 9200:9200 -p 9300:9300  --network esnet -e "discovery.type=single-node" 5acf0e8da90b`

`curl -XGET localhost:9200`
```json
{
  "name" : "VorDi0U",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "iNndv1V9SqSVDso7foubWQ",
  "version" : {
    "number" : "5.6.12",
    "build_hash" : "cfe3d9f",
    "build_date" : "2018-09-10T20:12:43.732Z",
    "build_snapshot" : false,
    "lucene_version" : "6.6.1"
  },
  "tagline" : "You Know, for Search"
}

```

查看健康状态
`curl -X GET 127.0.0.1:9200/_cat/health?v`

查询当前es集群中所有的indices
`curl -X GET 127.0.0.1:9200/_cat/indices?v`

创建索引
`curl -X PUT 127.0.0.1:9200/www`

删除索引
`curl -X DELETE 127.0.0.1:9200/www`

插入记录
```shell
curl -H "ContentType:application/json" -X POST 127.0.0.1:9200/user/person -d '
{
	"name": "dsb",
	"age": 9000,
	"married": true
}'
```
也可以使用PUT方法，但是需要传入id

```shell
curl -H "ContentType:application/json" -X PUT 127.0.0.1:9200/user/person/4 -d '
{
	"name": "sb",
	"age": 9,
	"married": false
}'
```

检索

Elasticsearch的检索语法比较特别，使用GET方法携带JSON格式的查询条件。

全检索：

`curl -X GET 127.0.0.1:9200/user/person/_search`

按条件检索：

```shell
curl -H "ContentType:application/json" -X PUT 127.0.0.1:9200/user/person/4 -d '
{
	"query":{
		"match": {"name": "sb"}
	}	
}'
```
ElasticSearch默认一次最多返回10条结果，可以像下面的示例通过size字段来设置返回结果的数目。

```shell
curl -H "ContentType:application/json" -X PUT 127.0.0.1:9200/user/person/4 -d '
{
	"query":{
		"match": {"name": "sb"},
		"size": 2
	}	
}'
```
