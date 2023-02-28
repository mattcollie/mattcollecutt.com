export async function load({ params }) {
    const photos = await import.meta.glob(`../../../lib/galleries/*/*.webp`);
    return {
        name: params.slug.replaceAll('_', ' '),
        photos: Object.keys(photos).filter(p => p.indexOf(params.slug.replaceAll('_', ' ')) !== -1)
    };
}