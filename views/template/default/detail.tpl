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
.layui-box .content {
  padding: 20px 0;
}
</style>

<div class="layui-box">
  <h1>{{.article.Title}}</h1>
  <div class="article-meta">
    类别：<a href="#"><i class="layui-icon layui-icon-list"></i> {{.article.FkCategoryId | getCategory}}</a>
    <span class="date"> {{.article.UpdatedAt}}</span>
  </div>

  <div class="content">{{str2html .article.Content}}</div>
</div>
