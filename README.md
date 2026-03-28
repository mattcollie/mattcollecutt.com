# mattcollecutt.com

Personal portfolio website built with SvelteKit 5, Tailwind CSS v4, and shadcn-svelte. Deployed via Cloudflare Pages.

**Live site:** [mattcollecutt.com](https://mattcollecutt.com)

## Tech Stack

- **Framework:** SvelteKit 5 (Svelte runes)
- **Styling:** Tailwind CSS v4, shadcn-svelte, tailwind-variants
- **Deployment:** Cloudflare Pages via `adapter-cloudflare`
- **Language:** TypeScript

## Getting Started

```bash
cd ui/portfolio
npm install
npm run dev
```

## Build

```bash
cd ui/portfolio
npm run build
npm run preview  # preview production build locally
```

## Project Structure

All source code lives under `ui/portfolio/`.

```
ui/portfolio/src/
├── routes/          # SvelteKit file-based routing
├── lib/components/  # Reusable Svelte components (Sidebar, Banner, etc.)
├── lib/components/ui/  # shadcn-svelte primitives (button, sheet)
├── lib/utils.ts     # Tailwind class merging helper
├── app.css          # Tailwind v4 theme config
└── app.html         # HTML shell with meta tags and analytics
```
