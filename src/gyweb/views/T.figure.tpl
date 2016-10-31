<!--图片新闻Item-->
{{define "figure"}}
<figure class="effect-zoe">
    <img src="/{{(index .Images 0).Url}}" alt="{{(index .Images 0).Url}}" />
    <figcaption>
        <h2>标题: <span>{{.Title}}</span></h2>
        <p class="icon-links">
            <a href="#"><span class="icon-heart"></span></a>
            <a href="/browse_category?id={{.Id}}"><span class="icon-eye"></span></a>
            <a href="#"><span class="icon-paper-clip"></span></a>
        </p>
        <p class="description">{{.Describe}}</p>
    </figcaption>
</figure>
{{end}}