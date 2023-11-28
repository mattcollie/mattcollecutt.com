<script>
    export let border = true;
</script>


<div class="w-full !h-[50px] " class:border-b-[1px]={border} class:border-b-stroke={border} >
    <slot/>
</div>