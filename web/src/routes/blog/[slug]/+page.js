export const prerender = 'auto';

export async function load({ params }) {
    const post = await import(`./web/src/lib/posts/${params.slug}.md`);
    const {title, date} = post.metadata;
    const content = post.default;

    return {
        content,
        title,
        date
    }
}