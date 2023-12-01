<script>
    import Header from "./Header.svelte"
    import Teaser from "./Teaser.svelte";

    export let stories = []
    console.log(stories.length)


</script>
<svelte:head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" href="/style/style.css">
    <title></title>
</svelte:head>

<Header/>
<main>
    <h1 class="page-width">Index</h1>

    <div class="section-article page-width">

        {#if stories.length > 0 }
            <div class="article-container">
                {#each stories as article}
                    <Teaser {article}/>
                {/each}
            </div>
        {/if}
    </div>

</main>

<footer>
    <div class="page-width">
        <p>&copy; 2023 Your Blog. All rights reserved.</p>
    </div>
</footer>
<style>

    main {
        min-height: 100%;
    }

    .article-container {

        display: flex;
        justify-content: space-between;
        flex-wrap: wrap;
        gap: 30px;
    }

    @media only screen and (max-width: 600px) {

        main {
            padding: 10px;
        }

        footer {
            margin-top: auto;
        }
    }
</style>
