import { memo, type FC } from "react";

const Recipes: FC = () => {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-gray-900">
      <h1 className="text-3xl font-semibold dark:text-gray-100">
        🍳 Your Recipes
      </h1>
    </div>
  );
};

export default memo(Recipes);
