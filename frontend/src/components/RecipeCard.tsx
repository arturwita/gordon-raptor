import React from "react";
import type { RecipeDto } from "../api/recipes";

interface Props {
  recipe: RecipeDto;
}

export const RecipeCard: React.FC<Props> = ({ recipe }) => {
  return (
    <div className="group relative bg-background-light dark:bg-background-dark rounded-xl overflow-hidden transition-all duration-300 ease-in-out hover:shadow-xl">
      <div className="relative w-full h-72">
        <div
          className="w-full h-full bg-cover bg-center transition-transform duration-300 group-hover:scale-105"
          style={{ backgroundImage: `url(${recipe.imageUrl})` }}
        ></div>
        <div className="absolute inset-0 bg-gradient-to-t from-black/80 to-transparent"></div>
        <h3 className="absolute bottom-4 left-4 text-xl font-bold text-white z-10">
          {recipe.name}
        </h3>
      </div>

      {/* Hover overlay */}
      <div className="absolute inset-0 bg-black/80 flex flex-col justify-end p-4 transition-opacity duration-300 opacity-0 group-hover:opacity-100">
        <h3 className="text-xl font-bold text-white mb-2">{recipe.name}</h3>
        <div className="flex flex-wrap gap-2 mb-4">
          {Object.entries(recipe.ingredients).map(([ingredient, quantity]) => (
            <div>
              <span
                key={ingredient}
                className="bg-primary/20 text-primary text-xs px-2 py-1 rounded-full font-medium"
              >
                {ingredient}
              </span>
              <span
                key={quantity}
                className="bg-primary/20 text-primary text-xs px-2 py-1 rounded-full font-medium"
              >
                {ingredient}
              </span>
            </div>
          ))}
        </div>

        <div className="flex justify-between items-center">
          <a
            className="text-sm font-bold text-primary hover:underline"
            href="#"
          >
            View Details
          </a>
          <div className="flex space-x-4 opacity-75">
            <button className="text-white hover:text-primary transition-colors text-sm flex items-center">
              <span className="material-symbols-outlined text-base">edit</span>
            </button>
            <button className="text-white hover:text-red-500 transition-colors text-sm flex items-center">
              <span className="material-symbols-outlined text-base">
                delete
              </span>
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};
