/** @type {import('tailwindcss').Config} */
export default {
  content: ["./output/**/*.html"],
  theme: {
    fontFamily: {
      main: ["Fira Sans", "Trebuchet MS", "Helvetica", "sans-serif"],
    },
    extend: {
      colors: {
        primary: "#022a2a",
      },
    },
  },
  plugins: [],
};
