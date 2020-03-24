<style>
.layui-box {
  margin-left: 10px;
  box-shadow: rgba(0, 0, 0, 0.05) 0px 1px 2px;
  background: rgb(255, 255, 255);
  padding: 20px 20px 0px;
  margin-bottom: 15px;
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

{{range $key, $value := .articleList}}
<div class="layui-box">
  <h1>
    <a href="/detail/{{$value.Id}}" class="article-title">{{$value.Title}}</a>
  </h1>
  <div class="article-meta">
    类别：<a href="#"><i class="layui-icon layui-icon-list"></i> {{$value.FkCategoryId | getCategory}}</a>
    <span class="date"> {{$value.UpdatedAt}}</span>
  </div>

  <div class="remark">{{$value.Remark}}</div>

  <div class="more">
    <a href="/detail/{{$value.Id}}">查看更多 <i class="layui-icon layui-icon-next"></i></a>
  </div>
</div>
{{end}}
