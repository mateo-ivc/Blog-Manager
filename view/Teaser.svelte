<script>
    export let article;
    let story = article.Story
    let post = story.info
</script>
<div class="media">
    <a href="{post.link}">
        <div class="content">
            <img src="{post.teaser}" width="100%" height="100%" style="object-fit: cover" alt="">
            <div class="description">
                <div>
                    <h1>{post.title}</h1>
                    <a>{post.date} | {post.author}</a>
                </div>
            </div>
        </div>
    </a>

</div>
<style>
    .media {
        width: 100%;
        height: 350px;
        flex: 0 0 calc(50% - 15px);
        box-sizing: border-box;
        border: 1px solid #ddd;
        border-radius: 15px;
        overflow: hidden;
    }

    a, a:visited, a:hover, a:active {
        color: inherit;
    }

    .content {
        position: relative;
        height: 100%;
        width: 100%;
    }

    h1 {
        margin: 0;
        font-size: 25px;
    }

    .description {
        position: absolute;
        width: 100%;
        bottom: 0;
        background: rgba(255, 255, 255, 0.8);
        height: fit-content;
    }

    .description > * {
        margin: 5px;
    }

</style>