<!-- section hero -->
<section class="hero background parallax shadow-dark d-flex align-items-center" id="home" style="background-image: url(/static/image/t1.png);">
    <div class="cta mx-auto mt-2">
        <h1 class="mt-0 mb-4">I’m {{.site.Name}}</h1>
        <p class="mb-4">
            Speaking in silence is the most intimate communication two beings can achieve.
            <br>
            无声的交谈是两个生命所能达到的最亲密的交流方式。
        </p>
        <a href="https://github.com/Gurudin" class="btn btn-default btn-lg mr-3">
            <i class="icon-grid"></i> View GitHub
        </a>
    </div>
    <div class="overlay"></div>
</section>

<!-- section skills -->
{{range $key, $value := .articleList}}
<section id="skills" class="shadow-blue white-bg padding">
    <h3 class="section-title"><a href="/detail/{{$value.Id}}" class="article-title">{{$value.Title}}</a></h3>
    <div class="spacer" data-height="80"></div>

    <p class="mt-3">{{$value.Remark}}</p>
</section>
{{end}}
