
namespace go http

struct BizRequest {
    1: i64 vint64(api.query = 'vint64', api.vd = "$>0&&$<200")
    2: string text(api.body = 'text')
    // 3: i32 token(api.header = 'token')
    // 6: list<string> items(api.query = 'items')
    // 7: i32 version(api.path = 'version')
}

struct BizResponse {
    # 1: i32 token(api.header = 'token')
    2: string text(api.body = 'text')
    5: i32 http_code(api.http_code = '')
}

service BizService {
    BizResponse BizMethod1(1: BizRequest req)(api.get = '/life/client/:version', api.baseurl = 'example.com', api.param = 'true')
    # BizResponse BizMethod2(1: BizRequest req)(api.post = '/life/client/:version', api.baseurl = 'example.com', api.param = 'true', api.serializer = 'form')
    # BizResponse BizMethod3(1: BizRequest req)(api.post = '/life/client/:version/other', api.baseurl = 'example.com', api.param = 'true', api.serializer = 'json')
}
