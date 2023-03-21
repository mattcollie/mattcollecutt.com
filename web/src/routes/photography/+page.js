export async function load({params}) {
    const photoGlobs = import.meta.glob('../../lib/galleries/*/*.webp');
    const photos = [];
    for (const p in photoGlobs) {
        const [gallery, date] = p.slice(20, p.lastIndexOf('/')).split('.');
        const name = p.slice(p.lastIndexOf('/') + 1);
        photos.push({gallery: gallery, name: name, date: date, photo: await photoGlobs[p]()})
    }

    return {photos: photos};
}