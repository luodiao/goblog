<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
  <title>{{.site.Title}}</title>
  {{assets_css "/static/layui/css/layui.css"}}
  {{assets_js "/static/js/jquery.min.js"}}
  {{assets_js "/static/js/vue.js"}}
  {{assets_js "/static/js/site.js"}}
  {{assets_js "/static/layui/layui.all.js"}}
  <style>
    body {
      color: rgb(86, 90, 95);
      text-size-adjust: 100%;
      background: rgb(245, 248, 249);
      font: 16px "open sans", "Helvetica Neue", "Microsoft Yahei", Helvetica, Arial, sans-serif;
    }
    .layui-main {
      width: 80%;
    }
    #header {
      background: #fff;
      position: relative;
      height: 64px;
      box-shadow: rgba(0, 0, 0, 0.05) 0px 1px 2px;
    }
    #header #logo .site-title {
      font-size: 16px;
      display: inline-block;
      vertical-align: top;
      font-weight: 600;
    }
    #header #logo, #main-nav {
      float: left;
      height: 40px;
      line-height: 40px;
      padding: 12px 15px;
    }
    .main-nav-link {
      font-weight: 300;
      hite-space: nowrap;
      color: rgb(86, 90, 95);
      transition: all 0.2s ease 0s;
      float: left;
      display: block;
      padding: 0px 15px;
    }
    a:hover {
      color: rgb(56, 183, 234);
    }
    #search-form-wrap {
      float: right;
      eight: 40px;
      line-height: 40px;
      padding: 12px 15px;
    }
    #search-form-wrap .search-form {
      position: relative;
      border: 1px solid rgb(236, 239, 242);
      border-radius: 21px;
    }
    #search-form-wrap .search-form .search-form-input {
      width: 115px;
      height: 30px;
      border-radius: 21px;
      border: none;
      padding: 0 0 0 15px;
      margin-right: 45px;
    }
    #search-form-wrap .search-form-submit {
      float: right;
      position: absolute;
      right: 15px;
      top: 10px;
      border: none;
    }

    .outer {
      position: relative;
      margin-top: 30px;
    }

    .profile {
      padding: 15px 20px;
      border-bottom: 1px solid rgb(236, 239, 242);
      background: white;
      display: table-row;
    }
    .profile .avatar {
      display: block;
      margin: 20px auto;
    }
    .profile .name {
      font-size: 20px;
      font-weight: 600;
      display: block;
      text-align: center;
    }
    .profile .location {
      font-size: 12px;
      margin-top: 5px;
      color: rgb(154, 158, 163);
      text-align: center;
    }
    .profile .stat {
      border-top: 1px solid rgb(236, 239, 242);
      margin-top: 30px;
    }
    .profile .stat p {
      position: relative;
      float: left;
      width: 49.5%;
      text-align: center;
      border-bottom: 1px solid rgb(236, 239, 242);
      padding: 30px 0;
      font-size: 16px;
    }
    .profile .stat p.first{
      border-right: 1px solid rgb(236, 239, 242);
    }
    .profile .stat-bottom div {
      text-align: center;
      padding: 30px 0 15px 0;
    }
    .profile .stat-bottom div i{
      font-size: 30px;
    }
    
    .sidebar .recent-post{
      margin-top: 15px;
      margin-left: 15px;
    }
    .sidebar .recent-post a {
      display: block;
      font-size: 14px;
      margin-bottom: 3px;
    }
    .sidebar .recent-post a.first {
      color: rgb(56, 183, 234);
    }
    .sidebar .recent-post p {
      font-size: 13px;
      color: #c2c2c2;
    }
  </style>
</head>
<body>
  <header id="header">
    <div class="layui-main">
      <a href="/" id="logo">
        <i class="logo"></i>
        <span class="site-title">{{.site.Title}}</span>
      </a>
      <nav id="main-nav">
        {{range $key, $value := .site.Navs}}
        <a class="main-nav-link" href="{{$value.url}}">{{$value.title}}</a>
        {{end}}
      </nav>

      <div id="search-form-wrap">
        <form class="search-form">
          <input type="text" class="ins-search-input search-form-input" placeholder="搜索">
          <button type="button" class="search-form-submit"><i class="layui-icon layui-icon-search"></i></button>
        </form>
      </div>
    </div>
  </header>
  
  <div class="outer layui-main">
   <div class="layui-row">
      <div class="layui-col-md2">
        <div class="profile">
          <img src="{{.site.Avatar}}" class="avatar" width="70%">
          <h2 class="name">{{.site.Name}}</h2>
          <p class="location"><i class="layui-icon layui-icon-location"></i> {{.site.Location}}</p>

          <div class="stat">
            <p class="first">{{.total}}<br>文章</p>
            <p>53<br>标签</p>
          </div>

          <div class="stat-bottom">
            <div class="layui-col-md4">
              <a href="#"><i class="layui-icon layui-icon-login-wechat"></i></a>
            </div>
            <div class="layui-col-md4">
              <a href="#"><i class="layui-icon layui-icon-rss"></i></a>
            </div>
            <div class="layui-col-md4">
              <a href="#"><i class="layui-icon layui-icon-share"></i></a>
            </div>
          </div>
        </div>
      </div>

      <div class="layui-col-md7 main">
        {{.LayoutContent}}
      </div>

      <div class="layui-col-md3 sidebar">
        <h5 class="widget-title" style="margin-left: 15px;">最新文章</h5>
        {{range $key, $value := .articleTop5}}
        <div class="recent-post">
          <a href="/?category={{$value.FkCategoryId}}" class="first">{{$value.FkCategoryId | getCategory}}</a>
          <a href="/detail/{{$value.Id}}">{{$value.Title}}</a>
          <p>2020-02-10</p>
        </div>
        {{end}}
      </div>
    </div>
  </div>

</body>
</html>