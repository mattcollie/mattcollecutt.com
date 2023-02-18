const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  purge: [],
  darkMode: 'class',
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      fontFamily: {
        serif: [
          "garamond",
          ...defaultTheme.fontFamily.serif
        ]
      }
    },
    colors: {
      black: '#282828',
      grey: '#767676',
      white: '#F8F8F8'
    }
  },
  plugins: [],
}
