
# import importlib
import importlib
from plugins.sms import sms_pb2, sms_pb2_grpc


class SMSServicer(sms_pb2_grpc.SMSServicer):
    """Implementation of KV service."""

    impl = importlib.import_module("impl.sms_impl")
    conf = importlib.import_module("conf")

    def Send(self, request, context):
        print("-------- SMS ---------")
        res = sms_pb2.Res()
        try:
            srv = self.conf.json_data['services']["sms"]
            if srv['banned']==False:
                impl = importlib.reload(self.impl)
                item = self.impl.send(request)
                res.result = item[0]
                res.msg = item[1]
            else:
                print("功能禁用")
                res.result = False
                res.msg = "功能禁用"
        except Exception as e:
            print("功能被卸载")
            res.result = False
            res.msg = "功能被卸载"
        return res
        

