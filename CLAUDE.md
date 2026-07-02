# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project

Personal portfolio website built with SvelteKit 5, Tailwind CSS v4, and shadcn-svelte. Deployed via Cloudflare Pages.

## Development Commands

```bash
# Run dev server
cd ui/portfolio && npm run dev

# Build for production
cd ui/portfolio && npm run build

# Preview production build
cd ui/portfolio && npm run preview
```

## Architecture

All source code lives under `ui/portfolio/`.

- **src/routes/** — SvelteKit file-based routing. `+page.svelte` (home, RFC-document style single page), `+layout.svelte` (root layout: document top bar, theme toggle, PostHog init), `api/health/` and `api/ping/` (JSON endpoints)
- **src/lib/components/ui/** — shadcn-svelte primitives (`button`, `sheet`) — currently unused by the page
- **src/lib/utils.ts** — `cn()` helper for Tailwind class merging
- **src/app.css** — Theme tokens as CSS custom properties (light/dark via `prefers-color-scheme` + `[data-theme]` override), mapped into Tailwind v4 via `@theme inline`; document utilities (`doc-a`, `doc-h2`, `placeholder-note`)
- **src/app.html** — HTML shell with meta tags, Google Fonts (Jost, Newsreader, IBM Plex Mono), pre-paint theme resolution script, PostHog analytics
- **static/** — Favicon, SVG icons, robots.txt

## Design system

- RFC/document aesthetic: single ~62ch column, numbered sections, hairline rules, no cards
- Fonts: Jost (masthead/display — to be replaced by a hand-traced SVG version), Newsreader (body), IBM Plex Mono (dates, domains, metadata)
- One accent color only: bottle green `#1F4A38` (light) / sage `#8FBCA5` (dark) — from the Robertson Accounting sign in Paeroa
- Theme: OS preference by default; `theme` key in localStorage overrides (`light`/`dark`); `?theme=` URL param forces a theme for previews
- Italic grey `placeholder-note` spans mark copy the owner still needs to write in his own words

## Key Patterns

- SvelteKit 5 with Svelte runes (`$state`, `$derived`, `$props`, `$bindable`)
- Tailwind CSS v4 with `@theme inline` mapping CSS custom properties to utility tokens
- `adapter-cloudflare` for Cloudflare Pages deployment

## Deployment

- Deployed via Cloudflare Pages (connected to repo, auto-deploys on push)
- Build command: `npm run build` with root directory `ui/portfolio`
