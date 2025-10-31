import { type FC } from "react";
import { Button } from "@/components/ui/button";
import type { RecipeDto } from "../../types/recipe";

interface RecipeCardProps {
  recipe: RecipeDto;
  onDetails: () => void;
}

const defaultPictureUrl = "public/logo.png";

export const RecipeCard: FC<RecipeCardProps> = ({ recipe, onDetails }) => {
  return (
    <div className="bg-white dark:bg-gray-800 rounded-xl shadow hover:shadow-lg transition p-4 flex flex-col">
      <img
        src={recipe.picture || defaultPictureUrl}
        alt={recipe.name}
        className="w-full h-40 object-cover rounded-md mb-3"
      />

      <h2 className="font-semibold text-lg mb-4">{recipe.name}</h2>

      <Button onClick={onDetails} className="mt-auto">
        Details
      </Button>
    </div>
  );
};
