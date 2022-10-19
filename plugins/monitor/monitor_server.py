import importlib
from plugins.monitor import monitor_pb2,monitor_pb2_grpc
from impl import sms_impl
import conf


# install 每一次重载插件的时候，必须重载配置文件
# update 
# banned 插件 禁用/启用
# uninstall 每一次修改配置时，必须在修改后写入配置文件

class MonitorServicer(monitor_pb2_grpc.MonitorServicer):
    
    importlib.import_module

    # 这个管理禁用和卸载
    # 因为禁用和卸载总是对单个服务操作
    def Command(self, request, context):
        print("---------   command   ---------")

        self.service(request.cmd,request.path)

        callback = monitor_pb2.CallBack()
        callback.result = True
        callback.msg = request.cmd +' '+ request.path + ' success'
        print("--------- command end ---------")
        return callback

    # 这个管理安装和更新
    # 因为安装和更新总是带着多文件(至少带有配置文件)
    def Commands(self, request, context):
        print("---------   commands   ---------")

        self.service(request.cmd,request.paths)
        callback = monitor_pb2.CallBack()
        callback.result = True
        callback.msg = request.cmd +' '+ ",".join(request.paths) + ' success'
        return callback
        
    def install(list):
        print('do install')
        print(list)
        for name in list:
            if name == "sms":
                importlib.reload(sms_impl)
            if name == "conf":
                importlib.reload(conf)

    def update(list):
        print('do update')
        for name in list:
            if name == "sms":
                importlib.reload(sms_impl)
            if name == "conf":
                importlib.reload(conf)

    def banned(name):
        print('do banned')
        print(conf.json_data['services'][name]['banned'])
        conf.json_data['services'][name]['banned'] = not conf.json_data['services'][name]['banned']
        conf.setJson()

    def uninstall(name):
        print('do uninstall')
        conf.json_data['services'].pop(name)
        conf.setJson()
        

    operator = {'update':update,'install':install,'uninstall':uninstall,'banned':banned}

    def service(self,command,name):
        print('command srv:'+command)
        self.operator.get(command)(name)