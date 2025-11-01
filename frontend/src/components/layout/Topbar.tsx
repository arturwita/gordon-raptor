import { type FC } from "react";
import { useNavigate } from "react-router";

import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { ThemeButton } from "../ThemeButton";
import type { UserJwtPayload } from "../../types/user";
import { AppRoutes } from "../../routes/AppRoutes";

interface TopbarProps {
  user?: UserJwtPayload;
  onLogout: () => void;
}

export const Topbar: FC<TopbarProps> = ({ user, onLogout }) => {
  const navigate = useNavigate();

  return (
    <header className="w-full h-16 px-6 flex items-center justify-between bg-white dark:bg-gray-900 border-b dark:border-gray-700 mb-4">
      <div
        className="flex items-center gap-3 cursor-pointer select-none"
        onClick={() => navigate(AppRoutes.Recipes)}
      >
        <img
          src="/logo-transparent.png"
          alt="Logo"
          className="w-9 h-9 object-cover"
        />
        <span className="text-xl font-bold text-gray-900 dark:text-gray-100">
          Gordon Raptor
        </span>
      </div>

      <div className="flex items-center gap-6">
        <ThemeButton />

        {user ? (
          <DropdownMenu>
            <DropdownMenuTrigger className="flex items-center gap-3 focus:outline-none">
              <div className="text-right leading-tight">
                <div className="font-semibold text-gray-900 dark:text-gray-100">
                  {getUserDisplayName(user)}
                </div>
                <div className="text-xs text-gray-600 dark:text-gray-400">
                  {user?.email}
                </div>
              </div>

              <img
                src={user?.picture}
                alt="picture"
                className="w-10 h-10 rounded-full border dark:border-gray-700"
              />
            </DropdownMenuTrigger>

            <DropdownMenuContent align="end" className="w-40">
              <DropdownMenuItem onClick={onLogout}>Log out</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <div className="w-10 h-10 rounded-full bg-gray-300 dark:bg-gray-700 animate-pulse" />
        )}
      </div>
    </header>
  );
};

const getUserDisplayName = (user?: {
  firstName?: string;
  lastName?: string;
}) => {
  const defaultValue = "Anonymous";
  if (!user) return defaultValue;

  const first = user.firstName?.trim();
  const last = user.lastName?.trim();
  const fullName = [first, last].filter(Boolean).join(" ");

  return fullName || defaultValue;
};
