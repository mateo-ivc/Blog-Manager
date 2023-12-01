<script>
    import Header from "../Header.svelte";

    export let stories = [];
    console.log(stories.length);
</script>

<svelte:head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" type="text/css" href="/style/style.css">
    <title></title>
</svelte:head>

<Header/>

<div id="parent">
    <div id="overview">
        <h2>Overview</h2>
        <p>This is the left-side overview div.</p>
        <ul>
            <li>Component 1</li>
            <li>Component 2</li>
            <li>Component 3</li>

        </ul>
    </div>
    <div id="preview">

        <iframe class="frame" style="">

        </iframe>

    </div>
    <div class="info">

    </div>
</div>


<style>
    #overview {
        height: 100%;
        width: 25%;
    }

    #preview {
        height: 100%;
        width: 50%;
    }

    .info {
        width: 25%;
    }

    #parent {
        height: calc(100% - 100px);
        display: flex;
        flex-direction: row;
        padding-bottom: 5px;
    }

    .frame {
        width: 100%;
        height: 100%;

    }

</style>