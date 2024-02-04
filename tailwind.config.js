const defaultTheme = require('tailwindcss/defaultTheme');

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{html,js,templ}"],
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
      neutral: {
        50: "#fafafa",
        100: "#f5f5f5",
        200: "#e5e5e5",
        300: "#d4d4d4",
        400: "#a3a3a3",
        500: "#737373",
        600: "#525252",
        700: "#404040",
        750: "#323232",
        800: "#262626",
        900: "#171717",
        950: "#0a0a0a"
      }
    }
  },
  plugins: []
}

