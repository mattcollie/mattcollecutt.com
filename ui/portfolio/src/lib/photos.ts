export type Photo = {
	src: string;
	full: string;
	alt: string;
	caption: string;
	/** mosaic hint: 'wide' spans two columns, 'std' one */
	size: 'wide' | 'std';
};

// Add new photos here: drop the RAF/export in the repo root, ask Claude to
// process it, append an entry. Order is display order.
export const photos: Photo[] = [
	{
		src: '/photos/pheasant.webp',
		full: '/photos/pheasant-full.webp',
		alt: 'A cock pheasant seen through a gap in dark bush',
		caption: 'pheasant, Te Aroha wetlands, September 2023',
		size: 'wide'
	},
	{
		src: '/photos/ship.webp',
		full: '/photos/ship-full.webp',
		alt: 'A Maersk container ship on dark water at dusk',
		caption: 'container ship, Tauranga, March 2025',
		size: 'std'
	}
];
