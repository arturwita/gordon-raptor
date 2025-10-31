import { Button } from "@/components/ui/button";
import type { FC } from "react";

interface GoogleLoginButtonProps {
  onClick?: () => void;
}

export const GoogleLoginButton: FC<GoogleLoginButtonProps> = ({ onClick }) => {
  return (
    <Button
      onClick={onClick}
      variant="outline"
      className="flex items-center justify-center gap-2 w-full bg-white text-gray-700 border border-gray-300 hover:bg-gray-100 dark:bg-gray-800 dark:text-gray-100 dark:hover:bg-gray-700 dark:border-gray-700 transition-all"
    >
      {GoogleIcon()}
      <span className="font-medium">Login with Google</span>
    </Button>
  );
};

const GoogleIcon = () => (
  <svg
    className="w-5 h-5"
    xmlns="http://www.w3.org/2000/svg"
    viewBox="0 0 48 48"
  >
    <path
      fill="#EA4335"
      d="M24 9.5c3.94 0 7.02 1.63 9.19 3.37l6.77-6.58C36.42 2.63 30.78 0 24 0 14.63 0 6.47 5.35 2.56 13.09l7.89 6.12C12.18 13.53 17.56 9.5 24 9.5z"
    />
    <path
      fill="#34A853"
      d="M46.98 24.55c0-1.6-.15-3.11-.43-4.55H24v9.09h13.02c-.56 2.86-2.24 5.29-4.77 6.94l7.38 5.73C43.98 38.17 46.98 31.97 46.98 24.55z"
    />
    <path
      fill="#4A90E2"
      d="M9.45 28.79a14.45 14.45 0 0 1-.75-4.24c0-1.47.27-2.88.75-4.24l-7.89-6.12A23.936 23.936 0 0 0 0 24.55c0 3.81.9 7.43 2.56 10.36l6.89-6.12z"
    />
    <path
      fill="#FBBC05"
      d="M24 48c6.48 0 11.91-2.14 15.87-5.82l-7.38-5.73c-2.05 1.38-4.71 2.17-8.49 2.17-6.44 0-11.82-4.03-13.55-9.7l-7.89 6.12C6.47 42.65 14.63 48 24 48z"
    />
  </svg>
);
