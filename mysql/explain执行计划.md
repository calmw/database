#### 介绍

    通过explain命令获取Mysql如何执行查询语句的信息，包括在查询语句执行过程中表如何连接和连接的顺序。[**使用示例**](../images/explain_select1.png)
        select_type表示查询语句的类型，常见取值：
            SIMPLE 简单查询，不包含子查询或者集合运算UNION
            PRIMARY 
        table 表述输出结果的表或者表明
        partitions 表示查询语句访问到了表中的哪个区
        type 表示Mysql在表中找到所需行的访问方式或者连接类型
            system 
        possible_keys表示查询可能使用的索引
        key表示查询实际使用的索引
        key_len 表示使用索引字段的长度
        ref 表示查询语句使用索引时用到了表中的那一列
        rows 表示查询语句根据表统计信息及索引使用情况，进行行扫描数量。这个值越小越好。
        filtered表示满足查询要求的行占整个存储引擎返回数量的百分比。
        filtered表示查询语句其他情况的说明和描述，包含不适合在其他列中显示但是对执行计划非常重要的额外信息。
        
      