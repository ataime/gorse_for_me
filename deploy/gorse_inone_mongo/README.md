# gorse 推荐系统搭建

```
docker-compose up 
```

测试
```
mysql 执行写入测试数据没有问题
clickhouse 写入数据item的 isHidden 值有问题，需要尝试指定值.

数据准备
curl -X POST "http://localhost:8088/api/feedback" \
     -H "Content-Type: application/json" \
     -d @feedback.json

获取推荐
curl -X GET "http://localhost:8088/api/recommend/user_1" \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer <YOUR_API_KEY>"

```

集成应用
```

func getRecommendations(userId string) ([]Recommendation, error) {
    url := fmt.Sprintf("http://localhost:8088/api/recommend/%s", userId)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Authorization", "Bearer <YOUR_API_KEY>")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var recommendations []Recommendation
    if err := json.NewDecoder(resp.Body).Decode(&recommendations); err != nil {
        return nil, err
    }
    return recommendations, nil
}
```

控制台:
http://localhost:8088/overview


使用postgres存储数据，然后执行testData灌入测试数据， 访问： http://localhost:8088/api/recommend/user20?n=30， 查看推荐结果。
