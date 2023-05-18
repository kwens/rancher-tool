# rancher

rancher 工具，调用rancher的api，目前仅提供deployment的update

## Usage
1. build：执行对应系统的脚本即可
```
./build_linux.sh
```
2. 执行update命令
   - update参数说明：
   - host: rancher api地址
   - token: 请求rancher api所需要的token，从rancher里面的`账号&API密钥` 可创建
   - project: 项目名称
   - namespace: 命名空间名称 
   - deployment: 工作负载名称
   - container: 容器名称
   - tag: 镜像tag

使用示例(请自行替换下面的参数)：
```
./rancher-tools update --host=https://192.168.xx.xx:xxx --token=access_key:access_secret -p $DEPLOY_PROJECT_NAME -n $DEPLOY_NAMESPACE_NAME -d $DEPLOY_DEPLOYMENT_NAME -c $DEPLOY_CONTAINER_NAME -t $BUILD_TAG
```
