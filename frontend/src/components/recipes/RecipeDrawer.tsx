import {
  Drawer,
  DrawerContent,
  DrawerHeader,
  DrawerTitle,
  DrawerFooter,
  DrawerDescription,
} from "@/components/ui/drawer";
import { Button } from "@/components/ui/button";
import { type FC, useState } from "react";
import { useAuth } from "../../context/AuthContext";
import { X } from "lucide-react";
import type { RecipeDto } from "../../types/recipe";

interface RecipeDrawerProps {
  recipe: RecipeDto;
  onClose: () => void;
}

export const RecipeDrawer: FC<RecipeDrawerProps> = ({ recipe, onClose }) => {
  const { token } = useAuth();
  const [loading, setLoading] = useState(false);

  const startCooking = async () => {
    try {
      setLoading(true);

      // await fetch("/recipes/cook", {
      //   method: "POST",
      //   headers: {
      //     Authorization: `Bearer ${token}`,
      //     "Content-Type": "application/json",
      //   },
      //   body: JSON.stringify({ recipeId: recipe.id }),
      // });
      console.log("Clicked start cooking");

      onClose();
    } finally {
      setLoading(false);
    }
  };

  return (
    <Drawer open={true} direction="right" onOpenChange={onClose}>
      <DrawerContent className="w-[420px] ml-auto border-l dark:border-gray-700">
        <Button
          variant="ghost"
          size="icon"
          className="absolute top-4 right-4"
          onClick={onClose}
        >
          <X className="w-5 h-5" />
        </Button>

        <DrawerHeader>
          <DrawerTitle>{recipe.name}</DrawerTitle>
          <DrawerDescription>{recipe.description}</DrawerDescription>
        </DrawerHeader>

        <div className="px-6 py-4 space-y-4">
          <h3 className="font-medium text-lg">Ingredients</h3>

          <ul className="space-y-1">
            {Object.entries(recipe.ingredients).map(([ingredient, qty]) => (
              <li
                key={ingredient}
                className="flex justify-between border-b dark:border-gray-700 pb-1"
              >
                <span>{ingredient}</span>
                <span className="text-gray-600 dark:text-gray-300">{qty}</span>
              </li>
            ))}
          </ul>
        </div>

        <DrawerFooter>
          <Button onClick={startCooking} disabled={loading}>
            {loading ? "Starting..." : "Start Cooking"}
          </Button>
        </DrawerFooter>
      </DrawerContent>
    </Drawer>
  );
};
