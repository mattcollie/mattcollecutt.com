<script>
    export let title = "";
</script>

<div class="flex flex-col p-4 gap-1">
    <span class="text-xs font-semibold">{title}</span>
    <slot/>
</div>