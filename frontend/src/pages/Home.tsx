import { type FC, memo, useEffect, useState } from "react";
import { Toggle } from "@/components/ui/toggle";
import { Sun, Moon } from "lucide-react";
import { GoogleLoginButton } from "../components/GoogleLoginButton";
import { appConfig } from "../lib/config";

const Home: FC = () => {
  const [darkMode, setDarkMode] = useState(true);

  const handleGoogleLogin = () => {
    window.location.href = `${appConfig.baseApiUrl}/auth/google/login`;
  };

  useEffect(() => {
    const root = window.document.documentElement;
    if (darkMode) {
      root.classList.add("dark");
    } else {
      root.classList.remove("dark");
    }
  }, [darkMode]);

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-white text-gray-900 transition-colors duration-300 dark:bg-gray-900 dark:text-gray-100 px-6">
      <div className="absolute top-6 right-6">
        <Toggle
          pressed={darkMode}
          onPressedChange={setDarkMode}
          aria-label="Toggle theme"
          className="border border-gray-300 dark:border-gray-700 bg-white dark:bg-gray-800 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-full p-2 transition"
        >
          {darkMode ? (
            <Sun className="w-5 h-5 text-yellow-400" />
          ) : (
            <Moon className="w-5 h-5 text-gray-700" />
          )}
        </Toggle>
      </div>

      <main className="text-center max-w-md">
        <h1 className="text-4xl font-bold mb-4">Gordon Raptor</h1>
        <p className="text-lg mb-4 text-gray-600 dark:text-gray-300">
          Cook along with an AI assistant in real time. <br /> Get started now.
        </p>

        <GoogleLoginButton onClick={() => handleGoogleLogin()} />
      </main>

      <footer className="absolute bottom-4 text-sm text-gray-500 dark:text-gray-400">
        © {new Date().getFullYear()} Gordon Raptor. All rights reserved.
      </footer>
    </div>
  );
};

export default memo(Home);
