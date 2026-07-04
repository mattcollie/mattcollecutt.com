export type Photo = {
	src: string;
	full: string;
	alt: string;
	caption: string;
	/** mosaic hint: 'wide' spans two columns, 'std' one */
	size: 'wide' | 'std';
};

// Add new photos here: drop the export in the repo root, ask Claude to
// process it, append an entry. Order is display order; three 'wide' +
// six 'std' currently fills the 3-column mosaic with no gaps.
export const photos: Photo[] = [
	{
		src: '/photos/pheasant.webp',
		full: '/photos/pheasant-full.webp',
		alt: 'A cock pheasant seen through a gap in dark bush',
		caption: 'pheasant, Te Aroha wetlands, September 2023',
		size: 'wide'
	},
	{
		src: '/photos/stream.webp',
		full: '/photos/stream-full.webp',
		alt: 'A boulder stream running out of dark misty conifer forest',
		caption: 'stream, April 2023',
		size: 'std'
	},
	{
		src: '/photos/swan-fog.webp',
		full: '/photos/swan-fog-full.webp',
		alt: 'A black swan almost disappearing into thick fog on still water',
		caption: 'black swan in fog, August 2023',
		size: 'wide'
	},
	{
		src: '/photos/chaffinch.webp',
		full: '/photos/chaffinch-full.webp',
		alt: 'A chaffinch perched on dead flax stalks against warm gold',
		caption: 'chaffinch on flax, June 2023',
		size: 'std'
	},
	{
		src: '/photos/tui.webp',
		full: '/photos/tui-full.webp',
		alt: 'A tūī looking up through a blur of branches',
		caption: 'tūī, May 2023',
		size: 'std'
	},
	{
		src: '/photos/kaka.webp',
		full: '/photos/kaka-full.webp',
		alt: 'A kākā eating in the top of a pine tree',
		caption: 'kākā, July 2023',
		size: 'std'
	},
	{
		src: '/photos/floodlights.webp',
		full: '/photos/floodlights-full.webp',
		alt: 'Weathered floodlights above a rail on a grey morning',
		caption: 'floodlights, August 2023',
		size: 'std'
	},
	{
		src: '/photos/swan.webp',
		full: '/photos/swan-full.webp',
		alt: 'A black swan up close, red eye and water drops on black feathers',
		caption: 'black swan, September 2023',
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
