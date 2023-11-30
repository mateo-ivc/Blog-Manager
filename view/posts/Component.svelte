<script>
    export let section;


    function getProps(values) {
        let props = {}
        for (const [key, value] of Object.entries(values)) {
            if (key.startsWith("@")) {
                const prop = key.replace("@", "")
                props[prop] = value
            }
        }
        return props
    }

    function getContent(values) {
        let content = "";
        for (const [key, value] of Object.entries(values)) {
            if (key.startsWith("#")) {
                return value
            }
        }
        return ""
    }
</script>


{#each Object.entries(section) as [key, values]}
    {#if (typeof values === 'object')}
        <svelte:element this="{key}" {...getProps(values)}>
            {getContent(values)}
            <svelte:self bind:section={values}/>
        </svelte:element>
    {/if}
{/each}

<style>

</style>