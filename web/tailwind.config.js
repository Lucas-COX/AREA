/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.tsx",
    "./components/**/*.tsx"
  ],
  theme: {
    extend: {
      colors: {
        "primary": "#1976d2",
        "secondary": "#9c27b0",
        "error": "#d32f2f",
        "warning": "#ed6c02",
        "success": "#2e7d32",
      }
    },
  },
  plugins: [],
}
