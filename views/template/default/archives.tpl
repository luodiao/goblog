<style>
.layui-box {
  background: white;
  margin-left: 10px;
  box-shadow: rgba(0, 0, 0, 0.05) 0px 1px 2px;
  background: rgb(255, 255, 255);
  padding: 20px 20px 0px;
}
</style>

<div class="layui-box">
  <ul class="layui-timeline">
    {{range $key,$value := .articleList}}
    <li class="layui-timeline-item">
      <i class="layui-icon layui-timeline-axis">&#xe63f;</i>
      <div class="layui-timeline-content layui-text">
        <h3 class="layui-timeline-title">{{$value.UpdatedAt}}</h3>
        <h5><a href="/detail/{{$value.Id}}">{{$value.Title}}</a></h5>
        <p>
          {{$value.Remark}}
        </p>
      </div>
    </li>
    {{end}}

    <li class="layui-timeline-item">
      <i class="layui-icon layui-timeline-axis">&#xe63f;</i>
      <div class="layui-timeline-content layui-text">
        <div class="layui-timeline-title">过去</div>
      </div>
    </li>
  </ul>
</div>