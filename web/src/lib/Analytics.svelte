<script>
    import { page } from '$app/stores'

    $: {
        if (typeof gtag !== 'undefined') {
            gtag('config', 'G-NHJJ1HY5YK', {
                page_title: document.title,
                page_path: $page.url.pathname,
            })
        }
    }
</script>

<svelte:head>
    <script
            async
            src="https://www.googletagmanager.com/gtag/js?id=G-NHJJ1HY5YK">
    </script>
    <script>
        window.dataLayer = window.dataLayer || []

        function gtag() {
            dataLayer.push(arguments)
        }

        gtag('js', new Date())
        gtag('config', 'G-NHJJ1HY5YK')
    </script>
</svelte:head>