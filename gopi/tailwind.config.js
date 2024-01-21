const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./*/*.{html,js,templ}"],
  theme: {
    extend: {
      fontFamily: {
        serif: [
          "Inter",
          ...defaultTheme.fontFamily.serif
        ],
        mono: [
          "consolas",
          ...defaultTheme.fontFamily.mono
        ]
      }
    },
    colors: {
      white: '#F8F8F8',
      black: '#111111',
      primary: '#111111',
      secondary: '#1B1B1B',
      stroke: '#222222',
      'stroke-light': '#444444',
      grey: '#999',
      transparent: '#FFFFFFFF',
    }
  },
  plugins: []
}

