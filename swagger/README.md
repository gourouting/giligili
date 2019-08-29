# G站文档说明

## 简介

文档使用swagger(OpenAPI 3.0)格式编写，但是为了结构化和层次化调整了一些细节。

## swagger-merger

为了把结构化的文档合并成一个单独的json文件，使用了以下工具:

https://www.npmjs.com/package/swagger-merger

```bash
npm install -g swagger-merger
cd swagger
swagger-merger -i index.yaml -o swagger.json
```