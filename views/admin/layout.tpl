<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>Go blog</title>
  <link rel="stylesheet" href="/static/layui/css/layui.css">
  <script src="/static/js/jquery.min.js"></script>
  <script src="/static/js/vue.js"></script>
  <script src="/static/js/site.js"></script>
  <script src="/static/layui/layui.all.js"></script>
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
  <div class="layui-header">
    <div class="layui-logo">Goblog</div>
    <!-- 头部区域（可配合layui已有的水平导航） -->
    <ul class="layui-nav layui-layout-left">
      <li class="layui-nav-item"><a href="/admin">控制台</a></li>
      <!-- <li class="layui-nav-item"><a href="">商品管理</a></li>
      <li class="layui-nav-item"><a href="">用户</a></li>
      <li class="layui-nav-item">
        <a href="javascript:;">其它系统</a>
        <dl class="layui-nav-child">
          <dd><a href="">邮件管理</a></dd>
          <dd><a href="">消息管理</a></dd>
          <dd><a href="">授权管理</a></dd>
        </dl>
      </li>-->
    </ul>
    <ul class="layui-nav layui-layout-right">
      <li class="layui-nav-item">
        <a href="javascript:;">
          <i class="layui-icon layui-icon-friends"></i>
          贤心
        </a>
        <dl class="layui-nav-child">
          <dd><a href="">修改密码</a></dd>
          <dd><a href="">退出</a></dd>
        </dl>
      </li>
    </ul>
  </div>
  
  <div class="layui-side layui-bg-black">
    <div class="layui-side-scroll">
      <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
      <ul class="layui-nav layui-nav-tree"  lay-filter="test">
        <li class="layui-nav-item"><a href=""><i class="layui-icon layui-icon-console"></i>&nbsp;&nbsp;&nbsp; Site 站点设置</a></li>
        <li class="layui-nav-item layui-nav-itemed">
          <a class="" href="javascript:;"><i class="layui-icon layui-icon-read"></i>&nbsp;&nbsp;&nbsp; Article 文章管理</a>
          <dl class="layui-nav-child">
            <dd><a href="/admin/category">类别设置</a></dd>
          </dl>
        </li>
        <!-- <li class="layui-nav-item">
          <a href="javascript:;">解决方案</a>
          <dl class="layui-nav-child">
            <dd><a href="javascript:;">列表一</a></dd>
            <dd><a href="javascript:;">列表二</a></dd>
            <dd><a href="">超链接</a></dd>
          </dl>
        </li> -->
      </ul>
    </div>
  </div>
  
  <div class="layui-body">
    <!-- 内容主体区域 -->
    <div style="padding: 15px;">{{.LayoutContent}}</div>
  </div>
  
  <div class="layui-footer">
    <!-- 底部固定区域 -->
    © Goblog
  </div>
  <script src="/static/layui/layui.all.js"></script>
</div>
</body>
</html>