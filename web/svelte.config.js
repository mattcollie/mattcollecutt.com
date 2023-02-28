import adapter from '@sveltejs/adapter-static';
import preprocess from "svelte-preprocess";
import {mdsvex} from "mdsvex";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
			pages: 'build',
			assets: 'build',
			fallback: 'index.html'
		}),
		prerender: {entries: ['/gallery/[slug]', '/blog/[slug]', '/', '/about', '/blog', '/gallery']}
	},
	extensions: ['.svelte', '.md'],
	preprocess: [
		preprocess({
			postcss: true
		}),
		mdsvex({ extensions: ['.md'] })
	]
};

export default config;
