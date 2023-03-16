<script>
    export let name;
    export let href;
</script>

<style>
    a {
        @apply
        flex flex-row gap-4
        border border-2 rounded-lg p-2 align-middle
        hover:border-grey hover:text-grey
    }
</style>

<a class="border-black dark:border-white" href={href} target="_blank" rel="noreferrer">
    <slot />
    {name}
    <svg class="border-black dark:border-white m-auto hover:border-grey" width="14" height="14" viewBox="0 0 14 14" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
        <path d="M12.7904 1L12.7904 9.38052M13 1.20951H4.6174M12.8548 1.14815L1 13" stroke="currentColor"/>
    </svg>
</a>