import { error } from '@sveltejs/kit';
import { notes } from '$lib/notes';

export function load({ params }) {
	const note = notes.find((n) => n.slug === params.slug);
	if (!note) error(404, 'No such note');
	return { note };
}
