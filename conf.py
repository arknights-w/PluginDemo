import json

from sympy import re

def getJson():
    with open("./impl/conf.json",encoding="utf-8") as fp:
        print("----------- get conf -----------")
        json_data = json.load(fp)
    # print('这是文件中的json数据: ',json_data)
    # print('这是读取到文件数据的数据类型：', type(json_data))
    # for v in json_data.get('services'):
        # print(v)
        # 取出特定数据
        # print("%s,%s"%(v['name'],v['description']))
    return json_data

def setJson():
     with open("./impl/conf.json","w",encoding="utf-8") as fp:
        try:
            json.dump(json_data,fp)
        except Exception as e:
            print(e)

json_data = getJson()