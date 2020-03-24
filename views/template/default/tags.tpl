<style>
.layui-box {
  margin-left: 10px;
  box-shadow: rgba(0, 0, 0, 0.05) 0px 1px 2px;
  background: rgb(255, 255, 255);
  padding: 20px;
  margin-bottom: 15px;
}
.layui-badge {
  padding: 5px 10px;
}
</style>
<div class="layui-box">
  {{range $key, $value := .tags}}
  <a href="{{$value.url}}"><span class="layui-badge {{$value.color}}">{{$key}}</span></a>
  {{end}}
</div>