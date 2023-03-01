<script>
    import {Swipeable} from 'thumb-ui'
    import { fade } from 'svelte/transition'
    import { page } from "$lib/store";
    $page = "/gallery";
    export let data;

    let loginProgress, introProgress;
</script>

<svelte:head>
    <title>{data.name}</title>
</svelte:head>

<div id="wrapper">
    <a class="absolute scale-50 cursor-pointer z-50 hover:scale-[0.60] transition top-5 left-5" href="/gallery">
        <svg width="59" height="60" viewBox="0 0 59 60" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
            <path d="M1.35352 1.64645L57.9221 58.215M1.35352 58.9221L57.9221 2.35356" stroke="currentColor" stroke-width="3"/>
        </svg>
    </a>
    <div class="slides fullpage" style="transform: scale({1 - $loginProgress*0.3})">
        <Swipeable numScreens={data.photos.length} let:current bind:progress={introProgress}>
            {#each data.photos as photo, i}
                <section class="absolute left-0 right-0 top-0 bottom-0 lg:mx-auto w-[100vw] h-full lg:w-[95vw] p-2 lg:p-10" class:current={current === i}>
                    <div class="h-full flex flex-col justify-around lg:justify-between py-[8vh]" style="opacity: {1 - Math.abs($introProgress - i)}">
                        <p>{photo.slice(photo.lastIndexOf('/')+1)}</p>
                        <img class="h-[80vh] lg:h-[70vh] lg:w-[90vw] object-cover lg:object-contain" src={photo} alt={photo}/>
                    </div>
                </section>
            {/each}

            <div class="dots">
                {#each data.photos as p, i}
                    <div class="dot" in:fade class:active={current === i}></div>
                {/each}
            </div>
        </Swipeable>
    </div>
</div>

<style>
    .fullpage {
        position: absolute;
        width: 100%;
        height: 100%;
    }
    .slides {
        text-align: center;
    }
    .current {
        transform: scale(1);
    }

    .dots {
        position: absolute;
        bottom: 5%;
        left: 50%;
        transform: translateX(-50%);
        line-height: 0;

        @apply
            flex flex-row gap-2
    }
    .dot {
        @apply
            w-2 h-2 bg-black dark:bg-white rounded-xl
    }
    .dot.active {
        @apply
            bg-white dark:bg-white border-black border-[1px]
    }
    #wrapper {
        position: absolute;
        top: 0;
        left: 0;
        height: 100vh;
        width: 100vw;
        margin: 0;

        @apply
            bg-white dark:bg-black
    }
</style>
