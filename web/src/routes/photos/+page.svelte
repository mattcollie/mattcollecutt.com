<svelte:head>
    <title>Photos</title>
    <meta name="description" content="Photos" />
</svelte:head>

<div  class="grid grid-cols-4 gap-4"
        hx-get="http://127.0.0.1:5000/instagram/media"
        hx-trigger="intersect once"
        hx-swap="beforeend">

</div>