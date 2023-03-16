<script>
    import {page} from "$lib/store";

    export let href;
    export let target = "_self";
    let text;
</script>

<style>
    a {
        @apply
        text-lg
    }

    span {
        @apply
        px-2 py-1
        text-grey-600
        hover:text-black
    }

    :global(.dark span) {
        @apply
        text-grey-200
        hover:text-white
    }

    .active-link {
        @apply
        text-black bg-grey-100 rounded-md
    }

    :global(.dark .active-link) {
        @apply
        text-white bg-grey-800
    }
</style>

<a class="text-black dark:text-white" href={href} target={target}>
    <span bind:this={text} class:active-link={$page === href}><slot/></span>
</a>