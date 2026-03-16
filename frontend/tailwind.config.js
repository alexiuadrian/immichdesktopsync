/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{svelte,ts,js}'],
  theme: {
    extend: {
      colors: {
        immich: {
          primary: '#4250af',
          dark: '#1e1e2e',
          surface: '#313244',
          overlay: '#45475a',
          text: '#cdd6f4',
          subtext: '#a6adc8',
        },
      },
    },
  },
  plugins: [],
};
