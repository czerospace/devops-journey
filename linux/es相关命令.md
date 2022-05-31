# 强制merge

url -X POST 'http://127.0.0.1:9200/ipanelcloudlog/_forcemerge?only_expunge_deletes'

# reindex
POST _reindex
{
  "source": {
    "index": "ipanelcloudlog1"
  },
  "dest": {
    "index": "ipanelcloudlog531"
  }
}

# 查看任务
GET _cat/tasks

# 取消任务
POST _tasks/qD8f05K7RZyWhducmitFAA:323420274/_cancel

# 查看reindex任务
GET _cat/tasks?actions=*reindex

# 删除 别名
POST /_aliases
{
    "actions": [
        {"remove": {"index": "ipanelcloudlog1", "alias": "ipanelcloudlog"}}
    ]
}

# 按日期查询
GET ipanelcloudlog530/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "time": {
              "lt": "2022-05-25T00:00:00.000+0800"
            }
          }
        }
      ]
    }
  }
}

# 按日期查询删除
POST ipanelcloudlog530/_delete_by_query?refresh
{
  "query": {
    "bool": {
      "must": [
        {
          "range": {
            "time": {
              "lt": "2022-05-25T00:00:00.000+0800"
            }
          }
        }
      ]
    }
  }
}