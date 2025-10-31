import { type FC, memo } from "react";
import { GoogleLoginButton } from "../components/GoogleLoginButton";
import { appConfig } from "../lib/config";
import { ThemeToggle } from "../components/ThemeToggle";
import Footer from "../components/layout/Footer";

const Home: FC = () => {
  const handleGoogleLogin = () => {
    window.location.href = `${appConfig.baseApiUrl}/auth/google/login`;
  };

  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-white text-gray-900 transition-colors duration-300 dark:bg-gray-900 dark:text-gray-100 px-6">
      <ThemeToggle />

      <main className="text-center max-w-md">
        <h1 className="text-4xl font-bold mb-4">Gordon Raptor</h1>
        <p className="text-lg mb-4 text-gray-600 dark:text-gray-300">
          Cook along with an AI assistant in real time. <br /> Get started now.
        </p>

        <GoogleLoginButton onClick={() => handleGoogleLogin()} />
      </main>

      <Footer />
    </div>
  );
};

export default memo(Home);
