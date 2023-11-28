<script>
    import { page } from '$app/stores';

    export let href = "/";
    export let external = false;
    export let target;
</script>

<a href={href} target={target} class="flex flex-row items-center gap-3 w-full h-fit px-2 py-1.5 rounded-lg hover:bg-stroke-light justify-between" class:!bg-stroke-light={$page.url.pathname === href}>
    <div class="flex flex-row items-center gap-3">
        <slot name="icon"/>
        <span class="text-xs font-semibold"><slot/></span>
    </div>
    {#if external}
        <svg width="10" height="9" viewBox="0 0 13 12" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M12.1121 1.15128C12.1121 0.519938 11.5994 0.00728558 10.9681 0.00728558H2.88685C2.2555 0.00728558 1.74285 0.519938 1.74285 1.15128C1.74285 1.78263 2.2555 2.29528 2.88685 2.29528L8.21288 2.29276L0.465006 10.0406C0.0180137 10.4876 0.0180137 11.2099 0.465006 11.6569C0.911999 12.1039 1.63426 12.1039 2.08125 11.6569L9.8266 3.91153L9.82912 9.2325C9.82912 9.86385 10.3418 10.3765 10.9731 10.3765C11.6045 10.3765 12.1171 9.86385 12.1171 9.2325L12.1171 1.15128H12.1121Z" fill="white"/>
        </svg>
    {/if}
</a>