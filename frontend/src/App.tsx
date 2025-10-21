import { useState } from "react";
import { useNavigate } from "react-router";

export default function App() {
  const [darkMode, setDarkMode] = useState(false);

  const navigate = useNavigate();

  return (
    <div className={darkMode ? "dark" : ""}>
      <div className="min-h-screen flex flex-col items-center justify-center bg-background dark:bg-dark-background text-text dark:text-dark-text transition-colors duration-300">
        <h1 className="text-4xl font-heading mb-6">Tailwind + Dark Mode</h1>
        <p className="text-3xl mb-6 bg-red-500 rounded-2xl p-4">
          This is an example of using custom fonts and colors.
        </p>
        <button
          className="px-6 py-3 rounded-md bg-button hover:bg-button-hover dark:bg-dark-button dark:hover:bg-dark-button-hover text-white transition-colors duration-300"
          onClick={() => setDarkMode(!darkMode)}
        >
          Toggle Dark Mode
        </button>

        <button onClick={() => navigate("/abc")}>Go to abc</button>
      </div>
    </div>
  );
}
