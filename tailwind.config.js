/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*.html", "./views/*/*.html"],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}