# note
http://blog.csdn.net/liuxinmingcode/article/details/50480097

### git教程
http://backlogtool.com/git-guide/cn/intro/intro1_1.html

- [x] [exec](https://github.com/luopengift/transport/blob/master/plugins/input/exec/README.md): 执行程序/脚本
- [x] 实现数据处理接口 
- [x] 已实现的数据处理组件
    - [x] Default,直接连接input,output
    - [x] 在行尾加入换行符,例子:写文件
    
    
    http://bbs.51cto.com/thread-1071001-1-1.html
    
http://www.bootcss.com/p/git-guide/


使用Golang实现的无锁队列，性能与Disruptor相当达到1400万/秒
https://zhuanlan.zhihu.com/p/24432607
https://myronliu347.github.io/vue-carbon/#!/



![nginx](https://github.com/luopengift/transport/blob/master/Image/png/nginx_log_analysis.png)
![nginx](https://github.com/luopengift/transport/blob/master/static/png/nginx_log_analysis.png)


```
flow
st=>start: Start|past:>http://www.google.com[blank]
e=>end: End:>http://www.google.com
op1=>operation: get_hotel_ids|past
op2=>operation: get_proxy|current
sub1=>subroutine: get_proxy|current
op3=>operation: save_comment|current
op4=>operation: set_sentiment|current
op5=>operation: set_record|current

cond1=>condition: ids_remain空?
cond2=>condition: proxy_list空?
cond3=>condition: ids_got空?
cond4=>condition: 爬取成功??
cond5=>condition: ids_remain空?

io1=>inputoutput: ids-remain
io2=>inputoutput: proxy_list
io3=>inputoutput: ids-got

st->op1(right)->io1->cond1
cond1(yes)->sub1->io2->cond2
cond2(no)->op3
cond2(yes)->sub1
cond1(no)->op3->cond4
cond4(yes)->io3->cond3
cond4(no)->io1
cond3(no)->op4
cond3(yes, right)->cond5
cond5(yes)->op5
cond5(no)->cond3
op5->e作者：Jlan链接：http://www.jianshu.com/p/02a5a1bf1096來源：简书著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```
