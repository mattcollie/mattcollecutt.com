const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'class',
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontFamily: {
        serif: [
          ...defaultTheme.fontFamily.serif
        ],
        mono: [
          "consolas",
          ...defaultTheme.fontFamily.mono
        ]
      }
    },
    colors: {
      black: '#282828',
      white: '#F8F8F8',
      transparent: '#FFFFFFFF',
      grey: {
        DEFAULT: '#767676',
        50: '#E4E4E4',
        100: '#D1D1D1',
        200: '#BFBFBF',
        300: '#ADADAD',
        400: '#9B9B9B',
        500: '#888888',
        600: '#676767',
        700: '#595959',
        800: '#4A4A4A',
        900: '#3B3B3B'
      }
    }
  },
  plugins: [],
}
