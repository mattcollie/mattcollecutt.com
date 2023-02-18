<script>
    import { page } from "$lib/store.js";
    export let href;
    export let target = "_self";
    let text;
</script>

<style>
    a {
        @apply
            text-xl lg:text-2xl
            tracking-wider
            hover:font-bold
    }

    a::before {
        display: block;
        content: attr(data-text);
        font-weight: bold;
        height: 0;
        overflow: hidden;
        visibility: hidden;
    }
</style>

<a class="text-black dark:text-white" class:font-bold={$page === href} href={href} target={target} data-text={text? text.textContent : ""}>
    <span bind:this={text}><slot/></span>
</a>