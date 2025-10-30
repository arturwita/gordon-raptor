import { Button } from "@/components/ui/button";
import type { FC } from "react";
import { FcGoogle } from "react-icons/fc";

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
      <FcGoogle className="w-5 h-5" />
      <span className="font-medium">Login with Google</span>
    </Button>
  );
};
