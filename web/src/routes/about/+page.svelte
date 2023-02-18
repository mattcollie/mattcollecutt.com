<script>
    import { page } from "$lib/store";
    const age = Math.abs(new Date(Date.now() - new Date(1996, 6, 24).getTime()).getUTCFullYear() - 1970)
    $page = "/about";
</script>

<h1>about</h1>

<div class="pt-12 flex flex-col gap-4">
    <p>Hi, I’m Matt.</p>

    <p>
        I’m currently {age} years old, working full time as a Software Developer at
        <a class="font-bold" href="http://mustard.co.nz" target="_blank">Mustard Technology Limited</a>.
        I’m very passionate about my job and always look for ways to further my knowledge
        and grow as a person personally and professionally.
    </p>

    <p>
        I’m very passionate about; cooking, brewing, hiking and coding. I enjoy experimenting and
        I think that’s why I’ve easily developed a taste for cooking and brewing. While with hiking
        I get to enjoy the serenity of nature and the wildlife that I come across on a journey.
    </p>

    <p>
        Outside of those hobbies I’d like to improve my writing and a blog is the obvious choice.
        It’s also a good place to collate and share any ideas that I have. It’s been on my life of
        things to do for a very long time and now is a good of time as any.
    </p>
</div>