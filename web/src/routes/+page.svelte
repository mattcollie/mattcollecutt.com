<script>
    import Banner from "$lib/components/Banner.svelte";
</script>

<style>
    a {
        @apply
            text-[#1F6FEB]
            hover:underline
        ;
    }
</style>

<svelte:head>
    <title>Matt Collecutt</title>
    <meta name="description" content="Home page" />
</svelte:head>


<Banner></Banner>

<div class="mx-auto w-full py-12 max-w-prose">
    <article class="flex flex-col gap-5 text-sm font-light leading-6 tracking-wide max-w-prose">
        <p>
            Hi, I’m Matt! I'm a software developer, <a href="/photos">amateur photographer</a>,
            <a href="/writing">writer</a>. Currently I'm working at
            <a href="http://www.mustard.co.nz/" target="_blank">mustard technology</a>, a small development company
            based in New Zealand.
        </p>

        <p>
            Before Mustard, I worked at <a href="https://www.theaccessgroup.com/">Ezitracker</a>, where I started my
            software development journey learning the ropes as an intern.
        </p>

        <p>
            In my spare time I've been working on a few projects like <a href="https://mattcollecutt.co.nz">this</a>
            website, <a href="https://trivino.xyz" target="_blank">TriVino</a> and
            <a href="https://nookly.co.nz" target="_blank">Nookly</a>
        </p>
    </article>
</div>
