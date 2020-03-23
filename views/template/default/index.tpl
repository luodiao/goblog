<style>
.layui-box {
  background: white;
  margin-left: 10px;
  box-shadow: rgba(0, 0, 0, 0.05) 0px 1px 2px;
  background: rgb(255, 255, 255);
  padding: 20px 20px 0px;
}
a:hover {
  color: rgb(56, 183, 234);
}
a {
  color: rgb(86, 90, 95);
}
.layui-box .date {
  font-size: 13px;
  color: #c2c2c2;
  margin-left: 20px;
}
.layui-box .article-meta {
  font-size: 15px;
}
.layui-box .remark {
  margin-top: 20px;
}
.layui-box .more {
  margin-top: 20px;
  padding: 20px; 0;
  border-top: 1px solid rgb(236, 239, 242);
  text-align: right;
}
</style>

<div class="layui-box">
  <h1>
    <a href="#" class="article-title">Arthas | 定位线上 Dubbo 线程池满异常</a>
    
  </h1>
  <div class="article-meta">
    类别：<a href="#"><i class="layui-icon layui-icon-list"></i> Arthas</a>
    <span class="date"> 2020-01-01</span>
  </div>

  <div class="remark">
    前言
    本文是 Arthas 系列文章的第二篇。

    Dubbo 线程池满异常应该是大多数 Dubbo 用户都遇到过的一个问题，本文以 Arthas 3.1.7 版本为例，介绍如何针对该异常进行诊断，主要使用到 dashboard/thread 两个指令。
  </div>

  <div class="more">
    <a href="#">查看更多 <i class="layui-icon layui-icon-next"></i></a>
  </div>
</div>