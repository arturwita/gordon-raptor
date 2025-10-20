// tailwind.config.js
module.exports = {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx,css}"],
  darkMode: "class",
  theme: {
    extend: {
      colors: {
        background: "#FAEBD7",
        text: "#4B2E05",
        button: "#D2691E",
        "button-hover": "#C45A1A",
        accent: "#B22222",
        "dark-background": "#1C1B1A",
        "dark-text": "#FAEBD7",
        "dark-button": "#C45A1A",
        "dark-button-hover": "#D2691E",
        "dark-accent": "#E57373",
      },
      fontFamily: {
        heading: ["Fredoka", "sans-serif"],
        body: ["Nunito Sans", "sans-serif"],
      },
    },
  },
  plugins: [],
};
