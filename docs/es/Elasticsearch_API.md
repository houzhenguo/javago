
## 低级版本

pom.xml
```xml
<dependencies>
        <dependency>
            <groupId>org.elasticsearch.client</groupId>
            <artifactId>elasticsearch-rest-client</artifactId>
            <version>6.5.4</version>
        </dependency>
        <dependency>
            <groupId>junit</groupId>
            <artifactId>junit</artifactId>
            <version>4.12</version>
        </dependency>
        <dependency>
            <groupId>com.fasterxml.jackson.core</groupId>
            <artifactId>jackson-databind</artifactId>
            <version>2.9.4</version>
        </dependency>
    </dependencies>

    <build>
    <plugins>
    <!-- java编译插件 不导入可能有 5版本的异常-->
    <plugin>
    <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-compiler-plugin</artifactId>
        <version>3.2</version>
        <configuration>
            <source>1.8</source>
            <target>1.8</target>
            <encoding>UTF-8</encoding>
        </configuration>
    </plugin>
    </plugins>
    </build>
```

```java
package com.houzhenguo.test;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.apache.http.HttpHost;
import org.apache.http.util.EntityUtils;
import org.elasticsearch.client.*;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

public class TestESRest {
    private static final ObjectMapper MAPPER = new ObjectMapper();
    private RestClient restClient;
    @Before
    public void init() {
        RestClientBuilder restClientBuilder = RestClient.builder(
                new HttpHost("ip", 9200, "http"),
                new HttpHost("ip", 9201, "http"),
                new HttpHost("ip", 9202, "http")
        );
        restClientBuilder.setFailureListener(new RestClient.FailureListener() {
            @Override
            public void onFailure(Node node) {
                System.out.println("出错了 ->" +node);
            }
        });
        this.restClient = restClientBuilder.build();
    }
    @After
    public void after() throws IOException {
        restClient.close();
    }
    // 查询集群的状态
    @Test
    public void testGetInfo() throws IOException {
        Request request = new Request("GET", "/_cluster/state");
        request.addParameter("pretty", "true");
        Response response = this.restClient.performRequest(request);
        System.out.println(response.getStatusLine());
        System.out.println(EntityUtils.toString(response.getEntity()));
    }
    // 新增数据
    @Test
    public void testCreateData() throws IOException {
        Request request = new Request("POST", "/haoke/house");
        Map<String, Object> data = new HashMap<>();
        data.put("id","2001");
        data.put("title","高高");
        data.put("price","3000");
        request.setJsonEntity(MAPPER.writeValueAsString(data));
        Response response = this.restClient.performRequest(request);
        System.out.println(response.getStatusLine());
        System.out.println(EntityUtils.toString(response.getEntity()));
    }
    // 根据id 查询数据
    @Test
    public void testQueryData() throws IOException{
        Request request = new Request("GET","/haoke/house/ZPg2mXABwUh2zxK_N_g5");
        Response response = this.restClient.performRequest(request);
        System.out.println(response.getStatusLine());
        System.out.println(EntityUtils.toString(response.getEntity()));
    }
    // 搜索数据
    @Test
    public void testSearchData() throws IOException {
        Request request = new Request("POST", "/haoke/house/_search");
        String searchJson = "{\"query\": {\"match\": {\"title\": \"拎包入住\"}}}";
        request.setJsonEntity(searchJson);
        request.addParameter("pretty","true");
        Response response = this.restClient.performRequest(request);
        System.out.println(response.getStatusLine());
        System.out.println(EntityUtils.toString(response.getEntity()));
    }
}

```
---

## REST 高级客户端

pom.xml

```xml
        <dependency>
            <groupId>org.elasticsearch.client</groupId>
            <artifactId>elasticsearch-rest-high-level-client</artifactId>
            <version>6.5.4</version>
        </dependency>
```

