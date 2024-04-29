/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./**/*.{razor,html}'],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
	  require('daisyui'),
  ],
  daisyui: {
    themes: ['nord', 'dim'],
    darkTheme: 'dim',
  },
}