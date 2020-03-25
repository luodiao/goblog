<!DOCTYPE html>
<html lang="en">
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>Bako - Personal Portfolio & Resume HTML Template</title>
    <meta name="description" content="Bako - Personal Portfolio & Resume HTML Template">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <link rel="shortcut icon" type="image/x-icon" href="images/favicon.png">
    <!-- STYLESHEETS -->
    {{assets_css "/static/layui/css/layui.css"}}
    {{assets_css "/static/css/bootstrap.min.css"}}
    <link rel="stylesheet" href="http://v.bootstrapmb.com/2020/2/3dfiy7434/css/style.css" type="text/css" media="all">
    <style>

    </style>
  </head>

  <body>
    <!-- site wrapper -->
    <div class="site-wrapper">
      <!-- mobile header -->
      <div class="mobile-header py-2 px-3 mt-4">
        <button class="menu-icon mr-2">
          <span></span>
          <span></span>
          <span></span>
        </button>
        <a href="/" class="logo">
          <img src="{{.site.Avatar}}" width="70" height="70" style="border-radius: 100% !important;" /></a>
        <a href="/" class="site-title dot ml-2">{{.site.Name}}</a></div>

      <!-- header -->
      <header class="left float-left shadow-dark" id="header">
        <button type="button" class="close" aria-label="Close">
          <span aria-hidden="true">&times;</span></button>
        <div class="header-inner d-flex align-items-start flex-column">
          <a href="/">
            <img src="{{.site.Avatar}}" class="rounded" width="70" height="70" style="border-radius: 100% !important;" />
          </a>
          <a href="/" class="site-title dot mt-3">{{.site.Name}}</a>
          <span class="site-slogan">{{.site.Location}}</span>
          <!-- navigation menu -->
          <nav>
            <ul class="vertical-menu scrollspy">
              {{range $key, $value := .site.Navs}}
              <li>
                <a href="{{$value.url}}">
                  <i class="{{$value.icon}}"></i> {{$value.title}}
                </a>
              </li>
              {{end}}
            </ul>
          </nav>

          <!-- footer -->
          <div class="footer mt-auto">
            <!-- social icons -->
            <ul class="social-icons list-inline">
              <li class="list-inline-item">
                <a href="#">
                  <i class="layui-icon layui-icon-login-wechat"></i>
                </a>
              </li>
              <li class="list-inline-item">
                <a href="#">
                  <i class="layui-icon layui-icon-login-qq"></i>
                </a>
              </li>
              <li class="list-inline-item">
                <a href="#">
                  <i class="layui-icon layui-icon-login-weibo"></i>
                </a>
              </li>
              <li class="list-inline-item">
                <a href="#">
                  <i class="layui-icon layui-icon-rss"></i>
                </a>
              </li>
              <li class="list-inline-item">
                <a href="#">
                  <i class="layui-icon layui-icon-share"></i>
                </a>
              </li>
            </ul>
            <!-- copyright -->
            <span class="copyright">© 2020 Gurudin Template</span></div>
        </div>
      </header>

      <!-- main content area -->
      <main class="content float-right">
        {{.LayoutContent}}
      </main>
    </div>
    <!-- Go to top button -->
    <a href="javascript:" id="return-to-top">
      <i class="fa fa-chevron-up"></i>
    </a>
  </body>
</html>