```java
package com.houzhenguo.test;

import org.apache.http.HttpHost;
import org.elasticsearch.action.ActionListener;
import org.elasticsearch.action.delete.DeleteRequest;
import org.elasticsearch.action.delete.DeleteResponse;
import org.elasticsearch.action.get.GetRequest;
import org.elasticsearch.action.get.GetResponse;
import org.elasticsearch.action.index.IndexRequest;
import org.elasticsearch.action.index.IndexResponse;
import org.elasticsearch.action.search.SearchRequest;
import org.elasticsearch.action.search.SearchResponse;
import org.elasticsearch.action.update.UpdateRequest;
import org.elasticsearch.action.update.UpdateResponse;
import org.elasticsearch.client.*;
import org.elasticsearch.common.Strings;
import org.elasticsearch.common.unit.TimeValue;
import org.elasticsearch.index.query.QueryBuilders;
import org.elasticsearch.search.SearchHit;
import org.elasticsearch.search.SearchHits;
import org.elasticsearch.search.builder.SearchSourceBuilder;
import org.elasticsearch.search.fetch.subphase.FetchSourceContext;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeUnit;

public class TestHighLevel {
    private RestHighLevelClient client;

    @Before
    public void init() {
        RestClientBuilder restClientBuilder = RestClient.builder(
                new HttpHost("ip", 9200, "http"),
                new HttpHost("ip", 9201, "http"),
                new HttpHost("ip", 9202, "http")
        );
        this.client = new RestHighLevelClient(restClientBuilder);
    }
    @After
    public void after() throws IOException {
        this.client.close();
    }
    // 新增文档 同步操作
    @Test
    public void testCreate() throws Exception {
        Map<String, Object> data = new HashMap<>();
        data.put("id", "2002");
        data.put("title", "南京西路 拎包入住 一室一厅");
        data.put("price", "4500");
        IndexRequest indexRequest = new IndexRequest("haoke", "house")
                .source(data);
        IndexResponse indexResponse = this.client.index(indexRequest,
                RequestOptions.DEFAULT);
        System.out.println("id->" + indexResponse.getId());
        System.out.println("index->" + indexResponse.getIndex());
        System.out.println("type->" + indexResponse.getType());
        System.out.println("version->" + indexResponse.getVersion());
        System.out.println("result->" + indexResponse.getResult());
        System.out.println("shardInfo->" + indexResponse.getShardInfo());
    }
    @Test
    public void testCreateAsync() throws Exception {
        Map<String, Object> data = new HashMap<>();
        data.put("id", "2003");
        data.put("title", "南京东路 最新房源 二室一厅");
        data.put("price", "5500");
        IndexRequest indexRequest = new IndexRequest("haoke", "house")
                .source(data);
        this.client.indexAsync(indexRequest, RequestOptions.DEFAULT, new
                ActionListener<IndexResponse>() {
                    @Override
                    public void onResponse(IndexResponse indexResponse) {
                        System.out.println("id->" + indexResponse.getId());
                        System.out.println("index->" + indexResponse.getIndex());
                        System.out.println("type->" + indexResponse.getType());
                        System.out.println("version->" + indexResponse.getVersion());
                        System.out.println("result->" + indexResponse.getResult());
                        System.out.println("shardInfo->" + indexResponse.getShardInfo());
                    } @
                            Override
                    public void onFailure(Exception e) {
                        System.out.println(e);
                    }
                });
        System.out.println("ok");
        Thread.sleep(20000);
    }
    @Test
    public void testQuery() throws Exception {
        GetRequest getRequest = new GetRequest("haoke", "house",
                "ZfhDmXABwUh2zxK_JvjK");
    // 指定返回的字段
        String[] includes = new String[]{"title", "id"};
        String[] excludes = Strings.EMPTY_ARRAY;
        FetchSourceContext fetchSourceContext =
                new FetchSourceContext(true, includes, excludes);
        getRequest.fetchSourceContext(fetchSourceContext);
        GetResponse response = this.client.get(getRequest, RequestOptions.DEFAULT);
        System.out.println("数据 -> " + response.getSource());
    }
    /**
     * 判断是否存在
     * *
     @throws Exception
     */
    @Test
    public void testExists() throws Exception {
        GetRequest getRequest = new GetRequest("haoke", "house",
                "ZfhDmXABwUh2zxK_JvjK");
        // 不返回的字段
        getRequest.fetchSourceContext(new FetchSourceContext(false));
        boolean exists = this.client.exists(getRequest, RequestOptions.DEFAULT);
        System.out.println("exists -> " + exists);
    }
    // 删除数据
    @Test
    public void testDelete() throws Exception {
        DeleteRequest deleteRequest = new DeleteRequest("haoke", "house",
                "ZfhDmXABwUh2zxK_JvjK");
        DeleteResponse response = this.client.delete(deleteRequest,
                RequestOptions.DEFAULT);
        System.out.println(response.status());// OK or NOT_FOUND
    }
    // 更新数据
    @Test
    public void testUpdate() throws Exception {
        UpdateRequest updateRequest = new UpdateRequest("haoke", "house",
                "ZvhEmXABwUh2zxK_Bvjn");
        Map<String, Object> data = new HashMap<>();
        data.put("title", "北京西路");
        data.put("price", "5000");
        updateRequest.doc(data);
        UpdateResponse response = this.client.update(updateRequest,
                RequestOptions.DEFAULT);
        System.out.println("version -> " + response.getVersion());
    }
    /**
     * 测试搜索
     * *
     @throws Exception
     */
    @Test
    public void testSearch() throws Exception {
        SearchRequest searchRequest = new SearchRequest("haoke");
        searchRequest.types("house");
        SearchSourceBuilder sourceBuilder = new SearchSourceBuilder();
        sourceBuilder.query(QueryBuilders.matchQuery("title", "西路"));
        sourceBuilder.from(0);
        sourceBuilder.size(5);
        sourceBuilder.timeout(new TimeValue(60, TimeUnit.SECONDS));
        searchRequest.source(sourceBuilder);
        SearchResponse search = this.client.search(searchRequest,
                RequestOptions.DEFAULT);
        System.out.println("搜索到 " + search.getHits().totalHits + " 条数据.");
        SearchHits hits = search.getHits();
        for (SearchHit hit : hits) {
            System.out.println(hit.getSourceAsString());
        }
    }
}


```