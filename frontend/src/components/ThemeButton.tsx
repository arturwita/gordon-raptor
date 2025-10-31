import { type FC, useState } from "react";
import { Sun, Moon } from "lucide-react";
import { Button } from "./ui/button";
import { Theme, type ThemeType } from "../lib/consts";

export const ThemeButton: FC = () => {
  const [theme, setTheme] = useState<ThemeType>(() =>
    document.documentElement.classList.contains("dark")
      ? Theme.Dark
      : Theme.Light
  );

  const toggleTheme = () => {
    const newTheme = theme === Theme.Dark ? Theme.Light : Theme.Dark;
    setTheme(newTheme);

    if (newTheme === Theme.Dark) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }

    localStorage.setItem("theme", newTheme);
  };

  return (
    <Button variant="outline" size="icon" onClick={toggleTheme}>
      {theme === Theme.Dark ? (
        <Sun className="h-5 w-5" />
      ) : (
        <Moon className="h-5 w-5" />
      )}
    </Button>
  );
};
