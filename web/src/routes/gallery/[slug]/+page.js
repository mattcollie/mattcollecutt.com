export async function load({ params }) {
    const photos = await import.meta.glob(`../../../lib/galleries/*/*.webp`);
    const galleryPhotos = [];
    for(const p in photos) {
        if(p.indexOf(params.slug.replaceAll('_', ' ')) !== -1) {
            const photo = await photos[p]();
            galleryPhotos.push(photo.default);
        }
    }
    return {
        name: params.slug.replaceAll('_', ' '),
        photos: galleryPhotos
    };
}