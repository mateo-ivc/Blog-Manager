<script>
    export let section;


    console.log(section)

    function getCSS(settings) {
        let style = ""
        for (const setting in settings) {
            style += `${setting}:${section.settings[setting]};`
        }
        console.log(style)
        return style
    }

</script>
<div>
    <img src="{section.link}" width="100%" alt="" style="{getCSS(section.settings)}">
</div>
<style></style>
