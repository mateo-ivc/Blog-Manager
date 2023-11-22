<script>
    import Header from "../Header.svelte";
    import Text from "./Text.svelte";

    export let story
    let post = story.Story

    console.log(post.sections.text)


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
            <img src="{post.teaser}" width="100%" height="800px" style="object-fit: cover" alt="">
        </div>
        <div class="description">
            <h1>{post.title}</h1>
            <a>{post.author} | {post.date}</a>
        </div>
    </div>
    {#each Object.keys(post.sections) as section (section)}


        {#if section === "text"}
            <!-- Todo: Create Component to render Text -->

            <Text bind:section={post.sections[section]}/>
        {/if}
        <!--{section}-->
    {/each}
</div>

<style>
    .media {
        position: relative;
        display: flex;
    }

    .media-wrapper {
        width: 100%;
    }

    .description {
        position: absolute;
        top: 50%;
        left: 50%;

        transform: translate(-50%, -50%);

        display: flex;
        flex-direction: column;
        text-align: center;
    }
</style>