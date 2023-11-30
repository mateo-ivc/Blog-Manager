<script>
    import Header from "../Header.svelte";
    import Component from "./Component.svelte";

    export let story
    let post = story.Story
    let info = post.info


</script>

<header>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" href="/style/style.css">
</header>

<Header/>
<div class="page-width">
    <div class="media">
        <div class="media-wrapper">
            <img src="{info.teaser}" width="100%" style="object-fit: cover" alt="">
        </div>
        <div class="description">
            <div class="title"><h1>{info.title}</h1></div>
            <div class="author">
                <a>Written by</a>
                <a>{info.author} | {info.date}</a>
            </div>
        </div>
    </div>
    <hr>


    <Component bind:section={post.layout}/>


</div>

<style>
    .media {
        margin-bottom: 20px;
    }

    .media-wrapper {
        width: 100%;
        margin-bottom: 10px;
    }

    .media-wrapper img {
        height: 75vh;
    }

    .description {
        display: flex;
        justify-content: space-between;
    }

    .title {
        align-self: end;
    }

    .author {
        display: flex;
        flex-direction: column;
        font-weight: 300;
        color: rgb(128, 128, 128);
        text-align: end;
    }
</style>