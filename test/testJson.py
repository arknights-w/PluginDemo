import json


with open("./test/test.json",encoding="utf-8") as fp:
    json_data = json.load(fp)
    # print('这是文件中的json数据: ',json_data)
    print('这是读取到文件数据的数据类型：', type(json_data))
    for v in json_data.get('services'):
        # print(v)
        # 取出特定数据
        print("%s,%s"%(v['name'],v['description']))