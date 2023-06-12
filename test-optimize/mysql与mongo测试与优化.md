#### 插入数据对比

- 数据格式 [**数据样本**](../static/data-demo.png)
- 插入一百万条数据对比
    - mysql：58m4.782290875s, mysql本次平均每秒插入287条
    - mongodb: 16m58.691230958s, mongo本次平均每秒插入980条
    - 本次测试两者都未加索引，在相同的资源环境下，数据中json键值对数量随机，从3-15不等，最终插入速度大约相差3.5倍
- 插入五百万条数据对比
    - mysql：5h8m9.844109167s, mysql本次平均每秒插入270条
    - mongodb: 1h28m39.358680791s, mongo本次平均每秒插入940条
    - 本次测试两者都未加索引，在相同的资源环境下，数据中json键值对数量随机，从3-15不等，最终插入速度大约相差3.5倍

#### mysql查询，以及优化

- 查询1
- ![select1.png](..%2Fimages%2Fselect1.png)
- 增加索引
- ![add_index1.png](..%2Fimages%2Fadd_index1.png)
- 添加索引后查询
- ![select2.png](..%2Fimages%2Fselect2.png)

#### 使用explain命令查看执行计划

- 在原来查询语句基础上使用explain
- ![explain_select1.png](..%2Fimages%2Fexplain_select1.png)
- [**使用explain执行计划**](../mysql/explain执行计划.md)