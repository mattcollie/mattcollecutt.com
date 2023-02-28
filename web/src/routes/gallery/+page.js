export async function load({ params }) {
    const galleryModules = import.meta.glob('../../lib/galleries/*/*.webp');
    const galleries = {};
    for(const p in galleryModules) {
        const [name, date] = p.slice(20, p.lastIndexOf('/')).split('.');
        if(!(name in galleries)) {
            galleries[name] = {
                date: date,
                photos: []
            };
        }

        galleries[name].photos.push(p.replace('../../', 'src/'));
    }

    return {galleries: galleries};
}