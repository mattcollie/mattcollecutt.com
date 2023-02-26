export async function load({ params }) {
    const posts = import.meta.glob('../../lib/posts/*.md');
    let blogPosts = [];
    for (const k of Object.keys(posts)) {
        let post = await posts[k]();
        blogPosts.push(post.metadata);
    }
    return {posts: blogPosts.sort((a, b) => new Date(b.date) - new Date(a.date))};
}